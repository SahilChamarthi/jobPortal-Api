package services

import (
	"errors"
	"project/internal/model"

	"github.com/rs/zerolog/log"
)

func (s *Services) JobCreate(nj model.CreateJob, id uint64) (model.Job, error) {
	job := model.Job{

		JobTitle:           nj.JobTitle,
		JobSalary:          nj.JobSalary,
		CompanyID:          id,
		Min_NoticePeriod:   nj.Min_NoticePeriod,
		Max_NoticePeriod:   nj.Max_NoticePeriod,
		Budget:             nj.Budget,
		Description:        nj.Description,
		Minimum_Experience: nj.Minimum_Experience,
		Maximum_Experience: nj.Maximum_Experience,
	}

	// Map foreign key IDs to related entities in the database
	for _, locationID := range nj.JobLocations {
		job.JobLocations = append(job.JobLocations, model.JobLocation{ID: locationID})
	}

	for _, techID := range nj.TechnologyStack {
		job.TechnologyStack = append(job.TechnologyStack, model.Technology{ID: techID})
	}

	for _, modeID := range nj.WorkMode {
		job.WorkModes = append(job.WorkModes, model.WorkMode{ID: modeID})
	}

	for _, qualificationID := range nj.Qualifications {
		job.Qualifications = append(job.Qualifications, model.Qualification{ID: qualificationID})
	}

	for _, shiftID := range nj.Shift {
		job.Shifts = append(job.Shifts, model.Shift{ID: shiftID})
	}

	for _, jobTypeID := range nj.Job_Type {
		job.JobTypes = append(job.JobTypes, model.JobType{ID: jobTypeID})
	}

	cu, err := s.r.CreateJob(job)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create job")
		return model.Job{}, errors.New("job creation failed")
	}

	return cu, nil
}

func (s *Services) GetJobsByCompanyId(id int) ([]model.Job, error) {
	AllCompanies, err := s.r.GetJobs(id)
	if err != nil {
		return nil, errors.New("job retreval failed")
	}
	return AllCompanies, nil
}

func (s *Services) FetchAllJobs() ([]model.Job, error) {

	AllJobs, err := s.r.GetAllJobs()
	if err != nil {
		return nil, err
	}
	return AllJobs, nil

}

func (s *Services) Getjobid(id uint64) (model.Job, error) {

	jobData, err := s.r.GetJobId(id)
	if err != nil {
		return model.Job{}, err
	}
	return jobData, nil
}

func (s *Services) ApplyJob_Service(cj model.CreateJob, id uint64) (model.Job, bool, error) {

	mj, ok, err := s.r.ApplyJob_Repository(cj, id)

	if err != nil || !ok {
		return model.Job{}, false, errors.New("not comes under job criteria")
	}

	return mj, true, nil

}
