package main

import (
	"context"
	"fmt"

	"github.com/fern91plugin/sdk"
)

type MyPlugin struct{}

func (p *MyPlugin) Execute(ctx context.Context, req *sdk.ExecuteRequest) (*sdk.ExecuteResponse, error) {
    return &sdk.ExecuteResponse{
        Success: true,
        Message: fmt.Sprintf("configured: %v", req.Config),
        Output:  "Executed with fernplugin SDK",
    }, nil
}

func main() {
    sdk.Serve(&MyPlugin{})
}
