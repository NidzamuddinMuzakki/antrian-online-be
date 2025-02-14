package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"time"
	_ "time/tzdata"

	commonCache "antrian-golang/common/cache"
	"antrian-golang/lib/security"
	"antrian-golang/repository"
	"antrian-golang/service"

	// commonGCP "bitbucket.org/moladinTech/go-lib-common/client/gcp"

	// calculationRepo "bitbucket.org/moladinTech/insurance-cofi/repository/calculation"

	"antrian-golang/config"

	// Import clients here
	// actLogClientLib "bitbucket.org/moladinTech/go-lib-activity-log/client"

	// Import repositories here
	// ...

	// Import services here
	serviceHealth "antrian-golang/service/health"

	// Import deliveries here

	httpDelivery "antrian-golang/delivery/http"
	httpDeliveryHealth "antrian-golang/delivery/http/health"

	// Import cmd here
	cmdHttp "antrian-golang/cmd/http"

	// Import common lib here
	commonDs "antrian-golang/common/data_source"
	"antrian-golang/common/logger"

	commonPanicRecover "antrian-golang/common/middleware/gin/panic_recovery"

	commonRegistry "antrian-golang/common/registry"

	commonTime "antrian-golang/common/time"

	commonValidator "antrian-golang/common/validator"

	// Import third parties here

	"github.com/golang-jwt/jwt"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	_ "github.com/spf13/viper/remote"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			time.Sleep(3 * time.Second)
		}
	}()
	ctx := context.Background()

	// Start Init //
	loc, err := time.LoadLocation(commonTime.LoadTimeZoneFromEnv())
	if err != nil {
		panic(err)
	}
	time.Local = loc
	// Configuration
	config.Init()
	// Logger
	logger.Init(logger.Config{
		AppName: config.Cold.AppName,
		Debug:   config.Hot.AppDebug,
	})
	// Validator
	validator := commonValidator.New()
	// Sentry

	// Database
	// - Master
	master, err := commonDs.NewDB(&commonDs.Config{
		Driver:                config.Cold.DBMysqlMasterDriver,
		Host:                  config.Cold.DBMysqlMasterHost,
		Port:                  config.Cold.DBMysqlMasterPort,
		DBName:                config.Cold.DBMysqlMasterDBName,
		User:                  config.Cold.DBMysqlMasterUser,
		Password:              config.Cold.DBMysqlMasterPassword,
		SSLMode:               config.Cold.DBMysqlMasterSSLMode,
		MaxOpenConnections:    config.Cold.DBMysqlMasterMaxOpenConnections,
		MaxLifeTimeConnection: config.Cold.DBMysqlMasterMaxLifeTimeConnection,
		MaxIdleConnections:    config.Cold.DBMysqlMasterMaxIdleConnections,
		MaxIdleTimeConnection: config.Cold.DBMysqlMasterMaxIdleTimeConnection,
	})
	if err != nil {
		panic(err)
	}
	// - Slave
	// slave, err := commonDs.NewDB(&commonDs.Config{
	// 	Driver:                config.Cold.DBMysqlSlaveDriver,
	// 	Host:                  config.Cold.DBMysqlSlaveHost,
	// 	Port:                  config.Cold.DBMysqlSlavePort,
	// 	DBName:                config.Cold.DBMysqlSlaveDBName,
	// 	User:                  config.Cold.DBMysqlSlaveUser,
	// 	Password:              config.Cold.DBMysqlSlavePassword,
	// 	SSLMode:               config.Cold.DBMysqlSlaveSSLMode,
	// 	MaxOpenConnections:    config.Cold.DBMysqlSlaveMaxOpenConnections,
	// 	MaxLifeTimeConnection: config.Cold.DBMysqlSlaveMaxLifeTimeConnection,
	// 	MaxIdleConnections:    config.Cold.DBMysqlSlaveMaxIdleConnections,
	// 	MaxIdleTimeConnection: config.Cold.DBMysqlSlaveMaxIdleTimeConnection,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// Activity Log Client

	// Auth

	// Tracer

	// Panic Recovery
	panicRecoveryMiddleware := commonPanicRecover.NewPanicRecovery(
		validator,
		commonPanicRecover.WithConfigEnv(config.Cold.AppEnv),
	)

	// Cache
	caches, err := commonCache.NewCache(
		commonCache.WithDriver(commonCache.RedisDriver),
		// commonCache.WithHost(config.Cold.RedisHost),
		commonCache.WithDatabase("0"),
		commonCache.WithPassword(""),
	)
	if err != nil {
		panic(err)
	}

	// Registry
	common := commonRegistry.NewRegistry(

		commonRegistry.WithValidator(validator),

		commonRegistry.WithPanicRecoveryMiddleware(panicRecoveryMiddleware),
		commonRegistry.WithCache(caches),
	)

	// End Init //

	// Start Clients //

	// End Clients //

	// Start Repositories //
	masterTx := commonDs.NewTransactionRunner(master)
	// masterUtilTx := util.NewTransactionRunner(master)
	tipePasionRepo := repository.NewTipePasienRepo(common, master)
	userRepo := repository.NewUserRepo(common, master)

	repoRegistry := repository.NewRegistryRepository(
		masterTx,
		userRepo,
		tipePasionRepo,
	)
	// End Repositories //

	// Map for partner portion

	var privatecdsaKey *ecdsa.PrivateKey
	if privatecdsaKey, err = jwt.ParseECPrivateKeyFromPEM([]byte(config.Cold.SignaturePrivateKey)); err != nil {
		panic(err)
	}

	var publicecdsaKey *ecdsa.PublicKey
	if publicecdsaKey, err = jwt.ParseECPublicKeyFromPEM([]byte(config.Cold.SignaturePublicKey)); err != nil {
		panic(err)
	}

	sec := security.NewJwtUtils(privatecdsaKey, publicecdsaKey)

	// Start Services //
	healthService := serviceHealth.NewHealth()
	tipePasienServ := service.NewTipePasienService(common, repoRegistry)
	userServ := service.NewUserService(common, repoRegistry, sec)

	serviceRegistry := service.NewRegistry(
		healthService,
		tipePasienServ,
		userServ,
	)
	// End Deliveries //

	// Start Deliveries //
	healthDelivery := httpDeliveryHealth.NewHealth(common, healthService)
	tipePasienDeliv := httpDelivery.NewTipePasienDelivery(common, serviceRegistry)
	userDeliv := httpDelivery.NewUserDelivery(common, serviceRegistry)

	registryDelivery := httpDelivery.NewRegistry(healthDelivery, tipePasienDeliv, userDeliv)

	// End Deliveries //

	// Start HTTP Server //
	httpServer := cmdHttp.NewServer(
		common,
		registryDelivery,
	)
	// wac, err := whatsapp.NewConn(20 * time.Second)
	// if err != nil {
	// 	panic(err)
	// }
	// ss, _ := os.Hostname()
	// sss := os.Environ()
	// validEnv := 0
	// for _, sVal := range sss {
	// 	if sVal == "nidzam=ganteng" {
	// 		validEnv = 1
	// 	}
	// }
	// if validEnv != 1 {
	// 	fmt.Println("maaf aplikasi anda terkena copyright")
	// 	time.Sleep(3 * time.Second)
	// 	panic("")

	// }
	// if ss != "DESKTOP-M6S1G24" {
	// 	panic("maaf aplikasi anda terkena copyright")
	// }
	// var input string
	// fmt.Print("Enter Password: ")
	// fmt.Scanln(&input)
	// if input != "nidzamganteng" {
	// 	// fmt.Println(input)
	// 	fmt.Println("maaf aplikasi anda terkena copyright")
	// 	time.Sleep(3 * time.Second)
	// 	panic("")
	// }

	httpServer.Serve(ctx)

	// End HTTP Server //
}
