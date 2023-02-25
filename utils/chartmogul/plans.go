package chartmogul

import (
	"fmt"

	cm "github.com/chartmogul/chartmogul-go/v3"
)

// GetMonthlyPlans returns all the monthly plans from ChartMogul API
func GetMonthlyPlans() []*cm.Plan {
	// Get all the plans from ChartMogul API
	plans, err := api.ListPlans(&cm.ListPlansParams{})

	// Check for errors
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var monthlyPlans []*cm.Plan

	// Iterate over the plans
	for _, plan := range plans.Plans {
		if plan.IntervalCount == 1 && plan.IntervalUnit == "month" {
			monthlyPlans = append(monthlyPlans, plan)
		}
	}

	return monthlyPlans
}
