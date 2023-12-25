package main

import (
	"fmt"
	"math"
)

type TemperatureRange struct {
	Min   float64
	Max   float64
	Items []float64
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	rangeTemps := binTemperatures(temps, 10.0)

	fmt.Println(rangeTemps)
}

func binTemperatures(temps []float64, Trange float64) []TemperatureRange {
	minTemp, maxTemp := math.Inf(1), math.Inf(-1) // определяем большую и меньшую температуру
	// обновляем значения в цикле проходя по всем температурам в массиве
	for _, temp := range temps {
		if temp < minTemp {
			minTemp = temp
		}
		if temp > maxTemp {
			maxTemp = temp
		}
	}
	// определчем на сколько диапазонов можно разделить наши температуру по параметру binSize
	numTrange := int(math.Ceil((maxTemp - minTemp) / Trange))
	rangeTemps := make([]TemperatureRange, numTrange)

	// проходимся по каждому элементу массива определяем все подходящие диапазоны и их температуры
	for _, temp := range temps {
		binIndex := int(math.Floor((temp - minTemp) / Trange))
		rangeTemps[binIndex].Min = minTemp + float64(binIndex)*Trange
		rangeTemps[binIndex].Max = minTemp + float64(binIndex+1)*Trange
		rangeTemps[binIndex].Items = append(rangeTemps[binIndex].Items, temp)
	}

	return rangeTemps
}
