package main

import (
	"github.com/distributed/bp/bputil"
	"github.com/distributed/i2cm"
	"io/ioutil"
	"log"
)

func main() {
	fnb, err := ioutil.ReadFile("serialport")
	if err != nil {
		log.Fatal(err)
	}

	m, _, err := bputil.DialNonStrictI2c(string(fnb))
	if err != nil {
		log.Fatal(err)
	}

	_, _, err = m.Transact8x8(i2cm.Addr7(0x40>>1), AddrIODIR, []byte{0xfe}, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = TurnOnLED(m)
	if err != nil {
		log.Fatal(err)
	}
}

const (
	AddrIODIR = 0x00
	AddrGPIO  = 0x09
)

func TurnOnLED(m i2cm.I2CMaster) error {
	err := m.Start()
	if err != nil {
		return err
	}
	err = m.WriteByte(0x40) // address device (MCP23008) for writing
	if err != nil {
		return err
	}
	err = m.WriteByte(AddrGPIO) // select register
	if err != nil {
		return err
	}
	// pin 0 needs to be configured as output already
	err = m.WriteByte(1 << 0) // write register
	if err != nil {
		return err
	}
	err = m.Stop()
	if err != nil {
		return err
	}
	return nil
}
