package controllers

import "github.com/tecnosam/rapid/internal/schema"

type AbstractStepController interface {
	Prepare() schema.StepStatus
	Run() schema.StepStatus
}
