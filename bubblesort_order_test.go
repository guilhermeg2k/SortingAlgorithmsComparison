package main

import (
	"fmt"
	"testing"
	"time"
)

func TestBubblesortOrder10e4(t *testing.T) {
	v := generateVector(10000, false, false)
	startTime := time.Now()
	bubbleSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestBubblesortOrder10e5(t *testing.T) {
	v := generateVector(100000, false, false)
	startTime := time.Now()
	bubbleSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestBubblesortOrder10e6(t *testing.T) {
	v := generateVector(1000000, false, false)
	startTime := time.Now()
	bubbleSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestBubblesortOrder10e7(t *testing.T) {
	PrintMemUsage()
	v := generateVector(10000000, false, false)
	startTime := time.Now()
	bubbleSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestBubblesortOrder10e8(t *testing.T) {
	v := generateVector(100000000, false, false)
	startTime := time.Now()
	bubbleSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

