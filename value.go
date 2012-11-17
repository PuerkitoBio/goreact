package goreact

import ()

// Make it so a zero-Value is a valid, usable Value that has never been set (
// no initial value)
type Value struct {
	v       interface{}
	chans   []chan interface{}
	setOnce bool
}

func NewValue(val interface{}) *Value {
	return &Value{val, make([]chan interface{}, 0, 1), true}
}

func (this *Value) Get() <-chan interface{} {
	// Create a unique channel for this communication
	c := make(chan interface{}, 1)
	this.chans = append(this.chans, c)
	// Send the current value, since Set() will not be called for this value-channel
	// combination.
	if this.setOnce {
		c <- this.v
	}
	return c
}

func (this *Value) Set(val interface{}) {
	this.v = val
	broadcastValue(this.chans, this.v)
}

/*
// Implement the Closer interface
func (this *Value) Close() error {
	for _, c := range this.chans {
		close(c)
	}
	this.chans = nil
	return nil // TODO: Cannot fail, but still returns an error so that it is a Closer. Good idea?
}

func (this *Value) Inject(ctx *Context) {
	if ctx != nil {
		ctx.Bind(this)
	}
}
*/

func broadcastValue(chans []chan interface{}, val interface{}) {
	// TODO:  Naive implementation, what if it blocks? 
	for _, c := range chans {
		c <- val
	}
}
