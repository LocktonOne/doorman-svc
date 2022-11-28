package config

import (
	"github.com/ethereum/go-ethereum/common"
	vault "github.com/hashicorp/vault/api"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"golang.org/x/net/context"
	"log"
	"os"
)

type VaultConfiger interface {
	VaultConfig() *VaultConfig
	RegistryConfig() *RegistryConfig
}

type VaultConfig struct {
	Endpoint        string `json:"endpoint"`
	Token           string `json:"token"`
	MountPath       string `json:"mount_path"`
	PrivateKeyPath  string `json:"private_key_path"`
	AddressRegistry string `json:"address_registry"`
}

type RegistryConfig struct {
	Address common.Address `json:"address"`
}

func NewVaultConfiger(getter kv.Getter) VaultConfiger {
	return &vaultConfig{
		getter: getter,
	}
}

type vaultConfig struct {
	getter    kv.Getter
	once      comfig.Once
	vaultOnce comfig.Once
}

func (c *vaultConfig) VaultConfig() *VaultConfig {
	return c.once.Do(func() interface{} {
		raw := kv.MustGetStringMap(c.getter, "vault")
		config := VaultConfig{}
		err := figure.Out(&config).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return &config
	}).(*VaultConfig)
}

func (c *vaultConfig) RegistryConfig() *RegistryConfig {
	rg := RegistryConfig{}
	rg.Address = c.getContractAddress()
	return &rg
}

func (c *vaultConfig) getContractAddress() common.Address {
	config := struct {
		Address string `fig:"address,required"` //todo change when will be created vault file
	}{}

	err := c.getVaultSecret(c.VaultConfig().AddressRegistry, &config)
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress(config.Address)

	return address
}

func (c *vaultConfig) getVaultSecret(key string, out interface{}) error {
	vaultClient := c.vaultClient()
	secret, err := vaultClient.KVv2(c.VaultConfig().MountPath).Get(context.Background(), key)
	if err != nil {
		return errors.Wrap(err, "failed to get secret data")
	}

	return figure.
		Out(out).
		With(figure.BaseHooks).
		From(secret.Data).
		Please()
}

func (c *vaultConfig) vaultClient() *vault.Client {
	return c.vaultOnce.Do(func() interface{} {
		conf := vault.DefaultConfig()
		conf.Address = c.VaultConfig().Endpoint
		client, err := vault.NewClient(conf)
		if err != nil {
			log.Panicf("unable to initialize Vault client: %v", err)
		}
		client.SetToken(os.Getenv(c.VaultConfig().Token))
		return client
	}).(*vault.Client)
}
