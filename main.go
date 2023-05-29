package main

import "fmt"

const ebulicaoK = 373

func main() {
	tempK := ebulicaoK
	tempC := ebulicaoK - 273

	fmt.Printf("A temperatura de ebulição da água em °K = %v, e a temperarura de ebulição da água em °C = %v.", tempK, tempC)
}