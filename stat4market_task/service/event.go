package service

import (
	"context"
	"stat4market_task/model"
	"stat4market_task/repository"
)

type Event interface {
	Create(ctx context.Context, event model.Event) error
}

type event struct {
	repo repository.Event
}

func EventService(repo repository.Event) Event {
	return &event{repo: repo}
}

func (s *event) Create(ctx context.Context, event model.Event) error {
	return s.repo.Insert(ctx, event)
}
