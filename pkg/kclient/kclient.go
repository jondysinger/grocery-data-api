package kclient

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jondysinger/grocery-data-api/pkg/models"
)

// Struct for interaction with Kroger API
type KClient struct {
	baseUrl   string
	id        string
	secret    string
	chain     string
	token     string
	netClient *http.Client
}

// Creates a new KClient
func New(baseUrl string, id string, secret string, chain string) (*KClient, error) {
	if baseUrl == "" {
		return nil, errors.New("parameter 'baseUrl' is required")
	} else if id == "" {
		return nil, errors.New("parameter 'id' is required")
	} else if secret == "" {
		return nil, errors.New("parameter 'secret' is required")
	} else if chain == "" {
		return nil, errors.New("parameter 'chain' is required")
	}

	var netClient = &http.Client{
		Timeout: time.Second * 180,
	}

	client := KClient{
		baseUrl:   baseUrl,
		id:        id,
		secret:    secret,
		chain:     chain,
		netClient: netClient,
	}
	return &client, nil
}

// Gets the error description from a Kroger API error response (status codes 400 or 500)
func getApiErrorDesc(body []byte) (string, error) {
	var errRes models.ApiErrorResponse
	if err := json.Unmarshal(body, &errRes); err != nil {
		return "", fmt.Errorf("failed to deserialize JSON response body: %v", err)
	}
	return errRes.Errors.Reason, nil
}

// Gets the error description from a Kroger Authorization error response (status code 401)
func getAuthErrorDesc(body []byte) (string, error) {
	var errRes models.AuthErrorResponse
	if err := json.Unmarshal(body, &errRes); err != nil {
		return "", fmt.Errorf("failed to deserialize JSON response body: %v", err)
	}
	return errRes.ErrorDescription, nil
}

// Gets the error description from the response body of a Kroger API call based on the status code
func getResponseError(statusCode int, status string, body []byte) error {
	switch statusCode {
	case 404:
		return errors.New("URL endpoint not found")
	case 400, 500:
		desc, err := getApiErrorDesc(body)
		if err != nil {
			desc = err.Error()
		}
		return fmt.Errorf("request failed with status '%s', Error description: %s", status, desc)
	case 401:
		desc, err := getAuthErrorDesc(body)
		if err != nil {
			desc = err.Error()
		}
		return fmt.Errorf("request failed with status '%s', Error description: %s", status, desc)
	}
	return fmt.Errorf("unknown error with status '%s'", status)
}

// Retrieves a client authentication OAuth2 token
func (client *KClient) GetAuthToken() error {
	reqUrl := fmt.Sprintf("%s/connect/oauth2/token", client.baseUrl)
	payload := strings.NewReader("grant_type=client_credentials&scope=product.compact")

	// API Reference: https://developer.kroger.com/reference#operation/accessToken
	req, err := http.NewRequest("POST", reqUrl, payload)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", client.id, client.secret)))))
	res, err := client.netClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	} else if res.StatusCode != 200 {
		return getResponseError(res.StatusCode, res.Status, body)
	}

	var authRes models.AuthorizationResponse
	if err := json.Unmarshal(body, &authRes); err != nil {
		return fmt.Errorf("failed to deserialize JSON response body: %v", err)
	}

	client.token = authRes.AccessToken
	return nil
}

// Gets Kroger locations by zip code
func (client *KClient) GetLocations(zipCode string, filterLimit int) (*models.LocationsResponse, error) {
	if client.token == "" {
		return nil, errors.New("client has no OAuth2 token, call GetAuthToken first")
	} else if zipCode == "" {
		return nil, errors.New("parameter 'zipCode' is required")
	} else if digits := countDigits(zipCode); digits < 5 || digits != len(zipCode) {
		return nil, fmt.Errorf("parameter 'zipCode' value '%s' is invalid. Must be a number with 5 digits", zipCode)
	} else if filterLimit < 0 || filterLimit > 200 {
		return nil, fmt.Errorf("parameter 'filterLimit' value %d is invalid. Valid values are 0 to 200", filterLimit)
	}

	reqUrl := fmt.Sprintf("%s/locations?filter.chain=%s&filter.zipCode.near=%s", client.baseUrl, client.chain, url.QueryEscape(zipCode))

	if filterLimit > 0 {
		reqUrl = fmt.Sprintf("%s&filter.limit=%d", reqUrl, filterLimit)
	}

	// API Reference: https://developer.kroger.com/reference#operation/SearchLocations
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.token))
	res, err := client.netClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	} else if res.StatusCode != 200 {
		return nil, getResponseError(res.StatusCode, res.Status, body)
	}

	var locsResp models.LocationsResponse
	if err := json.Unmarshal(body, &locsResp); err != nil {
		return nil, fmt.Errorf("failed to deserialize JSON response body: %v", err)
	}

	return &locsResp, nil
}

// Gets Kroger products based on a given search term. A locationId is optional and if given the product information
// will contain stock levels and pricing.
func (client *KClient) GetProducts(filterTerm string, locationId string, filterOffset int, filterLimit int) (*models.ProductsResponse, error) {
	if client.token == "" {
		return nil, errors.New("client has no OAuth2 token, call GetAuthToken first")
	} else if filterTerm == "" {
		return nil, errors.New("parameter 'filterTerm' is required")
	} else if filterOffset < 0 || filterOffset > 1000 {
		return nil, fmt.Errorf("parameter 'filterOffset' value %d is invalid. Valid values are 0 to 1000", filterOffset)
	} else if filterLimit < 0 || filterLimit > 50 {
		return nil, fmt.Errorf("parameter 'filterLimit' value %d is invalid. Valid values are 0 to 50", filterLimit)
	}

	reqUrl := fmt.Sprintf("%s/products?filter.term=%s", client.baseUrl, url.QueryEscape(filterTerm))
	if locationId != "" {
		reqUrl = fmt.Sprintf("%s&filter.locationId=%s", reqUrl, url.QueryEscape(locationId))
	}

	if filterOffset > 0 {
		reqUrl = fmt.Sprintf("%s&filter.start=%d", reqUrl, filterOffset)
	}

	if filterLimit > 0 {
		reqUrl = fmt.Sprintf("%s&filter.limit=%d", reqUrl, filterLimit)
	}

	// API Reference: https://developer.kroger.com/reference#operation/productGet
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.token))
	res, err := client.netClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	} else if res.StatusCode != 200 {
		return nil, getResponseError(res.StatusCode, res.Status, body)
	}

	var prodResp models.ProductsResponse
	if err := json.Unmarshal(body, &prodResp); err != nil {
		return nil, fmt.Errorf("failed to deserialize JSON response body: %v", err)
	}

	return &prodResp, nil
}
