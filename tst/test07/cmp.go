// main.go
package main

func HigherTemperature(temp1, temp2 float64) float64 {
	if temp1 > temp2 {
		return temp1
	}
	return temp2
}

func LowerTemperature(temp1, temp2 float64) float64 {
	if temp1 < temp2 {
		return temp1
	}
	return temp2
}
