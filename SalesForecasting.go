/* Sales Forecasting using Go. */

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var num_inputs int = 2
	var num_hidden int = 4
	var num_outputs int = 1
	var num_examples int = 4
	var num_epochs int = 60000
	var learning_rate float64 = 0.01
	var inputs [][]float64 = [][]float64{
		[]float64{0, 0},
		[]float64{0, 1},
		[]float64{1, 0},
		[]float64{1, 1},
	}
	var outputs []float64 = []float64{0, 1, 1, 0}
	var hidden_weights [][]float64 = make([][]float64, num_inputs)
	var hidden_biases []float64 = make([]float64, num_hidden)
	var output_weights [][]float64 = make([][]float64, num_hidden)
	var output_biases []float64 = make([]float64, num_outputs)
	for i := 0; i < num_inputs; i++ {
		hidden_weights[i] = make([]float64, num_hidden)
		for j := 0; j < num_hidden; j++ {
			hidden_weights[i][j] = rand.Float64()
		}
	}
	for i := 0; i < num_hidden; i++ {
		hidden_biases[i] = rand.Float64()
	}
	for i := 0; i < num_hidden; i++ {
		output_weights[i] = make([]float64, num_outputs)
		for j := 0; j < num_outputs; j++ {
			output_weights[i][j] = rand.Float64()
		}
	}
	for i := 0; i < num_outputs; i++ {
		output_biases[i] = rand.Float64()
	}
	for i := 0; i < num_epochs; i++ {
		var example int = rand.Intn(num_examples)
		var input []float64 = inputs[example]
		var output float64 = outputs[example]
		var hidden_layer_output []float64 = make([]float64, num_hidden)
		for j := 0; j < num_hidden; j++ {
			var sum float64 = hidden_biases[j]
			for k := 0; k < num_inputs; k++ {
				sum += input[k] * hidden_weights[k][j]
			}
			hidden_layer_output[j] = sigmoid(sum)
		}
		var output_layer_output []float64 = make([]float64, num_outputs)
		for j := 0; j < num_outputs; j++ {
			var sum float64 = output_biases[j]
			for k := 0; k < num_hidden; k++ {
				sum += hidden_layer_output[k] * output_weights[k][j]
			}
			output_layer_output[j] = sigmoid(sum)
		}
		var output_error []float64 = make([]float64, num_outputs)
		for j := 0; j < num_outputs; j++ {
			output_error[j] = output_layer_output[j] - output
		}
		var hidden_error []float64 = make([]float64, num_hidden)
		for j := 0; j < num_hidden; j++ {
			var sum float64 = 0.0
			for k := 0; k < num_outputs; k++ {
				sum += output_error[k] * output_weights[j][k]
			}
			hidden_error[j] = sum
		}
		for j := 0; j < num_hidden; j++ {
			for k := 0; k < num_outputs; k++ {
				output_weights[j][k] -= learning_rate * output_error[k] * hidden_layer_output[j]
			}
		}
		for j := 0; j < num_outputs; j++ {
			output_biases[j] -= learning_rate * output_error[j]
		}
		for j := 0; j < num_inputs; j++ {
			for k := 0; k < num_hidden; k++ {
				hidden_weights[j][k] -= learning_rate * hidden_error[k] * input[j]
			}
		}
		for j := 0; j < num_hidden; j++ {
			hidden_biases[j] -= learning_rate * hidden_error[j]
		}
	}
	for i := 0; i < num_examples; i++ {
		var input []float64 = inputs[i]
		var output float64 = outputs[i]
		var hidden_layer_output []float64 = make([]float64, num_hidden)
		for j := 0; j < num_hidden; j++ {
			var sum float64 = hidden_biases[j]
			for k := 0; k < num_inputs; k++ {
				sum += input[k] * hidden_weights[k][j]
			}
			hidden_layer_output[j] = sigmoid(sum)
		}
		var output_layer_output []float64 = make([]float64, num_outputs)
		for j := 0; j < num_outputs; j++ {
			var sum float64 = output_biases[j]
			for k := 0; k < num_hidden; k++ {
				sum += hidden_layer_output[k] * output_weights[k][j]
			}
			output_layer_output[j] = sigmoid(sum)
		}
		fmt.Printf("%v %v %v\n", input, output_layer_output, output)
	}
}

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}
