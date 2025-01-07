package components

import "boxes/utils"

type SteppedMovementComponent struct {
	StepStartX, StepStartY float32
	StepDirection          utils.Facing
	StepDistance           float32
	StepSpeed              float32
}
