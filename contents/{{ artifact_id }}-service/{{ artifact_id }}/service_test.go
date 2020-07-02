package {{artifact_id}}

import (
	"context"
	"testing"
)

func TestStatus(t *testing.T) {
	srv, ctx := setup()

	s, err := srv.Status(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// testing status
	ok := s == "ok"
	if !ok {
		t.Errorf("expected service to be ok")
	}
}

func setup() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}
