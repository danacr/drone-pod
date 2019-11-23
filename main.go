package main

import (
	"fmt"
	"os"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {

	droneip := getDroneIP()

	drone := tello.NewDriver(droneip, "8888")

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
