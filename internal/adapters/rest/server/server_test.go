package server_test

import (
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/rest/server"
	"testing"
)

func TestServeStartsServerSuccessfully(t *testing.T) {
	rs := server.NewRestServer(&server.RestServerOptions{})

	go func() {
		rs.Serve()
	}()
}
