package helpers

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/tokene/doorman/contracts/MasterAccessManagement"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokene/doorman/internal/config"
	gosdk "gitlab.com/tokene/go-sdk"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	serviceConfigCtxKey
	nodeAdminsCtxKey
	regestryConfigCtxKey
	ethrpcConfigCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}
func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxServiceConfig(entry *config.ServiceConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, serviceConfigCtxKey, entry)
	}
}
func ServiceConfig(r *http.Request) *config.ServiceConfig {
	return r.Context().Value(serviceConfigCtxKey).(*config.ServiceConfig)
}

func CtxRegistryConfig(entry *config.RegistryConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, regestryConfigCtxKey, entry)
	}
}
func RegistryConfig(r *http.Request) *config.RegistryConfig {
	return r.Context().Value(regestryConfigCtxKey).(*config.RegistryConfig)
}

func CtxNodeAdmins(entry gosdk.NodeAdminsI) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, nodeAdminsCtxKey, entry)
	}
}
func NodeAdmins(r *http.Request) gosdk.NodeAdminsI {
	return r.Context().Value(nodeAdminsCtxKey).(gosdk.NodeAdminsI)
}

func CtxEthRPCConfig(entry *config.EthRPCConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ethrpcConfigCtxKey, entry)
	}
}

func EthRPCConfig(r *http.Request) *config.EthRPCConfig {
	return r.Context().Value(ethrpcConfigCtxKey).(*config.EthRPCConfig)
}

func CheckPermissionsByAddress(contractAddress, userAddress common.Address, coreEndpoint string) (bool, error) {
	client, err := ethclient.Dial(coreEndpoint)
	contract, err := MasterAccessManagement.NewMain(contractAddress, client)
	if err != nil {
		return false, err
	}
	_, err = contract.MainCaller.HasPermission(&bind.CallOpts{}, userAddress, "", "") //todo Add resource and permission
	if err != nil {
		return false, err
	}
	//return success, err
	return true, nil
}
