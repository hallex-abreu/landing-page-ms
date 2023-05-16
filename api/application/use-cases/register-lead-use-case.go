package usecases

import (
	"errors"

	"github.com/hallex-abreu/landing-page-ms/api/application/respositories"
	"github.com/hallex-abreu/landing-page-ms/api/domain"
)

type RegisterLeadRequest struct {
	Name  string
	Email string
	Phone string
}

func RegisterLeadUseCase(registerLeadRequest RegisterLeadRequest, leadRepository respositories.LeadRepository) (*domain.Lead, error) {
	exist_lead_with_email := leadRepository.FindLeadByEmail(registerLeadRequest.Email)

	if exist_lead_with_email != nil {
		return nil, errors.New("Já existe um lead com esse email")
	}

	lead := &domain.Lead{
		Id:    0,
		Name:  registerLeadRequest.Name,
		Email: registerLeadRequest.Email,
		Phone: registerLeadRequest.Phone,
	}

	lead = leadRepository.RegisterLead(*lead)

	return lead, nil
}
