package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatCapacity = 10
var arrivalRate = 100 // milliseconds
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed random number generator
	rand.Seed(time.Now().UnixNano())

	// print welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	// create channels
	clientChan := make(chan string, seatCapacity)
	doneChan := make(chan bool)

	// create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		BarbersDoneChan: doneChan,
		ClientsChan:     clientChan,
		Open:            true, // the barbershop is open when this programer starts
	}

	// add barbers
	color.Green("The shop is open for the day!")

	// start the barbershop as a goroutine

	// add clients

	// block until the barbershop is closed
}
