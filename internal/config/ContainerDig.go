package config

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/dig"
	"pokemon-back/internal/core/usecase/trainer"
	"pokemon-back/internal/dataprovider"
	"pokemon-back/internal/presentation"
	trainercontroller "pokemon-back/internal/presentation/trainer"
	"pokemon-back/internal/repository"
)

type ContainerDig struct {}


func (cd *ContainerDig) BuildContainer() (*dig.Container, error) {
	c := dig.New()
	if err := c.Provide(trainer.NewCrearUseCase); err != nil {
		return nil, err
	}
	if err := c.Provide(trainer.NewLogin); err != nil {
		return nil, err
	}

	if err := c.Provide(repository.NewRepository); err != nil {
		return nil, err
	}
	if err := c.Provide(dataprovider.NewTrainerDataPostgreSQL); err != nil {
		return nil, err
	}
	if err := c.Provide(NewDatabase); err != nil {
		return nil, err
	}
	if err := c.Provide(trainercontroller.NewController); err != nil {
		return nil, err
	}
	if err := c.Provide(presentation.NewServer); err != nil {
		return nil, err
	}
	if err := c.Provide(gin.New); err != nil {
		return nil, err
	}
	if err := c.Provide(validator.New); err != nil {
		return nil, err
	}
	return c, nil
}
