package strings_vs_buffer

import "testing"

func BenchmarkStringsJoin(b *testing.B) {
	b.StartTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		stringsJoin(data)
	}
}

func BenchmarkBufferJoin(b *testing.B) {
	b.StartTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		stringsJoin(data)
	}
}
