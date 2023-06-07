package meent

import (
	"context"
	"errors"
	"log"
	"math/big"
	"mintapi/internal/core/domain"
	"mintapi/internal/core/port/resource"
	"mintapi/internal/infra/config"
	"mintapi/internal/infra/quicknode"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ContractManager struct {
	EthClient       quicknode.EthBaseClient
	ContractAddress string
	Meent           *Meent
}

func NewContractManager(ethClient quicknode.EthBaseClient, address string) resource.Contract {
	eventTicketContract, err := NewMeent(common.HexToAddress(address), ethClient.Client)
	if err != nil {
		log.Fatal(err)
	}

	return &ContractManager{
		EthClient:       ethClient,
		ContractAddress: address,
		Meent:           eventTicketContract,
	}
}

func (cm *ContractManager) CreateEvent(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	ownerPrivateKey := cm.EthClient.Configs.ConfigurationAccess.Get(config.OwnerPrivateKey).(string)

	auth := cm.EthClient.GetOwnerData(cm.EthClient.Client, ownerPrivateKey[2:], 0, 0)

	tx, err := cm.Meent.CreateEvent(
		auth,
		event.Name,
		big.NewInt(int64(event.Price)),
		big.NewInt(int64(event.TotalTicket)),
		common.HexToAddress(event.RealOwner),
		big.NewInt(int64(event.TotalReward)),
	)
	if err != nil {
		log.Fatal(err)
	}

	rec, err := bind.WaitMined(context.Background(), cm.EthClient.Client, tx)
	if err != nil {
		log.Fatal(err)
	}

	eventID, err := getEventIDFromReceipt(rec, cm.Meent)
	event.Id = eventID

	return event, err
}

func getEventIDFromReceipt(receipt *types.Receipt, eventTicketContract *Meent) (uint64, error) {
	for _, log := range receipt.Logs {
		event, err := eventTicketContract.ParseEventCreated(*log)
		if err != nil {
			return 0, err
		}

		return event.EventId.Uint64(), nil
	}

	return 0, errors.New("No Event event found in receipt")
}

func (cm *ContractManager) GetEvent(ctx context.Context, eventID uint64) (*domain.Event, error) {
	data, err := cm.Meent.Events(&bind.CallOpts{}, big.NewInt(int64(eventID)))

	return &domain.Event{
		Id:          eventID,
		Name:        data.EventName,
		Description: data.EventName,
	}, err
}
