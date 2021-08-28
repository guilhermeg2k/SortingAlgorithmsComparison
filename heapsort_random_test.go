package main

import (
	"fmt"
	"testing"
	"time"
)

func TestHeapsortRandom10e4(t *testing.T) {
	v := generateVector(10000, true, false)
	startTime := time.Now()
	PrintMemUsage()
	heapSort(v, len(v))
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestHeapsortRandom10e5(t *testing.T) {
	v := generateVector(100000, true, false)
	startTime := time.Now()
	heapSort(v, len(v))
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestHeapsortRandom10e6(t *testing.T) {
	v := generateVector(1000000, true, false)
	startTime := time.Now()
	heapSort(v, len(v))
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestHeapsortRandom10e7(t *testing.T) {
	v := generateVector(10000000, true, false)
	startTime := time.Now()
	heapSort(v, len(v))
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestHeapsortRandom10e8(t *testing.T) {
	v := generateVector(100000000, true, false)
	startTime := time.Now()
	heapSort(v, len(v))
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

