// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"time"

	"github.com/Brownie79/gobot"
	"github.com/Brownie79/gobot/drivers/i2c"
	"github.com/Brownie79/gobot/platforms/intel-iot/edison"
)

func main() {
	board := edison.NewAdaptor()
	accel := i2c.NewGroveAccelerometerDriver(board)

	work := func() {
		gobot.Every(500*time.Millisecond, func() {
			if x, y, z, err := accel.XYZ(); err == nil {
				fmt.Println(x, y, z)
				fmt.Println(accel.Acceleration(x, y, z))
			} else {
				fmt.Println(err)
			}
		})
	}

	robot := gobot.NewRobot("accelBot",
		[]gobot.Connection{board},
		[]gobot.Device{accel},
		work,
	)

	robot.Start()
}
