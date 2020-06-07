package cmd

import (
	"errors"
	"pokemon-back/internal/config"
	"pokemon-back/internal/presentation"
)

func StartApp() error {
	container := &config.ContainerDig{}
	c, err := container.BuildContainer()
	if err != nil {
		return err
	}
	erroServer := false
	if err := c.Invoke(func(server *presentation.Server) {
		if err := server.StartServer(); err != nil {
			erroServer = true
			return
		}
	}); err != nil {
		return err
	}
	if erroServer {
		return errors.New("Error al iniciar el server")
	}
	return nil
}
