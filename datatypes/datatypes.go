package datatypes

const (
	STRENGTH = iota + 1
	DEXTERITY
	CONSTITUTION
	INTELLIGENCE
	WISDOM
	CHARISMA
)

const (
	HUMAN = iota + 1
	ELF
	DWARF
	HALFLING
)

type Charstats struct {
	Base  [7]int
	Race  int
	Class int
	Level int
	HP    int
	AC    int
}
