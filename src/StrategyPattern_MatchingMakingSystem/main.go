package main

import (
	"fmt"
	"strategy_pattern/matching"
)

func main() {
	me := matching.Individual{ /* 這裡要帶入值 */ }
	others := []matching.Individual{matching.Individual{ /* 這裡要帶入值 */ }, matching.Individual{ /* 這裡要帶入值 */ }, matching.Individual{ /* 這裡要帶入值 */ }}

	system := &matching.MatchingSystem{}

	distAlgo := &matching.DistanceBasedStrategy{}
	system.SetMatchingAlog(distAlgo)
	distResult := system.Match(me, others)
	fmt.Println(distResult)

	habitAlgo := &matching.DistanceBasedStrategy{}
	system.SetMatchingAlog(habitAlgo)
	habitResult := system.Match(me, others)
	fmt.Println(habitResult)

}
