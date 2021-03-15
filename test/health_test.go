// build +e2e

package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v0"
)

func TestHealthEndpoint(t *testing.T) {
	resp, err := resty.R().Get(BASE_URL + "/api/health")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}
