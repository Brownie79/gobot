// +build example
//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/Brownie79/gobot"
	"github.com/Brownie79/gobot/drivers/aio"
	"github.com/Brownie79/gobot/platforms/dexter/gopigo3"
	"github.com/Brownie79/gobot/platforms/raspi"
)

func main() {
	raspiAdaptor := raspi.NewAdaptor()
	gpg3 := gopigo3.NewDriver(raspiAdaptor)
	sensor := aio.NewGroveLightSensorDriver(gpg3, "AD_1_1")

	work := func() {
		sensor.On(sensor.Event("data"), func(data interface{}) {
			fmt.Println("sensor", data)
		})
	}

	robot := gobot.NewRobot("gopigo3sensor",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{gpg3, sensor},
		work,
	)

	robot.Start()
}
