package services

import (
	"context"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	pb "DB_GORM/pb_file"
	"DB_GORM/utils"
)

func init() {
	utils.ErrorLog = log.New(os.Stderr, "[TEST-ERROR] ", log.LstdFlags)
	utils.InfoLog = log.New(os.Stdout, "[TEST-INFO] ", log.LstdFlags)
}

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	dialector := mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	return gdb, mock, err
}

func TestCreateUser_Success(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users`")).
		WithArgs("Nandhini", "9321456789", "Go", 25, 2, "B.Tech").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	s := &User{DB: db}

	req := &pb.UserRequest{
		Name:            "Nandhini",
		Contact:         "9321456789",
		Skills:          "Go",
		Age:             25,
		ExperienceYears: 2,
		Education:       "B.Tech",
	}

	resp, err := s.CreateUser(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "User Created Successfully!!", resp.Message)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUser_Success(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)

	rows := sqlmock.NewRows([]string{"user_id", "name", "contact", "skills", "age", "experience_years", "education"}).
		AddRow(1, "Bala", "8768915468", "Python", 30, 5, "M.Tech")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`user_id` = ? ORDER BY `users`.`user_id` LIMIT ?")).
		WithArgs(1, 1).
		WillReturnRows(rows)

	s := &User{DB: db}

	req := &pb.UserID{Id: 1}

	resp, err := s.GetUser(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "Bala", resp.Name)
	assert.Equal(t, "8768915468", resp.Contact)
	assert.Equal(t, "Python", resp.Skills)
	assert.Equal(t, int32(30), resp.Age)
	assert.Equal(t, int32(5), resp.ExperienceYears)
	assert.Equal(t, "M.Tech", resp.Education)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser_Success(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)

	rows := sqlmock.NewRows([]string{"user_id", "name", "contact", "skills", "age", "experience_years", "education"}).
		AddRow(1, "Anu", "8765432197", "Golang", 35, 10, "M.E")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`user_id` = ? ORDER BY `users`.`user_id` LIMIT ?")).
		WithArgs(1, 1).
		WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `users` WHERE `users`.`user_id` = ?")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	s := &User{DB: db}

	resp, err := s.DeleteUser(context.Background(), &pb.UserID{Id: 1}) // <- Fix here
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestListUser_Success(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)

	rows := sqlmock.NewRows([]string{"user_id", "name", "contact", "skills", "age", "experience_years", "education"}).
		AddRow(1, "Anu", "9876543210", "Go", 25, 2, "B.Tech").
		AddRow(2, "Bala", "8897654328", "Python", 30, 5, "M.Tech")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(rows)

	s := &User{DB: db}

	resp, err := s.ListUser(context.Background(), &pb.Empty{})

	assert.NoError(t, err)
	assert.Len(t, resp.Users, 2)
	assert.Equal(t, "Anu", resp.Users[0].Name)
	assert.Equal(t, "Bala", resp.Users[1].Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

//failure cases

func TestCreateUser_Failure(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users`")).
		WithArgs("Nandhini", "9321456789", "Go", 25, 2, "B.Tech").
		WillReturnError(assert.AnError)
	mock.ExpectRollback()

	s := &User{DB: db}

	req := &pb.UserRequest{
		Name:            "Nandhini",
		Contact:         "9321456789",
		Skills:          "Go",
		Age:             25,
		ExperienceYears: 2,
		Education:       "B.Tech",
	}

	resp, err := s.CreateUser(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUser_Failure(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`user_id` = ? ORDER BY `users`.`user_id` LIMIT ?")).
		WithArgs(100, 1).
		WillReturnError(assert.AnError)

	s := &User{DB: db}

	req := &pb.UserID{Id: 100}

	resp, err := s.GetUser(context.Background(), req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser_Failure(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`user_id` = ? ORDER BY `users`.`user_id` LIMIT ?")).
		WithArgs(99, 1).
		WillReturnError(assert.AnError)

	s := &User{DB: db}

	resp, err := s.DeleteUser(context.Background(), &pb.UserID{Id: 99})

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestListUser_Failure(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnError(assert.AnError)

	s := &User{DB: db}

	resp, err := s.ListUser(context.Background(), &pb.Empty{})

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
