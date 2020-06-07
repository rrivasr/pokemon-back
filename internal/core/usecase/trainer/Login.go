package trainer

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"pokemon-back/internal/config"
	"pokemon-back/internal/core/domain"
	"pokemon-back/internal/repository"
	"time"
)

type ILogin interface {
	Autenticar(trainerDTO *TrainerDTO ) (*TrainerDTO, error)
}

type Login struct {
	Repository repository.IRepository
	TokenConfig *config.TokenConfig
}

func (l *Login) Autenticar(trainerDTO *TrainerDTO) (*TrainerDTO, error) {
	trainer, _ := trainerDTO.fromThis()
	contrasenaActual := *trainer.Contrasena
	err := l.Repository.ObtenerTrainerByNombreAndPassword(trainer)
	if err != nil {
		return nil, err
	}
	if !*trainer.Habilitado {
		return nil, errors.New("El usuario está deshabilitado")
	}
	if err := l.validarContrasena(contrasenaActual, *trainer.Contrasena); err != nil {
		return nil, err
	}
	trainer.Contrasena = nil
	result, _ := ToThis(trainer)
	token, err := l.crearToken(trainer)
	if err != nil {
		return nil, err
	}
	result.Token = token
	return result, nil
}

func (l *Login) crearToken(trainer *domain.Trainer) (*string, error) {
	claims := jwt.MapClaims{
		"nombre": *trainer.Nombre,
		"exp": time.Now().Add(time.Minute * time.Duration(l.TokenConfig.Time)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	token.Header["kid"] = l.TokenConfig.Kid
	tokenString, err := token.SignedString([]byte(l.TokenConfig.Key))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (l *Login) validarContrasena(loginContrasena string, contrasena string) error {
	if err :=bcrypt.CompareHashAndPassword([]byte(contrasena), []byte(loginContrasena)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return errors.New("Usuario o contraseña inválida")
		}
		return err
	}
	return nil
}

func NewLogin(r *repository.Repository, tc *config.TokenConfig) *Login {
	return &Login{
		Repository: r,
		TokenConfig: tc,
	}
}
