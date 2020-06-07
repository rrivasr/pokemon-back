package trainer

import (
	"golang.org/x/crypto/bcrypt"
	"pokemon-back/internal/repository"
)

type ICrear interface {
	CrearTrainer(trainerDTO *TrainerDTO) (*TrainerDTO, error)
}

type Crear struct {
	Repository repository.IRepository
}

func (ct *Crear) CrearTrainer(trainerDTO *TrainerDTO) (*TrainerDTO, error) {
	trainer, _ := trainerDTO.fromThis()
	c, err := bcrypt.GenerateFromPassword([]byte(*trainer.Contrasena), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	contrasena := string(c)
	trainer.Contrasena = &contrasena
	if err := ct.Repository.CrearTrainer(trainer); err != nil {
		return nil, err
	}
	trainer.Contrasena = nil
	result, _ := ToThis(trainer)
	return result, nil
}

func NewCrearUseCase(repository *repository.Repository) *Crear {
	return &Crear{
		Repository: repository,
	}
}
