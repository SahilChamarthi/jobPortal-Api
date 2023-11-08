package model

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	CompanyName string `json:"company_name" validate:"required"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

type CreateCompany struct {
	CompanyName string `json:"company_name" validate:"required"`
	Adress      string `json:"company_adress" validate:"required"`
	Domain      string `json:"domain" validate:"required"`
}

// type Job struct {
// 	gorm.Model
// 	JobTitle  string  `json:"job_title" validate:"required"`
// 	JobSalary string  `json:"job_salary" validate:"required"`
// 	Company   Company `gorm:"ForeignKey:uid"`
// 	Uid       uint64  `JSON:"uid, omitempty"`
// }

// type CreateJob struct {
// 	JobTitle  string `json:"job_title" validate:"required"`
// 	JobSalary string `json:"job_salary" validate:"required"`
// }

type Job struct {
	gorm.Model
	ID                 uint `gorm:"primaryKey;autoIncrement"`
	JobTitle           string
	JobSalary          uint
	Description        string
	CompanyID          uint64
	Min_NoticePeriod   uint
	Max_NoticePeriod   uint
	Budget             uint
	Minimum_Experience uint
	Maximum_Experience uint

	Qualifications []Qualification `gorm:"many2many:qualification_jobs"`
	Shifts         []Shift         `gorm:"many2many:shift_jobs"`
	JobTypes       []JobType       `gorm:"many2many:jobtype_jobs"`

	JobLocations    []JobLocation `gorm:"many2many:location_jobs"`
	TechnologyStack []Technology  `gorm:"many2many:technology_jobs"`
	WorkModes       []WorkMode    `gorm:"many2many:workmode_jobs"`
}

type CreateJob struct {
	JobTitle           string `json:"title" validate:"required"`
	JobSalary          uint   `json:"job_salary" validate:"required"`
	Description        string `json:"description" validate:"required"`
	CompanyID          uint64
	Min_NoticePeriod   uint `json:"min_np" validate:"required"`
	Max_NoticePeriod   uint `json:"max_np" validate:"required"`
	Budget             uint `json:"budget" validate:"required"`
	Minimum_Experience uint `json:"min_exp" validate:"required"`
	Maximum_Experience uint `json:"max_exp" validate:"required"`

	Qualifications []uint `json:"qualifications" validate:"required"`
	Shift          []uint `json:"shifts" validate:"required"`
	Job_Type       []uint `json:"job_type" validate:"required"`

	JobLocations    []uint `json:"job_locations" validate:"required"`
	TechnologyStack []uint `json:"technology_stack" validate:"required"`
	WorkMode        []uint `json:"workmode" validate:"required"`
}

type Qualification struct {
	gorm.Model
	ID   uint
	Name string
}

type Shift struct {
	gorm.Model
	ID   uint
	Name string
}

type JobType struct {
	gorm.Model
	ID   uint
	Name string
}

type JobLocation struct {
	gorm.Model
	ID   uint
	Name string
}

type Technology struct {
	gorm.Model
	ID   uint
	Name string
}

type WorkMode struct {
	gorm.Model
	ID   uint
	Name string
}

type JobApplication struct {
	Name               string `json:"name" validate:"required"`
	Gmail              string `json:"gmail" validate:"required"`
	Age                uint   `json:"age" validate:"required"`
	JobTitle           string `json:"title" validate:"required"`
	JobSalary          string `json:"job_salary" validate:"required"`
	Min_NoticePeriod   string `json:"min_np" validate:"required"`
	Max_NoticePeriod   string `json:"max_np" validate:"required"`
	Budget             string `json:"budget" validate:"required"`
	JobLocations       []uint `json:"job_locations" validate:"required"`
	Technology_stack   []uint `json:"technology_stack" validate:"required"`
	WorkMode           []uint `json:"workmode" validate:"required"`
	Description        string `json:"description" validate:"required"`
	Minimum_Experience string `json:"min_exp" validate:"required"`
	Maximum_Experience string `json:"max_exp" validate:"required"`
	Qualifications     []uint `json:"qualifications" validate:"required"`
	Shift              []uint `json:"shifts" validate:"required"`
	Job_Type           []uint `json:"job_type" validate:"required"`
}
