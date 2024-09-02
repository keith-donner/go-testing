package main

import (
	"fmt"
	"or/core"
	"or/o"
)

func test2() {
	pnr := o.Pnr{}
	pnr.AirSegments = append(pnr.AirSegments, o.AirSegment{})
	pnr.AirSegments[0].IsPassive = true
	pnr.AirSegments[0].Carrier = "AC"

	if core.AllMatch(pnr.AirSegments, func(s o.AirSegment) bool {
		return s.IsPassive && s.Carrier == "AC"
	}) {
		fmt.Println("All segments match the criteria")
	} else {
		fmt.Println("Not all segments match the criteria")
	}

	if match := core.AnyMatch(pnr.AirSegments, func(s o.AirSegment) bool {
		return s.Carrier == "AC" || s.Carrier == "BA"
	}); match != nil {
		// Match found
		airSegment := *match
		if airSegment.Carrier == "AC" {
			fmt.Println("Carrier is AC")
		} else {
			fmt.Println("Carrier is BA")
		}
	} else {
		// No match found
		fmt.Println("No matching carrier found")
	}
}
