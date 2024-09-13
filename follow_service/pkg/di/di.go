package di

import (
	server "follow/pkg/api"
	"follow/pkg/api/service"
	"follow/pkg/config"
	"follow/pkg/db"
	"follow/pkg/repository"
	"follow/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	// Connect to the database
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	cr := repository.NewConnectionRepository(gormDB)
	cu := usecase.NewCompanyUseCase(cr)
	cs := service.NewCompanyServer(cu)

	grpcServer, err := server.NewGRPCServer(cfg, cs)
	if err != nil {
		return nil, err
	}

	return grpcServer, nil
}