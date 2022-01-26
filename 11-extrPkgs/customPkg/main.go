package main

import (
	"github.com/ncostamagna/matrix/lineal"
	"github.com/ncostamagna/matrix/matrix"
)

func main() {
	input := matrix.New([]float64{2, 2, 2},
		[]float64{3, 3, 4},
		[]float64{4, 4, 2},
		[]float64{2, 5, 6},
		[]float64{1, 7, 3})

	target := matrix.New([]float64{7}, []float64{14}, []float64{21}, []float64{22}, []float64{30})

	lr := lineal.NewLR(*input, *target)
	lr.LearningRate(0.001)
	lrModel := lr.Train(100000)
	lrModel.Print()

}
