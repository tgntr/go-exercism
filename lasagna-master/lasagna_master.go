package lasagna

func PreparationTime(layers []string, averageTime int) int {
	if averageTime == 0 {
		averageTime = 2
	}

	return len(layers) * averageTime
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, ingredient := range layers {
		switch ingredient {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendIngredients, myIngredients []string) []string {
	result := make([]string, len(myIngredients))
	friendLastIngredient := len(friendIngredients) - 1
	copy(result, myIngredients)
	return append(result, friendIngredients[friendLastIngredient:]...)
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	quantityMultiplier := float64(portions) / 2
	scaledQuantities := make([]float64, len(quantities))
	for i, quantity := range quantities {
		scaledQuantities[i] = quantity * quantityMultiplier
	}

	return scaledQuantities
}
