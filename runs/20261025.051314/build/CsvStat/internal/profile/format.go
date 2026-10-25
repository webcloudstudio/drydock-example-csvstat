package profile

import (
	"fmt"
	"math/big"
)

// FormatDecimal writes r with exactly two decimal places. The exact value is rounded once, half
// away from zero, using integer arithmetic. A result of zero carries no minus sign.
func FormatDecimal(r *big.Rat) string {
	// Scale |r| by 100 and round the quotient: a remainder of at least half the
	// denominator rounds the magnitude up, which is half away from zero.
	scaled := new(big.Int).Mul(new(big.Int).Abs(r.Num()), big.NewInt(100))
	denom := r.Denom()
	cents, rem := new(big.Int).QuoRem(scaled, denom, new(big.Int))
	if new(big.Int).Lsh(rem, 1).Cmp(denom) >= 0 {
		cents.Add(cents, big.NewInt(1))
	}
	if cents.Sign() == 0 {
		return "0.00"
	}

	whole, fraction := new(big.Int).QuoRem(cents, big.NewInt(100), new(big.Int))
	sign := ""
	if r.Sign() < 0 {
		sign = "-"
	}
	return fmt.Sprintf("%s%s.%02d", sign, whole.String(), fraction.Int64())
}

// FormatInteger writes an integer-valued r as a plain integer, with no decimal point and no
// separator. The caller passes only integer-valued rationals.
func FormatInteger(r *big.Rat) string {
	return r.Num().String()
}
