package main

import "testing"

func Test_sphere_intersect_from_inside(t *testing.T) {
	s := sphere{V0, 1}
	r := ray{V0, X}
	hit, ok := s.intersect(r)
	if !ok {
		t.Fatal("Missed the sphere")
	}
	if !hit.pt.eq(X) {
		t.Fatal("Expected hit at", X, "but got", hit.pt)
	}
}

func Test_glass_sphere_passthru(t *testing.T) {

}
