package process

import (
	"fmt"

	"github.com/ElrondNetwork/elrond-go/core"
	"github.com/ElrondNetwork/elrond-go/core/check"
	"github.com/ElrondNetwork/elrond-proxy-go/data"
)

// AddressPath defines the address path at which the nodes answer
const AddressPath = "/address/"

// AccountProcessor is able to process account requests
type AccountProcessor struct {
	connector       ExternalStorageConnector
	proc            Processor
	pubKeyConverter core.PubkeyConverter
}

// NewAccountProcessor creates a new instance of AccountProcessor
func NewAccountProcessor(proc Processor, pubKeyConverter core.PubkeyConverter, connector ExternalStorageConnector) (*AccountProcessor, error) {
	if check.IfNil(proc) {
		return nil, ErrNilCoreProcessor
	}
	if check.IfNil(pubKeyConverter) {
		return nil, ErrNilPubKeyConverter
	}
	if check.IfNil(connector) {
		return nil, ErrNilDatabaseConnector
	}

	return &AccountProcessor{
		proc:            proc,
		pubKeyConverter: pubKeyConverter,
		connector:       connector,
	}, nil
}

// GetAccount resolves the request by sending the request to the right observer and replies back the answer
func (ap *AccountProcessor) GetAccount(address string) (*data.Account, error) {
	observers, err := ap.getObserversForAddress(address)
	if err != nil {
		return nil, err
	}

	for _, observer := range observers {
		responseAccount := &data.ResponseAccount{}

		_, err = ap.proc.CallGetRestEndPoint(observer.Address, AddressPath+address, responseAccount)
		if err == nil {
			log.Info("account request", "address", address, "shard ID", observer.ShardId, "observer", observer.Address)
			return &responseAccount.AccountData, nil
		}

		log.Error("account request", "observer", observer.Address, "address", address, "error", err.Error())
	}

	return nil, ErrSendingRequest
}

// GetValueForKey returns the value for the given address and key
func (ap *AccountProcessor) GetValueForKey(address string, key string) (string, error) {
	observers, err := ap.getObserversForAddress(address)
	if err != nil {
		return "", err
	}

	for _, observer := range observers {
		apiResponse := make(map[string]interface{})
		apiPath := AddressPath + address + "/key/" + key
		_, err = ap.proc.CallGetRestEndPoint(observer.Address, apiPath, &apiResponse)
		if err != nil {
			log.Error("account request", "observer", observer.Address, "address", address, "error", err.Error())
			continue
		}

		return getValueOrError(apiResponse)
	}

	return "", ErrSendingRequest
}

func getValueOrError(response map[string]interface{}) (string, error) {
	if value, ok := response["value"]; ok {
		return fmt.Sprintf("%v", value), nil
	}

	if err, ok := response["error"]; ok {
		return "", fmt.Errorf("%v", err)
	}

	return "", fmt.Errorf("unexpected response")
}

// GetTransactions resolves the request and returns a slice of transaction for the specific address
func (ap *AccountProcessor) GetTransactions(address string) ([]data.DatabaseTransaction, error) {
	return ap.connector.GetTransactionsByAddress(address)
}

func (ap *AccountProcessor) getObserversForAddress(address string) ([]*data.Observer, error) {
	addressBytes, err := ap.pubKeyConverter.Decode(address)
	if err != nil {
		return nil, err
	}

	shardID, err := ap.proc.ComputeShardId(addressBytes)
	if err != nil {
		return nil, err
	}

	observers, err := ap.proc.GetObservers(shardID)
	if err != nil {
		return nil, err
	}

	return observers, nil
}
