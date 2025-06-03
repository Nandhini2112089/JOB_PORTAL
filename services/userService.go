package services

import (
	db "DB_GORM/DB"
	"DB_GORM/models"
	pb "DB_GORM/pb_file"
	"DB_GORM/utils"
	"context"
)

type User struct {
	pb.UnimplementedUserserviceServer
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
	data := db.DB.Create(&user)

	if data.Error != nil {
		utils.ErrorLog.Println("Error in User Creation", data.Error)
		return nil, data.Error
	}
	return &pb.UserResponse{
		Message: "User Created Successfully!!",
	}, nil
}

func (u *User) GetUser(ctx context.Context, req *pb.UserID) (*pb.GetResponse, error) {
	var user models.User
	data := db.DB.First(&user, req.Id)
	if data.Error != nil {
		utils.ErrorLog.Println("User Not Found", data.Error)
		return nil, data.Error
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
	data := db.DB.First(&user, req.Id)
	if data.Error != nil {
		utils.ErrorLog.Println("User Not Found", data.Error)
		return nil, data.Error
	}
	user.Name = req.Name
	user.Contact = req.Contact
	user.Age = int(req.Age)
	user.Skills = req.Skills
	user.ExperienceYears = int(req.ExperienceYears)
	user.Education = req.Education

	data1 := db.DB.Save(&user)

	if data1.Error != nil {
		utils.ErrorLog.Println("User Not able to Update..", data1.Error)
		return nil, data.Error
	}

	return &pb.UserResponse{
		Message: "User Updated Successfully...",
	}, nil
}

func (u *User) DeleteUser(ctx context.Context, req *pb.UserID) (*pb.UserResponse, error) {
	var user models.User

	data := db.DB.First(&user, req.Id)
	if data.Error != nil {
		utils.ErrorLog.Println("User Not Found", data.Error)
		return nil, data.Error
	}

	data1 := db.DB.Delete(&user)
	if data1.Error != nil {
		utils.ErrorLog.Println("User Not able to Update..", data1.Error)
		return nil, data.Error
	}

	return &pb.UserResponse{
		Message: "User Deleted Successfully...",
	}, nil
}

func (u *User) ListUser(ctx context.Context, req *pb.Empty) (*pb.ListResponse, error) {
	var users []models.User
	data := db.DB.Find(&users)
	if data.Error != nil {
		utils.ErrorLog.Println("Users Not Found", data.Error)
		return nil, data.Error
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
