package services

import (
	db "DB_GORM/DB"
	"DB_GORM/models"
	pb "DB_GORM/pb_file"
	"DB_GORM/utils"
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Application struct {
	pb.UnimplementedApplicationServiceServer
}

func (a *Application) CreateApplication(ctx context.Context, req *pb.ApplicationRequest) (*pb.ApplicationResponse, error) {
	var user models.User
	if err := db.DB.First(&user, req.UserId).Error; err != nil {
		utils.ErrorLog.Println("User Not Found", err)
		return nil, fmt.Errorf("User with ID %d not found", req.UserId)
	}
	var job models.Job
	if err := db.DB.First(&job, req.JobId).Error; err != nil {
		utils.ErrorLog.Println("Job Not Found", err)
		return nil, fmt.Errorf("Job with ID %d not found", req.JobId)
	}
	var count int64
	db.DB.Model(&models.Application{}).Where("user_id = ? AND job_id = ?", req.UserId, req.JobId).Count(&count)
	if count > 0 {
		utils.ErrorLog.Println("User has already applied for this job")
		return nil, fmt.Errorf("User has already applied for this job")
	}

	userSkills := strings.Split(strings.ToLower(user.Skills), ",")
	jobSkills := strings.Split(strings.ToLower(job.SkillsRequired), ",")

	skillSet := make(map[string]bool)
	var status string = "Rejected"

	for _, skill := range userSkills {
		skillSet[skill] = true
	}

	for _, skill := range jobSkills {
		if skillSet[skill] {
			status = "Shortlisted"
			break
		}
	}

	application := models.Application{
		UserID:    uint(req.UserId),
		JobID:     uint(req.JobId),
		Status:    status,
		AppliedAt: time.Now(),
	}
	data := db.DB.Create(&application)

	if data.Error != nil {
		utils.ErrorLog.Println("Error in Application Creation", data.Error)
		return nil, data.Error
	}

	if status == "Shortlisted" {
		var recruiter models.Recruiter
		if err := db.DB.First(&recruiter, job.RecruiterID).Error; err != nil {
			utils.ErrorLog.Println("Error fetching recruiter details:", err)
			return nil, err
		}
		interview := models.Interview{
			ApplicationID: application.ApplicationID,
			UserID:        application.UserID,
			JobID:         application.JobID,
			CompanyName:   recruiter.CompanyName,
			RecruiterID:   job.RecruiterID,
			RecruiterName: recruiter.Name,
			InterviewTime: time.Now().Add(48 * time.Hour),
		}

		if err := db.DB.Create(&interview).Error; err != nil {
			utils.ErrorLog.Println("Error in Interview Scheduling", err)
			return nil, err
		}
	}
	return &pb.ApplicationResponse{
		Message: "Application Created Successfully!!",
	}, nil
}

func (a *Application) GetApplication(ctx context.Context, req *pb.ApplicationID) (*pb.ApplicationGetResponse, error) {
	var application models.Application
	data := db.DB.First(&application, req.Id)
	if data.Error != nil {
		utils.ErrorLog.Println("Application Not Found", data.Error)
		return nil, data.Error
	}

	return &pb.ApplicationGetResponse{
		UserId:    int32(application.UserID),
		JobId:     int32(application.JobID),
		Status:    application.Status,
		AppliedAt: timestamppb.New(application.AppliedAt),
	}, nil
}
