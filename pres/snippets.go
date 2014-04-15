package main

// sersifacestart OMIT
func Open(fn string) (SerialPort, error)

type SerialPort interface {
	io.Reader
	io.Writer
	io.Closer

	// SetMode sets the frame format and handshaking configuration.
	// baudrate may be freely chosen, the driver is allowed to reject
	// unachievable baud rates. databits may be any number of data bits
	// supported by the driver. parity is one of (N|O|E) for none, odd
	// or even parity. handshake is either NO_HANDSHAKE or
	// RTSCTS_HANDSHAKE.
	SetMode(baudrate, databits, parity, stopbits, handshake int) error

	// SetReadParams sets the minimum number of bits to read and a read
	// timeout in seconds. These parameters roughly correspond to the
	// UNIX termios concepts of VMIN and VTIME.
	SetReadParams(minread int, timeout float64) error
}

// sersifacestop OMIT

// eestart OMIT
// EEPROM24 represents an I2C EEPROM device. The memory array is made
// available via a file-like interface. The file's size is fixed to
// the memory array size and writes past the end of the array result
// in an error.
type EEPROM24 interface {
	io.Reader
	io.Seeker
	io.Writer
}

// NewEEPROM24 constructs an I2C EEPROM driver for a device with base
// address devaddr residing on m's bus. The EEPROM driver parameters
// are passed in conf. Invalid configurations are rejected.
func NewEEPROM24(m I2CMaster, devaddr Addr, conf EEPROM24Config) (EEPROM24, error)

// eestop OMIT

// transactstart OMIT
// Implements a write-then-read transaction with 8 bit register
// addresses and 8 bit data. The transaction always writes data
// to the device, as the register address is always written.
// The read part of the transaction is not executed if len(r) == 0.
// ...
type Transactor8x8 interface {
	Transact8x8(addr Addr, regaddr uint8, w []byte, r []byte) (nw, nr int, err error)
}

// transactstop OMIT

// transactexamplestart OMIT
_, _, err = t.Transact8x8(i2cm.Addr7(0x40>>1), AddrGPIO, []byte{1<<0}, nil)
// transactexamplestop OMIT


// i2cmasterstart OMIT
type I2CMaster interface {
    // Start sends a start or repeated start condition,
    // depending on bus state.
    Start() error

    // Sends a stop condition.
    Stop() error

    // ReadByte reads one byte and sends an ACK
    // if ack is true. Application code is responsible
    // to ensure that ack is false for the last read
    // before a stop bit is sent.
    ReadByte(ack bool) (recvb byte, err error)

    // WriteByte writes one byte to the device. If
    // device does not ACK, it returns NACKReceived.
    WriteByte(b byte) error
}
// i2cmasterstop OMIT