package celestia

import (
	"fmt"
	"net/url"

	"github.com/rollkit/celestia-openrpc/types/share"
	"github.com/urfave/cli/v2"

	opservice "github.com/ethereum-optimism/optimism/op-service"
)

const (
	DaRpcFlagName = "da.rpc"
)

var (
	defaultDaRpc = "localhost:26650"
)

func CLIFlags(envPrefix string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    DaRpcFlagName,
			Usage:   "dial address of data availability grpc client",
			Value:   defaultDaRpc,
			EnvVars: opservice.PrefixEnvVar(envPrefix, "DA_RPC"),
		},
	}
}

type Config struct {
	DaRpc     string
	AuthToken string
	Namespace share.Namespace
}

func (c Config) Check() error {
	if c.DaRpc == "" {
		c.DaRpc = defaultDaRpc
	}

	if _, err := url.Parse(c.DaRpc); err != nil {
		return fmt.Errorf("invalid da rpc: %w", err)
	}

	return nil
}

type CLIConfig struct {
	DaRpc     string
	AuthToken string
	Namespace share.Namespace
}

func (c CLIConfig) Check() error {
	if c.DaRpc == "" {
		c.DaRpc = defaultDaRpc
	}

	if _, err := url.Parse(c.DaRpc); err != nil {
		return fmt.Errorf("invalid da rpc: %w", err)
	}

	if _, err := share.NamespaceFromBytes([]byte(c.Namespace)); err != nil {
		return fmt.Errorf("invalid namespace: %w", err)
	}

	return nil
}

func NewCLIConfig() CLIConfig {
	return CLIConfig{
		DaRpc: defaultDaRpc,
	}
}

func ReadCLIConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		DaRpc: ctx.String(DaRpcFlagName),
	}
}
