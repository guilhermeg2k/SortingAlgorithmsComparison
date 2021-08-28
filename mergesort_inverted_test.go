package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMergesortInverted10e4(t *testing.T) {
	v := generateVector(10000, false, true)
	startTime := time.Now()
	mergeSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestMergesortInverted10e5(t *testing.T) {
	v := generateVector(100000, false, true)
	startTime := time.Now()
	mergeSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestMergesortInverted10e6(t *testing.T) {
	v := generateVector(1000000, false, true)
	startTime := time.Now()
	mergeSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestMergesortInverted10e7(t *testing.T) {
	v := generateVector(10000000, false, true)
	startTime := time.Now()
	mergeSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestMergesortInverted10e8(t *testing.T) {
	v := generateVector(100000000, false, true)
	startTime := time.Now()
	mergeSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}

func TestMergesortInverted10e9(t *testing.T) {
	v := generateVector(1000000000, false, true)
	startTime := time.Now()
	mergeSort(v)
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
	PrintMemUsage()
}
