package {{artifact_id}}

import (
	"context"
)

type Service interface {
	Status(ctx context.Context) (string, error)
}

type {{artifact_id}}Service struct{} //todo: check camelcase

// NewService makes a new Service.
func NewService() Service {
	return {{artifact_id}}Service{}
}

// Status only tell us that our service is ok!
func ({{artifact_id}}Service) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

