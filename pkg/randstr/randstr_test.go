package randstr

import "testing"

func BenchmarkCharsetAlphanumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Charset(Alphanumeric, 100)
	}
}

func BenchmarkCharsetNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Charset(Numeric, 10)
	}
}
