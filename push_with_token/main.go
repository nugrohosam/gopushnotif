package main

import (
	"fmt"

	"github.com/NaySoftware/go-fcm"
)

const (
	serverKey = "AAAAglKZV0Q:APA91bEBP6_9l038hjti70Pv8PIu9LMawrIUEv_SF8hgNCV28RAKenNrBwC39QAS8W6NuLxsbrSUUrR19aBoHt495JJpO3ice5-A8BQmej3KjdoEGbVrxfRXinp-SSenqHihMTpKD2Ke"
)

func main() {

	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}

	ids := []string{
		"token1",
	}

	xds := []string{
		"token5",
		"token6",
		"token7",
	}

	c := fcm.NewFcmClient(serverKey)
	c.NewFcmRegIdsMsg(ids, data)
	c.AppendDevices(xds)

	status, err := c.Send()

	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}

}
