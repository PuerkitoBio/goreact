package goreact

import (
	"runtime"
)

type Context struct {
	Value
	f func(interface{}) interface{}
	c chan (<-chan interface{})
}

func NewContext(f func(interface{}) interface{}, v ValueProvider) *Context {
	ctx := new(Context)

	// Initialize
	ctx.f = f
	ctx.c = make(chan (<-chan interface{}), 1)
	ctx.Bind(v)

	// Launch goroutine to evaluate function
	go evaluate(ctx.c, f, ctx)

	return ctx
}

func (this *Context) Bind(v ValueProvider) {
	if v != nil {
		this.c <- v.Get()
	}
}

func evaluate(recv chan (<-chan interface{}), f func(interface{}) interface{}, retVal ValueProvider) {
	var chans = make([]<-chan interface{}, 1)
	for {
		select {
		// See if new values got bound
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
						retVal.Set(f(v))
					}
				default:
					// Continue with next channel
				}
			}
			runtime.Gosched()
		}
	}
}
