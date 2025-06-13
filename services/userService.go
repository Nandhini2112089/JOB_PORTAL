package services

import (
	"DB_GORM/models"
	pb "DB_GORM/pb_file"
	"DB_GORM/utils"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type User struct {
	pb.UnimplementedUserserviceServer
	DB *gorm.DB
}

func (u *User) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user := models.User{
		Name:            req.Name,
		Contact:         req.Contact,
		Skills:          req.Skills,
		Age:             int(req.Age),
		ExperienceYears: int(req.ExperienceYears),
		Education:       req.Education,
	}
	result := u.DB.Create(&user)
	if result.Error != nil {
		utils.ErrorLog.Println("Error in User Creation:", result.Error)
		return nil, result.Error
	}
	return &pb.UserResponse{Message: "User Created Successfully!!"}, nil
}

func (u *User) GetUser(ctx context.Context, req *pb.UserID) (*pb.GetResponse, error) {
	var user models.User
	result := u.DB.First(&user, req.Id)
	if result.Error != nil {
		utils.ErrorLog.Println("User Not Found:", result.Error)
		return nil, result.Error
	}
	return &pb.GetResponse{
		Name:            user.Name,
		Contact:         user.Contact,
		Skills:          user.Skills,
		Age:             int32(user.Age),
		ExperienceYears: int32(user.ExperienceYears),
		Education:       user.Education,
	}, nil
}

func (u *User) UpdateUser(ctx context.Context, req *pb.UpdateRequest) (*pb.UserResponse, error) {
	var user models.User
	result := u.DB.First(&user, req.Id)
	if result.Error != nil {
		utils.ErrorLog.Println("User Not Found:", result.Error)
		return nil, result.Error
	}

	user.Name = req.Name
	user.Contact = req.Contact
	user.Skills = req.Skills
	user.Age = int(req.Age)
	user.ExperienceYears = int(req.ExperienceYears)
	user.Education = req.Education

	saveResult := u.DB.Save(&user)
	if saveResult.Error != nil {
		utils.ErrorLog.Println("User Not Updated:", saveResult.Error)
		return nil, saveResult.Error
	}

	return &pb.UserResponse{Message: "User Updated Successfully!!"}, nil
}

func (u *User) DeleteUser(ctx context.Context, req *pb.UserID) (*pb.UserResponse, error) {
	var user models.User
	if err := u.DB.First(&user, req.Id).Error; err != nil {
		log.Println("Error finding user:", err)
		return nil, status.Errorf(codes.NotFound, "User Not Found!!")
	}

	if err := u.DB.Delete(&user).Error; err != nil {

		log.Println("Error deleting user:", err)
		return nil, status.Errorf(codes.Internal, "Failed to delete user")
	}

	log.Println("Deleted user:", user)
	return &pb.UserResponse{Message: "User Deleted Successfully!!"}, nil
}

func (u *User) ListUser(ctx context.Context, req *pb.Empty) (*pb.ListResponse, error) {
	var users []models.User
	result := u.DB.Find(&users)
	if result.Error != nil {
		utils.ErrorLog.Println("Users Not Found:", result.Error)
		return nil, result.Error
	}

	var userResponses []*pb.GetResponse
	for _, user := range users {
		userResponses = append(userResponses, &pb.GetResponse{
			Name:            user.Name,
			Contact:         user.Contact,
			Skills:          user.Skills,
			Age:             int32(user.Age),
			ExperienceYears: int32(user.ExperienceYears),
			Education:       user.Education,
		})
	}

	return &pb.ListResponse{Users: userResponses}, nil
}
