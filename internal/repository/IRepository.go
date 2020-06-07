package repository

import "pokemon-back/internal/core/domain"

type IRepository interface {
	CrearTrainer(nuevoTrainer *domain.Trainer) (error)
	ObtenerTrainerByNombreAndPassword(trainer *domain.Trainer) (error)
}
