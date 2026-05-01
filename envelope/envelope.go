package envelope

const SchemaVersion = "offload.telemetry.v1"

type Input struct {
	RecordedAt          string
	Tier                string
	Decision            string
	Route               string
	Model               string
	LatencyMS           int64
	TokensPerSecond     float64
	TimeToFirstTokenMS  int64
	CostUSD             float64
	StatusCode          int
	ParentTaskID        string
	Sender              string
	Prompt              string
	Body                string
	Secret              string
	ProviderToken       string
	AuthorizationHeader string
}

type Event struct {
	RecordedAt         string  `json:"recorded_at"`
	SchemaVersion      string  `json:"schema_version"`
	Tier               string  `json:"tier"`
	Decision           string  `json:"decision"`
	Route              string  `json:"route"`
	Model              string  `json:"model"`
	LatencyMS          int64   `json:"latency_ms"`
	TokensPerSecond    float64 `json:"tokens_per_second"`
	TimeToFirstTokenMS int64   `json:"time_to_first_token_ms"`
	CostUSD            float64 `json:"cost_usd"`
	StatusCode         int     `json:"status_code"`
	ParentTaskID       string  `json:"parent_task_id"`
	Sender             string  `json:"sender"`
}

func NewEvent(input Input) Event {
	return Event{
		RecordedAt:         input.RecordedAt,
		SchemaVersion:      SchemaVersion,
		Tier:               input.Tier,
		Decision:           input.Decision,
		Route:              input.Route,
		Model:              input.Model,
		LatencyMS:          input.LatencyMS,
		TokensPerSecond:    input.TokensPerSecond,
		TimeToFirstTokenMS: input.TimeToFirstTokenMS,
		CostUSD:            input.CostUSD,
		StatusCode:         input.StatusCode,
		ParentTaskID:       input.ParentTaskID,
		Sender:             input.Sender,
	}
}
