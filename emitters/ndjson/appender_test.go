package ndjson

import (
	"strings"
	"testing"

	"github.com/nfsarch33/offload-telemetry/envelope"
)

func TestMarshalLine_AppendsRedactedEventAsNDJSON(t *testing.T) {
	t.Parallel()

	line, err := MarshalLine(envelope.NewEvent(envelope.Input{
		Model:        "gpt-5.5",
		LatencyMS:    50,
		StatusCode:   200,
		ParentTaskID: "parent",
		Prompt:       "never emit this prompt",
	}))
	if err != nil {
		t.Fatalf("MarshalLine returned error: %v", err)
	}

	if !strings.HasSuffix(line, "\n") {
		t.Fatalf("line %q does not end with newline", line)
	}
	if strings.Contains(line, "never emit this prompt") {
		t.Fatalf("line leaked prompt: %s", line)
	}
}
