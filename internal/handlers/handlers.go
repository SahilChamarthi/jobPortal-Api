package handlers

import (
	"project/internal/auth"
	"project/internal/middlewear"
	"project/internal/services"

	"github.com/gin-gonic/gin"
)

func Api(a *auth.Auth, s *services.Services) *gin.Engine {
	r := gin.New()
	h, _ := NewHandler(a, s)
	m, _ := middlewear.NewMiddleWear(a)
	r.Use(m.Log(), gin.Recovery())
	r.POST("/api/signup", h.userSignin)
	r.POST("/api/login", h.userLoginin)
	r.POST("/api/createCompany", h.companyCreation)
	r.GET("/api/getAllCompany", h.getAllCompany)
	r.GET("/api/getCompany/:company_id", h.getCompany)
	r.POST("/api/companies/:company_id/postjob", h.postJob)
	r.GET("/api/companies/:company_id/jobs", h.getJob)
	r.GET("/api/jobs", h.getAllJob)
	r.GET("/api/jobs/:id", h.GetJobById)
	r.POST("/api/applyjob/:id", h.applyJob)
	return r
}
