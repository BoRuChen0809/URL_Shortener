package mypkg_test

import (
	mypkg "URL_Shortener/package"
	"testing"
)

func Test_baser62(t *testing.T) {
	e := mypkg.GetEncodeing()
	e.Set_encode()
	e.Set_decode()
	for i := 1; i <= 18; i++ {
		t.Log(mypkg.ByteArray2String(e.Uint2ByteArray(uint64(i))))
	}
}

func Benchmark_base62(b *testing.B) {
	test_record := make([]string, b.N)
	encode := mypkg.GetEncodeing()
	encode.Set_encode()
	encode.Set_decode()
	for i := 0; i < b.N; i++ {
		test_record[i] = mypkg.ByteArray2String(encode.Uint2ByteArray(uint64(i)))
	}

	for i := 0; i < b.N; i++ {
		num, err := encode.ByteArray2Uint(mypkg.String2ByteArray(test_record[i]))
		if err != nil {
			b.Errorf("ByteArray to Uint64 ERROR : %v\n", err)
		}
		if uint64(i) != num {
			b.Errorf("the results of encoding and decoding are not same, %v : %v\n", i, num)
		}
	}
}
