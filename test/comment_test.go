// build +e2e

package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v0"
)

func TestGetComment(t *testing.T) {
	resp, err := resty.R().Get(BASE_URL + "/api/comment")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostComment(t *testing.T) {
	resp, err := resty.R().
		SetBody(`{"slug": "/"}, "author": "Test Author", "body": "Test Comment"}`).
		Post(BASE_URL + "/api/comment")
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
}
