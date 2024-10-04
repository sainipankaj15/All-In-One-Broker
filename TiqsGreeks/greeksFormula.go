package tiqs_greeks_socket

import (
	"fmt"
	"math"
)

const tThreshold = 1e-4 // Threshold for T (in years)

func black76Call(S, K, T, r, sigma float64) float64 {
	if T <= tThreshold {
		// fmt.Println("returning from if conditon")
		return math.Max(0, S-K)
	}
	// fmt.Println("came down")

	// fmt.Println("returning from values from black76ccall", S, K, T, r, sigma)
	d1 := (math.Log(S/K) + (0.5*sigma*sigma)*T) / (sigma * math.Sqrt(T))
	d2 := d1 - sigma*math.Sqrt(T)
	cdf := func(x float64) float64 { return 0.5 * (1 + math.Erf(x/math.Sqrt(2))) }
	finalAnswer := S*cdf(d1) - K*math.Exp(-r*T)*cdf(d2)
	// fmt.Println("returning from black76Call", finalAnswer)
	return finalAnswer
}

func black76ImpliedVol(S, K, T, r, price float64) float64 {
	if T <= tThreshold {
		if price <= S-K {
			return 0.0
		} else {
			return 1.0
		}
	}

	black76CallPrice := func(sigma float64) float64 {
		return black76Call(S, K, T, r, sigma)
	}

	// Initial values for a and b
	a := 0.01
	b := 2.0

	// fa := black76CallPrice(a)
	// fb := black76CallPrice(b)

	// Add logging to check the signs of f(a) and f(b)
	// log.Printf("Initial values: f(a) = %v, f(b) = %v", fa, fb)

	// Use numerical optimization to find the implied volatility
	impliedVol, err := bisection(func(x float64) float64 { return black76CallPrice(x) - price }, a, b, 1e-8)
	if err != nil {
		panic(err)
	}
	return impliedVol
}

func black76Greeks(S, K, T, r, sigma float64) (delta, theta, gamma, vega float64) {
	if T <= tThreshold {
		delta = 0.0
		theta = 0.0
		gamma = 0.0
		vega = 0.0
		return
	}

	d1 := (math.Log(S/K) + (0.5*sigma*sigma)*T) / (sigma * math.Sqrt(T))
	pdf := func(x float64) float64 { return math.Exp(-x*x/2) / math.Sqrt(2*math.Pi) }
	cdf := func(x float64) float64 { return 0.5 * (1 + math.Erf(x/math.Sqrt(2))) }
	delta = cdf(d1)
	theta = (-(S*pdf(d1)*sigma)/(2*math.Sqrt(T)) - r*K*math.Exp(-r*T)*cdf(d1)) / 365
	gamma = pdf(d1) / (S * sigma * math.Sqrt(T))
	vega = S * pdf(d1) * math.Sqrt(T) / 100
	return delta, theta, gamma, vega
}

// bisection implements the bisection method to find the root of a function.
func bisection(f func(float64) float64, a, b, tol float64) (float64, error) {
	fa := f(a)
	if fa == 0 {
		return a, nil
	}
	fb := f(b)
	if fb == 0 {
		return b, nil
	}
	if fa*fb > 0 {
		return 0, fmt.Errorf("bisection: f(a) and f(b) have the same sign")
	}
	for math.Abs(b-a) > tol {
		c := (a + b) / 2
		fc := f(c)
		if fc == 0 {
			return c, nil
		}
		if fa*fc < 0 {
			b = c
			// fb = fc
		} else {
			a = c
			fa = fc
		}
	}
	return (a + b) / 2, nil
}
