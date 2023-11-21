package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/dirgadm/fithub-api/internal/model"
)

type IUser interface {
	GetDetail(ctx context.Context, id int) (user model.User, err error)
	GetByEmail(ctx context.Context, email string) (user model.User, err error)
	Create(ctx context.Context, user *model.User) (err error)
}

type user struct {
	opt ROption
}

func NewUser(opt ROption) IUser {
	return &user{
		opt: opt,
	}
}

func (m *user) GetDetail(ctx context.Context, id int) (user model.User, err error) {
	db := m.opt.Common.Database.Read

	row := db.QueryRowContext(ctx, "SELECT id, email, password, name, phone, created_at, updated_at FROM db_fithub.`user` WHERE id = ?", id)

	err = row.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user with ID %d not found", id)
		}
		return user, fmt.Errorf("error scanning user details: %v", err)
	}

	return user, nil
}

func (m *user) GetByEmail(ctx context.Context, email string) (user model.User, err error) {
	db := m.opt.Common.Database.Read

	row := db.QueryRowContext(ctx, "SELECT id, email, password, name, phone, created_at, updated_at FROM db_fithub.`user` WHERE email = ?", email)

	err = row.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Phone, &user.CreatedAt, &user.UpdatedAt)

	return user, nil
}

func (m *user) Create(ctx context.Context, user *model.User) (err error) {
	db := m.opt.Common.Database.Write

	_, err = db.ExecContext(
		ctx,
		"INSERT INTO db_fithub.`user` (email, password, name, phone, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		user.Email, user.Password, user.Name, user.Phone, user.CreatedAt, user.UpdatedAt,
	)

	return err
}
