package celestia

import (
	"context"
	"fmt"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/rollkit/celestia-openrpc/types/da"
	"github.com/rollkit/celestia-openrpc/types/namespace"
)

const AuthKey = "Authorization"

type Client struct {
	DA da.API
}

type DAClient struct {
	*Client
	Namespace namespace.Namespace
}

func NewDAClient(rpc string, authtoken string, namespace namespace.Namespace) (*DAClient, error) {
	var authHeader http.Header
	if authtoken != "" {
		authHeader = http.Header{AuthKey: []string{fmt.Sprintf("Bearer %s", authtoken)}}
	}

	var client Client
	_, err := jsonrpc.NewClient(context.TODO(), rpc, "da", client.DA, authHeader)
	if err != nil {
		return nil, err
	}

	return &DAClient{
		Client:    &client,
		Namespace: namespace,
	}, nil
}
