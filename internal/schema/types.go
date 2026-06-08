package schema

type StepType string
type StepStatus string
type ValidJSONSchemaType string
type ValidJSONType string

type DataType string

const (
	LLMCall          StepType = "llm_call"
	SQLQuery         StepType = "sql_query"
	LambdaExpression StepType = "lambda"
	HTTPRequest      StepType = "http_request"
	NativeMath       StepType = "native_math"
	StringProcessing StepType = "string_processing"
	ConditionRouter  StepType = "conditional_router"
)

const (
	JSONObject ValidJSONSchemaType = "object"
	JSONArray  ValidJSONSchemaType = "array"
)

const (
	JSONString ValidJSONType = "string"
	JSONNumber ValidJSONType = "number"
)

const (
	Pending StepStatus = "pending"
	Queued  StepStatus = "queued" // Ready to run but can't for some reason
	Running StepStatus = "running"
	Skipped StepStatus = "skipped"
	Failed  StepStatus = "failed"  // failed for user, calculation, or Third Party reason
	Crashed StepStatus = "crashed" // Runtime error
)

const (
	TypeBoolean DataType = "BOOLEAN"
	TypeNumber  DataType = "NUMBER"
	TypeString  DataType = "STRING"
	TypeObject  DataType = "OBJECT"
)
