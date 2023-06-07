package handler

import (
	"bytes"
	"encoding/json"
	"mintapi/internal/core/domain"
	"mintapi/internal/core/port/service"
	"mintapi/internal/core/port/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func buildGinContext(req *http.Request, rec *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	return c
}

func buildJSONRequest(t *testing.T, body interface{}) *http.Request {
	bodyBytes, err := json.Marshal(&body)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(bodyBytes))
	req.Header.Add("Content-Type", "application/json")

	return req
}

func TestEvent_HandleCreate(t *testing.T) {
	type args struct {
		c CreateEventRequest
	}

	tests := []struct {
		name             string
		e                *Event
		args             args
		wantStatus       int
		wantResponseBody string
	}{
		{
			name: "should work correctly",
			e: &Event{
				EventSrv: func() service.Event {
					eventServiceMock := mocks.Event{}
					eventServiceMock.On("Create", mock.Anything, mock.Anything).
						Return(&domain.Event{
							Id: 12,
						}, nil)

					return &eventServiceMock
				}(),
			},
			args: args{c: CreateEventRequest{
				Name: "test_event",
			}},
			wantStatus:       http.StatusCreated,
			wantResponseBody: `{"id":{"value":"3fa85f64-5717-4562-b3fc-2c963f66afa6"}}`,
		},
	}

	for _, tt := range tests {

		rec := httptest.NewRecorder()
		c := buildGinContext(buildJSONRequest(t, &tt.args.c), rec)

		t.Run(tt.name, func(t *testing.T) {
			tt.e.HandleCreate(c)

			assert.Equal(t, tt.wantStatus, rec.Code)
			assert.Equal(t, tt.wantResponseBody, rec.Body.String())
		})
	}
}
