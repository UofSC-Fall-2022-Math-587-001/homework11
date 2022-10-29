package hw10

import (
	"math"
	"testing"
	"github.com/UofSC-Fall-2022-Math-587-001/homework10/library"
)

func Test1(t *testing.T) {
	a := 10
	p := 11 
	b, _ := TonelliShanks(p,a)
	if b {
		t.Errorf("There is no solution to x^2 = %d mod %d but said there is\n",a,p)
	}
}

func Test2(t *testing.T) {
	a := 9
	p := 11 
	b, roots := TonelliShanks(p,a)
	if !b {
		t.Errorf("There are no solutions to x^2 = %d mod %d but said there are\n",a,p)
	}
	for _, r := range(roots) {
		if library.FastPower(uint(p),r,2) != a {
			t.Errorf("You said %d^2 = %d mod %d but it is not\n",r,a,p)
		}
	}
}

func Test3(t *testing.T) {
	a := 317
	p := 1021
	b, roots := TonelliShanks(p,a)
	if !b {
		t.Errorf("There are no solutions to x^2 = %d mod %d but said there are\n",a,p)
	}
	for _, r := range(roots) {
		if library.FastPower(uint(p),r,2) != a {
			t.Errorf("You said %d^2 = %d mod %d but it is not\n",r,a,p)
		}
	}
}

func Test4(t *testing.T) {
	a := 118
	p := 8675309
	b, roots := TonelliShanks(p,a)
	if !b {
		t.Errorf("There are no solutions to x^2 = %d mod %d but said there are\n",a,p)
	}
	for _, r := range(roots) {
		if library.FastPower(uint(p),r,2) != a {
			t.Errorf("You said %d^2 = %d mod %d but it is not\n",r,a,p)
		}
	}
}

func Test5(t *testing.T) {
	a := 77111
	p := 5
	e := 7
	b, roots := GenTonelliShanks(p,e,a)
	if !b {
		t.Errorf("There are no solutions to x^2 = %d mod %d but said there are\n",a,p)
	}
	for _, r := range(roots) {
		if library.FastPower(uint(math.Pow(5,7)),r,2) != a {
			t.Errorf("You said %d^2 = %d mod %d but it is not\n",r,a,int(math.Pow(5,7)))
		}
	}
}
