package di

import (
	server "github.com/ahdaan67/JobQuest/pkg/api"
	"github.com/ahdaan67/JobQuest/pkg/api/handler"
	"github.com/ahdaan67/JobQuest/pkg/client"
	"github.com/ahdaan67/JobQuest/pkg/config"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {

	adminClient := client.NewAdminClient(cfg)
	adminHandler := handler.NewAdminHandler(adminClient)

	employerClient := client.NewEmployerClient(cfg)
	employerHandler := handler.NewEmployerHandler(employerClient)

	jobSeekerClient := client.NewJobSeekerClient(cfg)
	jobSeekerHandler := handler.NewJobSeekerHandler(jobSeekerClient)

	jobClient := client.NewJobClient(cfg)
	jobHandler := handler.NewJobHandler(jobClient)

	
	serverHTTP := server.NewServerHTTP(adminHandler, employerHandler, jobSeekerHandler, jobHandler)

	return serverHTTP, nil
}