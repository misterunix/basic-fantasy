package class

import (
	"basic-fantasy/datatypes"
)

func CanBeMagicUser(c datatypes.Charstats) bool {
	return c.Base[datatypes.INTELLIGENCE] >= 9
}
