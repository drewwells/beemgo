package beemgo

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
)

type (
	Controller struct {
		beego.Controller
		Session *mgo.Session
	}
)

var (
	singleton Service
)

func (c *Controller) Prepare() {
	if singleton.spark == nil {
		singleton.Url = beego.AppConfig.String("mgourl")
		singleton.New()
	}
	c.Session = singleton.Session()
}

func (c *Controller) Finish() {
	defer func() {
		if c.Session != nil {
			singleton.Close(c.Session)
		}
	}()
}
