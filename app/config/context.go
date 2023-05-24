package config

import (
	"github.com/docker/docker/client" // paquete encargado de gestionar la conexión con Docker

	"github.com/gabivega362/gabthebox/app/docker"
)

// Context es la estructura que contiene las variables globales de la aplicación (argumentos, variables de entorno, clientes, etc...)
type Context struct {
	// Args es la estructura que contiene los argumentos pasados por terminal o por variables de entorno
	Args *Args
	// Docker es el cliente usado para gestionar contenedores a través de la API de Docker
	Docker *client.Client
}

// NewContext devuelve un nuevo contexto de la aplicación
func NewContext() (*Context, error) {
	// obtenemos la conexión con  el socket de docker
	docker, err := docker.NewDocker()
	if err != nil {
		return nil, err
	}

	// devolvemos el contexto
	return &Context{
		Args:   parseArgs(),
		Docker: docker,
	}, nil
}
