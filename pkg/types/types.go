package types

const (
	PERIODICITY = iota + 10010
	SINGULARITY
	TIMELINESS
)

type CreatedTaskRequest struct {
	TaskID            string `json:"task_id"`
	TaskName          string `json:"task_name"`
	TaskType          int    `json:"task_type"`
	ConcurrencyNumber int    `json:"concurrencry_number"`
	DepthNumber       int    `json:"depth_number"`
	Targets           string `json:"targers"`
}
