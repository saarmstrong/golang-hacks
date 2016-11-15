package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

type Robots struct {
	gbot Gobot
	adaptor SpheroAdaptor
	driver SpheroDriver
}

func NewRobots()(r *Robots)  {
	r.gbot = gobot.NewGobot()
	r.adaptor = sphero.NewSpheroAdaptor("sphero", "/dev/tty.Sphero-YBG-AMP-SPP")
	r.driver = sphero.NewSpheroDriver(adaptor, "sphero")
	return r
}

func (r *Robots)Append(work interface{})  {
	robot := r.gobot.NewRobot("sphero",
		[]r.gobot.Connection{adaptor},
		[]r.gobot.Device{driver},
		work,
	)
	r.gbot.AddRobot(robot)
}

func (r *Robots)Start()  {
	r.gbot.Start()
}

func (r *Robots)Stop()  {
	r.gbot.Stop()
}
