package goreact

import (
	"fmt"
	"testing"
	"time"
)

func TestBasicBehavior(t *testing.T) {
	var v Value

	f := func(val interface{}) {
		fmt.Printf("Eval=%v...\n", val)
	}
	NewContext(f, &v)

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		v.Set(i)
	}
	time.Sleep(time.Second)
}
