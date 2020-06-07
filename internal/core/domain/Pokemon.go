package domain

type Pokemon struct {
	Id *int64
	Nombre *string
	Tipo *string
	Peso *float32
	Altura *float32
	EvolucionID *int64
	Habiltado *bool
	Ataque []*Ataque
	ImagenFrente *string
	ImagenAtras *string
	Shiny *bool
	Mote *string
}

func (p *Pokemon) agregarMote(nuevoMote *string) {
	p.Mote = nuevoMote
}
