package associate

type Preload interface {
	GetPreload() string
}

type preload struct {
	preload string
}

func (p *preload) GetPreload() string {
	return p.preload
}

func NewPreload(joins string) Preload {
	return &preload{preload: joins}
}
