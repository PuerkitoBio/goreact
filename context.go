package goreact

type Context struct {
	*Value
	f    func(interface{}) *Value
	vals []*Value
}

func NewContext(f func(interface{}) *Value, v *Value) *Context {
	c := &Context{nil, f, nil}
	c.Bind(v)
	// TODO: Launch goroutine to evaluate function?
	return c
}

func (this *Context) Bind(v *Value) {
	// TODO: Would need to lock when setting the bound values, because goroutine
	// may be running
	this.vals = append(this.vals, v)
}
