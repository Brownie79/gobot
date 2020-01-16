// +build example
//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/Brownie79/gobot"
	"github.com/Brownie79/gobot/drivers/gpio"
	"github.com/Brownie79/gobot/platforms/intel-iot/edison"
)

func main() {
	e := edison.NewAdaptor()
	button := gpio.NewGroveButtonDriver(e, "2")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("On!")
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("Off!")
		})

	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{e},
		[]gobot.Device{button},
		work,
	)

	robot.Start()
}
