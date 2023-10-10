package main

import (
	"fmt"
	"homework/Lecture02/task3/bow"
	"homework/Lecture02/task3/gun"
	"homework/Lecture02/task3/sword"
	"homework/Lecture02/task3/weapon"
)

type Attacker interface {
	Attack() string
}

func describeAttack(a Attacker) {
	fmt.Println(a.Attack())
}

func describeSpecialEffect(w weapon.Weapon) {
	fmt.Printf("Special Effect: %s\n", w.SpecialEffect)
}

func main() {
	mySword := sword.Sword{Weapon: weapon.Weapon{Name: "Excalibur", Range: 1, Damage: 50, SpecialEffect: "Fire"}, Material: "Steel"}
	myGun := gun.Gun{Weapon: weapon.Weapon{Name: "AK-47", Range: 300, Damage: 100, SpecialEffect: "Rapid Fire"}, Caliber: 7.62}
	myBow := bow.Bow{Weapon: weapon.Weapon{Name: "Longbow", Range: 250, Damage: 40, SpecialEffect: "Piercing"}, ArrowType: "Broadhead"}

	var myWeapon Attacker

	myWeapon = mySword
	describeAttack(myWeapon)
	describeSpecialEffect(mySword.Weapon)

	myWeapon = myGun
	describeAttack(myWeapon)
	describeSpecialEffect(myGun.Weapon)

	myWeapon = myBow
	describeAttack(myWeapon)
	describeSpecialEffect(myBow.Weapon)
}
