package main

import (
	"github.com/distributed/bp/bputil"
	"github.com/distributed/i2cm"
	"io/ioutil"
	"log"
	"time"
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

	err = LoopAD0(m)
	if err != nil {
		log.Fatal(err)
	}
}

func LoopAD0(t i2cm.Transactor8x8) error {
	tick := time.Tick(100 * time.Millisecond)
	devaddr := i2cm.Addr7(0x48)
	var r [2]byte

	for {
		// "regaddr" 0x40 selects A/D channel 0
		_, _, err := t.Transact8x8(devaddr, 0x40, nil, r[:])
		if err != nil {
			return err
		}

		// r[0] is sample from last iteration
		// r[1] is sample from this iteration
		log.Printf("ad0 %#02x\n", r[1])
		<-tick
	}
}
