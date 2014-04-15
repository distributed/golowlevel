This is a repository for a talk I gave at the Go Meetup on 
April 14, 2014 at Google Zürich.

## Slides

The slides were presented with the [Go present tool](https://godoc.org/code.google.com/p/go.talks).
You can find all the necessary files in the `pres` subfolder. The
code examples in the slides all use the content of the file
pres/serialport to decide on which serial port to talk to
the I2C bus master. Change it to match your system, do not
put a newline in the file.

The talk concluded with a little gag, the Googly Eyed Gopher,
which is described towards the end of this file.

## Hardware

The talk was supported by a number of code examples which
communicated with hardware on a breadboard. On the breadboard
I wired 3 devices onto an I2C bus. The I2C bus lines were pulled
up to VCC by 5kOhm resistors. The voltage of the power
supply was 3.3V.

The I2C bus was driven by a [Bus Pirate](http://dangerousprototypes.com/bus-pirate-manual/).
The Bus Pirate contains a USB-to-serial converter, thus ultimately
the software interface to the Bus Pirate is a serial port, albeit
a relatively virtual one.

The circuit on the breadboard consisted of the following.
All chips are relatively easy to obtain from electronics
distributors.
Schematics will follow.

### EEPROM

A 24C02 2kbit EEPROM with all address lines (A2:A0) pulled to GND.
Thus, the EEPROM responds to the 7 bit address 0x50 a.k.a. the
8 bit address pair 0xa0/0xa1.

### I/O Expander

An MCP23008 8-bit I/O expander. All address lines (A2:A0) were pulled
to GND. Thus, the I/O expanders responds to the 7 bit address 0x20, 
a.k.a. the 8 bit address pair 0x40/0x41. To Pin 0, a resistor and an
LED were connected in series in positive logic: a high level on Pin 0
turned on the LED.

### A/D Converter

A PCF8591P 4 channel 8 bit A/D and 1 channel 8 bit D/A converter. All
address lines (A2:A0) were pulled to GND, making the chip respond
to the 7 bit address 0x48, a.k.a. the 8 bit address pair 0x90/0x91.
The reference voltage was VCC=3.3V. To A/D channel 0, the middle tap
of a potentiometer was connected. The other taps of the potentiometer
were at VCC and GND, respectively.


## The Googly Eyed Gopher

The subfodler `gg` contains the source code for the Googly Eyed Gopher.
The Go program in `gg.go` continously reads A/D samples from A/D channel
0 of the A/D converter on the breadboard. The Go program also serves
a static index page, `index.html` and the 2 images for the Gopher. The
index page is peppered with some Javascript that opens a websocket to the
Go server. On this connection, the Go server continously sends fresh A/D samples
which the JS side translates into rotations of the Gopher's pupils.
