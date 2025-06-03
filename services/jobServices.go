package services

import (
	db "DB_GORM/DB"
	"DB_GORM/models"
	pb "DB_GORM/pb_file"
	"DB_GORM/utils"
	"context"
	"fmt"
)

type Job struct {
	pb.UnimplementedJobServiceServer
}

func (j *Job) CreateJob(ctx context.Context, req *pb.JobRequest) (*pb.JobResponse, error) {
	var recruiter models.Recruiter
	if err := db.DB.First(&recruiter, req.RecruiterId).Error; err != nil {
		utils.ErrorLog.Println("Recruiter Not Found", err)
		return nil, fmt.Errorf("recruiter with ID %d not found", req.RecruiterId)
	}
	job := models.Job{
		RecruiterID:    uint(req.RecruiterId),
		Title:          req.Title,
		Description:    req.Description,
		SkillsRequired: req.SkillsRequired,
		Location:       req.Location,
		Salary:         req.Salary,
		JobType:        req.JobType,
	}
	data := db.DB.Create(&job)

	if data.Error != nil {
		utils.ErrorLog.Println("Error in Job Creation", data.Error)
		return nil, data.Error
	}
	return &pb.JobResponse{
		Message: "Job Created Successfully!!",
	}, nil
}

func (j *Job) CreateJobs(ctx context.Context, req *pb.BatchJobRequest) (*pb.BatchJobResponse, error) {
	var jobs []models.Job
	var errors []string

	for _, jobReq := range req.Jobs {
		var recruiter models.Recruiter
		if err := db.DB.First(&recruiter, jobReq.RecruiterId).Error; err != nil {
			errMsg := fmt.Sprintf("Recruiter ID %d not found", jobReq.RecruiterId)
			utils.ErrorLog.Println(errMsg)
			errors = append(errors, errMsg)
			continue
		}

		jobs = append(jobs, models.Job{
			RecruiterID:    uint(jobReq.RecruiterId),
			Title:          jobReq.Title,
			Description:    jobReq.Description,
			SkillsRequired: jobReq.SkillsRequired,
			Location:       jobReq.Location,
			Salary:         jobReq.Salary,
			JobType:        jobReq.JobType,
		})
	}

	if err := db.DB.Create(&jobs).Error; err != nil {
		utils.ErrorLog.Println("Error in Batch Job Creation", err)
		return nil, err
	}

	return &pb.BatchJobResponse{
		Message: "Batch Job Processing Completed",
		Errors:  errors,
	}, nil
}

func (j *Job) GetJob(ctx context.Context, req *pb.JobID) (*pb.JobRequest, error) {
	var job models.Job
	data := db.DB.First(&job, req.Id)
	if data.Error != nil {
		utils.ErrorLog.Println("Job Not Found", data.Error)
		return nil, data.Error
	}

	return &pb.JobRequest{
		Title:          job.Title,
		Description:    job.Description,
		SkillsRequired: job.SkillsRequired,
		Location:       job.Location,
		Salary:         job.Salary,
		JobType:        job.JobType,
	}, nil
}

func (j *Job) DeleteJob(ctx context.Context, req *pb.JobID) (*pb.JobResponse, error) {
	var job models.Job

	data := db.DB.First(&job, req.Id)
	if data.Error != nil {
		utils.ErrorLog.Println("Job Not Found", data.Error)
		return nil, data.Error
	}

	data1 := db.DB.Delete(&job)
	if data1.Error != nil {
		utils.ErrorLog.Println("Job Not able to Delete", data1.Error)
		return nil, data.Error
	}

	return &pb.JobResponse{
		Message: "Job Deleted Successfully...",
	}, nil
}
