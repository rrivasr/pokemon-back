package response

import "pokemon-back/internal/core/usecase/trainer"

type CreateResponse struct {
	Id          *int64  `json:"id,omitempty"`
	Nombre      *string `json:"nombre,omitempty"`
	Habilitado  *bool   `json:"habilitado,omitempty"`
	Experiencia *int    `json:"experiencia,omitempty"`
}

func NewCreateResponse(trainer *trainer.TrainerDTO) *CreateResponse {
	return &CreateResponse{
		Id:          trainer.Id,
		Nombre:      trainer.Nombre,
		Habilitado:  trainer.Habilitado,
		Experiencia: trainer.Experiencia,
	}
}
