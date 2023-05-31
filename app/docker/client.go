package docker

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
) // paquete encargado de gestionar la conexión con Docker

type DockerClient struct {
	// definimos el contexto que usaremos para gestionar docker en background
	ctx context.Context
	// cliente real de docker
	client *client.Client
}

// NewDocker devuelve un cliente que se comunica con docker a través de la API de Docker
func NewDockerClient() (*DockerClient, error) {
	// obtenemos el cliente de docker usando las variables de entorno propias de docker y activamos la compatibilidad con todas ñas versiones de docker para evitar posibles errores
	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	// devolvemos el cliente de docker
	return &DockerClient{
		ctx:    context.Background(),
		client: docker,
	}, nil
}

// Funcion para levantar el laboratorio
func (d DockerClient) StartLab(user string, image string) error {
	// TODO: check if user has already deployed this lab
	// forzamos la descarga de la imagen
	if _, err := d.client.ImagePull(d.ctx, image, types.ImagePullOptions{}); err != nil {
		return err
	}
	// creamos el contenedor
	id, err := d.createContainer(user, image)
	if err != nil {
		return err
	}
	// iniciamos el contenedor
	return d.client.ContainerStart(d.ctx, id, types.ContainerStartOptions{})

}

// Funcion que detiene el laboratorio
func (d DockerClient) StopLab(user string, image string) error {
	var id string // TODO: Get the ID from the database
	// Paramos el contenedor si esta en ejecucion
	if d.containerIsRunning(id) {
		if err := d.client.ContainerStop(d.ctx, id, container.StopOptions{}); err != nil {
			return err
		}
	}
	// eliminamos el contenedor
	return d.client.ContainerRemove(d.ctx, id, types.ContainerRemoveOptions{
		Force:         true,
		RemoveLinks:   true,
		RemoveVolumes: true,
	})
}

// Creamos el contenedor
func (d DockerClient) createContainer(user string, image string) (string, error) {

	// creamos el contenedor
	name := fmt.Sprintf("%s-%s", user, strings.ReplaceAll(image, "/", "_"))
	response, err := d.client.ContainerCreate(d.ctx, &container.Config{
		Image: image,
	}, nil, nil, nil, name)
	if err != nil {
		return "", err
	}
	// devolvemos el ID del nuevo contenedor creado
	return response.ID, nil
}

// Buscamos si el contenedor con el ID especificado esta en ejecución
func (d DockerClient) containerIsRunning(id string) bool {
	// obtenemos todos los contenedores en ejecucion
	containers, err := d.client.ContainerList(d.ctx, types.ContainerListOptions{})
	if err != nil {
		return false
	}
	// buscamos el contenedor que queremos comprobar
	for _, container := range containers {
		if container.ID == id {
			return true
		}
	}
	// si no lo encontramos devolvemos false
	return false

}
