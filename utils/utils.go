package utils

import (
	"fmt"
	"log"

	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func FormatDecimalWithCommas(d decimal.Decimal) string {
	str := strings.Split(d.StringFixed(2), ".")
	i, err := strconv.Atoi(str[0])
	if err != nil {
		log.Println(err)
	}
	return strings.Join([]string{Comma(int64(i)), str[1]}, ".")
}
func Comma(v int64) string {
	sign := ""
	if v < 0 {
		sign = "-"
		v = 0 - v
	}

	parts := []string{"", "", "", "", "", "", "", ""}
	j := len(parts) - 1

	for v > 999 {
		parts[j] = strconv.FormatInt(v%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		v = v / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(v))
	return sign + strings.Join(parts[j:], ",")
}

func GetMonth(monthInt int) (time.Month, error) {
	// Validate if the month integer is within the valid range (1-12)
	if monthInt < 1 || monthInt > 12 {
		return 0, fmt.Errorf("invalid month integer: %d", monthInt)
	}

	// Cast the integer to time.Month
	return time.Month(monthInt), nil
}
