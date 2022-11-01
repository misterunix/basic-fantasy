package class

import (
	"basic-fantasy/datatypes"
)

func CanBeThief(c datatypes.Charstats) bool {
	return c.Base[datatypes.DEXTERITY] >= 9
}
