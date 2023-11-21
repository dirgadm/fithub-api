package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/dirgadm/fithub-api/internal/common"
	"github.com/dirgadm/fithub-api/internal/model"
	"github.com/dirgadm/fithub-api/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetDetail(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Close()

	repo := repository.NewUser(
		repository.ROption{
			Common: common.Options{
				Database: common.Database{
					Read:  db,
					Write: db,
				},
			},
		},
	)

	user := model.User{
		Id:        1,
		Email:     "test@example.com",
		Password:  "hashed_password",
		Name:      "Test User",
		Phone:     "1234567890",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = db.Exec("INSERT INTO user (id, email, password, name, phone, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		user.Id, user.Email, user.Password, user.Name, user.Phone, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		t.Fatalf("Error inserting sample user: %v", err)
	}

	ctx := context.Background()
	resultUser, err := repo.GetDetail(ctx, user.Id)
	if err != nil {
		t.Fatalf("Error calling GetDetail: %v", err)
	}

	assert.Equal(t, user.Id, resultUser.Id)
	assert.Equal(t, user.Email, resultUser.Email)
	assert.Equal(t, user.Password, resultUser.Password)
	assert.Equal(t, user.Name, resultUser.Name)
	assert.Equal(t, user.Phone, resultUser.Phone)
	assert.Equal(t, user.CreatedAt.Format(time.RFC3339), resultUser.CreatedAt.Format(time.RFC3339))
	assert.Equal(t, user.UpdatedAt.Format(time.RFC3339), resultUser.UpdatedAt.Format(time.RFC3339))
}

func TestGetByEmail(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Close()

	repo := repository.NewUser(
		repository.ROption{
			Common: common.Options{
				Database: common.Database{
					Read:  db,
					Write: db,
				},
			},
		},
	)

	user := model.User{
		Id:        1,
		Email:     "test@example.com",
		Password:  "hashed_password",
		Name:      "Test User",
		Phone:     "1234567890",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = db.Exec("INSERT INTO user (id, email, password, name, phone, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		user.Id, user.Email, user.Password, user.Name, user.Phone, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		t.Fatalf("Error inserting sample user: %v", err)
	}

	ctx := context.Background()
	resultUser, err := repo.GetByEmail(ctx, user.Email)
	if err != nil {
		t.Fatalf("Error calling GetByEmail: %v", err)
	}

	assert.Equal(t, user.Id, resultUser.Id)
	assert.Equal(t, user.Email, resultUser.Email)
	assert.Equal(t, user.Password, resultUser.Password)
	assert.Equal(t, user.Name, resultUser.Name)
	assert.Equal(t, user.Phone, resultUser.Phone)
	assert.Equal(t, user.CreatedAt.Format(time.RFC3339), resultUser.CreatedAt.Format(time.RFC3339))
	assert.Equal(t, user.UpdatedAt.Format(time.RFC3339), resultUser.UpdatedAt.Format(time.RFC3339))
}

func TestCreateUser(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Close()

	repo := repository.NewUser(
		repository.ROption{
			Common: common.Options{
				Database: common.Database{
					Read:  db,
					Write: db,
				},
			},
		},
	)

	user := model.User{
		Email:     "test@example.com",
		Password:  "hashed_password",
		Name:      "Test User",
		Phone:     "1234567890",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	ctx := context.Background()
	err = repo.Create(ctx, &user)
	if err != nil {
		t.Fatalf("Error calling CreateUser: %v", err)
	}

	resultUser, err := repo.GetByEmail(ctx, user.Email)
	if err != nil {
		t.Fatalf("Error retrieving created user: %v", err)
	}

	assert.Equal(t, user.Email, resultUser.Email)
	assert.Equal(t, user.Password, resultUser.Password)
	assert.Equal(t, user.Name, resultUser.Name)
	assert.Equal(t, user.Phone, resultUser.Phone)
	assert.Equal(t, user.CreatedAt.Format(time.RFC3339), resultUser.CreatedAt.Format(time.RFC3339))
	assert.Equal(t, user.UpdatedAt.Format(time.RFC3339), resultUser.UpdatedAt.Format(time.RFC3339))
}
