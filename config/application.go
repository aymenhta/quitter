package config

import (
	"log"

	"github.com/aymenhta/quitter/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type application struct {
	InfoLog, ErrorLog *log.Logger
	UsersModel        *models.UserModel
}

var G application

func InitApplication(
	db *pgxpool.Pool,
	infoLog, errorlog *log.Logger,
) {
	G = application{
		InfoLog:    infoLog,
		ErrorLog:   errorlog,
		UsersModel: &models.UserModel{DB: db},
	}
}
