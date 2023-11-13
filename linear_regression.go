// linear_regression.go
package main

import (

)

// LinearRegression struct to represent a simple linear regression model
type LinearRegression struct {
	SumX   float64
	SumY   float64
	SumXY  float64
	SumXX  float64
	Count float64
}

// NewLinearRegression function to create a new LinearRegression instance
func NewLinearRegression() *LinearRegression {
	return &LinearRegression{}
}

// Train function to train the linear regression model with data points
func (lr *LinearRegression) Train(data []DataPoint) {
	for _, point := range data {
		lr.SumX += point.X
		lr.SumY += point.Y
		lr.SumXY += point.X * point.Y
		lr.SumXX += point.X * point.X
		lr.Count++
	}
}

// Predict function to predict the output for a given input
func (lr *LinearRegression) Predict(x float64) float64 {
	if lr.Count == 0 {
		return 0
	}

	m := (lr.Count*lr.SumXY - lr.SumX*lr.SumY) / (lr.Count*lr.SumXX - lr.SumX*lr.SumX)
	b := (lr.SumY - m*lr.SumX) / lr.Count

	return m*x + b
}
