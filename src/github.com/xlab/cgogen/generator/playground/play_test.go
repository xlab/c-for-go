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

func BenchmarkCube(b *testing.B) {
	cube := [][][]string{
		[][]string{
			[]string{"a1", "b1"},
			[]string{"c1", "d1"},
			[]string{"i1", "j1"},
			[]string{"k1", "l1"},
		},
		[][]string{
			[]string{"a2", "b2"},
			[]string{"c2", "d2"},
			[]string{"i2", "j2"},
			[]string{"k2", "l2"},
		},
		[][]string{
			[]string{"a3", "b3"},
			[]string{"c3", "d3"},
			[]string{"i3", "j3"},
			[]string{"k3", "l3"},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		passCube(cube, 3, 4, 2)
	}
	b.ReportAllocs()
}
