package request

import "pokemon-back/internal/core/usecase/trainer"

type LoginRequest struct {
	Nombre     *string `json:"nombre" validate:"required,min=3,max=16"`
	Contrasena *string `json:"contrasena" validate:"required,min=5,max=30"`
}

func (lr *LoginRequest) FromThis() *trainer.TrainerDTO {
	return &trainer.TrainerDTO{
		Nombre:      lr.Nombre,
		Contrasena:  lr.Contrasena,
	}
}
