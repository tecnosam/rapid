package schema

type StepCore struct {
	StepType        StepType
	Args            []string
	ExecutionLedger StepExecutionLedger
	next            []*Step // nil means terminal step
}

func NewStepCore(stepType StepType, args []string, next []*Step) StepCore {
	return StepCore{
		StepType:        stepType,
		Args:            args,
		next:            next,
		ExecutionLedger: StepExecutionLedger{},
	}
}

func (s *StepCore) GetNext() []*Step {
	return s.next
}

func (s *StepCore) ReadVariable(key string) (any, error) {
	return s.ExecutionLedger.memoryCell.ReadVariable(key)
}

func (s *StepCore) UpdateStatus(newStatus StepStatus, statusMessage string) {
	s.ExecutionLedger.Status = newStatus
	s.ExecutionLedger.StatusMessage = statusMessage
}
