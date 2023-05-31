package handler

import (
	"mintapi/internal/core/domain"
	"mintapi/internal/core/port/service"
)

type CreateTicketRequest struct {
	ShowID   uint64 `json:"show_id"`
	Metadata string `json:"metadata"`
	OwnerID  string `json:"owner_id"`
	Price    uint64 `json:"price"`
}

type CreateTicketVenueRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

func (r CreateTicketRequest) toDomainEvent() *domain.Ticket {
	return &domain.Ticket{
		Owner:    r.OwnerID,
		Price:    r.Price,
		ShowID:   r.ShowID,
		Metadata: r.Metadata,
	}
}

type Ticket struct {
	TicketSrv service.Ticket
}

func NewTicket(ticketSrv service.Ticket) *Ticket {
	return &Ticket{
		TicketSrv: ticketSrv,
	}
}
