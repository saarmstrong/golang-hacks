package lib

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

type LilRobots struct {
	gbot *gobot.Gobot
	adaptor *sphero.SpheroAdaptor
	driver *sphero.SpheroDriver
}

func NewGRobots() *LilRobots {
    r := new(LilRobots)
	r.gbot = gobot.NewGobot()
	r.adaptor = sphero.NewSpheroAdaptor("sphero", "/dev/tty.Sphero-YBG-AMP-SPP")
	r.driver = sphero.NewSpheroDriver(r.adaptor, "sphero")
	return r
}

func (r *LilRobots)Append(work interface{})  {
	robot := gobot.NewRobot("sphero",
		[]gobot.Connection{r.adaptor},
		[]gobot.Device{r.driver},
		work,
	)
	r.gbot.AddRobot(robot)
}

func (r *LilRobots)Start()  {
	r.gbot.Start()
}

func (r *LilRobots)Stop()  {
	r.gbot.Stop()
}
