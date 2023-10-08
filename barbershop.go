package main

import (
	"time"

	"github.com/fatih/color"
)

// BarberShop represents the barbershop in the Sleeping Barber Problem
type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

// addBarber adds a barber to the shop
func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients", barber)

		for {
			// if there are no clients, the barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do, so %s takes a nap.", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan
			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed, so send the barber home and close the channel
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

// cutHair simulates the time it takes to cut a client's hair
func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is done cutting %s's hair.", barber, client)
}

// sendBarberHome sends the barber home for the day
func (shop *BarberShop) sendBarberHome(barber string) {
	color.Blue("%s goes home for the day.", barber)
	shop.BarbersDoneChan <- true
}
