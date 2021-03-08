package main

import (
	"covid-riteaid/internal"
	"covid-riteaid/pkg"
	"log"
	"os"
	"time"
)

var zipCodes []string = []string{"08003", "08057", "08009"}

func main() {
	phone := os.Getenv("COVID_PHONE")
	testSms := os.Getenv("COVID_TEST_SMS")

	times := 0
	for {
		stores := make(map[int]pkg.Store)
		for _, zip := range zipCodes {
			log.Printf("find store in %s\n", zip)
			sst, err := internal.FindStore(zip)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("found %d stores\n", len(sst))
			for _, st := range sst {
				stores[st.StoreNumber] = st
				log.Printf("add store %d\n", st.StoreNumber)
			}
		}

		var goodStores []pkg.Store
		for _, store := range stores {
			sts := internal.Check(store.StoreNumber)
			log.Printf("%v\n", sts)
			if sts.Slots.Num1 || sts.Slots.Num2 || testSms != "" {
				goodStores = append(goodStores, stores[sts.Store])
			}
		}
		var sent bool
		var err error
		if sent, err = internal.SendAvailability(phone, goodStores); err != nil {
			log.Printf("fail to sent sms %s", err.Error())
		}
		if sent {
			log.Printf("sent SMS")
			return
		}
		time.Sleep(5 * time.Minute)
		times = times + 1
		if times%7 == 0 {
			internal.NotifyHealthStatus(phone)
		}
	}
}
