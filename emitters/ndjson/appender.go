package ndjson

import (
	"encoding/json"
	"fmt"

	"github.com/nfsarch33/offload-telemetry/envelope"
)

func MarshalLine(event envelope.Event) (string, error) {
	encoded, err := json.Marshal(event)
	if err != nil {
		return "", fmt.Errorf("marshal event: %w", err)
	}

	return string(encoded) + "\n", nil
}
