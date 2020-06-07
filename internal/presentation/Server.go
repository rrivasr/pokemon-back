package presentation

import (
	"github.com/gin-gonic/gin"
	"pokemon-back/internal/presentation/trainer"
)

type Server struct {
	Engine *gin.Engine
	TrainerController *trainer.Controller
}

func (s *Server) StartServer() error {
	if err := s.serveTrainer(); err != nil {
		return err
	}
	if err := s.Engine.Run(":8080"); err != nil {
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