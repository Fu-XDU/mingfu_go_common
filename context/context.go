package context

import (
	"context"
	"github.com/Fu-XDU/mingfu_go_common/constants"
	"google.golang.org/grpc/metadata"
)

func NewContextWithUuid(uuid string) (ctx context.Context) {
	header := metadata.Pairs(constants.UUID, uuid)
	ctx = metadata.NewOutgoingContext(context.Background(), header)

	return
}

func GetUuidFromContext(ctx context.Context) (uuid string) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md, ok = metadata.FromOutgoingContext(ctx)
		if !ok {
			return
		}
	}

	values := md.Get(constants.UUID)
	if values == nil || len(values) == 0 {
		return
	}

	uuid = values[0]
	return
}
