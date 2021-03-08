package pkg

type Response struct {
	Data struct {
		Slots Slots `json:"slots"`
	} `json:"Data"`
	Status    string      `json:"Status"`
	ErrCde    interface{} `json:"ErrCde"`
	ErrMsg    interface{} `json:"ErrMsg"`
	ErrMsgDtl interface{} `json:"ErrMsgDtl"`
}
type Slots struct {
	Num1 bool `json:"1"`
	Num2 bool `json:"2"`
}

type CheckResult struct {
	Store int
	Slots Slots
	Err   string
}

type Store struct {
	StoreNumber         int           `json:"storeNumber"`
	Address             string        `json:"address"`
	City                string        `json:"city"`
	State               string        `json:"state"`
	Zipcode             string        `json:"zipcode"`
	TimeZone            string        `json:"timeZone"`
	FullZipCode         string        `json:"fullZipCode"`
	FullPhone           string        `json:"fullPhone"`
	LocationDescription string        `json:"locationDescription"`
	StoreHoursMonday    string        `json:"storeHoursMonday"`
	StoreHoursTuesday   string        `json:"storeHoursTuesday"`
	StoreHoursWednesday string        `json:"storeHoursWednesday"`
	StoreHoursThursday  string        `json:"storeHoursThursday"`
	StoreHoursFriday    string        `json:"storeHoursFriday"`
	StoreHoursSaturday  string        `json:"storeHoursSaturday"`
	StoreHoursSunday    string        `json:"storeHoursSunday"`
	RxHrsMon            string        `json:"rxHrsMon"`
	RxHrsTue            string        `json:"rxHrsTue"`
	RxHrsWed            string        `json:"rxHrsWed"`
	RxHrsThu            string        `json:"rxHrsThu"`
	RxHrsFri            string        `json:"rxHrsFri"`
	RxHrsSat            string        `json:"rxHrsSat"`
	RxHrsSun            string        `json:"rxHrsSun"`
	StoreType           string        `json:"storeType"`
	Latitude            float64       `json:"latitude"`
	Longitude           float64       `json:"longitude"`
	Name                string        `json:"name"`
	MilesFromCenter     float64       `json:"milesFromCenter"`
	SpecialServiceKeys  []string      `json:"specialServiceKeys"`
	Event               interface{}   `json:"event"`
	HolidayHours        []interface{} `json:"holidayHours"`
	PickupDateAndTimes  struct {
		RegularHours []string `json:"regularHours"`
		SpecialHours struct {
			Two0210306 string `json:"2021-03-06"`
			Two0210307 string `json:"2021-03-07"`
		} `json:"specialHours"`
		DefaultTime string `json:"defaultTime"`
		Earliest    string `json:"earliest"`
	} `json:"pickupDateAndTimes"`
}

type FindStoreResponse struct {
	Data struct {
		Stores          []Store     `json:"stores"`
		GlobalZipCode   interface{} `json:"globalZipCode"`
		ResolvedAddress struct {
			AddressLine       interface{} `json:"addressLine"`
			AdminDistrict     string      `json:"adminDistrict"`
			Altitude          int         `json:"altitude"`
			Confidence        string      `json:"confidence"`
			CalculationMethod string      `json:"calculationMethod"`
			CountryRegion     string      `json:"countryRegion"`
			DisplayName       string      `json:"displayName"`
			District          string      `json:"district"`
			FormattedAddress  string      `json:"formattedAddress"`
			GeocodeBestView   struct {
				NorthEastElements struct {
					Altitude  int     `json:"altitude"`
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"northEastElements"`
				SouthWestElements struct {
					Altitude  int     `json:"altitude"`
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"southWestElements"`
			} `json:"geocodeBestView"`
			Latitude   float64 `json:"latitude"`
			Locality   string  `json:"locality"`
			Longitude  float64 `json:"longitude"`
			PostalCode string  `json:"postalCode"`
			PostalTown string  `json:"postalTown"`
		} `json:"resolvedAddress"`
		AmbiguousAddresses interface{}   `json:"ambiguousAddresses"`
		Warnings           []interface{} `json:"warnings"`
	} `json:"Data"`
	Status    string      `json:"Status"`
	ErrCde    interface{} `json:"ErrCde"`
	ErrMsg    interface{} `json:"ErrMsg"`
	ErrMsgDtl interface{} `json:"ErrMsgDtl"`
}
