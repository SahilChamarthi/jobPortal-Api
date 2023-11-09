package services

import (
	"errors"
	"project/internal/model"
	"project/internal/repository"
	"reflect"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

func TestServices_JobCreate(t *testing.T) {
	type args struct {
		nj model.CreateJob
		id uint64
	}
	tests := []struct {
		name             string
		args             args
		want             model.Job
		wantErr          bool
		mockRepoResponse func() (model.Job, error)
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{nj: model.CreateJob{
				JobTitle:  "gwyug",
				JobSalary: 514212,
			},
				id: 15,
			},
			want: model.Job{
				JobTitle:  "gwyug",
				JobSalary: 514212,

				CompanyID: 15,
			},
			wantErr: false,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{
					JobTitle:  "gwyug",
					JobSalary: 514212,

					CompanyID: 15,
				}, nil

			},
		},

		{
			name: "failure",
			args: args{
				nj: model.CreateJob{},
				id: 15,
			},
			want:    model.Job{},
			wantErr: true,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{}, errors.New("failed to crerate job")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateJob(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo)
			//got, err := tt.s.CompanyCreate(tt.args.nc)
			got, err := s.JobCreate(tt.args.nj, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Services.JobCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.JobCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_GetJobsByCompanyId(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name             string
		args             args
		want             []model.Job
		wantErr          bool
		mockRepoResponse func() ([]model.Job, error)
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				id: 12,
			},
			want: []model.Job{
				{
					JobTitle:  "hcwc",
					JobSalary: 25778,
					CompanyID: 12,
				},
				{
					JobTitle:  "sajchbuc",
					JobSalary: 25778,
					CompanyID: 13,
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]model.Job, error) {
				return []model.Job{
					{
						JobTitle:  "hcwc",
						JobSalary: 25778,
						CompanyID: 12,
					},
					{
						JobTitle:  "ygyudsgf",
						JobSalary: 25778,
						CompanyID: 13,
					},
				}, nil
			},
		},
		{
			name: "failure",
			args: args{
				id: 12,
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]model.Job, error) {
				return nil, errors.New("job not found")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().GetJobs(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo)
			//got, err := tt.s.CompanyCreate(tt.args.nc)
			got, err := s.GetJobsByCompanyId(tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("Services.GetJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.GetJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_FetchAllJobs(t *testing.T) {
	tests := []struct {
		name             string
		want             []model.Job
		wantErr          bool
		mockRepoResponse func() ([]model.Job, error)
	}{
		// TODO: Add test cases.
		{
			name: "sucesss",
			want: []model.Job{
				{
					JobTitle:  "hcwc",
					JobSalary: 25778,
					CompanyID: 12,
				},
				{
					JobTitle:  "ygyudsgf",
					JobSalary: 25778,
					CompanyID: 13,
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]model.Job, error) {
				return []model.Job{
					{
						JobTitle:  "hcwc",
						JobSalary: 25778,
						CompanyID: 12,
					},
					{
						JobTitle:  "ygyudsgf",
						JobSalary: 25778,
						CompanyID: 13,
					},
				}, nil
			},
		},
		{
			name:    "failure",
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]model.Job, error) {
				return nil, errors.New("fail to found jobs")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().GetAllJobs().Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo)
			got, err := s.FetchAllJobs()
			if (err != nil) != tt.wantErr {
				t.Errorf("Services.FetchAllJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.FetchAllJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_Getjobid(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name             string
		args             args
		want             model.Job
		wantErr          bool
		mockRepoResponse func() (model.Job, error)
	}{
		// TODO: Add test cases.

		{
			name: "success",
			args: args{
				id: 15,
			},
			want: model.Job{

				JobTitle:  "eyhe",
				JobSalary: 162522,
				CompanyID: 17,
			},
			wantErr: false,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{

					JobTitle:  "eyhe",
					JobSalary: 162522,
					CompanyID: 17,
				}, nil
			},
		},
		{
			name: "failure",
			args: args{
				id: 12,
			},
			want:    model.Job{},
			wantErr: true,
			mockRepoResponse: func() (model.Job, error) {
				return model.Job{}, errors.New("job not found")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().GetJobId(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo)
			got, err := s.Getjobid(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Services.Getjobid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.Getjobid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_ApplyJob_Service(t *testing.T) {
	type args struct {
		ja model.JobApplication
		id uint64
	}
	tests := []struct {
		name             string
		args             args
		want             model.ApprovedApplication
		wantErr          bool
		mockRepoResponse func() (model.ApprovedApplication, error)
	}{

		{
			name: "successs",
			args: args{
				ja: model.JobApplication{
					Name:           "bumesh",
					Gmail:          "bumesh@gmail.com",
					Age:            23,
					Phone:          9018373973,
					JobTitle:       "software testing",
					ExpectedSalary: 26000,
					NoticePeriod:   40,
					Experience:     3,

					Qualifications:   []uint{1, 2},
					Shift:            []uint{1, 2, 3},
					JobType:          []uint{2, 5, 1},
					JobLocations:     []uint{1, 2, 4},
					Technology_stack: []uint{2, 4},
					WorkMode:         []uint{1, 4},
				}, id: 2,
			},

			want: model.ApprovedApplication{
				Name:  "bumesh",
				Gmail: "bumesh@gmail.com",
				Phone: 9018373973,
			},
			wantErr: false,
			mockRepoResponse: func() (model.ApprovedApplication, error) {
				return model.ApprovedApplication{
					Name:  "bumesh",
					Gmail: "bumesh@gmail.com",
					Phone: 9018373973,
				}, nil
			},
		},

		// {
		// 	name:    "failure",
		// 	args:    args{},
		// 	want:    model.ApprovedApplication{},
		// 	wantErr: true,
		// 	mockRepoResponse: func() (model.ApprovedApplication, error) {
		// 		return model.ApprovedApplication{}, errors.New("job cannot be applied did not meet criteria")
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)

			mockRepo := repository.NewMockAllInRepo(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ApplyJob_Repository(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s, _ := NewServices(mockRepo)
			got, err := s.ApplyJob_Service(tt.args.ja, tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("Services.ApplyJob_Service() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Services.ApplyJob_Service() = %v, want %v", got, tt.want)
			}
		})
	}
}
