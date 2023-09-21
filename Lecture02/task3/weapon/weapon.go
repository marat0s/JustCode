package weapon

type Weapon struct {
	Name          string
	Range         int
	Damage        int
	SpecialEffect string
}

func (w Weapon) Attack() string {
	return "This weapon attacks."
}
