package models

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

type JsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
