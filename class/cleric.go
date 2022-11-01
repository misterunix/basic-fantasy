package class

import "basic-fantasy/datatypes"

func CanBeCleric(c datatypes.Charstats) bool {
	return c.Base[datatypes.WISDOM] >= 9
}
