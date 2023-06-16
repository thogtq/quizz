package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "12:12"
	str2 := "12:30"
	fmt.Println("Sol", Solution(str, str2))
}

func roundStart(sMin int, max int) int {
	for {
		if sMin == 0 || sMin == 15 || sMin == 30 || sMin == 45 {
			return sMin
		}
		if sMin == max {
			break
		}
		sMin++
	}
	return sMin
}

func roundEnd(eMinute int, min int) int {
	for {
		if eMinute == 0 || eMinute == 15 || eMinute == 30 || eMinute == 45 {
			return eMinute
		}
		if eMinute == min {
			break
		}
		eMinute--
	}
	return eMinute
}

func Solution(A string, B string) int {
	rounds := 0
	startArr := strings.Split(A, ":")
	endArr := strings.Split(B, ":")
	sHour, _ := strconv.Atoi(startArr[0])
	sMinute, _ := strconv.Atoi(startArr[1])
	eHour, _ := strconv.Atoi(endArr[0])
	eMinute, _ := strconv.Atoi(endArr[1])
	sMax := eMinute
	eMin := sMinute
	if sHour != eHour {
		sMax = 60
		eMin = 0
	}
	sMinute = roundStart(sMinute, sMax)
	if sMinute == 60 {
		sHour++
		sMinute = 0
	}
	eMinute = roundEnd(eMinute, eMin)
	if sHour == eHour && eMinute-sMinute < 15 && eHour < sHour {
		return 0
	}
	fmt.Println(roundStart(sMinute, sMax))
	fmt.Println(sHour, sMinute, eHour, eMinute)
	eHourTemp := eHour
	sHourTemp := sHour
	for i := sHour; i != eHour && eHourTemp < sHourTemp || i <= eHour && eHourTemp >= sHourTemp; {
		fmt.Println("trong vong", sHour, eHour)
		if sHour == eHour {
			rounds += (eMinute - sMinute) / 15
			break
		} else {
			rounds += (60 - sMinute) / 15
			sHour++
			if sHour == 24 {
				sHour = 0
			}
		}
		fmt.Println(rounds, sHour)
	}
	return rounds
}
