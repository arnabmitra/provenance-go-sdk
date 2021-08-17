package main

import (
"context"
"fmt"

"google.golang.org/grpc"

minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

func main() {
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}

	defer grpcConn.Close()

	mintClient := minttypes.NewQueryClient(grpcConn)
	var inf, error = mintClient.Inflation(
		context.Background(),
		&minttypes.QueryInflationRequest{},
	)
	if error == nil {
		fmt.Printf("no error")
	} else {
		fmt.Printf("error", error)
	}

	fmt.Printf(inf.Inflation.String())
}
