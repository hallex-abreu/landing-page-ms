package mysql

import (
	"log"

	"github.com/hallex-abreu/landing-page-ms/api/domain"
	dbconfig "github.com/hallex-abreu/landing-page-ms/api/infra/database/mysql/db-config"
)

type MysqlLeadRepository struct{}

func (m MysqlLeadRepository) FindAllLeads() ([]*domain.Lead, error) {
	var leads []*domain.Lead

	result := dbconfig.Mysql.Find(&leads)

	if result.Error != nil {
		log.Printf("Failed to find leads %v", result.Error)
		return nil, result.Error
	}

	return leads, nil
}

func (m MysqlLeadRepository) FindLeadByEmail(email string) *domain.Lead {
	var lead *domain.Lead

	result := dbconfig.Mysql.Where("email = ?", email).First(&lead)

	if result.Error != nil {
		log.Printf("Failed to find lead by email %v", result.Error)
		return nil
	}

	return lead
}

func (m MysqlLeadRepository) RegisterLead(lead domain.Lead) *domain.Lead {
	dbconfig.Mysql.Create(&lead)
	return &lead
}
