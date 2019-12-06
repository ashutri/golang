package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Validation(input int64) bool {
  //To convert int64 input into string
	result := strconv.FormatInt(int64(input), 10)
  //To get string as an array of char
	actualInput := strings.Split(result, "")
  //To get the length of the String input
	inputLength := len(actualInput)
	if inputLength < 13 || inputLength > 16 {
		return false
	}
	var finalSum int64
  //To get first digit of the card number
	firstDigit, _ := strconv.ParseInt(actualInput[0], 10, 8)
	secondDigit, _ := strconv.ParseInt(actualInput[1], 10, 8)
  // validation for Visa, Master and Discover cards
	if firstDigit != 4 && firstDigit != 5 && firstDigit != 6 && firstDigit != 3 {

		return false
	}
  //validation for American Express cards 
	if firstDigit == 3 && secondDigit != 7 {
		return false
	}
  //Reduced the iteration to n/2 iterations
	for i := inputLength - 2; i >= 0; i -= 2 {
		currDigit, _ := strconv.ParseInt(actualInput[i], 10, 8)
		nextDigit, _ := strconv.ParseInt(actualInput[i+1], 10, 8)
    //Using luhn algorithm
		currDigit *= 2
		finalSum += (currDigit/10 + currDigit%10 + nextDigit)
    //To add first digit of the input
		if i == 1 {
			finalSum += firstDigit
		}
	}
	return finalSum%10 == 0
}

func main() {
  //to get stdin from the user 
	var input int64
	fmt.Scanf("%d", &input)
  
	isValid := Validation(input)	
	if isValid {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}
