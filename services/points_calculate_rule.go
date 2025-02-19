package services

import (
	"fetch-rewards-takehome/models"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type Rule interface {
	Calculate(receipt *models.Receipt) int
}

// RetailerNameRule, adds 1 point for every alphanumeric character in the retailer name.
type RetailerNameRule struct{}

func (r *RetailerNameRule) Calculate(receipt *models.Receipt) int {
	// check the alphanumeric characters in the retailer name
	points := 0

	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}

	fmt.Printf("[INFO] RetailerNameRule | Retailer: %s | Points: %d\n", receipt.Retailer, points)
	return points
}

// RoundDollarRule, adds 50 points if the total is a round dollar amount with no cents.
type RoundDollarRule struct{}

func (r *RoundDollarRule) Calculate(receipt *models.Receipt) int {
	// check if the total is a round dollar amount
	points := 0

	// change the string to a float
	floatTotal, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		// error
		fmt.Println("[ERROR] RoundDollarRule | ParseFloat: ", err.Error())
		return 0
	}

	if floatTotal == float64(int(floatTotal)) {
		points += 50
	}
	fmt.Printf("[INFO] RoundDollarRule | Total: %s | Points: %d\n", receipt.Total, points)
	return points
}

// MultipleOfQuarterRule, adds 25 points if the total is a multiple of 0.25.
type MultipleOfQuarterRule struct{}

func (r *MultipleOfQuarterRule) Calculate(receipt *models.Receipt) int {
	// check if the total is a multiple of 0.25
	points := 0
	// change the string to a float
	floatTotal, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		// error
		fmt.Println("[ERROR] MultipleOfQuarterRule | ParseFloat: ", err.Error())
		return 0
	}

	if floatTotal/0.25 == float64(int(floatTotal/0.25)) {
		points += 25
	}

	fmt.Printf("[INFO] MultipleOfQuarterRule | Total: %s | Points: %d\n", receipt.Total, points)
	return points
}

// TwoItemsRule, adds 5 points for every two items on the receipt.
type TwoItemsRule struct{}

func (r *TwoItemsRule) Calculate(receipt *models.Receipt) int {
	// check if the total is a multiple of 0.25
	points := 0

	if len(receipt.Items)/2 > 0 {
		points += 5 * (len(receipt.Items) / 2)
	}
	fmt.Printf("[INFO] TwoItemsRule | Length: %d | Points: %d\n", len(receipt.Items), points)
	return points
}

// ItemDescriptionRule, adds a point for every description containing a trimmed length of a multiple of 3.
type ItemDescriptionRule struct{}

func (r *ItemDescriptionRule) Calculate(receipt *models.Receipt) int {
	// check the length of the item description
	points := 0
	for _, item := range receipt.Items {
		// handling the tailing and leading spaces
		item.ShortDescription = strings.TrimSpace(item.ShortDescription)
		if len(item.ShortDescription)%3 == 0 {
			// multiply the price by 0.2
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				// error
				fmt.Println("[ERROR] ItemDescriptionRule | Price: ", err.Error())
				return 0
			}

			// round up to the nearest integer
			addedPoints := int(math.Ceil(price * 0.2))
			// fmt.Printf("[INFO] ItemDescriptionRule | Price: %s | Added Points: %d\n", item.Price, addedPoints)
			points += addedPoints
		}
	}
	fmt.Printf("[INFO] ItemDescriptionRule | Points: %d\n", points)
	return points
}

// OddDayRule, add 6 points if the day in the purchase date is odd.
type OddDayRule struct{}

func (r *OddDayRule) Calculate(receipt *models.Receipt) int {
	// check if the day in the purchase date is odd
	points := 0
	if receipt.PurchaseDate[9]%2 == 1 {
		points += 6
	}
	fmt.Printf("[INFO] OddDayRule | Day: %s | Points: %d\n", receipt.PurchaseDate, points)
	return points
}

// TimeRule, add 10 points for purchases after 2:00pm and before 4:00pm.
type TimeRule struct{}

func (r *TimeRule) Calculate(receipt *models.Receipt) int {
	// check if the purchase time is after 2:00pm and before 4:00pm
	points := 0
	// check hour
	hour, err := strconv.Atoi(receipt.PurchaseTime[:2])
	if err != nil {
		// error
		fmt.Println("[ERROR] TimeRule | Check Hour: ", err.Error())
		return 0
	}

	// check minute
	minute, err := strconv.Atoi(receipt.PurchaseTime[3:])
	if err != nil {
		// error
		fmt.Println("[ERROR] TimeRule | Check Minute: ", err.Error())
		return 0
	}

	if hour >= 14 && hour < 16 {
		if hour == 14 && minute == 0 {
			points += 0
		} else if hour == 16 && minute == 0 {
			points += 0
		} else {
			points += 10
		}
	}
	fmt.Printf("[INFO] TimeRule | Time: %s | Points: %d\n", receipt.PurchaseTime, points)
	return points
}
