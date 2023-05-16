package usecases

import (
	"errors"

	"github.com/hallex-abreu/landing-page-ms/api/application/respositories"
	"github.com/hallex-abreu/landing-page-ms/api/domain"
)

func FindAllLeadsUseCase(leadRepository respositories.LeadRepository) ([]*domain.Lead, error) {
	leads, err := leadRepository.FindAllLeads()

	if err != nil {
		return nil, errors.New("Erro ao tentar listar leads")
	}

	return leads, nil
}
