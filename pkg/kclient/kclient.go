package kclient

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type AuthorizationResponse struct {
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type ApiErrorResponse struct {
	Errors struct {
		TimeStamp int    `json:"timestamp"`
		Code      string `json:"code"`
		Reason    string `json:"reason"`
	} `json:"errors"`
}

type AuthErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type Location struct {
	LocationId string `json:"locationId"`
	Chain      string `json:"chain"`
	Address    struct {
		AddressLine1 string `json:"addressLine1"`
		City         string `json:"city"`
		State        string `json:"state"`
		ZipCode      string `json:"zipCode"`
		County       string `json:"county"`
	} `json:"address"`
	Geolocation struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
		LatLng    string  `json:"latLng"`
	} `json:"geolocation"`
	Name  string `json:"Name"`
	Hours struct {
		Timezone  string `json:"timezone"`
		GmtOffset string `json:"gmtOffset"`
		Open24    bool   `json:"open24"`
		Monday    struct {
			Open   string `json:"open"`
			Close  string `json:"close"`
			Open24 bool   `json:"open24"`
		} `json:"monday"`
		Tuesday struct {
			Open   string `json:"open"`
			Close  string `json:"close"`
			Open24 bool   `json:"open24"`
		} `json:"tuesday"`
		Wednesday struct {
			Open   string `json:"open"`
			Close  string `json:"close"`
			Open24 bool   `json:"open24"`
		} `json:"wednesday"`
		Thursday struct {
			Open   string `json:"open"`
			Close  string `json:"close"`
			Open24 bool   `json:"open24"`
		} `json:"thursday"`
		Friday struct {
			Open   string `json:"open"`
			Close  string `json:"close"`
			Open24 bool   `json:"open24"`
		} `json:"friday"`
		Saturday struct {
			Open   string `json:"open"`
			Close  string `json:"close"`
			Open24 bool   `json:"open24"`
		} `json:"saturday"`
		Sunday struct {
			Open   string `json:"open"`
			Close  string `json:"close"`
			Open24 bool   `json:"open24"`
		} `json:"sunday"`
	} `json:"hours"`
	Phone       string `json:"phone"`
	Departments []struct {
		DepartmentID string `json:"departmentId"`
		Name         string `json:"name"`
	} `json:"departments"`
}

type LocationsResponse struct {
	Data []Location `json:"data"`
	Meta struct {
		Pagination struct {
			Total int `json:"total"`
			Start int `json:"start"`
			Limit int `json:"limit"`
		} `json:"pagination"`
		Warnings []string `json:"warnings"`
	} `json:"meta"`
}

type Product struct {
	ProductId      string `json:"productId"`
	AisleLocations []struct {
		BayNumber          string `json:"bayNumber"`
		Description        string `json:"description"`
		Number             string `json:"number"`
		NumberOfFacings    string `json:"numberOfFacings"`
		SequenceNumber     string `json:"sequenceNumber"`
		Side               string `json:"side"`
		ShelfNumber        string `json:"shelfNumber"`
		ShelfPositionInBay string `json:"shelfPositionInBay"`
	} `json:"aisleLocations"`
	Brand         string   `json:"brand"`
	Categories    []string `json:"categories"`
	CountryOrigin string   `json:"countryOrigin"`
	Description   string   `json:"description"`
	Items         []struct {
		ItemId    string `json:"itemId"`
		Inventory struct {
			StockLevel string `json:"stockLevel"`
		} `json:"inventory"`
		Favorite    bool `json:"favorite"`
		Fulfillment struct {
			Curbside   bool `json:"curbside"`
			Delivery   bool `json:"delivery"`
			InStore    bool `json:"instore"`
			ShipToHome bool `json:"shiptohome"`
		} `json:"fulfillment"`
		Price struct {
			Regular                float32 `json:"regular"`
			Promo                  float32 `json:"promo"`
			RegularPerUnitEstimate float32 `json:"regularPerUnitEstimate"`
			PromoPerUnitEstimate   float32 `json:"promoPerUnitEstimate"`
		} `json:"price"`
		NationalPrice struct {
			Regular                float32 `json:"regular"`
			Promo                  float32 `json:"promo"`
			RegularPerUnitEstimate float32 `json:"regularPerUnitEstimate"`
			PromoPerUnitEstimate   float32 `json:"promoPerUnitEstimate"`
		} `json:"nationalPrice"`
		Size   string `json:"size"`
		SoldBy string `json:"soldBy"`
	} `json:"items"`
	ItemInformation struct {
		Depth  string `json:"depth"`
		Height string `json:"height"`
		Width  string `json:"width"`
	} `json:"itemInformation"`
	Temperature struct {
		Indicator     string `json:"indicator"`
		HeatSensitive bool   `json:"heatSensitive"`
	} `json:"temperature"`
	Images []struct {
		Id          string `json:"id"`
		Perspective string `json:"perspective"`
		Default     bool   `json:"default"`
		Sizes       []struct {
			Id   string `json:"id"`
			Size string `json:"size"`
			Url  string `json:"url"`
		} `json:"sizes"`
	} `json:"images"`
	Upc string `json:"upc"`
}

