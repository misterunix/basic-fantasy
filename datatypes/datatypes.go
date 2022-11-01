package datatypes

const (
	STRENGTH = iota
	DEXTERITY
	CONSTITUTION
	INTELLIGENCE
	WISDOM
	CHARISMA
)

const (
	HUMAN = iota
	ELF
	DWARF
	HALFLING
)

type Charstats struct {
	Base  [6]int
	Race  int
	Class int
	Level int
	HP    int
	AC    int
}
