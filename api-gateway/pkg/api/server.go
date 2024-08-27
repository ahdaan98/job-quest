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

	jobSeekerRoutes := router.Group("/")
	jobSeekerRoutes.Use(middleware.JobSeekerAuthMiddleware())
	{
		jobSeekerRoutes.POST("/job-seeker/apply-job", jobHandler.ApplyJob)
		jobSeekerRoutes.GET("/job-seeker/view-jobs", jobHandler.ViewAllJobs)
		jobSeekerRoutes.GET("/job-seeker/jobs", jobHandler.GetJobDetails)
		jobSeekerRoutes.GET("/job-seeker/saved-jobs", jobHandler.GetASavedJob)
		jobSeekerRoutes.POST("/job-seeker/save-jobs", jobHandler.SaveAJob)
		jobSeekerRoutes.DELETE("/job-seeker/saved-jobs", jobHandler.DeleteSavedJob)
	}

	// Employer authenticated routes
	employerRoutes := router.Group("/")
	employerRoutes.Use(middleware.EmployerAuthMiddleware())
	{
		employerRoutes.POST("/employer/job-post", jobHandler.PostJobOpening)
		employerRoutes.GET("/employer/all-job-postings", jobHandler.GetAllJobs)
		employerRoutes.GET("/employer/job-postings", jobHandler.GetAJob)
		employerRoutes.DELETE("/employer/job-postings", jobHandler.DeleteAJob)
		employerRoutes.PUT("/employer/job-postings", jobHandler.UpdateAJob)

		employerRoutes.GET("/employer/company", employerHandler.GetCompanyDetails)
		employerRoutes.PUT("/employer/company", employerHandler.UpdateCompany)
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