package main

import (
	"fmt"
	"reflect"
	"testing"
)

var functions = []rollResult{carre, un, deux, petiteSuite, grandeSuite, yahzee}

type Roll []int
type rollResult func(Roll) (int, string)

func testRollResult(t *testing.T, expectedName string, fun rollResult, roll Roll, expectedResult int) {
	result, name := fun(roll)
	if result != expectedResult || name != expectedName {
		println(result)
		fmt.Println(expectedName, " should return ", expectedResult, ",", expectedName)
		t.Fail()
	}
}

func TestRollResults(t *testing.T) {
	var values = []struct {
		expectedName string
		fun rollResult
		roll Roll
		expectedResult int}{
			{"paire", paire, Roll{3, 6, 6, 4, 5}, 12},
			{"doublePaire", doublePaire, Roll{3, 6, 6, 4, 4}, 20},
			{"brelan", brelan, Roll{6, 6, 6, 4, 4}, 18},
			{"carre", carre, Roll{6, 6, 6, 6, 4}, 24},
			{"yahzee", yahzee, Roll{6, 6, 6, 6, 6}, 50},
			{"trois", trois, Roll{3, 2, 2, 2, 2}, 3},
			{"Tiote suite", petiteSuite, Roll{5, 4, 3, 2, 1}, 15},
			{"full", full, Roll{5, 5, 5, 3, 3}, 21},
		}
	for _, val := range values {
		testRollResult(t, val.expectedName, val.fun, val.roll, val.expectedResult)
	}
}

func TestCheck(t *testing.T) {
	roll := []int{6, 5, 4, 1, 1}

	name := check(roll)

	println(name)
}

func sameValuesInRoll(nb int, name string) rollResult {
	return func(sortedRoll Roll) (int, string) {
		same := make([]int, nb-1)

	for _, val := range sortedRoll {
		if  same[0]!= 0 && sameValues(same){
				return nb * val, name
			}
		push(same, val)
	}
	return 0, name}
}

var paire = sameValuesInRoll(2, "paire")
var brelan = sameValuesInRoll(3, "brelan")
var carre = sameValuesInRoll(4, "carre")


func push(tab []int, v int) {
	for i := 1; i < len(tab); i++ {
		tab[i-1] = tab[i]
	}
	tab[len(tab)-1] = v
}

func sum(tab []int) int {
	var result int
	for _, val := range tab {
		result += val
	}
	return result
}

func sameValues(tab []int) bool {
	v := tab[0]
	for _, val := range tab {
		if v != val {
			return false
		}
	}
	return true
}

func yahzee(sortedRoll Roll) (int, string) {
	name := "yahzee"

	p := sortedRoll[0]
	for _, val := range sortedRoll {
		if p != val {
			return 0, name
		}
	}
	return 50, name
}


func unit(nb int, name string) rollResult {
	return func(roll Roll) (int, string) {
		var result int
		for _, v := range roll {
			if v == nb {
				result = result + nb
			}
		}
		return result,name
	}
}

var un = unit(1, "un")
var deux = unit(2, "deux")
var trois = unit(3, "trois")
var quatre = unit(4, "quatre")
var cinq = unit(5, "cinq")
var six = unit(6, "six")

func sameRoll(roll Roll, name string, score int) rollResult {
	return func(nr Roll) (int, string) {
		if (reflect.DeepEqual(roll, nr)){
			return score, name
		}
		return 0, name
	}
}

var petiteSuite = sameRoll(Roll{5,4,3,2,1},"Tiote suite", 15)
var grandeSuite = sameRoll(Roll{6,5,4,3,2},"grande suite", 20)

func chance(sortedRoll Roll) (int, string) {
	name := "chance"
	var result int
	for _, val := range sortedRoll {
		result += val
	}
	return result, name
}

func full(sortedRoll Roll) (int, string) {
	name := "full"

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

func doublePaire(sortedRoll Roll) (int, string) {
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

func check(roll Roll) string {
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
