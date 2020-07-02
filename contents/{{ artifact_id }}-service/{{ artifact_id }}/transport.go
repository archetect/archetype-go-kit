package {{artifact_id}}

import (
	"context"
	"encoding/json"
	"net/http"
)

type statusRequest struct{}

type statusResponse struct {
	Status string `json:"status"`
}


func decodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req statusRequest
	return req, nil
}

// Last but not least, we have the encoder for the response output
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
