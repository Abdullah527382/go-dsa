package main

import (
    "fmt"
	"strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

 func timeConversion(s string) string {
    // Write your code here
	// The last 2 chars AM or PM
	var suffix = s[len(s)-2:len(s)]
	// The first 2 characters
	var prefix, err = strconv.Atoi(s[0:2])

	if err != nil {
        // ... handle error
        panic(err)
    }

	// Remaining chars
	var middle = s[2:8]

	var time string = ""
	var factor = 0
	if (suffix == "PM"){
		// Consider edge case of 12+pm
		if (prefix != 12){
			factor = 12
		}
		time = strconv.Itoa(prefix + factor) + middle
	} else {
		if (prefix == 12){
			factor = 12
		}
		time = strconv.Itoa(prefix - factor) + middle
		if (prefix == 12 || prefix < 10){
			time = "0" + time
		}
	}
	return time
}

func main() {


    var stringPM string = "12:45:54PM"
    var stringAM string = "11:20:00AM"

    var resultPM = timeConversion(stringPM)
	var resultAM = timeConversion(stringAM)

    fmt.Printf("The result of time PM string %s\n", resultPM)
    fmt.Printf("The result of time AM string %s\n", resultAM)
}
