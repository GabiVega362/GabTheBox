package config

import (
	"github.com/docker/docker/client"
	"github.com/gabivega362/gabthebox/server/docker"
)

type Context struct {
	Args *Args
	Docker *client.Client
}

func NewContext() (*Context, error) { 
	//Get the Docker Client
	docker, err := docker.NewDocker()
	if err != nil {
		return nil, err
	}

	//Return the context
	return &Context{
		Args : parseArgs(),
		Docker: docker,
	}, nil
}
