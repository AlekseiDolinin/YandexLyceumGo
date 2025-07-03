package main

import (
	"slices"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func calculateRating(goals, misses, assists int) float64 {
	if misses == 0 {
		return (float64(goals) + float64(assists)/2)
	} else {
		return (float64(goals) + float64(assists)/2) / float64(misses)
	}
}

func NewPlayer(name string, goals, misses, assists int) Player {
	rating := calculateRating(goals, misses, assists)
	return Player{Name: name, Goals: goals, Misses: misses, Assists: assists, Rating: rating}
}

func goalsSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Goals != b.Goals {
			return b.Goals - a.Goals
		}
		if a.Name < b.Name {
			return -1
		} else if a.Name > b.Name {
			return 1
		}
		return 0
	})
	return players
}

func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		if a.Rating != b.Rating {
			if a.Rating > b.Rating {
				return -1
			}
			return 1
		}
		if a.Name < b.Name {
			return -1
		} else if a.Name > b.Name {
			return 1
		}
		return 0
	})
	return players
}

func gmSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		gmA := float64(a.Goals)
		if a.Misses != 0 {
			gmA = float64(a.Goals) / float64(a.Misses)
		}

		gmB := float64(b.Goals)
		if b.Misses != 0 {
			gmB = float64(b.Goals) / float64(b.Misses)
		}

		if gmA != gmB {
			if gmA > gmB {
				return -1
			}
			return 1
		}
		if a.Name < b.Name {
			return -1
		} else if a.Name > b.Name {
			return 1
		}
		return 0
	})
	return players
}

func main() {

}
