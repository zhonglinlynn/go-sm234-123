// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
	for i := 0; i <= 10000000; i++ {
		hw.Reset()
		hw.Write(msg)
		hw.Sum(nil)
		//fmt.Printf("sm3: ")
		//fmt.Printf("%s\n", byteToString(hash))
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}

func Test11Sm3(t *testing.T) {
	msg := []byte("abc")
	hw := New()
	start := time.Now()
	for i := 0; i <= 10; i++ {
		hw.Reset()
		hw.Write(msg)
		hash := hw.Sum(nil)
		fmt.Printf("sm3: ")
		fmt.Printf("%s\n", byteToString(hash))
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
