package docker

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
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
func (d DockerClient) StartLab(user string, image string) (string, uint16, error) {
	// forzamos la descarga de la imagen
	out, err := d.client.ImagePull(d.ctx, image, types.ImagePullOptions{})
	if err != nil {
		return "", 0, err
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
	// creamos el contenedor
	id, port, err := d.createContainer(image)
	if err != nil {
		return "", 0, err
	}
	// iniciamos el contenedor
	err = d.client.ContainerStart(d.ctx, id, types.ContainerStartOptions{})
	if err != nil {
		return "", 0, err
	}
	// devolvemos los datos
	return id, port, nil
}

// Funcion que detiene el laboratorio
func (d DockerClient) StopLab(id string) error {
	// Paramos el contenedor si esta en ejecucion
	if d.containerIsRunning(id) {
		if err := d.client.ContainerStop(d.ctx, id, container.StopOptions{}); err != nil {
			return err
		}
	}
	// eliminamos el contenedor
	return d.client.ContainerRemove(d.ctx, id, types.ContainerRemoveOptions{
		Force:         true,
		RemoveVolumes: true,
	})
}

// Creamos el contenedor
func (d DockerClient) createContainer(image string) (string, uint16, error) {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	port := uint16(seed.Intn(3000)) + 3000

	// creamos el contenedor
	response, err := d.client.ContainerCreate(d.ctx, &container.Config{
		Image: image,
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"80/tcp": {
				{
					HostIP:   "0.0.0.0",
					HostPort: fmt.Sprint(port),
				},
			},
		},
	}, nil, nil, "")
	if err != nil {
		return "", 0, err
	}
	// devolvemos el ID del nuevo contenedor creado
	return response.ID, port, nil
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
