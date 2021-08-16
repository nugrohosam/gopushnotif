package main

import (
	"fmt"

	"github.com/NaySoftware/go-fcm"
)

const (
	serverKey = "YOUR-KEY"
	topic     = "/topics/someTopic"
)

func main() {

	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}

	c := fcm.NewFcmClient(serverKey)
	c.NewFcmMsgTo(topic, data)

	status, err := c.Send()

	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}

}
