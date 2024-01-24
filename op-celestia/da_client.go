package celestia

import (
	"context"
	"fmt"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/rollkit/celestia-openrpc/types/blob"
	"github.com/rollkit/celestia-openrpc/types/share"
)

const AuthKey = "Authorization"

type Client struct {
	Blob blob.API
}

type DAClient struct {
	*Client
	Namespace share.Namespace
}

func NewDAClient(rpc string, authtoken string, namespace share.Namespace) (*DAClient, error) {
	var authHeader http.Header
	if authtoken != "" {
		authHeader = http.Header{AuthKey: []string{fmt.Sprintf("Bearer %s", authtoken)}}
	}

	var client Client
	_, err := jsonrpc.NewClient(context.TODO(), rpc, "blob", client.Blob, authHeader)
	if err != nil {
		return nil, err
	}

	return &DAClient{
		Client:    &client,
		Namespace: namespace,
	}, nil
}
