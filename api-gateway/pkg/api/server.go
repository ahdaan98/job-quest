package server

import (
	"github.com/ahdaan67/JobQuest/pkg/api/handler"
	"github.com/ahdaan67/JobQuest/pkg/api/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler, employerHandler *handler.EmployerHandler, jobSeekerHandler *handler.JobSeekerHandler, jobHandler *handler.JobHandler) *ServerHTTP {

	router := gin.New()

	router.Use(gin.Logger())
	// Route for admin auth
	router.POST("/admin/login", adminHandler.LoginHandler)
	router.POST("/admin/signup", adminHandler.AdminSignUp)

	// Route for employer auth
	router.POST("/employer/signup", employerHandler.EmployerSignUp)
	router.POST("/employer/login", employerHandler.EmployerLogin)

	router.POST("/job-seeker/signup", jobSeekerHandler.JobSeekerSignUp)
	router.POST("/job-seeker/login", jobSeekerHandler.JobSeekerLogin)

	router.Use(middleware.JobSeekerAuthMiddleware())
	{
		router.GET("/job-seeker/view-jobs", jobHandler.ViewAllJobs)
		router.GET("/job-seeker/jobs", jobHandler.GetJobDetails)

		router.GET("/job-seeker/saved-jobs", jobHandler.GetASavedJob)
		router.POST("/job-seeker/save-jobs", jobHandler.SaveAJob)
		router.DELETE("/job-seeker/saved-jobs", jobHandler.DeleteSavedJob)
	}

	router.Use(middleware.EmployerAuthMiddleware())
	{
		router.POST("/employer/job-post", jobHandler.PostJobOpening)
		router.GET("/employer/all-job-postings", jobHandler.GetAllJobs)
		router.GET("/employer/job-postings", jobHandler.GetAJob)
		router.DELETE("/employer/job-postings", jobHandler.DeleteAJob)
		router.PUT("/employer/job-postings", jobHandler.UpdateAJob)

		router.GET("/employer/company", employerHandler.GetCompanyDetails)
		router.PUT("/employer/company", employerHandler.UpdateCompany)

	}

	return &ServerHTTP{engine: router}
}

func (s *ServerHTTP) Start() {
	log.Printf("starting server on :8000")
	err := s.engine.Run(":8000")
	if err != nil {
		log.Printf("error while starting the server")
	}
}