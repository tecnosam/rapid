package schema

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

// WorkflowState manages the orchestration metadata and step ledger
type WorkflowState struct {
	mu                 sync.RWMutex
	SessionID          uuid.UUID `json:"session_id"`
	Status             string    `json:"status"`              // RUNNING, PAUSED, COMPLETED, FAILED
	InstructionPointer []*Step   `json:"instruction_pointer"` // Currently executing StepIDs (supports parallel fan-out)
}

// StepExecutionLedger encapsulates the metrics, inputs, and private memory cell for an individual step
type StepExecutionLedger struct {
	Status        StepStatus `json:"status"` // PENDING, RUNNING, SUCCESS, FAILED
	Timing        StepTiming `json:"timing"`
	StatusMessage string     `json:"status_message"`
	ErrorCode     int        `json:"error_code"`
	memoryCell    StepMemoryCell
}

type StepTiming struct {
	StartedAt  time.Time `json:"started_at"`
	StoppedAt  time.Time `json:"stopped_at"`
	DurationMs float64   `json:"duration_ms"`
}

type StepMemoryCell struct {
	mu        sync.RWMutex
	boolCells map[string]bool
	numCells  map[string]float64
	strCells  map[string]string

	boolSeqCells map[string][]bool
	numSeqCells  map[string][]float64
	strSeqCells  map[string][]string

	variableMapping map[string]DataType
}

func (s *StepMemoryCell) ReadVariable(key string) (any, error) {
	// 1. O(1) look up the data type register index
	dType, exists := s.variableMapping[key]

	if !exists {
		return nil, fmt.Errorf("variable %s not initialized in step memory", key)
	}

	// 2. Direct map retrieval without interface reflection or type assertion loops
	switch dType {
	case TypeBoolean:
		return s.boolCells[key], nil
	case TypeNumber:
		return s.numCells[key], nil
	case TypeString:
		return s.strCells[key], nil
	case TypeBooleanSequence:
		return s.boolSeqCells[key], nil
	case TypeNumSequence:
		return s.numSeqCells[key], nil
	case TypeStringSequence:
		return s.strSeqCells[key], nil
	default:
		return nil, fmt.Errorf("unknown data type mapping descriptor")
	}
}

func (s *StepMemoryCell) WriteStringResult(value string) {

	s.variableMapping[STEP_RESULT_KEY] = TypeString
	s.strCells[STEP_RESULT_KEY] = value
}

func (s *StepMemoryCell) WriteString(key, value string) error {
	if key == STEP_RESULT_KEY {
		return fmt.Errorf("Key cannot be result string: %s", STEP_RESULT_KEY)
	}
	s.strCells[STEP_RESULT_KEY] = value
	return nil
}

func (s *StepExecutionLedger) WriteStringResult(value string) {
	s.memoryCell.WriteStringResult(value)
}
