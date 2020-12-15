package main

import (
	"fmt"
)

type number struct {
	id          int
	isFirst     bool
	turnsSpoken []int
}

type game struct {
	lastNumberSpoken *number
	numberCache      map[int]*number
	turns            int
}

func newGame(seeds []int) *game {
	g := &game{numberCache: make(map[int]*number)}

	for i, seed := range seeds {
		n := &number{id: seed, isFirst: true, turnsSpoken: []int{i + 1}}
		g.numberCache[seed] = n
		g.lastNumberSpoken = n
		g.turns++
	}
	return g
}

func (g *game) addToCache(id, turn int) {
	var n *number

	if cached, ok := g.numberCache[id]; ok {
		n = cached
		n.isFirst = false
		n.turnsSpoken = append(n.turnsSpoken, turn)
	} else {
		n = &number{id: id, isFirst: true, turnsSpoken: []int{turn}}
	}

	g.numberCache[id] = n
	g.lastNumberSpoken = n
}

func (g *game) speak() {
	var id int
	g.turns++
	num := g.numberCache[g.lastNumberSpoken.id]

	if !num.isFirst {
		lastTurnSpoken := len(num.turnsSpoken) - 1
		id = num.turnsSpoken[lastTurnSpoken] - num.turnsSpoken[lastTurnSpoken-1]
	}

	g.addToCache(id, g.turns)
}

func partOne(seeds []int, tgt int) int {
	g := newGame(seeds)
	for i := len(seeds) + 1; i <= tgt; i++ {
		g.speak()
	}

	return g.lastNumberSpoken.id
}

func main() {
	fmt.Println("part 1:", partOne([]int{19, 20, 14, 0, 9, 1}, 2020))
	fmt.Println("part 2:", partOne([]int{19, 20, 14, 0, 9, 1}, 30000000))
}
