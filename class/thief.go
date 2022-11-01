package class

import (
	"basic-fantasy/datatypes"
)

func CanBeTheif(c datatypes.Charstats) bool {
	return c.Base[datatypes.DEXTERITY] >= 9
}
