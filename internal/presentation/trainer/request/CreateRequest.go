package request

import "pokemon-back/internal/core/usecase/trainer"

type CreateRequest struct {
	Nombre     *string `json:"nombre" validate:"required,min=3,max=16"`
	Contrasena *string `json:"contrasena" validate:"required,min=5,max=30"`
}

func (cr *CreateRequest) FromThis() *trainer.TrainerDTO {
	return &trainer.TrainerDTO{
		Nombre:      cr.Nombre,
		Contrasena:  cr.Contrasena,
	}
}
