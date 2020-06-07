package trainer

import (
	"pokemon-back/internal/core/domain"
	"pokemon-back/internal/core/usecase/pokemon"
)

type TrainerDTO struct {
	Id          *int64
	Nombre      *string
	Contrasena  *string
	Experiencia *int
	Atrapados   []*pokemon.PokemonDTO
	Equipo      []*pokemon.PokemonDTO
	Habilitado  *bool
	Token       *string
}

func (t *TrainerDTO) fromThis() (*domain.Trainer, error) {
	exp := 0
	hab := true
	id := (int64(0))
	return &domain.Trainer{
		Id:          &id,
		Nombre:      t.Nombre,
		Contrasena:  t.Contrasena,
		Experiencia: &exp,
		Habilitado:  &hab,
	}, nil
}

func ToThis(trainer *domain.Trainer) (*TrainerDTO, error) {
	return &TrainerDTO{
		Id:          trainer.Id,
		Nombre:      trainer.Nombre,
		Contrasena:  trainer.Contrasena,
		Experiencia: trainer.Experiencia,
		Atrapados:   nil,
		Equipo:      nil,
		Habilitado:  trainer.Habilitado,
	}, nil
}
