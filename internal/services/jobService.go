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

func (s *Services) ApplyJob_Service(ja model.JobApplication, id uint64) (model.ApprovedApplication, error) {

	job, err := s.r.ApplyJob_Repository(id)

	if err != nil {
		return model.ApprovedApplication{}, err
	}

	var count int

	user := model.ApprovedApplication{
		Name:  ja.Name,
		Gmail: ja.Gmail,
		Phone: ja.Phone,
	}

	err = errors.New("")
	if ja.ExpectedSalary <= job.Budget {
		log.Info().Str("Budget", "true").Send()
		count++
	} else {
		log.Info().Str("Budget", "false").Send()
		return model.ApprovedApplication{}, err
	}

	if ja.NoticePeriod >= job.Min_NoticePeriod && ja.NoticePeriod <= job.Max_NoticePeriod {
		log.Info().Str("Min_NP", "true").Send()
		count++
	} else {
		log.Info().Str("Min_NP", "false").Send()
		return model.ApprovedApplication{}, err
	}

	if ja.Experience >= job.Minimum_Experience && ja.Experience <= job.Maximum_Experience {
		log.Info().Str("MinExp", "true").Send()
		count++
	} else {
		log.Info().Str("MinExp", "false").Send()
		return model.ApprovedApplication{}, err
	}

	//comparing job criteria locations and application criteria locations
	var loc_job []uint
	var loc_app []uint
	for _, v := range job.JobLocations {
		loc_job = append(loc_job, v.ID)
	}

	loc_app = ja.JobLocations

	if sliceContainsAtLeastOne(loc_job, loc_app) {
		log.Info().Str("JobLocations", "true").Send()
		count++
	} else {
		log.Info().Str("JobLocations", "false").Send()
	}

	//comparing job criteria technologystack and application criteria technologystack
	var tech_job []uint
	var tech_app []uint
	for _, v := range job.TechnologyStack {
		tech_job = append(tech_job, v.ID)
	}

	tech_app = ja.Technology_stack
	if sliceContainsAtLeastOne(tech_job, tech_app) {
		log.Info().Str("TechnologyStack", "true").Send()
		count++
	} else {
		log.Info().Str("TechnologyStack", "false").Send()
	}

	//comparing job criteria technologystack and application criteria technologystack
	var mode_job []uint
	var mode_app []uint
	for _, v := range job.WorkModes {
		mode_job = append(mode_job, v.ID)
	}
	mode_app = ja.WorkMode
	if sliceContainsAtLeastOne(mode_job, mode_app) {
		log.Info().Str("WorkModes", "true").Send()
		count++
	} else {
		log.Info().Str("WorkModes", "false").Send()
	}

	//comparing job criteria qualification and application criteria qualification
	var q_job []uint
	var q_app []uint
	for _, v := range job.Qualifications {
		q_job = append(q_job, v.ID)
	}
	q_app = ja.Qualifications
	if sliceContainsAtLeastOne(q_job, q_app) {
		log.Info().Str("Qualificvations", "true").Send()
		count++
	} else {
		log.Info().Str("Qualifications", "false").Send()
	}

	//comparing job criteria shifts and application criteria shifts
	var shift_job []uint
	var shift_app []uint
	for _, v := range job.Shifts {
		shift_job = append(shift_job, v.ID)
	}
	shift_app = ja.Shift
	if sliceContainsAtLeastOne(shift_job, shift_app) {
		log.Info().Str("Shifts", "true").Send()
		count++
	} else {
		log.Info().Str("Shifts", "false").Send()
	}

	//comparing job criteria technologystack and application criteria technologystack
	var type_job []uint
	var type_app []uint
	for _, v := range job.JobTypes {
		type_job = append(type_job, v.ID)
	}
	type_app = ja.JobType
	if sliceContainsAtLeastOne(type_job, type_app) {
		log.Info().Str("JobTypes", "true").Send()
		count++
	} else {
		log.Info().Str("JobTypes", "false").Send()
	}

	if count >= 4 {
		return user, nil
	}

	return model.ApprovedApplication{}, err

}

// function to check the slices
func sliceContainsAtLeastOne(slice, subSlice []uint) bool {

	for i := 0; i < len(slice); i++ {

		for j := 0; j < len(subSlice); j++ {

			if slice[i] == subSlice[j] {
				return true
			}
		}

	}

	return false
}
