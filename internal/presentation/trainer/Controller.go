package trainer

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"pokemon-back/internal/core/usecase/trainer"
	"pokemon-back/internal/presentation/trainer/request"
	"pokemon-back/internal/presentation/trainer/response"
)

type Controller struct {
	TrainerUseCase trainer.ICrear
	LoginUseCase trainer.ILogin
	Validate *validator.Validate
}

func (c *Controller) Create(context *gin.Context) {
	body := &request.CreateRequest{}
	if err := context.ShouldBindJSON(body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := c.Validate.Struct(body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	t, err := c.TrainerUseCase.CrearTrainer(body.FromThis())
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, response.NewCreateResponse(t))
	return
}

func (c *Controller) Login(context *gin.Context) {
	body := &request.LoginRequest{}
	if err := context.ShouldBindJSON(body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := c.Validate.Struct(body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	l, err := c.LoginUseCase.Autenticar(body.FromThis())
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	result := response.NewLoginResponse(l.Token)
	context.JSON(http.StatusOK, result)
	return
}

func NewController(trainerUseCase *trainer.Crear, loginUseCase *trainer.Login, validator *validator.Validate) *Controller {
	return &Controller{
		TrainerUseCase: trainerUseCase,
		LoginUseCase: loginUseCase,
		Validate: validator,
	}
}
