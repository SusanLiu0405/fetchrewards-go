package services

import (
	"errors"
	"fetchrewards-go/models"
	"github.com/google/uuid"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var receiptStore sync.Map

func ProcessReceipt(receipt models.Receipt) (models.ReceiptIdResponse, error) {
	// Validate receipt structure
	if err := validateReceipt(receipt); err != nil {
		return models.ReceiptIdResponse{}, err
	}

	// Validate PurchaseDate format
	_, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		log.Printf("Invalid date format: %s", err)
		return models.ReceiptIdResponse{}, errors.New("400 bad request: invalid date format")
	}

	// Validate PurchaseTime format
	_, err = time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		log.Printf("Invalid time format: %s", err)
		return models.ReceiptIdResponse{}, errors.New("400 bad request: invalid time format")
	}

	id := uuid.New()
	points := calculatePoints(receipt)
	response := models.ReceiptResponse{ID: id, Points: points}

	receiptStore.Store(id, response)
	return models.ReceiptIdResponse{ID: id}, nil
}

// validateReceipt checks if the receipt contains all required fields
// Returns an error if any field is missing or invalid
func validateReceipt(receipt models.Receipt) error {
	if receipt.Retailer == "" {
		return errors.New("400 bad request: missing retailer")
	}
	if receipt.PurchaseDate == "" {
		return errors.New("400 bad request: missing purchaseDate")
	}
	if receipt.PurchaseTime == "" {
		return errors.New("400 bad request: missing purchaseTime")
	}
	if receipt.Total == "" {
		return errors.New("400 bad request: missing total")
	}
	if len(receipt.Items) == 0 {
		return errors.New("400 bad request: missing items")
	}

	// Validate each item
	for i, item := range receipt.Items {
		if item.ShortDescription == "" {
			return errors.New("400 bad request: missing shortDescription in item at index " + strconv.Itoa(i))
		}
		if item.Price == "" {
			return errors.New("400 bad request: missing price in item at index " + strconv.Itoa(i))
		}
	}

	return nil
}

func GetReceiptPoints(id uuid.UUID) (int, error) {
	if response, ok := receiptStore.Load(id); ok {
		if resp, ok := response.(models.ReceiptResponse); ok {
			return resp.Points, nil
		}
		return 0, errors.New("type assertion failed")
	}
	return 0, errors.New("receipt not found")
}

func calculatePoints(receipt models.Receipt) int {
	points := 0
	// Rule 1: One point for every alphanumeric character in the retailer name
	alphaNumRegex := regexp.MustCompile(`[a-zA-Z0-9]`)
	matches := alphaNumRegex.FindAllStringIndex(receipt.Retailer, -1)
	points += len(matches)
	log.Printf("Rule 1: Added %d points for alphanumeric characters in the retailer name.", len(matches))

	// Parse total as float, for mathematical operations
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		log.Printf("Error parsing total amount: %s", err)
		return points
	}

	remainderDollar := math.Mod(total, 1)
	remainderQuarter := math.Mod(total, 0.25)

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	if remainderDollar == 0 {
		points += 50
		log.Printf("Rule 2: Added 50 points for round dollar amount.")
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if remainderQuarter == 0 {
		points += 25
		log.Printf("Rule 3: Added 25 points for multiple of 0.25.")
	}

	// Rule 4: 5 points for every two items on the receipt
	pairPoints := (len(receipt.Items) / 2) * 5
	points += pairPoints
	log.Printf("Rule 4: Added %d points for every two items on the receipt.", pairPoints)

	// Rule 5: Additional points for item descriptions
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				log.Printf("Error parsing item price: %s", err)
				continue
			}
			itemPoints := int(math.Ceil(price * 0.2))
			points += itemPoints
			log.Printf("Rule 5: Added %d points for item description '%s'.", itemPoints, item.ShortDescription)
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		log.Printf("Error parsing purchase date: %s", err)
		return points
	}
	if purchaseDate.Day()%2 != 0 {
		points += 6
		log.Printf("Rule 6: Added 6 points for odd purchase day.")
	}

	// Rule 7: 10 points if the time of purchase is after 2:00 PM and before 4:00 PM
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		log.Printf("Error parsing purchase time: %s", err)
		return points
	}
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
		log.Printf("Rule 7: Added 10 points for purchase time between 2:00 PM and 4:00 PM.")
	}
	log.Printf("Total points: %d", points)

	return points
}
