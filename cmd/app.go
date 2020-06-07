package cmd

import (
	"errors"
	container2 "pokemon-back/internal/config/container"
	"pokemon-back/internal/presentation"
)

func StartApp() error {
	container := &container2.ContainerDig{}
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
