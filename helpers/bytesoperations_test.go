package helpers

import (
	"reflect"
	"testing"
)

func TestXorBytes(t *testing.T) {
	tests := []struct {
		a, b, want []byte
	}{
		{[]byte{0x01, 0x02, 0x03}, []byte{0x01, 0x02, 0x03}, []byte{0x00, 0x00, 0x00}},
		{[]byte{0xFF, 0x00, 0xAA}, []byte{0x00, 0xFF, 0x55}, []byte{0xFF, 0xFF, 0xFF}},
	}

	for _, test := range tests {
		got := XorBytes(test.a, test.b)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("XorBytes(%v, %v) = %v; want %v", test.a, test.b, got, test.want)
		}
	}
}

func TestUint16ToBits(t *testing.T) {
	tests := []struct {
		value uint16
		want  []int
	}{
		{0x0000, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{0xFFFF, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
	}

	for _, test := range tests {
		got := Uint16ToBits(test.value)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Uint16ToBits(%v) = %v; want %v", test.value, got, test.want)
		}
	}
}

func TestUint32ToBits(t *testing.T) {
	tests := []struct {
		value uint32
		want  []int
	}{
		{0x00000000, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{0xFFFFFFFF, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
	}

	for _, test := range tests {
		got := Uint32ToBits(test.value)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Uint32ToBits(%v) = %v; want %v", test.value, got, test.want)
		}
	}
}

func TestBitsToInt(t *testing.T) {
	tests := []struct {
		bits []int
		want int
	}{
		{[]int{0, 0, 0, 0}, 0},
		{[]int{1, 1, 1, 1}, 15},
	}

	for _, test := range tests {
		got := BitsToInt(test.bits)
		if got != test.want {
			t.Errorf("BitsToInt(%v) = %v; want %v", test.bits, got, test.want)
		}
	}
}
