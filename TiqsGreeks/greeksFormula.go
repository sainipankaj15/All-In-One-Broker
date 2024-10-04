package tiqs_greeks_socket

import (
	"fmt"
	"math"
)

// tThreshold is the minimum time to expiry (in years) below which special handling is applied
const tThreshold = 1e-4

// black76Call calculates the price of a call option using the Black-76 model
func black76Call(S, K, T, r, sigma float64) float64 {
	if T <= tThreshold {
		// For very short-term options, use intrinsic value
		return math.Max(0, S-K)
	}

	// Calculate d1 and d2
	d1 := (math.Log(S/K) + (0.5*sigma*sigma)*T) / (sigma * math.Sqrt(T))
	d2 := d1 - sigma*math.Sqrt(T)

	// Define the cumulative distribution function
	cdf := func(x float64) float64 { return 0.5 * (1 + math.Erf(x/math.Sqrt(2))) }

	// Calculate and return the option price
	return S*cdf(d1) - K*math.Exp(-r*T)*cdf(d2)
}

// black76ImpliedVol calculates the implied volatility using the Black-76 model
func black76ImpliedVol(S, K, T, r, price float64) float64 {
	if T <= tThreshold {
		// For very short-term options, use a simple approximation
		if price <= S-K {
			return 0.0
		} else {
			return 1.0
		}
	}

	// Define the Black-76 call price function
	black76CallPrice := func(sigma float64) float64 {
		return black76Call(S, K, T, r, sigma)
	}

	// Initial values for the bisection method
	a := 0.01
	b := 2.0

	// Use numerical optimization to find the implied volatility
	impliedVol, err := bisection(func(x float64) float64 { return black76CallPrice(x) - price }, a, b, 1e-8)
	if err != nil {
		// Return 0 if bisection fails
		return 0 
	}
	return impliedVol
}

// black76Greeks calculates the Greeks (delta, theta, gamma, vega) using the Black-76 model
func black76Greeks(S, K, T, r, sigma float64) (delta, theta, gamma, vega float64) {
	if T <= tThreshold {
		// For very short-term options, set all Greeks to 0
		return 0.0, 0.0, 0.0, 0.0
	}

	// Calculate d1
	d1 := (math.Log(S/K) + (0.5*sigma*sigma)*T) / (sigma * math.Sqrt(T))

	// Define probability density function and cumulative distribution function
	pdf := func(x float64) float64 { return math.Exp(-x*x/2) / math.Sqrt(2*math.Pi) }
	cdf := func(x float64) float64 { return 0.5 * (1 + math.Erf(x/math.Sqrt(2))) }

	// Calculate Greeks
	delta = cdf(d1)
	theta = (-(S*pdf(d1)*sigma)/(2*math.Sqrt(T)) - r*K*math.Exp(-r*T)*cdf(d1)) / 365
	gamma = pdf(d1) / (S * sigma * math.Sqrt(T))
	vega = S * pdf(d1) * math.Sqrt(T) / 100

	return delta, theta, gamma, vega
}

// bisection implements the bisection method to find the root of a function
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
		} else {
			a = c
			fa = fc
		}
	}
	return (a + b) / 2, nil
}
