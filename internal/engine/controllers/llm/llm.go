package controllers

import (
	"github.com/tecnosam/rapid/internal/extensions/adapters/llm"
	"github.com/tecnosam/rapid/internal/schema"
)

type LLMCallStepController struct {
	step *schema.LLMCallStep

	adapter       *llm.LLMAdapter
	prompt        string
	substitutions map[string]string
}

func (c *LLMCallStepController) Prepare() schema.StepStatus {
	/*
		Prepare's the prompt with the required substitutions
		for processing.
		also loads the required adapter ()
	*/

	c.step.Core.UpdateStatus(
		schema.Ready,
		"",
	)

	return schema.Ready
}

func (c *LLMCallStepController) Run() schema.StepStatus {
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
	result := adapter.Invoke(&c.prompt, &c.substitutions)

	// TODO: add a cast operation for step.ReturnType
	// (user might want to cast result to specific d-type)
	c.step.Core.ExecutionLedger.WriteStringResult(result)

	c.step.Core.UpdateStatus(
		schema.Success,
		"Failed to load required LLM Adapter!",
	)

	return c.step.Core.ExecutionLedger.Status
}
