package domain

import "errors"

type Trainer struct {
	Id *int64
	Nombre *string
	Contrasena *string
	Experiencia *int
	Atrapados []*Pokemon
	Equipo []*Pokemon
	Habilitado *bool
}

func (t *Trainer) getCAntidadAtrapados() int {
	return len(t.Atrapados)
}

func (t *Trainer) agregarExperiencia(nuevaExperiencia *int) {
	sum := *t.Experiencia + *nuevaExperiencia
	t.Experiencia = &sum
}

func (t *Trainer) usarExperiencia(experiencia *int) error {
	if *t.Experiencia >= *experiencia {
		res := *t.Experiencia - *experiencia
		t.Experiencia = &res
		return nil
	}
	return errors.New("No tienes suficiente experiencia")
}

func (t *Trainer) agregarPokemon(pokemon *Pokemon) {
	t.Atrapados = append(t.Atrapados, pokemon)
}


