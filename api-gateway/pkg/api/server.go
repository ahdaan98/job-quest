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

func NewServerHTTP(adminHandler *handler.AdminHandler, employerHandler *handler.EmployerHandler, jobSeekerHandler *handler.JobSeekerHandler, jobHandler *handler.JobHandler, chatHandler *handler.ChatHandler, videocallHandler *handler.VideoCallHandler,followCompanyHandler *handler.FollowCompanyHandler) *ServerHTTP {

	router := gin.New()

	router.Use(gin.Logger())

	router.Static("/static", "./static")
	router.LoadHTMLGlob("template/*")

	router.GET("/exit", videocallHandler.ExitPage)
	router.GET("/error", videocallHandler.ErrorPage)
	router.GET("/index", videocallHandler.IndexedPage)
	// Route for admin auth
	router.POST("/admin/login", adminHandler.LoginHandler)
	router.POST("/admin/signup", adminHandler.AdminSignUp)

	// Route for employer auth
	router.POST("/employer/signup", employerHandler.EmployerSignUp)
	router.POST("/employer/login", employerHandler.EmployerLogin)

	router.POST("/job-seeker/signup", jobSeekerHandler.JobSeekerSignUp)
	router.POST("/job-seeker/login", jobSeekerHandler.JobSeekerLogin)
	router.GET("/job-seeker/linkedin/signin", jobSeekerHandler.LinkedinSignIn)
	router.GET("/job-seeker/linkedin/complete/signin", jobSeekerHandler.CompleteLinkedInSignIn)

	router.POST("/job-seeker/otp/signup", jobSeekerHandler.JobSeekerOTPSignUp)
	router.POST("/job-seeker/verify/otp", jobSeekerHandler.JobSeekerVerifyOTP)

	jobSeekerRoutes := router.Group("/")
	jobSeekerRoutes.Use(middleware.JobSeekerAuthMiddleware())
	{
		jobSeekerRoutes.POST("/job-seeker/apply-job", jobHandler.ApplyJob)
		jobSeekerRoutes.GET("/job-seeker/view-jobs", jobHandler.ViewAllJobs)
		jobSeekerRoutes.GET("/job-seeker/jobs", jobHandler.GetJobDetails)
		jobSeekerRoutes.GET("/job-seeker/saved-jobs", jobHandler.GetASavedJob)
		jobSeekerRoutes.POST("/job-seeker/save-jobs", jobHandler.SaveAJob)
		jobSeekerRoutes.DELETE("/job-seeker/saved-jobs", jobHandler.DeleteSavedJob)

		jobSeekerRoutes.GET("/job-seeker/activate/subscription", jobSeekerHandler.ActivateSubscriptionPlan)

		jobSeekerRoutes.POST("/follow", followCompanyHandler.FollowCompany)
		jobSeekerRoutes.POST("/unfollow", followCompanyHandler.UnfollowCompany)
		jobSeekerRoutes.GET("/isfollowing/:userID/:companyID", followCompanyHandler.IsFollowingCompany)
		jobSeekerRoutes.GET("/followedcompanies/:userID", followCompanyHandler.GetFollowedCompanies)
		jobSeekerRoutes.GET("/checkfollowrequest/:userID/:companyID", followCompanyHandler.CheckFollowRequestExists)
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
		employerRoutes.GET("/employer/get-applicants", jobHandler.GetApplicants)

		employerRoutes.GET("/employer/company", employerHandler.GetCompanyDetails)
		employerRoutes.PUT("/employer/company", employerHandler.UpdateCompany)

		router.GET("/employer/chat", chatHandler.EmployerMessage)
		employerRoutes.POST("/employer/chats", chatHandler.GetChat)
		router.GET("/group/:groupID/chat", chatHandler.GroupMessage)

		employerRoutes.PUT("/employer/update/apply/jobs", jobHandler.UpdateApplyJob)
		employerRoutes.GET("/employer/get/applicants", jobHandler.GetAcceptedApplicants)
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
