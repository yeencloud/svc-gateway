package clients

import (
	"context"

	rpc "github.com/yeencloud/lib-rpc"
	identityService "github.com/yeencloud/svc-identity/contract/proto/generated"
)

type IdentityClient struct {
	rpcClient      *rpc.RPCClient
	identityClient identityService.IdentityServiceClient
}

func NewIdentityClient(config string) *IdentityClient {
	client := rpc.NewRPCClient(config)

	return &IdentityClient{
		rpcClient:      client,
		identityClient: identityService.NewIdentityServiceClient(client.Connection),
	}
}

func (client *IdentityClient) checkConnection() error {
	err := client.rpcClient.Connect()
	if err != nil {
		return err
	}

	client.identityClient = identityService.NewIdentityServiceClient(client.rpcClient.Connection)
	return nil
}

func (client *IdentityClient) Register(ctx context.Context, in *identityService.RegisterObject) (*identityService.RegisterResponse, error) {
	err := client.checkConnection()
	if err != nil {
		return nil, err
	}

	return client.identityClient.Register(ctx, in)
}
