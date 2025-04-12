package sdk

import (
	"context"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type PluginImpl interface {
    Execute(ctx context.Context, req *ExecuteRequest) (*ExecuteResponse, error)
}

type PluginWrapperGRPC struct {
    impl PluginImpl
    UnimplementedPluginServer
}

func (s *PluginWrapperGRPC) Execute(ctx context.Context, req *ExecuteRequest) (*ExecuteResponse, error) {
    return s.impl.Execute(ctx, req)
}

type grpcClient struct {
    client PluginClient
}

func (c *grpcClient) Execute(ctx context.Context, req *ExecuteRequest) (*ExecuteResponse, error) {
    return c.client.Execute(ctx, req)
}

type PluginWrapper struct {
    Impl PluginImpl
}

func (p *PluginWrapper) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
    RegisterPluginServer(s, &PluginWrapperGRPC{impl: p.Impl})
    return nil
}

func (p *PluginWrapper) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, cc *grpc.ClientConn) (interface{}, error) {
    return &grpcClient{client: NewPluginClient(cc)}, nil
}

func (p *PluginWrapper) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
    return nil, nil
}

func (p *PluginWrapper) Server(broker *plugin.MuxBroker) (interface{}, error) {
	return nil, nil
}

func Serve(impl PluginImpl) {
    plugin.Serve(&plugin.ServeConfig{
        HandshakeConfig: HandshakeConfig,
        Plugins: map[string]plugin.Plugin{
            "executor": &PluginWrapper{Impl: impl},
        },
        GRPCServer: plugin.DefaultGRPCServer,
    })
}
