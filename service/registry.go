package service

import (
	"antrian-golang/service/health"
)

// @Notice: Register your services here

type IRegistry interface {
	GetHealth() health.IHealth
	GetTipePasienService() ITipePasienService
	GetUserService() IUserService
}

type Registry struct {
	health            health.IHealth
	tipePasienService ITipePasienService
	userService       IUserService
}

func NewRegistry(health health.IHealth, tipePasienService ITipePasienService, userService IUserService) *Registry {
	return &Registry{
		health:            health,
		tipePasienService: tipePasienService,
		userService:       userService,
	}
}

func (r *Registry) GetHealth() health.IHealth {
	return r.health
}

func (r *Registry) GetTipePasienService() ITipePasienService {
	return r.tipePasienService
}

func (r *Registry) GetUserService() IUserService {
	return r.userService
}
