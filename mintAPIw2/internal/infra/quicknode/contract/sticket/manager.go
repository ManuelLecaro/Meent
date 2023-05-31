package sticket

import (
	"context"
	"errors"
	"log"
	"math/big"
	"mintapi/internal/core/domain"
	"mintapi/internal/infra/config"
	"mintapi/internal/infra/quicknode"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ContractManager struct {
	EthClient       quicknode.EthBaseClient
	ContractAddress string
	Sticket         *Sticket
}

func NewContractManager(ethClient quicknode.EthBaseClient, address string) *ContractManager {
	eventTicketContract, err := NewSticket(common.HexToAddress(address), ethClient.Client)
	if err != nil {
		log.Fatal(err)
	}

	return &ContractManager{
		EthClient:       ethClient,
		ContractAddress: address,
		Sticket:         eventTicketContract,
	}
}

func (cm *ContractManager) CreateEvent(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	ownerPrivateKey := cm.EthClient.Configs.ConfigurationAccess.Get(config.OwnerPrivateKey).(string)

	auth := cm.EthClient.GetOwnerData(cm.EthClient.Client, ownerPrivateKey[2:], 0, 0)

	tx, err := cm.Sticket.CreateShow(
		auth,
		event.Name,
		big.NewInt(int64(event.Date)),
		event.Venue.Address,
		[]string{},
	)
	if err != nil {
		log.Fatal(err)
	}

	rec, err := bind.WaitMined(context.Background(), cm.EthClient.Client, tx)
	if err != nil {
		log.Fatal(err)
	}

	eventID, err := getShowIdFromReceipt(rec, cm.Sticket)
	event.Id = eventID

	return event, err
}

func (cm *ContractManager) CreateTicket(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error) {
	ownerPrivateKey := cm.EthClient.Configs.ConfigurationAccess.Get(config.OwnerPrivateKey).(string)

	auth := cm.EthClient.GetOwnerData(cm.EthClient.Client, ownerPrivateKey[2:], 0, 0)

	tx, err := cm.Sticket.CreateTicket(
		auth,
		big.NewInt(int64(ticket.ShowID)),
		ticket.Metadata,
		big.NewInt(int64(ticket.Price)),
	)
	if err != nil {
		log.Fatal(err)
	}

	rec, err := bind.WaitMined(context.Background(), cm.EthClient.Client, tx)
	if err != nil {
		log.Fatal(err)
	}

	ticketID, err := getTicketIdFromReceipt(rec, cm.Sticket)
	ticket.Id = ticketID

	return ticket, err
}

func getTicketIdFromReceipt(receipt *types.Receipt, eventTicketContract *Sticket) (uint64, error) {
	for _, log := range receipt.Logs {
		event, err := eventTicketContract.ParseTicketCreated(*log)
		if err != nil {
			return 0, err
		}

		return event.Id.Uint64(), nil
	}

	return 0, errors.New("No TicketCreated event found in receipt")
}

func getShowIdFromReceipt(receipt *types.Receipt, eventTicketContract *Sticket) (uint64, error) {
	for _, log := range receipt.Logs {
		event, err := eventTicketContract.ParseShowCreated(*log)
		if err != nil {
			return 0, err
		}

		return event.Id.Uint64(), nil
	}

	return 0, errors.New("No TicketCreated event found in receipt")
}
