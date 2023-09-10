package templates

const HOME_SERVICE = `package service

// @Service
type HomeService struct {
}

func (c *HomeService) GetHello() string {
	return "Hello World!"
}`
