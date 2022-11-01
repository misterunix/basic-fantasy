package class

import (
	"basic-fantasy/datatypes"
)

func CanBeFighter(c datatypes.Charstats) bool {
	return c.Base[datatypes.STRENGTH] >= 9
}
