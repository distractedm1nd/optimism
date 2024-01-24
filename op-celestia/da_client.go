package celestia

import (
	"context"

	celestia "github.com/rollkit/celestia-openrpc"
	"github.com/rollkit/celestia-openrpc/types/share"
)

type DAClient struct {
	*celestia.Client
	Namespace share.Namespace
}

func NewDAClient(rpc string, authtoken string, namespace share.Namespace) (*DAClient, error) {
	client, err := celestia.NewClient(context.TODO(), rpc, authtoken)
	if err != nil {
		return nil, err
	}
	return &DAClient{
		Client:    client,
		Namespace: namespace,
	}, nil
}
