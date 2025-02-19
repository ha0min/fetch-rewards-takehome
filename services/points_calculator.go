package services

import "fetch-rewards-takehome/models"

type PointsCalculator struct {
	rules []Rule
}

func NewPointsCalculator() *PointsCalculator {
	return &PointsCalculator{
		rules: []Rule{
			&RetailerNameRule{},
			&TwoItemsRule{},
			&ItemDescriptionRule{},
			&OddDayRule{},
			&TimeRule{},
			&RoundDollarRule{},
		},
	}
}

func (pc *PointsCalculator) CalculatePoints(receipt *models.Receipt) int {
	points := 0

	for _, rule := range pc.rules {
		points += rule.Calculate(receipt)
	}

	return points
}
