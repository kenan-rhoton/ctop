package config

import (
	"reflect"
)

// defaults
type Params struct {
	FilterStr string // Container Name or ID Filter
	SortField string // Container Sort Field
}

var params = []*Param{
	&Param{
		Key:   "filterStr",
		Val:   "",
		Label: "Container Name or ID Filter",
	},
	&Param{
		Key:   "sortField",
		Val:   "state",
		Label: "Container Sort Field",
	},
}

type Param struct {
	Key   string
	Val   string
	Label string
}

func initParams() {
	Config.FilterStr = ""
	Config.SortField = "state"
}

// Get Param by key
func Get(k string) reflect.Value {
	return reflect.ValueOf(&Config).Elem().FieldByName(k)
}

// Get Param value by key
func GetVal(k string) string {
	p := Get(k)
	if p.Kind().String() != "string" {
		log.Errorf("Tried to access a " + p.Kind().String() + " named " + k + " as a String Param")
		return ""
	}
	return p.String()
}

// Set param value
func Update(k, v string) {
	p := Get(k)
	if p.CanSet() {
		log.Noticef("config change: %s: %s -> %s", k, quote(p.String()), quote(v))
		p.SetString(v)
	} else {
		log.Errorf("ignoring update for non-existent parameter: %s", k)
	}
}
