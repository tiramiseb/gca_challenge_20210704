package main

import (
	"testing"
)

func BenchmarkFindValue_1000x1000(b *testing.B) {
	h := 1000
	w := 1000
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue_100x100(b *testing.B) {
	h := 100
	w := 100
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue_20x20(b *testing.B) {
	h := 20
	w := 20
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue_3x5(b *testing.B) {
	h := 3
	w := 5
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue2_1000x1000(b *testing.B) {
	h := 1000
	w := 1000
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue2(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue2_100x100(b *testing.B) {
	h := 100
	w := 100
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue2(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue2_20x20(b *testing.B) {
	h := 20
	w := 20
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue2(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue2_3x5(b *testing.B) {
	h := 3
	w := 5
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue2(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue3_1000x1000(b *testing.B) {
	h := 1000
	w := 1000
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue3(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue3_100x100(b *testing.B) {
	h := 100
	w := 100
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue3(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue3_20x20(b *testing.B) {
	h := 20
	w := 20
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue3(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue3_3x5(b *testing.B) {
	h := 3
	w := 5
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue3(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue4_1000x1000(b *testing.B) {
	h := 1000
	w := 1000
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue4(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue4_100x100(b *testing.B) {
	h := 100
	w := 100
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue4(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue4_20x20(b *testing.B) {
	h := 20
	w := 20
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue4(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue4_3x5(b *testing.B) {
	h := 3
	w := 5
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue4(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue5_1000x1000(b *testing.B) {
	h := 1000
	w := 1000
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue5(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue5_100x100(b *testing.B) {
	h := 100
	w := 100
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue5(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue5_20x20(b *testing.B) {
	h := 20
	w := 20
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue5(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue5_3x5(b *testing.B) {
	h := 3
	w := 5
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue5(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue6_1000x1000(b *testing.B) {
	h := 1000
	w := 1000
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue6(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue6_100x100(b *testing.B) {
	h := 100
	w := 100
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue6(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue6_20x20(b *testing.B) {
	h := 20
	w := 20
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue6(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue6_3x5(b *testing.B) {
	h := 3
	w := 5
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue6(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue7_1000x1000(b *testing.B) {
	h := 1000
	w := 1000
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue7(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue7_100x100(b *testing.B) {
	h := 100
	w := 100
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue7(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue7_20x20(b *testing.B) {
	h := 20
	w := 20
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue7(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue7_3x5(b *testing.B) {
	h := 3
	w := 5
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue7(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue8_1000x1000(b *testing.B) {
	h := 1000
	w := 1000
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue8(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue8_100x100(b *testing.B) {
	h := 100
	w := 100
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue8(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue8_20x20(b *testing.B) {
	h := 20
	w := 20
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue8(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue8_3x5(b *testing.B) {
	h := 3
	w := 5
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue8(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue8_10000x10000(b *testing.B) {
	h := 10000
	w := 10000
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue8(w, h, r, c)
			}
		}
	}
}

func BenchmarkFindValue8_100000x100000(b *testing.B) {
	h := 100000
	w := 100000
	for i := 0; i < b.N; i++ {
		for r := 1; r <= h; r++ {
			for c := 1; c <= w; c++ {
				findValue8(w, h, r, c)
			}
		}
	}
}