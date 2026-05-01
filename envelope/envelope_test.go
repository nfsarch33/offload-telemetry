package envelope

import (
	"encoding/json"
	"strings"
	"testing"
	"testing/quick"
)

func TestNewEvent_RedactsSensitiveInput(t *testing.T) {
	t.Parallel()

	event := NewEvent(Input{
		Model:               "claude-opus-4.7",
		LatencyMS:           1234,
		TokensPerSecond:     42.5,
		TimeToFirstTokenMS:  250,
		CostUSD:             0.12,
		StatusCode:          200,
		ParentTaskID:        "task-123",
		Prompt:              "launch secret prompt",
		Body:                "request body with token",
		Secret:              "sk-live-secret",
		ProviderToken:       "provider-token",
		AuthorizationHeader: "Bearer never-log-me",
	})

	encoded, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("Marshal returned error: %v", err)
	}

	for _, forbidden := range []string{
		"launch secret prompt",
		"request body with token",
		"sk-live-secret",
		"provider-token",
		"Bearer never-log-me",
	} {
		if strings.Contains(string(encoded), forbidden) {
			t.Fatalf("encoded event leaked %q: %s", forbidden, encoded)
		}
	}
}

func TestNewEvent_PropertyDoesNotLeakPromptBodyOrSecret(t *testing.T) {
	t.Parallel()

	property := func(prompt, body, secret string) bool {
		prompt = "prompt-marker-" + prompt
		body = "body-marker-" + body
		secret = "secret-marker-" + secret

		event := NewEvent(Input{
			Model:           "gpt-5.5",
			LatencyMS:       10,
			StatusCode:      200,
			ParentTaskID:    "task",
			Prompt:          prompt,
			Body:            body,
			Secret:          secret,
			ProviderToken:   "token-" + secret,
			TokensPerSecond: 1,
		})

		encoded, err := json.Marshal(event)
		if err != nil {
			return false
		}
		payload := string(encoded)
		return !containsNonEmpty(payload, prompt) &&
			!containsNonEmpty(payload, body) &&
			!containsNonEmpty(payload, secret)
	}

	if err := quick.Check(property, &quick.Config{MaxCount: 200}); err != nil {
		t.Fatal(err)
	}
}

func TestNewEvent_MatchesStableSchema(t *testing.T) {
	t.Parallel()

	event := NewEvent(Input{
		RecordedAt:         "2026-05-01T01:23:45Z",
		Tier:               "a",
		Decision:           "offloaded",
		Route:              "claude_code_subagent",
		Model:              "gpt-5.5",
		LatencyMS:          99,
		TokensPerSecond:    12.5,
		TimeToFirstTokenMS: 33,
		CostUSD:            0.07,
		StatusCode:         202,
		ParentTaskID:       "parent-1",
		Sender:             "cursor-ide",
		Prompt:             "do not include",
	})

	encoded, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		t.Fatalf("MarshalIndent returned error: %v", err)
	}

	want := `{
  "recorded_at": "2026-05-01T01:23:45Z",
  "schema_version": "offload.telemetry.v1",
  "tier": "a",
  "decision": "offloaded",
  "route": "claude_code_subagent",
  "model": "gpt-5.5",
  "latency_ms": 99,
  "tokens_per_second": 12.5,
  "time_to_first_token_ms": 33,
  "cost_usd": 0.07,
  "status_code": 202,
  "parent_task_id": "parent-1",
  "sender": "cursor-ide"
}`
	if string(encoded) != want {
		t.Fatalf("schema mismatch\nwant:\n%s\n\ngot:\n%s", want, encoded)
	}
}

func containsNonEmpty(payload, value string) bool {
	return value != "" && strings.Contains(payload, value)
}
