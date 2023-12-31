package uility

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
)

var E *casbin.Enforcer

type Casbin struct {
	Xorm     *xormadapter.Adapter
	Enforcer *casbin.Enforcer
}

func NewCasbin() *Casbin {
	xorm, err := xormadapter.NewAdapter("mysql", "root:admin123@tcp(127.0.0.1:3306)/cloud-k", true) // Your driver and data source.
	if err != nil {
		fmt.Println(err)
	}
	enforcer, err := casbin.NewEnforcer("./uility/model.conf", xorm)
	if err != nil {
		fmt.Println(err)
	}
	return &Casbin{
		Xorm:     xorm,
		Enforcer: enforcer,
	}
}

func init() {
	a, err := xormadapter.NewAdapter("mysql", "root:admin123@tcp(127.0.0.1:3306)/cloud-k", true) // Your driver and data source.
	if err != nil {
		fmt.Println(err)
		return
	}
	e, err := casbin.NewEnforcer("./uility/model.conf", a)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Load the policy from DB.
	err = e.LoadPolicy()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = e.SavePolicy()
	if err != nil {
		fmt.Println(err)
		return
	}
	E = e
}
