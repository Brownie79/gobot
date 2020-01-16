// +build example
//
// Do not build by default.

package main

import (
	"github.com/Brownie79/gobot"
	"github.com/Brownie79/gobot/api"
	"github.com/Brownie79/gobot/drivers/gpio"
	"github.com/Brownie79/gobot/platforms/digispark"
)

func main() {
	master := gobot.NewMaster()
	api.NewAPI(master).Start()

	digisparkAdaptor := digispark.NewAdaptor()
	led := gpio.NewLedDriver(digisparkAdaptor, "0")

	robot := gobot.NewRobot("digispark",
		[]gobot.Connection{digisparkAdaptor},
		[]gobot.Device{led},
	)

	master.AddRobot(robot)

	master.Start()
}
