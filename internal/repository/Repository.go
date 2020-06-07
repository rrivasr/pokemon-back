package repository

import (
	"pokemon-back/internal/core/domain"
	"pokemon-back/internal/dataprovider"
)

type TrainerDAO interface {
	Create(nuevoTrainer *domain.Trainer) (error)
	ObtenerByNombre(trainer *domain.Trainer) (error)
}

type Repository struct {
	TrainerDao TrainerDAO
}

func (r *Repository) CrearTrainer(nuevoTrainer *domain.Trainer) (error) {
	return r.TrainerDao.Create(nuevoTrainer)
}

func (r *Repository) ObtenerTrainerByNombreAndPassword(trainer *domain.Trainer) (error) {
	return r.TrainerDao.ObtenerByNombre(trainer)
}

func NewRepository(trainerDAO *dataprovider.TrainerDataPostgreSQL) *Repository {
	return &Repository{
		TrainerDao: trainerDAO,
	}
}
