// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package empty_struct

// Injectors from wire.go:

func NewService() *Service {
	service := &Service{}
	return service
}