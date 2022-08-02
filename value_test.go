package interface_value

import (
	"fmt"
	"testing"
)

func TestInt(t *testing.T) {
	fmt.Println(Int(int(256), 32))
	fmt.Println(Int(int32(128), 32))
	fmt.Println(Int(int16(64), 32))
	fmt.Println(Int(int8(32), 32))

	fmt.Println(Int(uint32(512), 16))
	fmt.Println(Int(uint(512), 8))
}
