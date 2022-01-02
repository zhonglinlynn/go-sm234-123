package sm3

import (
	"fmt"
	"testing"
	"time"
)

func byteToString(b []byte) string {
	ret := ""
	for i := 0; i < len(b); i++ {
		ret += fmt.Sprintf("%02x", b[i])
	}
	//fmt.Println(ret)
	return ret
}
func TestSm3(t *testing.T) {

	msg := []byte("abc")
	hw := New()
	start := time.Now()
	for i := 0; i <= 1000; i++ {
		hw.Write(msg)
		hash := hw.Sum(nil)
		fmt.Printf("sm3: ")
		fmt.Printf("%s\n", byteToString(hash))
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}

func Test11Sm3(t *testing.T) {
	msg := []byte("abc")
	hw := New()

	start := time.Now()
	for i := 0; i <= 10000000; i++ {
		hw.Reset()
		hw.Write(msg)
		hw.Sum(nil)
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}

func BenchmarkSm3(t *testing.B) {
	t.ReportAllocs()
	msg := []byte("abc")
	hw := New()
	for i := 0; i < t.N; i++ {
		hw.Write(msg)
		hw.Sum(nil)

		Sm3Sum(msg)
	}
}
