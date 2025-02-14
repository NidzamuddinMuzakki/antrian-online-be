package http

import (
	"antrian-golang/delivery/http/health"
)

// @Notice: Register your http deliveries here

type IRegistry interface {
	GetHealth() health.IHealth
	GetTipePasienDelivery() ITipePasienDelivery
	GetUserDelivery() IUserDelivery
}

type Registry struct {
	health             health.IHealth
	tipePasienDelivery ITipePasienDelivery
	userDelivery       IUserDelivery
}

func NewRegistry(
	health health.IHealth,
	tipePasienDelivery ITipePasienDelivery,
	userDelivery IUserDelivery,
) *Registry {
	return &Registry{
		health:             health,
		tipePasienDelivery: tipePasienDelivery,
		userDelivery:       userDelivery,
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
