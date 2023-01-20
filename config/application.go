package config

import (
	"log"

	"github.com/aymenhta/quitter/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	InfoLog, ErrorLog *log.Logger
	UsersModel        *models.UserModel
}

var App Application

func InitApplication(
	db *pgxpool.Pool,
	infoLog, errorlog *log.Logger,
) {
	App = Application{
		InfoLog:    infoLog,
		ErrorLog:   errorlog,
		UsersModel: &models.UserModel{DB: db},
	}
}
