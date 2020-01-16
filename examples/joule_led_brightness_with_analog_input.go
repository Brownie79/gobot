// +build example
//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/Brownie79/gobot"
	"github.com/Brownie79/gobot/drivers/aio"
	"github.com/Brownie79/gobot/drivers/gpio"
	"github.com/Brownie79/gobot/drivers/i2c"
	"github.com/Brownie79/gobot/platforms/intel-iot/joule"
)

func main() {
	e := joule.NewAdaptor()
	ads1015 := i2c.NewADS1015Driver(e)
	sensor := aio.NewAnalogSensorDriver(ads1015, "0")
	led := gpio.NewLedDriver(e, "J12_26")

	work := func() {
		sensor.On(aio.Data, func(data interface{}) {
			brightness := uint8(gobot.ToScale(gobot.FromScale(float64(data.(int)), 0, 1023), 0, 255))
			fmt.Println("sensor", data)
			fmt.Println("brightness", brightness)
			led.Brightness(brightness)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{e},
		[]gobot.Device{ads1015, sensor, led},
		work,
	)

	robot.Start()
}
