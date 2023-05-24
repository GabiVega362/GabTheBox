package docker

import "github.com/docker/docker/client" // paquete encargado de gestionar la conexión con Docker

// NewDocker devuelve un cliente que se comunica con docker a través de la API de Docker
func NewDocker() (*client.Client, error) {
	// obtenemos el cliente de docker usando las variables de entorno propias de docker y activamos la compatibilidad con todas ñas versiones de docker para evitar posibles errores
	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	// devolvemos el cliente de docker
	return docker, nil
}
