package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {

	droneip := getDroneIP()

	drone := tello.NewDriver(droneip, "8888")

	// capture sigterm and land drone
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for sig := range c {
			fmt.Println("Need to die, but I must land the drone first", sig)
			drone.Land()
			os.Exit(0)
		}
	}()

	work := func() {

		drone.On(tello.ConnectedEvent, func(data interface{}) {
			fmt.Println("Connected")
		})

		drone.TakeOff()

	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()

}

func getDroneIP() string {
	var droneip string
	switch node := os.Getenv("NODE"); node {
	case "rockpi0":
		droneip = "192.168.86.250"
	case "rockpi1":
		droneip = "192.168.86.251"
	case "rockpi2":
		droneip = "192.168.86.252"
	}
	return droneip
}
