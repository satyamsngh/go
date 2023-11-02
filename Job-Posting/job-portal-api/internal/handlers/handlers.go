package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/repository"
	"job-portal-api/internal/services"

	"time"

	"net/http"
)

// Define a function called API that takes an argument a of type *auth.Auth
// and returns a pointer to a gin.Engine

func API(a *auth.Auth, c *repository.Repo) *gin.Engine {

	// Create a new Gin engine; Gin is a HTTP web framework written in Go
	r := gin.New()

	m, err := middlewares.NewMid(a)
	ms, err := services.NewStore(c)

	h := handler{
		s: ms,
		a: a,
	}

	// If there is an error in setting up the middleware, panic and stop the application
	// then log the error message
	if err != nil {
		log.Panic().Msg("middlewares not set up")
	}

	// Attach middleware's Log function and Gin's Recovery middleware to our application
	// The Recovery middleware recovers from any panics and writes a 500 HTTP response if there was one.
	r.Use(m.Log(), gin.Recovery())

	// Define a route at path "/check"
	// If it receives a GET request, it will use the m.Authenticate(check) function.
	r.GET("api/check", m.Authenticate(check))
	r.POST("api/register", h.Register)
	r.POST("api/login", h.Login)
	r.POST("/api/companies", m.Authenticate(h.AddCompanies))
	r.GET("/api/view", m.Authenticate(h.ViewCompanies))
	r.GET("/api/companies/:companyID", m.Authenticate(h.ViewCompaniesById))
	r.POST("/companies/:companyID/jobs", m.Authenticate(h.CreateJob))
	r.GET("api/companies/:companyID/list-jobs", m.Authenticate(h.ListJobs))
	r.GET("api/jobs", m.Authenticate(h.AllJobs))
	r.GET("/api/jobs/:jobID", m.Authenticate(h.JobsByID))

	return r
}

func check(c *gin.Context) {

	time.Sleep(time.Second * 3)
	select {
	case <-c.Request.Context().Done():
		fmt.Println("user not there")
		return
	default:
		c.JSON(http.StatusOK, gin.H{"msg": "statusOk"})

	}

}
