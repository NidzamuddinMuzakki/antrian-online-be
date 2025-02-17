package repository

import commonDS "antrian-golang/common/data_source"

// @Notice: Register your repositories here

type IRegistry interface {
	GetTipePasienRepository() ITipePasienRepo
	GetUserRepository() IUserRepo
	GetRoleRepository() IRoleRepo
	GetLoketRepository() ILoketRepo
	GetAntrianRepository() IAntrianRepo
	GetUtilTx() *commonDS.TransactionRunner
}

type Registry struct {
	tipePasienRepository ITipePasienRepo
	userRepository       IUserRepo
	roleRepository       IRoleRepo
	loketRepository      ILoketRepo
	antrianRepository    IAntrianRepo
	masterUtilTx         *commonDS.TransactionRunner
}

func NewRegistryRepository(
	masterUtilTx *commonDS.TransactionRunner,
	userRepository IUserRepo,
	roleRepository IRoleRepo,
	loketRepository ILoketRepo,
	antrianRepository IAntrianRepo,
	tipePasienRepository ITipePasienRepo,
) *Registry {
	return &Registry{
		masterUtilTx:         masterUtilTx,
		tipePasienRepository: tipePasienRepository,
		userRepository:       userRepository,
		roleRepository:       roleRepository,
		loketRepository:      loketRepository,
		antrianRepository:    antrianRepository,
	}
}

func (r Registry) GetUtilTx() *commonDS.TransactionRunner {
	return r.masterUtilTx
}
func (r Registry) GetTipePasienRepository() ITipePasienRepo {
	return r.tipePasienRepository
}

func (r Registry) GetUserRepository() IUserRepo {
	return r.userRepository
}

func (r Registry) GetRoleRepository() IRoleRepo {
	return r.roleRepository
}

func (r Registry) GetLoketRepository() ILoketRepo {
	return r.loketRepository
}

func (r Registry) GetAntrianRepository() IAntrianRepo {
	return r.antrianRepository
}
