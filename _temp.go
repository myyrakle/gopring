package service

import "github.com/myyrakle/gopring/src/repository"

// @Service
type HomeService struct {
	userRepository *repository.UserRepository
}

func (c *HomeService) GetHello() string {
	c.userRepository.CreateUser("test")
	return "Hello World!"
}
