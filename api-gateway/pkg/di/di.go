package di

import (
	server "github.com/ahdaan67/JobQuest/pkg/api"
	"github.com/ahdaan67/JobQuest/pkg/api/handler"
	"github.com/ahdaan67/JobQuest/pkg/client"
	"github.com/ahdaan67/JobQuest/pkg/config"
	"github.com/ahdaan67/JobQuest/pkg/helper"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {

	adminClient := client.NewAdminClient(cfg)
	adminHandler := handler.NewAdminHandler(adminClient)

	employerClient := client.NewEmployerClient(cfg)
	employerHandler := handler.NewEmployerHandler(employerClient)

	jobSeekerClient := client.NewJobSeekerClient(cfg)
	jobSeekerHandler := handler.NewJobSeekerHandler(jobSeekerClient)

	jobClient := client.NewJobClient(cfg)
	jobHandler := handler.NewJobHandler(jobClient,jobSeekerClient,cfg)

	helper := helper.NewHelper(&cfg)
	chatClient := client.NewChatClient(cfg)
	chatHandler := handler.NewChatHandler(chatClient, helper)

	v := handler.NewVideoCallHandler()


	fc := client.NewfollowCompanyClient(cfg)
	fh := handler.NewFollowCompanyHandler(fc, employerClient)

	
	serverHTTP := server.NewServerHTTP(adminHandler, employerHandler, jobSeekerHandler, jobHandler, chatHandler, v, fh)

	return serverHTTP, nil
}