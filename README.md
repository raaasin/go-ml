﻿# Go-ML Web Application

This is a web application built with Go that allows users to perform machine learning tasks, specifically linear regression and deployed with Docker on render. Users can input data points, train the model, and test it with new values.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [Training the Model](#training-the-model)
  - [Testing the Model](#testing-the-model)
- [Docker Support](#docker-support)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites

Make sure you have the following installed on your machine:

- Go (Golang)
- Git

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/raaasin/go-ml
   cd go-ml
   ```
2. Build the application:

   ```bash
   go build -o main .
   ```
3. Run the application:

   ```bash
   ./main
   ```

   The application will be accessible at [http://localhost:8080](http://localhost:8080).

## Usage

### Training the Model

1. Open your web browser and navigate to [http://localhost:8080](http://localhost:8080).
2. Enter data points in the provided form and click "Train" to train the linear regression model.

### Testing the Model

1. After training the model, scroll down to the "Test the Model" section.
2. Enter a value in the input field and click "Test" to see the predicted output.

## Docker Support

To run the application using Docker:

1. Build the Docker image:

   ```bash
   docker build -t go-ml .
   ```
2. Run the Docker container:

   ```bash
   docker run -p 8080:8080 go-ml
   ```

   Access the application at [http://localhost:8080](http://localhost:8080).

## Contributing

Contributions are welcome! Feel free to open issues and pull requests.

## License

This project is licensed under the [MIT License](https://opensource.org/license/mit/).
