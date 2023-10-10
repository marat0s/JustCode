package bow

import (
	"homework/Lecture02/task3/weapon"
)

type Bow struct {
	weapon.Weapon
	ArrowType string
}

func (b Bow) ShootArrow() string {
	return "The bow shoots an arrow."
}

func (b Bow) Attack() string {
	return "The bow shoots an arrow far away."
}
