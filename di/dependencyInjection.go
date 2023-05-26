package di

import (
	"go.uber.org/dig"
	"log"
)

var Container = dig.New()

type RegisterFunction func(container *dig.Container) (*dig.Container, error)

func InitializeContainer(fs []RegisterFunction) {
	var err error

	for _, f := range fs {
		Container, err = f(Container)
		if err != nil {
			log.Fatalf("failed to register package: %v", err)
		}
	}
}
