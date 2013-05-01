package main

import (
	"testing"
)

func Test_cross(t *testing.T) {
	if X.cross(Y) != Z {
		t.Fail()
	}
}

func Test_scale_zero(t *testing.T) {
	if V0.scale(20) != V0 {
		t.Fail()
	}
}

func Test_scale_reverse(t *testing.T) {
	if X.scale(-1).scale(-1) != X {
		t.Fail()
	}
	if X.scale(2).scale(0.5) != X {
		t.Fail()
	}
}

func Test_norm(t *testing.T) {
	v := randVec3()
	if v.norm().mag() != 1.0 {
		t.Fail()
	}
}

func Test_inc_dec(t *testing.T) {
	a, b := randVec3(), randVec3()
	c := a
	a.inc(b)
	if a.eq(c) {
		t.Fail()
	}
	a.dec(b)
	if !a.eq(c) {
		t.Fail()
	}
}

func Test_reflect(t *testing.T) {
	v, n := randVec3(), randVec3().norm()
	if !v.eq(v.reflect(n).reflect(n)) {
		t.Fail()
	}
}
