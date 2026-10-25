package decimal

import (
	"fmt"
	"math/big"
)

// Format writes r with exactly two decimal places. It rounds once, from the exact
// value, half away from zero. A minus sign appears only when the rounded value is
// negative, so a value that rounds to zero is written 0.00.
func Format(r *big.Rat) string {
	// q = floor((2n + d) / 2d) is |r| * 100 rounded half up, where n/d = |r| * 100.
	scaled := new(big.Rat).Mul(new(big.Rat).Abs(r), big.NewRat(100, 1))
	numerator := new(big.Int).Lsh(scaled.Num(), 1)
	numerator.Add(numerator, scaled.Denom())
	denominator := new(big.Int).Lsh(scaled.Denom(), 1)
	q := new(big.Int).Quo(numerator, denominator)

	whole, cents := new(big.Int).QuoRem(q, big.NewInt(100), new(big.Int))
	sign := ""
	if r.Sign() < 0 && q.Sign() != 0 {
		sign = "-"
	}
	return fmt.Sprintf("%s%s.%02d", sign, whole, cents.Int64())
}
