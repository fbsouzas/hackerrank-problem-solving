package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(timeConversion("01:05:45AM"))
	fmt.Println(timeConversion("05:05:45AM"))
	fmt.Println(timeConversion("10:05:45AM"))
	fmt.Println(timeConversion("12:05:45PM"))
	fmt.Println(timeConversion("01:05:45PM"))
	fmt.Println(timeConversion("04:05:45PM"))
	fmt.Println(timeConversion("07:05:45PM"))
	fmt.Println(timeConversion("11:05:45PM"))
	fmt.Println(timeConversion("12:05:45AM"))
}

func timeConversion(s string) string {
	splitted := strings.Split(s, ":")
	hour, _ := strconv.Atoi(splitted[0])

	if strings.Contains(s, "PM") {
		if hour != 12 {
			splitted[0] = strconv.Itoa(hour + 12)
		}
		splitted[2] = strings.Replace(splitted[2], "PM", "", 1)
	}

	if strings.Contains(s, "AM") {
		if hour == 12 {
			splitted[0] = "0" + strconv.Itoa(hour-12)
		}
		splitted[2] = strings.Replace(splitted[2], "AM", "", 1)
	}

	return strings.Join(splitted, ":")
}
