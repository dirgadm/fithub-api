package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/dirgadm/fithub-api/internal/common"
	"github.com/dirgadm/fithub-api/internal/model"
	"github.com/dirgadm/fithub-api/internal/repository"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE task (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			duedate TIMESTAMP,
			status TEXT,
			user_id INTEGER
		);
	`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS user (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email VARCHAR(255) NULL,
			password VARCHAR(255) NULL,
			name VARCHAR(255) NULL,
			phone VARCHAR(255) NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT NULL
		);
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestGetList(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Close()

	repo := repository.NewTask(
		repository.ROption{
			Common: common.Options{
				Database: common.Database{
					Read:  db,
					Write: db,
				},
			},
		},
	)

	testTasks := []model.Task{
		{
			Title:       "Task 1",
			Description: "Description 1",
			Duedate:     time.Now(),
			Status:      "pending",
			UserId:      1,
		},
		{
			Title:       "Task 2",
			Description: "Description 2",
			Duedate:     time.Now(),
			Status:      "completed",
			UserId:      1,
		},
	}

	for _, task := range testTasks {
		err := repo.Create(context.Background(), &task)
		if err != nil {
			t.Fatalf("Error creating task: %v", err)
		}
	}

	retrievedTasks, count, err := repo.GetList(context.Background(), 0, 10, "Task")
	if err != nil {
		t.Fatalf("Error getting task list: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, int64(len(testTasks)), count)
	assert.Len(t, retrievedTasks, len(testTasks))
}

func TestGetByIdAndUserId(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Close()

	repo := repository.NewTask(
		repository.ROption{
			Common: common.Options{
				Database: common.Database{
					Read:  db,
					Write: db,
				},
			},
		},
	)

	testTask := &model.Task{
		Id:          1,
		Title:       "Test Task",
		Description: "Test Description",
		Duedate:     time.Now(),
		Status:      "pending",
		UserId:      1,
	}

	err = repo.Create(context.Background(), testTask)
	if err != nil {
		t.Fatalf("Error creating test task: %v", err)
	}

	retrievedTask, err := repo.GetByIdAndUserId(context.Background(), testTask.Id, testTask.UserId)
	if err != nil {
		t.Fatalf("Error getting task by ID and user ID: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, testTask.Id, retrievedTask.Id)
	assert.Equal(t, testTask.Title, retrievedTask.Title)
	assert.Equal(t, testTask.Description, retrievedTask.Description)
	assert.Equal(t, testTask.Duedate.Local(), retrievedTask.Duedate.Local())
	assert.Equal(t, testTask.Status, retrievedTask.Status)
	assert.Equal(t, testTask.UserId, retrievedTask.UserId)
}

func TestGetByTitle(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Close()

	repo := repository.NewTask(
		repository.ROption{
			Common: common.Options{
				Database: common.Database{
					Read:  db,
					Write: db,
				},
			},
		},
	)

	testTask := model.Task{
		Id:          1,
		Title:       "Test Task",
		Description: "Test Description",
		Duedate:     time.Now(),
		Status:      "pending",
		UserId:      1,
	}

	err = repo.Create(context.Background(), &testTask)
	if err != nil {
		t.Fatalf("Error creating test task: %v", err)
	}

	retrievedTask, err := repo.GetByTitle(context.Background(), testTask.Title)
	if err != nil {
		t.Fatalf("Error getting task by title: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, testTask.Id, retrievedTask.Id)
	assert.Equal(t, testTask.Title, retrievedTask.Title)
	assert.Equal(t, testTask.Description, retrievedTask.Description)
	assert.Equal(t, testTask.Duedate.Local(), retrievedTask.Duedate.Local())
	assert.Equal(t, testTask.Status, retrievedTask.Status)
	assert.Equal(t, testTask.UserId, retrievedTask.UserId)
}

func TestUpdate(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Close()

	repo := repository.NewTask(
		repository.ROption{
			Common: common.Options{
				Database: common.Database{
					Read:  db,
					Write: db,
				},
			},
		},
	)

	testTask := model.Task{
		Id:          1,
		Title:       "Test Task",
		Description: "Test Description",
		Duedate:     time.Now(),
		Status:      "pending",
		UserId:      1,
	}

	err = repo.Create(context.Background(), &testTask)
	if err != nil {
		t.Fatalf("Error creating test task: %v", err)
	}

	testTask.Title = "Updated Title"
	testTask.Description = "Updated Description"
	testTask.Duedate = time.Now()

	err = repo.Update(context.Background(), &testTask)
	if err != nil {
		t.Fatalf("Error updating task: %v", err)
	}

	updatedTask, err := repo.GetByIdAndUserId(context.Background(), testTask.Id, testTask.UserId)
	if err != nil {
		t.Fatalf("Error getting updated task: %v", err)
	}

	assert.Equal(t, "Updated Title", updatedTask.Title)
	assert.Equal(t, "Updated Description", updatedTask.Description)
	assert.Equal(t, testTask.Duedate.Local(), updatedTask.Duedate.Local())
	assert.Equal(t, "completed", updatedTask.Status)
}

func TestCreate(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Error setting up test database: %v", err)
	}
	defer db.Close()

	repo := repository.NewTask(
		repository.ROption{
			Common: common.Options{
				Database: common.Database{
					Read:  db,
					Write: db,
				},
			},
		},
	)

	testTask := model.Task{
		Id:          1,
		Title:       "Test Task",
		Description: "Test Description",
		Duedate:     time.Now(),
		Status:      "pending",
		UserId:      1,
	}

	err = repo.Create(context.Background(), &testTask)
	if err != nil {
		t.Fatalf("Error creating test task: %v", err)
	}

	createdTask, err := repo.GetByIdAndUserId(context.Background(), testTask.Id, testTask.UserId)
	if err != nil {
		t.Fatalf("Error getting created task: %v", err)
	}

	assert.Equal(t, testTask.Title, createdTask.Title)
	assert.Equal(t, testTask.Description, createdTask.Description)
	assert.Equal(t, testTask.Duedate.Local(), createdTask.Duedate.Local())
	assert.Equal(t, testTask.Status, createdTask.Status)
	assert.Equal(t, testTask.UserId, createdTask.UserId)
}
