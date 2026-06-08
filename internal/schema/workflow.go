package schema

import (
	"time"

	"github.com/google/uuid"
)

// Workflow is a declarative blueprint of a DAG-based execution topology.
// It is separated from the runtime state (WorkflowState) so that one
// definition can drive many concurrent pipeline runs.
type Workflow struct {
	ID          uuid.UUID         `json:"id"`
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Description string            `json:"description,omitempty"`
	EntryPoint  string            `json:"entry_point"` // Step ID that begins execution
	Steps       map[string]*Step  `json:"steps"`       // All steps keyed by unique ID
	Settings    map[string]string `json:"settings,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}
