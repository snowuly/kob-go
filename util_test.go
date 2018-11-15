package kob

import "testing"

func BenchmarkPathToRep(b *testing.B) {
	s := "/path/:name/to/:age/xx"
	for i := 0; i < b.N; i++ {
		PathToReg(s)
	}
}
