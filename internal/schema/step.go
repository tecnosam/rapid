package schema

type Step interface {
	GetNext() []*Step
}

type StepCore struct {
	StepType StepType
	Args     []string
	next     []*Step // nil means terminal step
}

func NewStepCore(stepType StepType, args []string, next []*Step) StepCore {
	return StepCore{
		StepType: stepType,
		Args:     args,
		next:     next,
	}
}

func (s *StepCore) GetNext() []*Step {
	return s.next
}

type JSONProperty struct {
	Type ValidJSONType
}

type JSONReturnSchema struct {
	SchemaType         ValidJSONSchemaType
	Properties         *map[string]JSONProperty
	RequiredProperties *[]string
}

type LLMCallStep struct {
	Core            StepCore
	Provider        string
	Model           *string
	Prompt_template *string
	ReturnType      string
	JSONSchema      *JSONReturnSchema
}

func (s *LLMCallStep) GetNext() []*Step {
	return s.Core.next
}

type SQLQueryStep struct {
	Core       StepCore
	Datasource string
	Query      string
	Args       []string
}

func (s *SQLQueryStep) GetNext() []*Step {
	return s.Core.next
}

type LambdaExpressionStep struct {
	Core       StepCore
	expression string
	Variables  map[string]string
}

func (s *LambdaExpressionStep) GetNext() []*Step {
	return s.Core.next
}

type HTTPRequestStep struct {
	Core      StepCore
	Method    string
	URL       string
	Headers   *map[string]string
	Body      *map[string]string
	TimeoutMS uint8
}

func (s *HTTPRequestStep) GetNext() []*Step {
	return s.Core.next
}

type NativeMathStep struct {
	Core      StepCore
	Operation string
	Args      []float64
}

func (s *NativeMathStep) GetNext() []*Step {
	return s.Core.next
}

type StringProcessingStep struct {
	Core      StepCore
	Operation string
	Args      []string
}

func (s *StringProcessingStep) GetNext() []*Step {
	return s.Core.next
}

type ConditionRouterStep struct {
	Core       StepCore
	Expression string
	Variables  map[string]string

	Branches map[string]*Step
}

func (s *ConditionRouterStep) GetNext() []*Step {
	return s.Core.next
}
