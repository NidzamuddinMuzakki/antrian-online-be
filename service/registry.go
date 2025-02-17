package service

import (
	"antrian-golang/service/health"
)

// @Notice: Register your services here

type IRegistry interface {
	GetHealth() health.IHealth
	GetTipePasienService() ITipePasienService
	GetUserService() IUserService
	GetLoketService() ILoketService
	GetAntrianService() IAntrianService
	GetRoleService() IRoleService
}

type Registry struct {
	health            health.IHealth
	tipePasienService ITipePasienService
	loketService      ILoketService
	userService       IUserService
	antrianService    IAntrianService
	roleService       IRoleService
}

func NewRegistry(health health.IHealth, tipePasienService ITipePasienService, userService IUserService, loketService ILoketService, antrianService IAntrianService, roleService IRoleService) *Registry {
	return &Registry{
		health:            health,
		tipePasienService: tipePasienService,
		userService:       userService,
		loketService:      loketService,
		antrianService:    antrianService,
		roleService:       roleService,
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

func (r *Registry) GetLoketService() ILoketService {
	return r.loketService
}

func (r *Registry) GetAntrianService() IAntrianService {
	return r.antrianService
}

func (r *Registry) GetRoleService() IRoleService {
	return r.roleService
}
