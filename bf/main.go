package main

import (
	"basic-fantasy/class"
	"basic-fantasy/datatypes"
	"basic-fantasy/dice"
	"fmt"
)

func main() {
	fmt.Println("Hello, basic-fantasy-rpg!!")

	var c datatypes.Charstats
	c.Base[datatypes.STRENGTH] = dice.Roll(3, 6)
	c.Base[datatypes.DEXTERITY] = dice.Roll(3, 6)
	c.Base[datatypes.CONSTITUTION] = dice.Roll(3, 6)
	c.Base[datatypes.INTELLIGENCE] = dice.Roll(3, 6)
	c.Base[datatypes.WISDOM] = dice.Roll(3, 6)
	c.Base[datatypes.CHARISMA] = dice.Roll(3, 6)

	fmt.Println(c)

	if class.CanBeCleric(c) {
		fmt.Println("You can be a cleric!")
	}
	if class.CanBeFighter(c) {
		fmt.Println("You can be a fighter!")
	}
	if class.CanBeMagicUser(c) {
		fmt.Println("You can be a magic-user!")
	}
	if class.CanBeThief(c) {
		fmt.Println("You can be a thief!")
	}
	if class.CanBeFighterMagicUser(c) {
		fmt.Println("You can be a fighter-magic-user!")
	}
	if class.CanBeMagicUserThief(c) {
		fmt.Println("You can be a magic-user-thief!")
	}

}
