package class

import (
	"basic-fantasy/datatypes"
)

func CanBeFighterMagicUser(c datatypes.Charstats) bool {
	if c.Base[datatypes.STRENGTH] >= 9 &&
		c.Base[datatypes.INTELLIGENCE] >= 9 &&
		(c.Race == datatypes.ELF || c.Race == 0) {
		return true
	}
	return false
}
