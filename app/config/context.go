package config

import (
	"github.com/docker/docker/client" // paquete encargado de gestionar la conexión con Docker

	"github.com/gabivega362/gabthebox/app/docker"
)

// Config es la estructura que contiene las variables globales de la aplicación (argumentos, variables de entorno, clientes, etc...)
type Config struct {
	// Args es la estructura que contiene los argumentos pasados por terminal o por variables de entorno
	Args *Args
	// Docker es el cliente usado para gestionar contenedores a través de la API de Docker
	Docker *client.Client
}

// NewConfig devuelve un nuevo contexto de la aplicación
func NewConfig() (*Config, error) {
	// obtenemos la conexión con  el socket de docker
	docker, err := docker.NewDocker()
	if err != nil {
		return nil, err
	}

	// devolvemos el contexto
	return &Config{
		Args:   parseArgs(),
		Docker: docker,
	}, nil
}
