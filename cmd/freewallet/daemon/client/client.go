package client

import (
	"context"
	"time"

	"github.com/hungyu99/freed/cmd/freewallet/daemon/server"

	"github.com/pkg/errors"

	"github.com/hungyu99/freed/cmd/freewallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the freewalletd server, and returns the client instance
func Connect(address string) (pb.KaspawalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("freewallet daemon is not running, start it with `freewallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewKaspawalletdClient(conn), func() {
		conn.Close()
	}, nil
}