// handlers.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// PageVariables struct to hold variables used in the template
type PageVariables struct {
	Number       string
	TrainingData []DataPoint
}

// DataPoint struct to represent a data point for training the model
type DataPoint struct {
	X float64
	Y float64
}

// homePage function to handle the root URL
func homePage(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	r.ParseForm()

	// Check if the request is a POST request with training data
	if r.Method == http.MethodPost {
		xValues := r.Form["x[]"]
		yValues := r.Form["y[]"]

		var trainingData []DataPoint

		for i := 0; i < len(xValues); i++ {
			x, errX := strconv.ParseFloat(xValues[i], 64)
			y, errY := strconv.ParseFloat(yValues[i], 64)

			if errX == nil && errY == nil {
				// Add the data point to the training data
				point := DataPoint{X: x, Y: y}
				trainingData = append(trainingData, point)
			}
		}

		if len(trainingData) > 0 {
			// Train the model with all the collected training data
			model.Train(trainingData)
		}

		renderTemplate(w, "index", PageVariables{Number: "", TrainingData: trainingData})
		return
	}

	// If it's not a POST request, render the default template
	renderTemplate(w, "index", PageVariables{Number: "", TrainingData: nil})
}

// testModel function to handle testing the trained model
func testModel(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	r.ParseForm()

	// Get the value of the "number" parameter from the form
	numberStr := r.Form.Get("number")
	number, err := strconv.ParseFloat(numberStr, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Use the trained model to predict the output for the given input
	predicted := model.Predict(number)

	// Render the HTML template with the predicted value
	renderTemplate(w, "index", PageVariables{Number: fmt.Sprintf("Prediction: %.2f", predicted), TrainingData: nil})
}

// renderTemplate function to render HTML templates
func renderTemplate(w http.ResponseWriter, tmpl string, data PageVariables) {
	t, err := template.New(tmpl + ".html").ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
