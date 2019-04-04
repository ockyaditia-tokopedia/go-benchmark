package main

import (
	"strconv"
	"testing"
)

var smallStr = "35"
var bigStr = "999999999999999"

func BenchmarkAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val, _ := strconv.Atoi(smallStr)
		_ = val
	}
}

func BenchmarkAtoiParseInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val, _ := strconv.ParseInt(smallStr, 0, 64)
		_ = val
	}
}

func BenchmarkAtoiBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val, _ := strconv.Atoi(bigStr)
		_ = val
	}
}

func BenchmarkAtoiParseIntBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val, _ := strconv.ParseInt(bigStr, 0, 64)
		_ = val
	}
}
