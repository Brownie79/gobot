// +build example
//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/Brownie79/gobot"
	"github.com/Brownie79/gobot/drivers/aio"
	"github.com/Brownie79/gobot/drivers/gpio"
	"github.com/Brownie79/gobot/platforms/intel-iot/edison"
)

func main() {
	e := edison.NewAdaptor()
	sensor := aio.NewAnalogSensorDriver(e, "0")
	led := gpio.NewLedDriver(e, "3")

	work := func() {
		sensor.On(aio.Data, func(data interface{}) {
			brightness := uint8(
				gobot.ToScale(gobot.FromScale(float64(data.(int)), 0, 4096), 0, 255),
			)
			fmt.Println("sensor", data)
			fmt.Println("brightness", brightness)
			led.Brightness(brightness)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{e},
		[]gobot.Device{sensor, led},
		work,
	)

	robot.Start()
}
