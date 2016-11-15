package lib

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

type Robots struct {
	gbot *gobot.Gobot
	adaptor *sphero.SpheroAdaptor
	driver *sphero.SpheroDriver
}

func NewRobots()(r *Robots)  {
	r.gbot = gobot.NewGobot()
	r.adaptor = sphero.NewSpheroAdaptor("sphero", "/dev/tty.Sphero-YBG-AMP-SPP")
	r.driver = sphero.NewSpheroDriver(r.adaptor, "sphero")
	return r
}

func (r *Robots)Append(work interface{})  {
	robot := gobot.NewRobot("sphero",
		[]gobot.Connection{r.adaptor},
		[]gobot.Device{r.driver},
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
