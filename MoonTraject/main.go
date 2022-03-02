package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
)

var (
	cycleCount = 170000
	GVal       = 6.66743e-11
	deltaT     = 540.0
	// Mass
	sunMass   = 1.989 * 1e30
	earthMass = 5.972 * 1e24
	// Distance
	earthMoonDistance = 384400 * 1e3
	sunEarthDistance  = 1.5 * 1e11
	//Acceleration
	earthMoon = GVal * earthMass / (earthMoonDistance * earthMoonDistance)
	sunEarth  = GVal * sunMass / (sunEarthDistance * sunEarthDistance)
	//Velocity
	initMoonVelocity  = math.Sqrt(earthMoon * earthMoonDistance)
	initEarthVelocity = math.Sqrt(sunEarth * sunEarthDistance)
)

type Earth struct {
	posX      float64
	posY      float64
	velocityX float64
	velocityY float64
}

type Moon struct {
	posX      float64
	posY      float64
	velocityX float64
	velocityY float64
}

func csvExport(data [][]string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		if err := writer.Write(value); err != nil {
			return err
		}
	}
	return nil
}

func convertToStringArray(posX float64, posY float64) []string {
	sx := fmt.Sprintf("%f", posX)
	sy := fmt.Sprintf("%f", posY)
	conv := []string{sx, sy}
	return conv
}

func midPoint() {
	var earthPosData [][]string
	var moonPosData [][]string
	var velocityX_2Earh float64
	var velocityY_2Earh float64
	var velocityX_2Moon float64
	var velocityY_2Moon float64
	earth := Earth{0, sunEarthDistance, initEarthVelocity, 0}
	moon := Moon{0, earthMoonDistance, initMoonVelocity, 0}

	i := 0
	for i < cycleCount {
		// Moon
		xMoon := 0 - moon.posX
		yMoon := 0 - moon.posY

		lengthMoon := math.Sqrt(xMoon*xMoon + yMoon*yMoon)

		accxMoon := xMoon / lengthMoon * earthMoon
		accyMoon := yMoon / lengthMoon * earthMoon

		velocityX_2Moon = moon.velocityX + accxMoon*deltaT/2.0
		velocityY_2Moon = moon.velocityY + accyMoon*deltaT/2.0

		moon.posX += velocityX_2Moon * deltaT
		moon.posY += velocityY_2Moon * deltaT

		moon.velocityX += accxMoon * deltaT
		moon.velocityY += accyMoon * deltaT

		// Earth
		xEarth := 0.0 - earth.posX
		yEarth := 0.0 - earth.posY

		lengthEarth := math.Sqrt(xEarth*xEarth + yEarth*yEarth)

		accXEarth := xEarth / lengthEarth * sunEarth
		accYEarth := yEarth / lengthEarth * sunEarth

		velocityX_2Earh = earth.velocityX + accXEarth*deltaT/2.0
		velocityY_2Earh = earth.velocityY + accYEarth*deltaT/2.0

		earth.posX += velocityX_2Earh * deltaT
		earth.posY += velocityY_2Earh * deltaT

		earth.velocityX += accXEarth * deltaT
		earth.velocityY += accYEarth * deltaT

		earthPosData = append(earthPosData, convertToStringArray(earth.posX, earth.posY))
		moonPosData = append(moonPosData, convertToStringArray((moon.posX+earth.posX), (moon.posY+earth.posY)))
		i++
	}
	csvExport(earthPosData, "earthPosData.csv")
	csvExport(moonPosData, "moonPosData.csv")
}

func main() {
	midPoint()
}
