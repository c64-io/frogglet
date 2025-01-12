package components

import (
	"boxes/utils"
)

type ColliderComponent struct {
	ColliderType utils.ColliderType
	CollidesWith utils.ColliderType
}
