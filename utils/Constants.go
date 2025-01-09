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
	ColliderTypeNone ColliderType = 1 << iota
	ColliderTypePlayer
	ColliderTypeEnemy
	ColliderTypeWall
	ColliderTypeProjectile
	ColliderTypePickup
)
