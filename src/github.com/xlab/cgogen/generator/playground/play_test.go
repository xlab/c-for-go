package main

import "testing"

func BenchmarkToC(b *testing.B) {
	names := []string{"Maxim", "Xlab", "Yo", "Lol"}
	d := Data{Names: names, Size: len(names)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		passData(d)
	}
	b.ReportAllocs()
}
