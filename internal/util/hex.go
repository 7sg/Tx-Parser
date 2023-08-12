package util

import (
	"fmt"
	"strconv"
	"strings"
)

func IntegerToHex(integer int64) string {
	return fmt.Sprintf("0x%x", integer)
}

func HexToInteger(hex string) (int64, error) {
	gohex := strings.TrimPrefix(hex, "0x")
	intVal, err := strconv.ParseInt(gohex, 16, 64)
	if err != nil {
		return 0, err
	}
	return intVal, nil
}
