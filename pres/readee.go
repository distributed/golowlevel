package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/distributed/bp/bputil"
	"github.com/distributed/i2cm"
	"io"
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

	err = ReadEEPROM(m)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadEEPROM(m i2cm.I2CMaster) error {
	ee, err := i2cm.NewEEPROM24(m, i2cm.Addr7(0xa0>>1), i2cm.Conf_24C02)
	if err != nil {
		return err
	}

	var r [48]byte
	_, err = io.ReadFull(ee, r[:])
	if err != nil {
		return err
	}

	spew.Dump(r)

	return nil
}
