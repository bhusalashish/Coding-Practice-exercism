package cars

const (
    costOfCarNotInGroup = 10000
    costOfCarInGroupOfTen = 95000
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return (successRate/100) * float64(productionRate);
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int(CalculateWorkingCarsPerHour(productionRate, successRate)/60.0)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    groupsOfTen := carsCount/10;
    notInGroups := carsCount % 10;
    totalCost := groupsOfTen*costOfCarInGroupOfTen + costOfCarNotInGroup*notInGroups;
    return uint(totalCost);
}
