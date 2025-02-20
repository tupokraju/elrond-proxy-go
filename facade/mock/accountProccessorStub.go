package mock

import (
	"github.com/ElrondNetwork/elrond-proxy-go/common"
	"github.com/ElrondNetwork/elrond-proxy-go/data"
)

// AccountProcessorStub --
type AccountProcessorStub struct {
	GetAccountCalled                        func(address string, options common.AccountQueryOptions) (*data.AccountModel, error)
	GetValueForKeyCalled                    func(address string, key string, options common.AccountQueryOptions) (string, error)
	GetShardIDForAddressCalled              func(address string) (uint32, error)
	GetTransactionsCalled                   func(address string) ([]data.DatabaseTransaction, error)
	ValidatorStatisticsCalled               func() (map[string]*data.ValidatorApiResponse, error)
	GetAllESDTTokensCalled                  func(address string, options common.AccountQueryOptions) (*data.GenericAPIResponse, error)
	GetESDTTokenDataCalled                  func(address string, key string, options common.AccountQueryOptions) (*data.GenericAPIResponse, error)
	GetESDTNftTokenDataCalled               func(address string, key string, nonce uint64, options common.AccountQueryOptions) (*data.GenericAPIResponse, error)
	GetESDTsWithRoleCalled                  func(address string, role string, options common.AccountQueryOptions) (*data.GenericAPIResponse, error)
	GetNFTTokenIDsRegisteredByAddressCalled func(address string, options common.AccountQueryOptions) (*data.GenericAPIResponse, error)
	GetKeyValuePairsCalled                  func(address string, options common.AccountQueryOptions) (*data.GenericAPIResponse, error)
	GetESDTsRolesCalled                     func(address string, options common.AccountQueryOptions) (*data.GenericAPIResponse, error)
}

// GetKeyValuePairs -
func (aps *AccountProcessorStub) GetKeyValuePairs(address string, options common.AccountQueryOptions) (*data.GenericAPIResponse, error) {
	return aps.GetKeyValuePairsCalled(address, options)
}

// GetAllESDTTokens -
func (aps *AccountProcessorStub) GetAllESDTTokens(address string, options common.AccountQueryOptions) (*data.GenericAPIResponse, error) {
	return aps.GetAllESDTTokensCalled(address, options)
}

// GetESDTTokenData -
func (aps *AccountProcessorStub) GetESDTTokenData(address string, key string, options common.AccountQueryOptions) (*data.GenericAPIResponse, error) {
	return aps.GetESDTTokenDataCalled(address, key, options)
}

// GetESDTNftTokenData -
func (aps *AccountProcessorStub) GetESDTNftTokenData(address string, key string, nonce uint64, options common.AccountQueryOptions) (*data.GenericAPIResponse, error) {
	return aps.GetESDTNftTokenDataCalled(address, key, nonce, options)
}

// GetESDTsWithRole -
func (aps *AccountProcessorStub) GetESDTsWithRole(address string, role string, options common.AccountQueryOptions) (*data.GenericAPIResponse, error) {
	return aps.GetESDTsWithRoleCalled(address, role, options)
}

// GetESDTsRoles -
func (aps *AccountProcessorStub) GetESDTsRoles(address string, options common.AccountQueryOptions) (*data.GenericAPIResponse, error) {
	if aps.GetESDTsRolesCalled != nil {
		return aps.GetESDTsRolesCalled(address, options)
	}

	return &data.GenericAPIResponse{}, nil
}

// GetNFTTokenIDsRegisteredByAddress -
func (aps *AccountProcessorStub) GetNFTTokenIDsRegisteredByAddress(address string, options common.AccountQueryOptions) (*data.GenericAPIResponse, error) {
	return aps.GetNFTTokenIDsRegisteredByAddressCalled(address, options)
}

// GetAccount --
func (aps *AccountProcessorStub) GetAccount(address string, options common.AccountQueryOptions) (*data.AccountModel, error) {
	return aps.GetAccountCalled(address, options)
}

// GetValueForKey --
func (aps *AccountProcessorStub) GetValueForKey(address string, key string, options common.AccountQueryOptions) (string, error) {
	return aps.GetValueForKeyCalled(address, key, options)
}

// GetShardIDForAddress --
func (aps *AccountProcessorStub) GetShardIDForAddress(address string) (uint32, error) {
	return aps.GetShardIDForAddressCalled(address)
}

// GetTransactions --
func (aps *AccountProcessorStub) GetTransactions(address string) ([]data.DatabaseTransaction, error) {
	return aps.GetTransactionsCalled(address)
}

// ValidatorStatistics --
func (aps *AccountProcessorStub) ValidatorStatistics() (map[string]*data.ValidatorApiResponse, error) {
	return aps.ValidatorStatisticsCalled()
}
