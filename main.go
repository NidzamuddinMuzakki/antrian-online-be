package main

import (
	"context"
	"time"
	_ "time/tzdata"

	commonCache "antrian-golang/common/cache"

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
	// commonDs "antrian-golang/common/data_source"
	"antrian-golang/common/logger"

	commonPanicRecover "antrian-golang/common/middleware/gin/panic_recovery"

	commonRegistry "antrian-golang/common/registry"

	commonTime "antrian-golang/common/time"

	commonValidator "antrian-golang/common/validator"

	// Import third parties here
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	_ "github.com/spf13/viper/remote"
)

func main() {
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
		// AppName: config.Cold.AppName,
		// Debug: config.Hot.AppDebug,
	})
	// Validator
	validator := commonValidator.New()
	// Sentry

	// Database
	// - Master
	// master, err := commonDs.NewDB(&commonDs.Config{
	// 	Driver:                config.Cold.DBMysqlMasterDriver,
	// 	Host:                  config.Cold.DBMysqlMasterHost,
	// 	Port:                  config.Cold.DBMysqlMasterPort,
	// 	DBName:                config.Cold.DBMysqlMasterDBName,
	// 	User:                  config.Cold.DBMysqlMasterUser,
	// 	Password:              config.Cold.DBMysqlMasterPassword,
	// 	SSLMode:               config.Cold.DBMysqlMasterSSLMode,
	// 	MaxOpenConnections:    config.Cold.DBMysqlMasterMaxOpenConnections,
	// 	MaxLifeTimeConnection: config.Cold.DBMysqlMasterMaxLifeTimeConnection,
	// 	MaxIdleConnections:    config.Cold.DBMysqlMasterMaxIdleConnections,
	// 	MaxIdleTimeConnection: config.Cold.DBMysqlMasterMaxIdleTimeConnection,
	// })
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
	// masterTx := commonDs.NewTransactionRunner(master)
	// masterUtilTx := util.NewTransactionRunner(master)
	// repoRegistry := repository.NewRegistryRepository(
	// 	masterTx,
	// 	masterUtilTx,
	// 	loadingFeeRepo,
	// 	insuranceTypeRepo,
	// 	carMappingRepo,
	// 	pricingOtrCategoryRepo,
	// 	vehicleCategoryRepo,
	// 	insuranceOrdeVehicleRepo,
	// 	insuranceOrdeVehicleDetailsRepo,
	// 	packageInsuranceRepo,
	// 	partnerRepo,
	// 	insuranceMappingRepo,
	// 	mappingAreaRepositoryRepo,
	// 	vehicleGroupRepo,
	// 	carModelRepo,
	// 	vehicleHistoryRepo,
	// 	packageInsuranceLifeRepo,
	// 	tenorRepo,
	// 	principalCategoryRepo,
	// 	insuranceOrdeLifeRepo,
	// 	insuranceOrdeLifeDetailRepo,
	// 	lifeHistoryRepo,
	// 	gcStorage,
	// 	insuranceTypeLifeRepo,
	// 	carTypeRepo,
	// 	insuranceAreaRepo,
	// )
	// End Repositories //

	// Map for partner portion

	// Start Services //
	healthService := serviceHealth.NewHealth()
	// loadingFeeService := calculationService.NewLoadingFeeService(common, repoRegistry, actLogClientLoadingFee)
	// internalCalculationService := calculationService.NewInternalCalculationService(common, repoRegistry, mapPartnerPortionVehicle, raksa, actLogClientRedisRestart)
	// insuranceTypeService := configService.NewInsuraceTypeService(common, repoRegistry, actLogClientInsuranceType)
	// packageInsuranceService := calculationService.NewPackageInsuranceService(common, repoRegistry)
	// partnerService := partnerService.NewPartnerInsuranceService(common, partnerRepo, insuranceMappingRepo, actLogClientPartnerInsuranceMapping, repoRegistry)
	// pricingOtrCategoryService := configService.NewPricingOtrCategoryService(common, repoRegistry, actLogClientPricingCategory)
	// packageVehicleService := packageService.NewPackageVehicleService(common, repoRegistry, actLogPackageVehicle)

	// callbackPolisService := callBackPolisService.NewCallbackPolisService(common, repoRegistry, apiInvoice, notifcationService)
	// packageLifeService := packageService.NewPackageLifeService(common, repoRegistry, actLogPackageLife)
	// internalLifeCalculationService := calculationService.NewInternalLifeCalculationService(common, repoRegistry, mapPartnerPortionLife, actLogClientRedisRestart)
	// insuranceAreaService := configService.NewInsuraceAreaService(common, repoRegistry, actLogClientInsuranceArea)
	// carTypeService := configService.NewCarTypeService(common, repoRegistry, actLogClientCarType)
	// principalCategoryService := configService.NewPrincipalCategoryService(common, repoRegistry, actLogClientPrincipalCategory)
	// securityCService := securityService.NewSignatureService(common)
	// serviceRegistry := service.NewRegistry(
	// 	healthService,
	// )
	// End Deliveries //

	// Start Deliveries //
	healthDelivery := httpDeliveryHealth.NewHealth(common, healthService)

	registryDelivery := httpDelivery.NewRegistry(healthDelivery)

	// End Deliveries //

	// Start HTTP Server //
	httpServer := cmdHttp.NewServer(
		common,
		registryDelivery,
	)
	httpServer.Serve(ctx)
	// End HTTP Server //
}
