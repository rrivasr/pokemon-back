package pokemon

import "pokemon-back/internal/core/usecase/ataque"

type PokemonDTO struct {
	Id *int64
	Nombre *string
	Tipo *string
	Peso *float32
	Altura *float32
	EvolucionID *int64
	Habiltado *bool
	Ataque []*ataque.AtaqueDTO
	ImagenFrente *string
	ImagenAtras *string
	Shiny *bool
	Mote *string
}
