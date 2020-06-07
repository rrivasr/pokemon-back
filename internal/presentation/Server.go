package presentation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"pokemon-back/internal/presentation/trainer"
)

type Server struct {
	Engine *gin.Engine
	TrainerController *trainer.Controller
	Port string `json:"-" envconfig:"SERVER_PORT" default:"8080"`
}

func (s *Server) StartServer() error {
	server := &Server{}
	if err := envconfig.Process("", server); err != nil {
		return err
	}
	if err := s.serveTrainer(); err != nil {
		return err
	}
	if err := s.Engine.Run(fmt.Sprintf(":%s", server.Port)); err != nil {
		return err
	}
	return nil
}

func (s *Server) serveTrainer() error {
	t := s.Engine.Group("/trainer")
	{
		t.POST("", s.TrainerController.Create)
		t.POST("login", s.TrainerController.Login)
	}
	return nil
}

func NewServer(ginEngine *gin.Engine, controller *trainer.Controller) *Server{
	return &Server{
		Engine: ginEngine,
		TrainerController: controller,
	}
}