package operator

func Div(x, y float64) float64 {

	if y <= 0 {
		panic("sarasa")
	}

	return x / y
}
