package utils

type Facing uint8

const (
	FacingNone Facing = iota
	FacingRight
	FacingDown
	FacingLeft
	FacingUp
)

type ColliderType uint8

const (
	ColliderTypeNone   ColliderType = 0
	ColliderTypePlayer ColliderType = 1 << (iota - 1)
	ColliderTypeEnemy
	ColliderTypeWall
	ColliderTypeProjectile
	ColliderTypePickup
)
