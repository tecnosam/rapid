package http

import (
	"fmt"

	http_adapter "github.com/tecnosam/rapid/internal/extensions/adapters/http"
	"github.com/tecnosam/rapid/internal/schema"
)

type HTTPStepController struct {
	step schema.HTTPRequestStep

	adapter *http_adapter.HTTPAdapter
	httpDTO *http_adapter.HTTPRequest
}

func (c *HTTPStepController) Prepare() schema.StepStatus {
	/*
		loads the required adapter and prepares the DTO
	*/

	c.step.Core.UpdateStatus(
		schema.Ready,
		"",
	)

	return schema.Ready
}

func (c *HTTPStepController) Run() schema.StepStatus {
	/*
		Calls the LLM
	*/
	if c.adapter == nil {
		c.step.Core.UpdateStatus(
			schema.Crashed,
			"Failed to load required LLM Adapter!",
		)
		return c.step.Core.ExecutionLedger.Status
	}

	// TODO: concurrency consideration
	adapter := *c.adapter
	result := adapter.Invoke(*c.httpDTO)

	// TODO: add a cast operation for step.ReturnType
	// (user might want to cast result to specific d-type)
	// TODO: we may need to store (or at least log) http headers and status

	if result.IsSuccess() {
		c.step.Core.ExecutionLedger.WriteStringResult(string(result.Data))
		c.step.Core.UpdateStatus(
			schema.Success,
			"Failed to load required LLM Adapter!",
		)
	} else {

		c.step.Core.UpdateStatus(
			schema.Failed,
			getHTTPFailureMessage(&result),
		)
	}

	return c.step.Core.ExecutionLedger.Status
}

func getHTTPFailureMessage(result *http_adapter.HTTPResponse) string {
	failureMessage := fmt.Sprintf("HTTP Failed with %d status", result.HTTPStatusCode)

	if result.Data != nil {
		failureMessage = fmt.Sprintf(
			"%s. Message: %s",
			failureMessage,
			string(result.Data),
		)
	}

	return failureMessage
}
