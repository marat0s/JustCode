package gun

import "homework/Lecture02/task3/weapon"

type Gun struct {
	weapon.Weapon
	Caliber float64
}

func (g Gun) Shoot() string {
	return "The gun fires a bullet."
}

func (g Gun) Attack() string {
	return "The gun shoots a bullet."
}
