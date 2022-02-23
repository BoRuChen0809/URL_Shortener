package mypkg

import (
	"fmt"
	"math"
)

type Encodeing struct {
	encode []byte
	decode map[byte]int
}

func GetEncodeing() *Encodeing {
	return &Encodeing{}
}

func (e *Encodeing) Set_encode() {
	encode := []byte{112, 77, 122, 65, 109, 70, 84, 108, 67, 103, 101, 111, 72, 66, 57, 71, 116, 115, 98, 56,
		88, 79, 110, 80, 119, 99, 78, 90, 54, 82, 48, 83, 51, 53, 75, 104, 97, 68, 73, 76, 49, 121,
		87, 52, 50, 55, 120, 105, 100, 89, 117, 114, 86, 113, 85, 106, 81, 107, 102, 69, 118, 74}
	e.encode = encode
}

func (e *Encodeing) Set_decode() {
	decode := make(map[byte]int)
	for i, v := range e.encode {
		decode[v] = i
	}
	e.decode = decode
}

func (e *Encodeing) Uint2ByteArray(num uint64) []byte {
	var i uint64
	ans := []byte{}
	for {
		i = num
		if i < 62 {
			ans = append(ans, e.encode[i])
			if len(ans) < 6 {
				for i := len(ans); i < 6; i++ {
					ans = append(ans, e.encode[0])
				}
			}
			return ans
		}
		num, i = num/62, num%62
		ans = append(ans, e.encode[i])
	}
}

func (e *Encodeing) ByteArray2Uint(b []byte) (uint64, error) {
	var ans uint64
	for i, v := range b {
		num, ok := e.decode[v]
		if !ok {
			return 0, fmt.Errorf("解碼錯誤")
		}
		ans += uint64(math.Pow(62, float64(i))) * uint64(num)
	}
	return ans, nil
}

func ByteArray2String(b []byte) string {
	return string(b[:])
}

func String2ByteArray(s string) []byte {
	return []byte(s)
}
