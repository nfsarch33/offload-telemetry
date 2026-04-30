package prometheus

import (
	"fmt"
	"strings"

	"github.com/nfsarch33/offload-telemetry/envelope"
)

func Render(event envelope.Event) string {
	status := fmt.Sprintf("%d", event.StatusCode)
	model := sanitizeLabel(event.Model)

	return fmt.Sprintf(
		"offload_requests_total{model=%q,status_code=%q} 1\n"+
			"offload_latency_ms{model=%q,status_code=%q} %d\n"+
			"offload_tokens_per_second{model=%q,status_code=%q} %.6f\n",
		model, status,
		model, status, event.LatencyMS,
		model, status, event.TokensPerSecond,
	)
}

func sanitizeLabel(value string) string {
	value = strings.ReplaceAll(value, "\\", "\\\\")
	return strings.ReplaceAll(value, `"`, `\"`)
}
