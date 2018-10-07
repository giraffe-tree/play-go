package main

func main() {

}

func sum(arr ... float64) float64 {
	sum := 0.0
	for _, x := range arr {
		sum += x
	}
	return sum
}
