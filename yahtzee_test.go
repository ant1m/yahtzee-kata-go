package main

import (
	"fmt"
	"reflect"
	"testing"
)

var functions = []rollResult{carre, un, deux, petiteSuite, grandeSuite, yahzee}

func TestCanGetPaire(t *testing.T) {
	roll := []int{3, 6, 6, 4, 5}

	if result, _ := paire(roll); result != 12 {
		t.Fail()
	}
}

func TestCanGetDoublePaire(t *testing.T) {
	roll := []int{3, 6, 6, 4, 4}

	if result, _ := doublePaire(roll); result != 20 {
		t.Fail()
	}
}

func TestCanGetBrelan(t *testing.T) {
	roll := []int{6, 6, 6, 4, 5}

	if result, _ := brelan(roll); result != 18 {
		t.Fail()
	}
}

func TestCanGetBrelanForFull(t *testing.T) {
	roll := []int{6, 6, 6}

	if result, _ := brelan(roll); result != 18 {
		t.Fail()
	}
}

func TestCanGetCarre(t *testing.T) {
	roll := []int{6, 6, 6, 6, 5}

	if result, _ := carre(roll); result != 24 {
		t.Fail()
	}
}

func TestCanGetYahzee(t *testing.T) {
	roll := []int{6, 6, 6, 6, 6}

	if result, _ := yahzee(roll); result != 50 {
		t.Fail()
	}
}

func TestCanGetTrois(t *testing.T) {
	roll := []int{3, 2, 2, 2, 2}

	if result, _ := trois(roll); result != 3 {
		t.Fail()
	}
}

func TestCanGetPetiteSuite(t *testing.T) {
	roll := []int{5, 4, 3, 2, 1}

	if result, _ := petiteSuite(roll); result != 15 {
		t.Fail()
	}
}

func TestCanGetFull(t *testing.T) {
	roll := []int{5, 5, 5, 3, 3}

	if result, _ := Full(roll); result != 21 {
		fmt.Println(Full(roll))
		t.Fail()
	}
}

func TestCheck2(t *testing.T) {
	roll := []int{6, 5, 4, 1, 1}

	name := check(roll)

	println(name)
}

type Result struct {
	Category string
}

func Check(roll []int) *Result {

	return &Result{Category: "PAIRE"}
}

func paire(sortedRoll []int) (int, string) {
	name := "paire"
	var p int = 0
	for _, val := range sortedRoll {
		if p != 0 {
			if val == p {
				return 2 * val, name
			}
		}
		p = val
	}
	return 0, name
}

func brelan(sortedRoll []int) (int, string) {
	name := "brelan"

	var p1 int = 0
	var p2 int = 0
	for _, val := range sortedRoll {
		if p1+p2 != 0 {
			if val == p1 && p1 == p2 {
				return 3 * val, name
			}
		}
		p1 = p2
		p2 = val
	}
	return 0, name
}

func carre(sortedRoll []int) (int, string) {
	name := "carre"

	var p1 int = 0
	var p2 int = 0
	var p3 int = 0
	for _, val := range sortedRoll {
		if p1+p2+p3 != 0 {
			if val == p1 && p1 == p2 && p2 == p3 {
				return 4 * val, name
			}
		}
		p1 = p2
		p2 = p3
		p3 = val
	}
	return 0, name
}

func yahzee(sortedRoll []int) (int, string) {
	name := "yahzee"

	p := sortedRoll[0]
	for _, val := range sortedRoll {
		if p != val {
			return 0, name
		}
	}
	return 50, name
}

func doublePaire(sortedRoll []int) (int, string) {
	name := "doublePaire"

	p, _ := paire(sortedRoll)
	if p == 0 {
		return 0, name
	}
	val := p / 2
	newRoll := make([]int, len(sortedRoll)-2)
	var index int
	for _, j := range sortedRoll {
		if j != val {
			newRoll[index] = j
			index++
		}
	}
	deuxiemePaire, _ := paire(newRoll)
	if deuxiemePaire == 0 {
		return 0, name
	}
	return p + deuxiemePaire, name
}

func un(sortedRoll []int) (int, string) {
	name := "un"
	var result int
	for _, j := range sortedRoll {
		if j == 1 {
			result++
		}
	}
	return result, name
}

func deux(sortedRoll []int) (int, string) {
	name := "deux"
	var result int
	for _, j := range sortedRoll {
		if j == 2 {
			result += 2
		}
	}
	return result, name
}

func trois(sortedRoll []int) (int, string) {
	name := "trois"
	var result int
	for _, j := range sortedRoll {
		if j == 3 {
			result += 3
		}
	}
	return result, name
}

func quatre(sortedRoll []int) (int, string) {
	name := "quatre"
	var result int
	for _, j := range sortedRoll {
		if j == 4 {
			result += 4
		}
	}
	return result, name
}

func cinq(sortedRoll []int) (int, string) {
	name := "cinq"
	var result int
	for _, j := range sortedRoll {
		if j == 5 {
			result += 5
		}
	}
	return result, name
}

func six(sortedRoll []int) (int, string) {
	name := "six"
	var result int
	for _, j := range sortedRoll {
		if j == 6 {
			result += 6
		}
	}
	return result, name
}

func petiteSuite(sortedRoll []int) (int, string) {
	name := "Tiote suite"
	if reflect.DeepEqual(sortedRoll, []int{5, 4, 3, 2, 1}) {
		return 15, name
	}
	return 0, name
}

func grandeSuite(sortedRoll []int) (int, string) {
	name := "grande suite"
	if reflect.DeepEqual(sortedRoll, []int{6, 5, 4, 3, 2}) {
		return 20, name
	}
	return 0, name
}

func chance(sortedRoll []int) (int, string) {
	name := "chance"
	var result int
	for _, val := range sortedRoll {
		result += val
	}
	return result, name
}

func Full(sortedRoll []int) (int, string) {
	name := "Full"

	b, _ := brelan(sortedRoll)
	if b == 0 {
		return 0, name
	}
	val := b / 3
	newRoll := make([]int, len(sortedRoll)-3)
	var index int
	for _, j := range sortedRoll {
		if j != val {
			newRoll[index] = j
			index++
		}
	}
	p, _ := paire(newRoll)
	if p == 0 {
		return 0, name
	}
	return b + p, name
}

type rollResult func([]int) (int, string)

func check(roll []int) string {
	var category = ""
	var max int
	var maxIndex int
	for i, fun := range functions {
		result, name := fun(roll)
		if result > max {
			max = result
			category = name
			maxIndex = i
		}
	}
	if len(functions) > 1 {
		functions = append(functions[:maxIndex], functions[maxIndex+1:]...)
		if reflect.DeepEqual(functions, []int{}) {
			return "chance"
		}
	}

	return category
}
