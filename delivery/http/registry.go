package http

import (
	"antrian-golang/delivery/http/health"
)

// @Notice: Register your http deliveries here

type IRegistry interface {
	GetHealth() health.IHealth
}

type Registry struct {
	health health.IHealth
}

func NewRegistry(
	health health.IHealth,

) *Registry {
	return &Registry{
		health: health,
	}
}

func (r *Registry) GetHealth() health.IHealth {
	return r.health
}
