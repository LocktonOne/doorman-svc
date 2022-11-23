package config

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	ServiceConfiger
	VaultConfiger
	EthRPCConfiger
	GetClient() *ethclient.Client
}

type config struct {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	getter kv.Getter
	ServiceConfiger
	VaultConfiger
	EthRPCConfiger
	EthClient *ethclient.Client
}

func New(getter kv.Getter) Config {
	return &config{
		getter:          getter,
		Copuser:         copus.NewCopuser(getter),
		Listenerer:      comfig.NewListenerer(getter),
		Logger:          comfig.NewLogger(getter, comfig.LoggerOpts{}),
		ServiceConfiger: NewServiceConfiger(getter),
		VaultConfiger:   NewVaultConfiger(getter),
		EthRPCConfiger:  NewEthRPCConfiger(getter),
		EthClient:       NewEthRPCConfiger(getter).EthRPCConfig().EthClient(),
	}
}
