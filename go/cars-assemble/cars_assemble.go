package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	successRate = successRate / 100
	return int(float64(productionRate) * successRate / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	costPerCar := 10000
	costPerTenCars := 95000 // cost for 10 cars

	groupsOfTen := carsCount / 10
	remainingCars := carsCount % 10

	totalCost := groupsOfTen*costPerTenCars + remainingCars*costPerCar
	return uint(totalCost)
}
