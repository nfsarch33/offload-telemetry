# offload-telemetry

Private Go module for redacted Claude Code and Codex offload telemetry.

The envelope deliberately keeps only operational fields:

- model
- latency in milliseconds
- tokens per second
- time to first token in milliseconds
- cost in USD
- upstream status code
- parent task id

Prompts, request bodies, provider tokens, bearer headers, and secrets are never part of the emitted event shape.

## Commands

```bash
make test
make vet
make lint
make sentrux
```
