package services

import (
	"errors"
	"project/internal/model"

	"github.com/rs/zerolog/log"
)

// type CompanyService interface {
// 	CompanyCreate(nc model.CreateCompany) (model.Company, error)
// 	GetAllCompanies() ([]model.Company, error)
// 	GetCompany(id int) (model.Company, error)
// 	JobCreate(nj model.CreateJob, id uint64) (model.Job, error)
// 	GetJobs(id int) ([]model.Job, error)
// 	GetAllJobs() ([]model.Job, error)
// }

func (s *Services) CompanyCreate(nc model.CreateCompany) (model.Company, error) {
	company := model.Company{CompanyName: nc.CompanyName, Adress: nc.Adress, Domain: nc.Domain}
	cu, err := s.r.CreateCompany(company)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create company")
		return model.Company{}, errors.New("company creation failed")
	}

	return cu, nil
}

func (s *Services) GetAllCompanies() ([]model.Company, error) {

	AllCompanies, err := s.r.GetAllCompany()
	if err != nil {
		log.Error().Err(err).Msg("couldnot get companies")
		return nil, err
	}
	return AllCompanies, nil

}

func (s *Services) GetCompanyById(id int) (model.Company, error) {

	AllCompanies, err := s.r.GetCompany(id)
	if err != nil {
		log.Error().Err(err).Msg("couldnot get company")
		return model.Company{}, err
	}
	return AllCompanies, nil

}
