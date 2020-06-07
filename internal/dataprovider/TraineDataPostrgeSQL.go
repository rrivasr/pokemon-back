package dataprovider

import (
	"pokemon-back/internal/core/domain"
	"database/sql"
)


type TrainerDataPostgreSQL struct {
	Db *sql.DB
}

func (t *TrainerDataPostgreSQL) Create(nuevoTrainer *domain.Trainer) (error) {
	p, err := t.Db.Prepare("INSERT INTO trainer(nombre,contrasena) VALUES ($1,$2) RETURNING ID")
	if err != nil {
		return err
	}
	defer p.Close()

	if err := p.QueryRow(*nuevoTrainer.Nombre, *nuevoTrainer.Contrasena).Scan(nuevoTrainer.Id); err != nil {
		return err
	}
	return nil
}

func (t *TrainerDataPostgreSQL) ObtenerByNombre(trainer *domain.Trainer) (error) {
	p, err := t.Db.Prepare("SELECT contrasena,habilitado FROM trainer WHERE nombre=$1")
	if err != nil {
		return err
	}
	defer p.Close()
	if err := p.QueryRow(*trainer.Nombre).Scan(trainer.Contrasena, trainer.Habilitado); err != nil {
		return err
	}
	return nil
}

func NewTrainerDataPostgreSQL(engine *sql.DB) *TrainerDataPostgreSQL {
	return &TrainerDataPostgreSQL{
		Db: engine,
	}
}
