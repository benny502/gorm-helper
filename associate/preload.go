package associate

type Preload interface {
	GetPreload() string
	GetArgs() []interface{}
}

type preload struct {
	preload string
	args    []interface{}
}

func (p *preload) GetPreload() string {
	return p.preload
}

func (p *preload) GetArgs() []interface{} {
	return p.args
}

func NewPreload(joins string, args ...interface{}) Preload {
	return &preload{preload: joins, args: args}
}
