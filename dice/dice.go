package dice

import (
	"crypto/rand"
	"log"
	"math/big"
	"sort"
)

// Roll returns the total of count, sided dice. ie 3d6
func Roll(count, sides int) int {
	var t int
	for i := 0; i < count; i++ {
		r := cryptoRandSecure(int64(sides)) + 1
		//fmt.Println("roll: ", r)
		t = t + int(r)
	}
	return t
}

// RollKeep returns the rolls of count,sided dice as a slice of size count
func RollKeep(count, sides int) []int {
	t := make([]int, count)
	for i := 0; i < count; i++ {
		r := cryptoRandSecure(int64(sides)) + 1
		t[i] = int(r)
	}
	return t
}

// RollTop returns the total of count, sided dice keeping the top keep. ie 4d6k3
func RollTop(count, sides, keep int) int {
	t := make([]int, count)
	for i := 0; i < count; i++ {
		r := cryptoRandSecure(int64(sides)) + 1
		t[i] = int(r)
	}
	sort.Ints(t)

	var total int
	for i := 1; i < keep+1; i++ {
		total = total + t[i]
	}
	return total
}

// Rollfor72 Will roll Roll until total of 6 rolls is 72 or greater. Returns slice with results.
func Rollfor72() []int {
	var total int
	s := make([]int, 6)
	for {
		total = 0
		s = nil
		for i := 0; i < 6; i++ {
			t := Roll(3, 6)
			total = total + t
			s = append(s, t)
		}
		if total >= 72 {
			break
		}
	}
	return s
}

func cryptoRandSecure(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Println(err)
	}
	return nBig.Int64()
}
