package class

import (
	"basic-fantasy/datatypes"
)

func CanBeMagicUserThief(c datatypes.Charstats) bool {
	if c.Base[datatypes.DEXTERITY] >= 9 &&
		c.Base[datatypes.INTELLIGENCE] >= 9 &&
		c.Race == datatypes.ELF {
		return true
	}
	return false
}
