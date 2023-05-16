package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/hallex-abreu/landing-page-ms/api/application/use-cases"
	"github.com/hallex-abreu/landing-page-ms/api/infra/database/mysql"
	"github.com/hallex-abreu/landing-page-ms/api/infra/http/controllers/lead-controller/dtos"
)

func FindAllLeads(c *gin.Context) {
	leadRepository := mysql.MysqlLeadRepository{}

	leads, err := usecases.FindAllLeadsUseCase(leadRepository)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, leads)
	}
}

func RegisterLead(c *gin.Context) {
	leadRepository := mysql.MysqlLeadRepository{}

	var body dtos.LeadDto

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lead := usecases.RegisterLeadRequest{
		Name:  body.Name,
		Email: body.Email,
		Phone: body.Phone,
	}

	_, err = usecases.RegisterLeadUseCase(lead, leadRepository)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, lead)
	}
}
