package main

import (
	"fmt"
	"testing"
	"time"
)

func TestBubblesortInverted10e4(t *testing.T) {
	v := generateVector(10000, false, true)
	startTime := time.Now()
	bubbleSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestBubblesortInverted10e5(t *testing.T) {
	v := generateVector(100000, false, true)
	startTime := time.Now()
	bubbleSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestBubblesortInverted10e6(t *testing.T) {
	v := generateVector(1000000, false, true)
	startTime := time.Now()
	bubbleSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestBubblesortInverted10e7(t *testing.T) {
	v := generateVector(10000000, false, true)
	startTime := time.Now()
	bubbleSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestBubblesortInverted10e8(t *testing.T) {
	v := generateVector(100000000, false, true)
	startTime := time.Now()
	bubbleSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

