package main

import (
	"fmt"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody("Hello there!")
	params.SetTo("whatsapp:+27649015322")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
}
