package services_test

import (
	"context"
	"testing"

	"DB_GORM/models"
	"DB_GORM/pb_file"
	"DB_GORM/services"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type testUserService struct {
	services.User
	DB *gorm.DB
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}
	db.AutoMigrate(&models.User{})
	return db
}

func getService(t *testing.T) *testUserService {
	db := setupTestDB(t)
	return &testUserService{
		User: services.User{},
		DB:   db,
	}
}

func TestCreateUser(t *testing.T) {
	svc := getService(t)

	req := &pb_file.UserRequest{
		Name:            "John",
		Contact:         "1234567890",
		Skills:          "Go, gRPC",
		Age:             30,
		ExperienceYears: 5,
		Education:       "B.Tech",
	}

	resp, err := svc.CreateUser(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "User Created Successfully!!", resp.Message)
}

func TestGetUser(t *testing.T) {
	svc := getService(t)

	user := models.User{
		Name:    "Alice",
		Contact: "111",
		Skills:  "Python",
		Age:     25,
	}
	svc.DB.Create(&user)

	resp, err := svc.GetUser(context.Background(), &pb_file.UserID{Id: int32(user.UserID)})
	assert.NoError(t, err)
	assert.Equal(t, "Alice", resp.Name)
}

func TestUpdateUser(t *testing.T) {
	svc := getService(t)

	user := models.User{
		Name:    "Old Name",
		Contact: "222",
		Skills:  "Java",
		Age:     28,
	}
	svc.DB.Create(&user)

	req := &pb_file.UpdateRequest{
		Id:              int32(user.UserID),
		Name:            "New Name",
		Contact:         "999",
		Skills:          "Java, Spring",
		Age:             29,
		ExperienceYears: 6,
		Education:       "MCA",
	}
	resp, err := svc.UpdateUser(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "User Updated Successfully...", resp.Message)
}

func TestDeleteUser(t *testing.T) {
	svc := getService(t)

	user := models.User{
		Name:    "ToDelete",
		Contact: "333",
	}
	svc.DB.Create(&user)

	resp, err := svc.DeleteUser(context.Background(), &pb_file.UserID{Id: int32(user.UserID)})
	assert.NoError(t, err)
	assert.Equal(t, "User Deleted Successfully...", resp.Message)
}

func TestListUser(t *testing.T) {
	svc := getService(t)

	svc.DB.Create(&models.User{Name: "User1", Contact: "123"})
	svc.DB.Create(&models.User{Name: "User2", Contact: "456"})

	resp, err := svc.ListUser(context.Background(), &pb_file.Empty{})
	assert.NoError(t, err)
	assert.Len(t, resp.Users, 2)
}
