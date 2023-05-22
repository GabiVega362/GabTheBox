package docker

import "github.com/docker/docker/client"

func NewDocker() (*client.Client, error) {
	//Create the docker client using env variables
	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	//Return the docker client
	return docker, nil
}
