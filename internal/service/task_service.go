package service

import (
	"context"

	"github.com/dirgadm/fithub-api/internal/dto"
	"github.com/dirgadm/fithub-api/internal/model"
	"github.com/dirgadm/fithub-api/pkg/constants"
	"github.com/dirgadm/fithub-api/pkg/ehttp"
	"github.com/dirgadm/fithub-api/pkg/log"
)

type ITaskService interface {
	GetList(ctx context.Context, offset int, limit int, search string) (res []dto.TaskResponse, total int64, err error)
	Update(ctx context.Context, id int) (res dto.TaskResponse, err error)
	Create(ctx context.Context, req dto.CreateTaskRequest) (res dto.TaskResponse, err error)
}

type TaskService struct {
	opt SOption
}

func NewTaskService(opt SOption) ITaskService {
	return &TaskService{
		opt: opt,
	}
}

func (s *TaskService) GetList(ctx context.Context, offset int, limit int, search string) (res []dto.TaskResponse, total int64, err error) {
	var tasks []model.Task
	tasks, total, err = s.opt.Repository.Task.GetList(ctx, offset, limit, search)
	if err != nil {
		s.opt.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return
	}

	for _, task := range tasks {
		var user model.User
		user, err = s.opt.Repository.User.GetDetail(ctx, task.UserId)
		if err != nil {
			s.opt.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
			return
		}

		res = append(res, dto.TaskResponse{
			Id:          task.Id,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			User: dto.UserResponse{
				Id:        user.Id,
				Name:      user.Name,
				Email:     user.Email,
				CreatedAt: user.CreatedAt,
			},
			DueDate: task.Duedate,
		})
	}

	return
}

func (s *TaskService) Update(ctx context.Context, id int) (res dto.TaskResponse, err error) {
	userId := ctx.Value(constants.KeyUserID).(int)

	var user model.User
	user, err = s.opt.Repository.User.GetDetail(ctx, id)
	if err != nil {
		s.opt.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return
	}

	var task model.Task
	task, err = s.opt.Repository.Task.GetByIdAndUserId(ctx, id, userId)
	if err != nil {
		s.opt.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return
	}
	if task == (model.Task{}) {
		err = ehttp.ErrorOutput("id", "The task is not belong to this user")
		return
	}

	if task.Status != "pending" {
		err = ehttp.ErrorOutput("status", "The task should be pending")
		return
	}

	var t model.Task
	t = model.Task{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      "completed",
		UserId:      user.Id,
		Duedate:     task.Duedate,
	}
	err = s.opt.Repository.Task.Update(ctx, &t)

	res = dto.TaskResponse{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      "completed",
		User: dto.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
		DueDate: task.Duedate,
	}

	return
}

func (s *TaskService) Create(ctx context.Context, req dto.CreateTaskRequest) (res dto.TaskResponse, err error) {
	userId := ctx.Value(constants.KeyUserID).(int)

	var task model.Task
	task, err = s.opt.Repository.Task.GetByTitle(ctx, req.Title)
	if err != nil {
		s.opt.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return
	}
	if task != (model.Task{}) {
		err = ehttp.ErrorOutput("id", "Title already exist")
		return
	}

	var t model.Task
	t = model.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      "pending",
		UserId:      userId,
		Duedate:     req.Duedate,
	}
	err = s.opt.Repository.Task.Create(ctx, &t)

	res = dto.TaskResponse{
		Id:          t.Id,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		DueDate:     t.Duedate,
	}

	return
}
