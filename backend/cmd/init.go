package main

import (
	"github.com/54m/api_gen-example/backend/domain/user"
	"github.com/54m/api_gen-example/backend/infrastructure/userimpl"
	"github.com/54m/api_gen-example/backend/interfaces/props"
)

type repositories struct {
	userRepository user.Repository
}

func initRepositories() *repositories {
	return &repositories{
		userRepository: userimpl.NewUserImpl(),
	}
}

func initControllerProps(repo *repositories) *props.ControllerProps {
	p := &props.ControllerProps{
		UserService: user.NewService(repo.userRepository),
	}

	return p
}
