package sword

import "homework/Lecture02/task3/weapon"

type Sword struct {
	weapon.Weapon
	Material string
}

func (s Sword) Slash() string {
	return "The sword slashes."
}

func (s Sword) Attack() string {
	return "The sword slashes through the air."
}
