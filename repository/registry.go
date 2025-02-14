package repository

import commonDS "antrian-golang/common/data_source"

// @Notice: Register your repositories here

type IRegistry interface {
	GetTipePasienRepository() ITipePasienRepo
	GetUserRepository() IUserRepo
	GetUtilTx() *commonDS.TransactionRunner
}

type Registry struct {
	tipePasienRepository ITipePasienRepo
	userRepository       IUserRepo
	masterUtilTx         *commonDS.TransactionRunner
}

func NewRegistryRepository(
	masterUtilTx *commonDS.TransactionRunner,
	userRepository IUserRepo,
	tipePasienRepository ITipePasienRepo,

) *Registry {
	return &Registry{
		masterUtilTx:         masterUtilTx,
		tipePasienRepository: tipePasienRepository,
		userRepository:       userRepository,
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