type ProductsResponse struct {
	Data []Product `json:"data"`
	Meta struct {
		Pagination struct {
			Total int `json:"total"`
			Start int `json:"start"`
			Limit int `json:"limit"`
		} `json:"pagination"`
		Warnings []string `json:"warnings"`
	} `json:"meta"`
}

// Struct for interaction with Kroger API
type KClient struct {
	baseUrl string
	id      string
	secret  string
	chain   string
	token   string
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

	client := KClient{
		baseUrl: baseUrl,
		id:      id,
		secret:  secret,
		chain:   chain,
	}
	return &client, nil
}

// Gets the error description from a Kroger API error response (status codes 400 or 500)
func getApiErrorDesc(body []byte) (string, error) {
	var errRes ApiErrorResponse
	if err := json.Unmarshal(body, &errRes); err != nil {
		return "", fmt.Errorf("failed to deserialize JSON response body: %v", err)
	}
	return errRes.Errors.Reason, nil
}

// Gets the error description from a Kroger Authorization error response (status code 401)
func getAuthErrorDesc(body []byte) (string, error) {
	var errRes AuthErrorResponse
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
	res, err := http.DefaultClient.Do(req)
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

	var authRes AuthorizationResponse
	if err := json.Unmarshal(body, &authRes); err != nil {
		return fmt.Errorf("failed to deserialize JSON response body: %v", err)
	}

	client.token = authRes.AccessToken
	return nil
}

// Counts the number of digits in a string
func countDigits(numberString string) (digits int) {
	for _, char := range strings.Split(numberString, "") {
		if _, err := strconv.Atoi(char); err == nil {
			digits++
		}
	}
	return digits
}

// Gets Kroger locations by zip code
func (client *KClient) GetLocations(zipCode string) ([]Location, error) {
	if client.token == "" {
		return nil, errors.New("client has no OAuth2 token, call GetAuthToken first")
	} else if zipCode == "" {
		return nil, errors.New("parameter 'zipCode' is required")
	} else if digits := countDigits(zipCode); digits < 5 || digits != len(zipCode) {
		return nil, fmt.Errorf("parameter 'zipCode' value '%s' is invalid. Must be a number with 5 digits", zipCode)
	}

	reqUrl := fmt.Sprintf("%s/locations?filter.chain=%s&filter.zipCode.near=%s", client.baseUrl, client.chain, url.QueryEscape(zipCode))

	// API Reference: https://developer.kroger.com/reference#operation/SearchLocations
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.token))
	res, err := http.DefaultClient.Do(req)
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

	var locsResp LocationsResponse
	if err := json.Unmarshal(body, &locsResp); err != nil {
		return nil, fmt.Errorf("failed to deserialize JSON response body: %v", err)
	}

	return locsResp.Data, nil
}

// Gets Kroger products based on a given search term. A locationId is optional and if given the product information
// will contain stock levels and pricing.
func (client *KClient) GetProducts(filterTerm string, locationId string, filterOffset int, filterLimit int) ([]Product, error) {
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
	res, err := http.DefaultClient.Do(req)
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

	var prodResp ProductsResponse
	if err := json.Unmarshal(body, &prodResp); err != nil {
		return nil, fmt.Errorf("failed to deserialize JSON response body: %v", err)
	}

	return prodResp.Data, nil
}
