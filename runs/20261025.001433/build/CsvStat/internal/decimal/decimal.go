package decimal

import (
	"fmt"
	"math/big"
)

// Format writes r with exactly two decimals. The absolute value is rounded half away from zero
// once, from the exact rational value. A negative value that rounds to zero carries no sign.
func Format(r *big.Rat) string {
	abs := new(big.Rat).Abs(r)

	// floor(abs*100 + 1/2) equals floor((200*num + den) / (2*den)) for non-negative abs.
	num := new(big.Int).Mul(abs.Num(), big.NewInt(200))
	num.Add(num, abs.Denom())
	den := new(big.Int).Mul(abs.Denom(), big.NewInt(2))
	hundredths := num.Quo(num, den)

	whole, frac := new(big.Int).QuoRem(hundredths, big.NewInt(100), new(big.Int))
	sign := ""
	if r.Sign() < 0 && hundredths.Sign() != 0 {
		sign = "-"
	}
	return fmt.Sprintf("%s%s.%02d", sign, whole.String(), frac)
}

// Whole writes an integral r as a plain integer with no decimal point.
func Whole(r *big.Rat) string {
	return r.Num().String()
}
