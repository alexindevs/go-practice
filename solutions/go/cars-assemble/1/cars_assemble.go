// Package cars gives you tools for calculating the statistics of car production
package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	pct := successRate / float64(100) // should give you a number less than 1, unless of course the success rate is 100. pct should also be of type float64
    return float64(productionRate) * pct // should give you the number of cars produced succesfully 
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsHr := CalculateWorkingCarsPerHour(productionRate, successRate)
    return int(workingCarsHr / float64(60))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOf10 := carsCount / 10
    individiual := carsCount % 10

	return uint((groupsOf10 * 95000) + (individiual * 10000))    
}
