package services

import (
	db "DB_GORM/DB"
	"DB_GORM/models"
	pb "DB_GORM/pb_file"
	"DB_GORM/utils"
	"context"
)

type Recruiter struct {
	pb.UnimplementedRecruiterServiceServer
}

func (r *Recruiter) CreateRecruiter(ctx context.Context, req *pb.RecruiterRequest) (*pb.RecruiterResponse, error) {
	recruiter := models.Recruiter{
		Name:        req.Name,
		Contact:     req.Contact,
		CompanyName: req.CompanyName,
	}
	data := db.DB.Create(&recruiter)

	if data.Error != nil {
		utils.ErrorLog.Println("Error in Recruiter Creation", data.Error)
		return nil, data.Error
	}
	return &pb.RecruiterResponse{
		Message: "Recruiter Created Successfully!!",
	}, nil
}

func (r *Recruiter) GetRecruiter(ctx context.Context, req *pb.RecruiterID) (*pb.RecruiterGetResponse, error) {
	var recruiter models.Recruiter
	data := db.DB.First(&recruiter, req.RecruiterId)
	if data.Error != nil {
		utils.ErrorLog.Println("Recruiter Not Found", data.Error)
		return nil, data.Error
	}

	return &pb.RecruiterGetResponse{
		RecruiterId: int32(recruiter.RecruiterID),
		Name:        recruiter.Name,
		Contact:     recruiter.Contact,
		CompanyName: recruiter.CompanyName,
	}, nil
}

func (u *User) UpdateRecruiter(ctx context.Context, req *pb.UpdateRecruiterRequest) (*pb.RecruiterResponse, error) {
	var recruiter models.Recruiter
	data := db.DB.First(&recruiter, req.RecruiterId)
	if data.Error != nil {
		utils.ErrorLog.Println("Recruiter Not Found", data.Error)
		return nil, data.Error
	}
	recruiter.Name = req.Name
	recruiter.Contact = req.Contact
	recruiter.CompanyName = req.CompanyName

	data1 := db.DB.Save(&recruiter)

	if data1.Error != nil {
		utils.ErrorLog.Println("Recruiter Not able to Update..", data1.Error)
		return nil, data.Error
	}

	return &pb.RecruiterResponse{
		Message: "User Updated Successfully...",
	}, nil
}

func (r *Recruiter) DeleteRecruiter(ctx context.Context, req *pb.RecruiterID) (*pb.RecruiterResponse, error) {
	var recruiter models.Recruiter

	data := db.DB.First(&recruiter, req.RecruiterId)
	if data.Error != nil {
		utils.ErrorLog.Println("Recruiter Not Found", data.Error)
		return nil, data.Error
	}

	data1 := db.DB.Delete(&recruiter)
	if data1.Error != nil {
		utils.ErrorLog.Println("Recruiter Not able to delete..", data1.Error)
		return nil, data.Error
	}

	return &pb.RecruiterResponse{
		Message: "Recruiter Deleted Successfully...",
	}, nil
}

func (r *Recruiter) ListRecruiters(ctx context.Context, req *pb.RecruiterEmpty) (*pb.RecruiterListResponse, error) {
	var recruiters []models.Recruiter
	data := db.DB.Find(&recruiters)
	if data.Error != nil {
		utils.ErrorLog.Println("Recruiters Not Found", data.Error)
		return nil, data.Error
	}

	var recruiterResponses []*pb.RecruiterGetResponse

	for _, recruiter := range recruiters {
		recruiterResponses = append(recruiterResponses, &pb.RecruiterGetResponse{
			RecruiterId: int32(recruiter.RecruiterID),
			Name:        recruiter.Name,
			Contact:     recruiter.Contact,
			CompanyName: recruiter.CompanyName,
		})
	}

	return &pb.RecruiterListResponse{Recruiters: recruiterResponses}, nil
}
