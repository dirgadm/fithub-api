package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dirgadm/fithub-api/internal/model"
)

type ITask interface {
	GetList(ctx context.Context, offset int, limit int, search string) (products []model.Task, count int64, err error)
	GetByIdAndUserId(ctx context.Context, id int, userId int) (t model.Task, err error)
	GetByTitle(ctx context.Context, title string) (t model.Task, err error)
	Update(ctx context.Context, task *model.Task) (err error)
	Create(ctx context.Context, task *model.Task) (err error)
}

type task struct {
	opt ROption
}

func NewTask(opt ROption) ITask {
	return &task{
		opt: opt,
	}
}

func (m *task) GetByIdAndUserId(ctx context.Context, id int, userId int) (t model.Task, err error) {
	db := m.opt.Common.Database.Read
	query := "SELECT id, title, description, duedate, status, user_id FROM task WHERE id = ? AND user_id = ?"

	var task model.Task
	err = db.QueryRow(query, id, userId).Scan(&task.Id, &task.Title, &task.Description, &task.Duedate, &task.Status, &task.UserId)

	return task, nil
}

func (m *task) GetByTitle(ctx context.Context, title string) (t model.Task, err error) {
	db := m.opt.Common.Database.Read
	query := "SELECT id, title, description, duedate, status, user_id FROM task WHERE title = ?"

	var task model.Task
	err = db.QueryRow(query, title).Scan(&task.Id, &task.Title, &task.Description, &task.Duedate, &task.Status, &task.UserId)

	return task, nil
}

func (m *task) GetList(ctx context.Context, offset int, limit int, search string) (ts []model.Task, count int64, err error) {
	db := m.opt.Common.Database.Read
	query := "SELECT id, title, description, duedate, status, user_id FROM task"

	var rows *sql.Rows
	if search != "" {
		query += " WHERE title LIKE ? ORDER BY id LIMIT ?, ?"

		rows, err = db.Query(query, "%"+search+"%", offset, limit)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to query tasks: %v", err)
		}
	} else {
		query += " ORDER BY id LIMIT ?, ?"

		rows, err = db.Query(query, offset, limit)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to query tasks: %v", err)
		}
	}

	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Duedate, &task.Status, &task.UserId); err != nil {
			return nil, 0, fmt.Errorf("failed to scan task: %v", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error in rows: %v", err)
	}

	return tasks, int64(len(tasks)), nil
}

func (m *task) Update(ctx context.Context, task *model.Task) (err error) {
	db := m.opt.Common.Database.Write

	_, err = db.Exec("UPDATE task SET status = 'completed', title = ?, description = ?, duedate = ? WHERE id = ?", task.Title, task.Description, task.Duedate, task.UserId)
	if err != nil {
		return fmt.Errorf("failed to update task details: %v", err)
	}

	return nil
}

func (m *task) Create(ctx context.Context, task *model.Task) (err error) {
	db := m.opt.Common.Database.Write

	result, err := db.Exec("INSERT INTO task (title, description, duedate, status, user_id) VALUES (?, ?, ?, ?, ?)",
		task.Title, task.Description, task.Duedate, "pending", task.UserId)
	if err != nil {
		return fmt.Errorf("failed to create task: %v", err)
	}

	_, err = result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %v", err)
	}

	return err
}
