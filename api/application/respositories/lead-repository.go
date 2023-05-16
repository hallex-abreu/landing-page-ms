package respositories

import "github.com/hallex-abreu/landing-page-ms/api/domain"

type LeadRepository interface {
	FindAllLeads() ([]*domain.Lead, error)
	FindLeadByEmail(email string) *domain.Lead
	RegisterLead(lead domain.Lead) *domain.Lead
}
