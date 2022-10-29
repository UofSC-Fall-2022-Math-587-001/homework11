package hw10

import (
	"math"
	"github.com/UofSC-Fall-2022-Math-587-001/homework10/library"
)

// "math"
// "github.com/UofSC-Fall-2022-Math-587-001/homework10/library"

// Use Euler's Criteria (Exercise 3.37b) to give an efficient algorithm
// to determine if x^2 = a mod p has a solution for odd p and gcd(a,p) = 1
// In other words, EulerCrit(p,a) determines if a is a quadratic residue
// modulo p
func EulerCrit(p, a int) bool {
	return false 
}

/// GetQuadNonRes(p) returns a quadratic non-residue 
func GetQuadNonRes(p int) int {
	return 0 
}

// TonelliShanks(p,a) returns the solutions of x^2 = a mod p for a prime p
func TonelliShanks(p, a int) (bool,[]int) {
	// If p = 2, then a^2 = a mod 2 so return true, [a]. Next 
	// check using Euler's criteria that a is a quadratic residue mod p 
	// if so return false and the empty slice
	// If p = 3 mod 4, then check that a^{(p+1)/4} is a square root of 
	// a. If not, return false and the empty slice. Otherwise 
	// return true, [r,(-r) mod p] 
	// Find a quadratic non-residue modulo p 
	// Factor p-1 into 2^s*q for q odd 
	// Initialize d = (q+1)/2, x = a^q mod p, c = z^q mod p and 
	// r = a^d mod p 
	// Loop while x != 1
	//  - if x = 0, then return true, [0]
	//  - else compute the minimal i such that x^{2^i} = 1 
	//    let b = c^{2*s - i -1} mod p, s = i, c = b^2 mod p, 
	//    x = x*c mod p, and r = r*b mod p 
	return false, []int{}
}

// GenTonelliShanks(p,e,a) returns the solutions of x^2 = a mod p^e for a 
// prime p and a with gcd(a,p) = 1
func GenTonelliShanks(p, e, a int) (bool,[]int) {
	// Use recursion. Assume you have a solution to z^2 = a mod p^{e-1} 
	// and search for a solution x = z + p^{e-1}y for 0 <= y < p to 
	// x^2 = a mod p^e 
	return false, []int{}
}


