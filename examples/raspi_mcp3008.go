// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"time"

	"github.com/Brownie79/gobot"
	"github.com/Brownie79/gobot/drivers/spi"
	"github.com/Brownie79/gobot/platforms/raspi"
)

func main() {
	a := raspi.NewAdaptor()
	adc := spi.NewMCP3008Driver(a)

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			result, err := adc.Read(0)
			fmt.Println("A0", result, err)
		})
	}

	robot := gobot.NewRobot("mcp3008bot",
		[]gobot.Connection{a},
		[]gobot.Device{adc},
		work,
	)

	robot.Start()
}
