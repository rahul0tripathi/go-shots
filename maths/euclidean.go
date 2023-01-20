package main

import "fmt"

/*
EuclideanGcd - euclidean algorithm calculates the gcd(a,b)
	r1 = a mod b
	r2 = b mod r1
	.
	.
	.
	r = r(prev-1) mod r(prev); if r == 0 then gcd(a,b) = r(prev)
*/
func EuclideanGcd(a, b int) int {
	if b == 0 {
		return a
	}
	return EuclideanGcd(b, a%b)

}

/*
ExtendedEuclidean - helps find (x,y) such that ax + by = gcd(a,b)
	r(-1) = a       					x(-1) = 1, y(-1) = 0
	r(0) = b       						x(0) = 0, y(0) = 1
	r(1) = a mod b		q(1) = a/b		x(1) = x(-1) - q*x(0), y(1) = y(-1) - q*y(0)
	r(2) = b mod r(1)	q(2) = b/r(1)	x(2) = x(0) - q*x(1), y(1) = y(0) - q*y(1)
	r(n) = r(n-2) mod r(n-1)	q(n) = r(n-2)/r(n-1)	x(n) = x(n-2) - q*x(n-1), y(n) = y(n-2) - q*y(n-1)
	r(n+1) = r(n-1) mod r(n) == 0;  x = x(n); y = y(n)
*/
func ExtendedEuclidean(a, b, xa, ya, xb, yb int) (int, int, int) {
	if b == 0 {
		return xa, ya, a
	}
	q := a / b
	xa, xb, ya, yb = xb, xa-q*xb, yb, ya-q*yb
	return ExtendedEuclidean(b, a%b, xa, ya, xb, yb)
}

func main() {
	fmt.Println(EuclideanGcd(18, 12))
	fmt.Println(ExtendedEuclidean(1759, 550, 1, 0, 0, 1))
}
