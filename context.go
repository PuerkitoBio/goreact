package goreact

import (
	"runtime"
)

type Context struct {
	f func(interface{})
	c chan (<-chan interface{})
}

func NewContext(f func(interface{}), v *Value) *Context {
	ctx := &Context{f, make(chan (<-chan interface{}), 1)}
	ctx.Bind(v)
	// Launch goroutine to evaluate function
	go evaluate(ctx.c, f)
	return ctx
}

func (this *Context) Bind(v *Value) {
	if v != nil {
		this.c <- v.Get()
	}
}

func evaluate(recv chan (<-chan interface{}), f func(interface{})) {
	var chans = make([]<-chan interface{}, 1)
	for {
		select {
		// Try to receive new values
		case vc := <-recv:
			chans = append(chans, vc)
		default:
			// Try reading a new value from one of the bound values
			for _, c := range chans {
				// Do not block on receive
				select {
				case v, ok := <-c:
					if ok {
						// Channel not closed, run function
						f(v)
					}
				default:
					// Continue with next channel
				}
			}
			runtime.Gosched()
		}
	}
}
