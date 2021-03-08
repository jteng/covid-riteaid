package internal

import (
	"covid-riteaid/pkg"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func SendAvailability(phone string, stores []pkg.Store) (bool, error) {
	if len(stores) == 0 {
		return false, nil
	}

	fmt.Println("creating session")
	sess := session.Must(session.NewSession())
	fmt.Println("session created")

	svc := sns.New(sess)
	fmt.Println("service created")
	msg := ""
	for _, store := range stores {
		if msg == "" {
			msg = fmt.Sprintf("Following stores have availability:\nStore:%d\rPhone:%s\rAddr:%s\n", store.StoreNumber, store.FullPhone, store.Address)
			continue
		}
		msg = fmt.Sprintf("%s\nStore:%d\rPhone:%s\rAddr:%s\n", msg, store.StoreNumber, store.FullPhone, store.Address)
	}

	params := &sns.PublishInput{
		Message:     aws.String(msg),
		PhoneNumber: aws.String(fmt.Sprintf("+1%s", phone)),
	}
	resp, err := svc.Publish(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return false, err
	}

	// Pretty-print the response data.
	fmt.Println(resp)
	return true, nil
}

func NotifyHealthStatus(phone string) (bool, error) {

	fmt.Println("creating session")
	sess := session.Must(session.NewSession())
	fmt.Println("session created")

	svc := sns.New(sess)
	fmt.Println("service created")
	msg := "running"

	params := &sns.PublishInput{
		Message:     aws.String(msg),
		PhoneNumber: aws.String(fmt.Sprintf("+1%s", phone)),
	}
	resp, err := svc.Publish(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return false, err
	}

	// Pretty-print the response data.
	fmt.Println(resp)
	return true, nil
}
