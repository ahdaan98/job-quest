package di

import (
	server "job_service/pkg/api"
	"job_service/pkg/api/service"
	"job_service/pkg/config"
	"job_service/pkg/db"
	"job_service/pkg/repository"
	"job_service/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	// Connect to the database
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	jobRepository := repository.NewJobRepository(gormDB)
	jobUseCase := usecase.NewJobUseCase(jobRepository)
	jobServiceServer := service.NewJobServer(jobUseCase)

	grpcServer, err := server.NewGRPCServer(cfg, jobServiceServer)
	if err != nil {
		return nil, err
	}

	return grpcServer, nil
}