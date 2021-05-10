package tiker

import "testing"

func BenchmarkLongTimeTickDo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongTimeTickDo()
	}
}

func BenchmarkLongTimeTickDoWithoutDefault(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongTimeTickDoWithoutDefault()

	}
}

func BenchmarkLongTimeTickDoWithSleep(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongTimeTickDoWithSleep()
	}
}
