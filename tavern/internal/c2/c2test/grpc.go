package c2test

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"realm.pub/tavern/internal/c2"
	"realm.pub/tavern/internal/c2/c2pb"
	"realm.pub/tavern/internal/ent"
	"realm.pub/tavern/internal/ent/enttest"
)

func New(t *testing.T) (c2pb.C2Client, *ent.Client, func()) {
	t.Helper()

	// TestDB Config
	var (
		driverName     = "sqlite3"
		dataSourceName = "file:ent?mode=memory&cache=shared&_fk=1"
	)

	// Ent Client
	graph := enttest.Open(t, driverName, dataSourceName, enttest.WithOptions())

	// gRPC Server
	lis := bufconn.Listen(1024 * 1024 * 10)
	baseSrv := grpc.NewServer()
	c2pb.RegisterC2Server(baseSrv, c2.New(graph))

	go func() {
		require.NoError(t, baseSrv.Serve(lis), "failed to serve grpc")
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)

	return c2pb.NewC2Client(conn), graph, func() {
		assert.NoError(t, lis.Close())
		baseSrv.Stop()
		assert.NoError(t, graph.Close())
	}
}
