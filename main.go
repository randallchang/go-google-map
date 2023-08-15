package main

import (
	"fmt"
	"math"
)

const R = 6371000

func main() {

	lonA := 25.049291680277307
	latA := 121.54540711534234
	lonB := 25.04734583519809
	latB := 121.5436348961113

	len := Sphere(lonA,
		latA,
		lonB,
		latB)

	fmt.Print(len)
}

func Sphere(lonA, latA, lonB, latB float64) float64 {
	c := math.Sin(latA)*math.Sin(latB)*math.Cos(lonA-lonB) + math.Cos(latA)*math.Cos(latB)
	return R * math.Acos(c) * math.Pi / 180
}
