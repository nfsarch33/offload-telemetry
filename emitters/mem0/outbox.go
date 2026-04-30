package mem0

import (
	"context"

	"github.com/nfsarch33/offload-telemetry/envelope"
)

type Outbox interface {
	Append(ctx context.Context, event envelope.Event) error
}
