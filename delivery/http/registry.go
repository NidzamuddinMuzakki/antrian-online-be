package http

import (
	"antrian-golang/delivery/http/health"
)

// @Notice: Register your http deliveries here

type IRegistry interface {
	GetHealth() health.IHealth
	GetTipePasienDelivery() ITipePasienDelivery
	GetUserDelivery() IUserDelivery
	GetLoketDelivery() ILoketDelivery
	GetAntrianDelivery() IAntrianDelivery
	GetRoleDelivery() IRoleDelivery
}

type Registry struct {
	health             health.IHealth
	tipePasienDelivery ITipePasienDelivery
	userDelivery       IUserDelivery
	loketDelivery      ILoketDelivery
	antrianDelivery    IAntrianDelivery
	roleDelivery       IRoleDelivery
}

func NewRegistry(
	health health.IHealth,
	tipePasienDelivery ITipePasienDelivery,
	userDelivery IUserDelivery,
	loketDelivery ILoketDelivery,
	antrianDelivery IAntrianDelivery,
	roleDelivery IRoleDelivery,

) *Registry {
	return &Registry{
		health:             health,
		tipePasienDelivery: tipePasienDelivery,
		userDelivery:       userDelivery,
		loketDelivery:      loketDelivery,
		antrianDelivery:    antrianDelivery,
		roleDelivery:       roleDelivery,
	}
}

func (r *Registry) GetHealth() health.IHealth {
	return r.health
}

func (r *Registry) GetTipePasienDelivery() ITipePasienDelivery {
	return r.tipePasienDelivery
}

func (r *Registry) GetUserDelivery() IUserDelivery {
	return r.userDelivery
}

func (r *Registry) GetLoketDelivery() ILoketDelivery {
	return r.loketDelivery
}

func (r *Registry) GetAntrianDelivery() IAntrianDelivery {
	return r.antrianDelivery
}

func (r *Registry) GetRoleDelivery() IRoleDelivery {
	return r.roleDelivery
}
