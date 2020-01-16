package ble

import (
	"testing"

	"github.com/Brownie79/gobot/gobottest"
)

func initTestBLESerialPort() *SerialPort {
	return NewSerialPort("TEST123", "123", "456")
}

func TestBLESerialPort(t *testing.T) {
	d := initTestBLESerialPort()
	gobottest.Assert(t, d.Address(), "TEST123")
}
