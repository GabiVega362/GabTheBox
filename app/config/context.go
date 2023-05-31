package config

import (
	"database/sql"

	"github.com/gabivega362/gabthebox/app/database"
	"github.com/gabivega362/gabthebox/app/docker"
)

// Config es la estructura que contiene las variables globales de la aplicación (argumentos, variables de entorno, clientes, etc...)
type Config struct {
	// Args es la estructura que contiene los argumentos pasados por terminal o por variables de entorno
	Args *Args
	// Docker es el cliente usado para gestionar contenedores a través de la API de Docker
	Docker *docker.DockerClient
	// Cliente usado para gestionar la base de datos
	Database *sql.DB
}

// NewConfig devuelve un nuevo contexto de la aplicación
func NewConfig() (*Config, error) {
	// obtenemos los argumentos
	args := parseArgs()

	// obtenemos la conexión con  el socket de docker
	dckr, err := docker.NewDockerClient()
	if err != nil {
		return nil, err
	}

	// obtenemos la conexión con la base de datos
	db, err := database.NewDatabaseConn(args.DBUser, args.DBPass, args.DBName)
	if err != nil {
		return nil, err
	}

	// devolvemos el contexto
	return &Config{
		Args:     args,
		Docker:   dckr,
		Database: db,
	}, nil
}
