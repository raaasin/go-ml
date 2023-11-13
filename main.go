// main.go
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

var model *LinearRegression

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

func main() {
	model = NewLinearRegression()

	// Handle requests to the root URL ("/")
	http.HandleFunc("/", homePage)

	// Handle requests to test the trained model
	http.HandleFunc("/test", testModel)

	// Start the web server on port 8080
	fmt.Println("Server listening on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
