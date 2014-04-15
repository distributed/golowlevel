package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/distributed/bp/bputil"
	"github.com/distributed/i2cm"
	"log"
	"math"
	"net/http"
	"time"
)

func main() {
	var wss websocket.Server

	var ps positStreamer
	ps.positreq = make(chan chan float64)

	m, _, err := bputil.DialNonStrictI2c("/dev/cu.usbserial-A90156F4")
	if err != nil {
		log.Fatal(err)
	}

	ps.positioner = &ADPositioner{m, i2cm.Addr7(0x48), 0}

	go ps.pump()

	wss.Handler = ps.wsHandler

	fmt.Println("navigate to http://localhost:7500/")
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/posit", wss.Handler)
	err = http.ListenAndServe(":7500", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type positStreamer struct {
	positreq   chan chan float64
	positioner Positioner // made pluggable in case you don't have hardware but still want to benefit from the terrible utility that is the googly eyed gopher
}

func (ps *positStreamer) pump() {
	tick := time.Tick(100 * time.Millisecond)
	var reqs []chan float64

	for {
		select {
		case <-tick:
			v, err := ps.positioner.Position()
			if err != nil {
				log.Fatal(err)
			}

			for _, req := range reqs {
				req <- v
			}
			reqs = reqs[:0]
		case req := <-ps.positreq:
			reqs = append(reqs, req)
		}
	}
}

type Update struct {
	Position float64
}

func (ps *positStreamer) wsHandler(conn *websocket.Conn) {
	rc := make(chan float64)

	for {
		ps.positreq <- rc
		pos := <-rc
		update := Update{pos}
		err := websocket.JSON.Send(conn, update)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

type Positioner interface {
	Position() (float64, error)
}

type SinAbsPositioner struct {
	x float64
}

func (s *SinAbsPositioner) Position() (float64, error) {
	v := math.Abs(math.Sin(s.x))
	s.x += 0.1
	return v, nil
}

type ADPositioner struct {
	T          i2cm.Transactor8x8
	DeviceAddr i2cm.Addr
	Channel    uint8
}

func (p *ADPositioner) Position() (float64, error) {
	var r [2]uint8
	_, _, err := p.T.Transact8x8(p.DeviceAddr, 0x40|(p.Channel&0x0f), nil, r[:])
	if err != nil {
		return 0, err
	}

	return float64(r[1]) / 256.0, nil
}
