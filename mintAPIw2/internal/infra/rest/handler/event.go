package handler

import (
	"context"
	"fmt"
	"mintapi/internal/core/domain"
	"mintapi/internal/core/port/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Event struct {
	EventSrv service.Event
}

type CreateEventRequest struct {
	Name        string                  `json:"name" binding:"required"`
	Date        time.Time               `json:"date" binding:"required"`
	Venue       CreateEventVenueRequest `json:"venue" binding:"required"`
	Description string                  `json:"description"`
}

type CreateEventVenueRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

func (r CreateEventRequest) toDomainEvent() *domain.Event {
	return &domain.Event{
		Name:        r.Name,
		Description: r.Description,
		Date:        uint64(r.Date.Unix()),
		Venue: &domain.Venue{
			Name:    r.Venue.Name,
			Address: r.Venue.Address,
		},
	}
}

func NewEvent(srv service.Event) *Event {
	return &Event{
		EventSrv: srv,
	}
}

func (e *Event) HandleCreate(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var req CreateEventRequest

	if err := e.getAndValidateBody(c, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	resultChan := make(chan struct {
		Value *domain.Event
		Error error
	})

	go func() {
		res, err := e.EventSrv.Create(ctx, req.toDomainEvent())

		resultChan <- struct {
			Value *domain.Event
			Error error
		}{Value: res, Error: err}
	}()

	handleResponse(c, resultChan, http.StatusCreated)
}

func (e *Event) HandleGet(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var req CreateEventRequest

	if err := e.getAndValidateBody(c, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	resultChan := make(chan struct {
		Value *domain.Event
		Error error
	})

	go func() {
		res, err := e.EventSrv.Create(ctx, req.toDomainEvent())

		resultChan <- struct {
			Value *domain.Event
			Error error
		}{Value: res, Error: err}
	}()

	handleResponse(c, resultChan, http.StatusCreated)
}

func (e *Event) getAndValidateBody(c *gin.Context, req interface{}) error {
	if err := c.ShouldBindJSON(&req); err != nil {
		return err
	}

	return nil
}

func handleResponse(c *gin.Context, resultChan <-chan struct {
	Value *domain.Event
	Error error
}, successStatus int) {
	select {
	case result := <-resultChan:
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})

			return
		}

		c.JSON(successStatus, result.Value)

	case <-c.Request.Context().Done():
		err := c.Request.Context().Err()
		if err == context.DeadlineExceeded {
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "Timeout exceeded"})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Request canceled: %v", err)})

		return
	}
}
