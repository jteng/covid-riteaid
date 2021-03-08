package internal

import (
	"covid-riteaid/pkg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{Timeout: 10 * time.Second}
}
func Check(store int) pkg.CheckResult {
	result := pkg.CheckResult{
		Store: store,
	}
	res, err := client.Get(fmt.Sprint("https://www.riteaid.com/services/ext/v2/vaccine/checkSlots?storeNumber=%d", store))
	if err != nil {
		result.Err = err.Error()
		return result
	}
	defer res.Body.Close()

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	riteaidRes := &pkg.Response{}
	err = json.Unmarshal(responseData, riteaidRes)
	if err != nil {
		result.Err = err.Error()
		return result
	}
	result.Slots = riteaidRes.Data.Slots
	return result
}

func FindStore(zip string) ([]pkg.Store, error) {
	var result []pkg.Store
	qry := fmt.Sprintf("https://www.riteaid.com/services/ext/v2/stores/getStores?address=%s&attrFilter=PREF-112&fetchMechanismVersion=2&radius=50", zip)

	res, err := client.Get(qry)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	riteaidRes := &pkg.FindStoreResponse{}
	err = json.Unmarshal(responseData, riteaidRes)
	if err != nil {
		return nil, err
	}

	result = riteaidRes.Data.Stores
	return result, nil
}
