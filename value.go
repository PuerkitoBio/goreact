package goreact

type Value struct {
	v     interface{}
	chans []chan interface{}
}

func (this *Value) Get() <-chan interface{} {
	c := make(chan interface{})
	this.chans = append(this.chans, c)
	return c
}

func (this *Value) Set(val interface{}) {
	this.v = val
	broadcastValue(this.chans, this.v)
}

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

func broadcastValue(chans []chan interface{}, val interface{}) {
	// TODO:  Naive implementation, what if it blocks?  
	for _, c := range chans {
		c <- val
	}
}
