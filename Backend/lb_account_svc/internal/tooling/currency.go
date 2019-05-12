package tooling

import (
	"encoding/binary"
	"fmt"
	"math"
)



// String returns a formatted USD value
func String(m int) string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("$%.2f", x)
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}