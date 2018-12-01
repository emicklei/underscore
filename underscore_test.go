package underscore

import (
	"testing"

	"github.com/emicklei/otto"
)

func TestMap(t *testing.T) {
	o := otto.New()
	src, err := Asset("underscore.js")
	check(t, err)
	_, err = o.Run(string(src))
	check(t, err)
	v, err := o.Run("_.map([1, 2, 3], function(num){ return num * 3; });")
	check(t, err)
	t.Log(v.IsObject())
	for _, k := range v.Object().Keys() {
		e, _ := v.Object().Get(k)
		t.Logf("%v=%v", k, e)
	}
}

func check(t *testing.T, e error) {
	if e != nil {
		t.Error(e.Error())
	}
}
