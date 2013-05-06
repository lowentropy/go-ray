package main

import "testing"

func Test_NewCamera(t *testing.T) {
	c := NewCamera(V0, Z, 2, 1, 1)
	if c.aspect != 2 {
		t.Fail()
	}
	if c.normal != Z {
		t.Fail()
	}
}

func Test_CameraShoot(t *testing.T) {
	c := NewCamera(V0, Z, 3, 3, 1)
	r := c.Shoot(1, 1, 0)
	expected := ray{V0, Z}
	if r != expected {
		t.Fail()
	}

	r = c.Shoot(0, 0, 0)
	if r.normal.x != r.normal.y {
		t.Fatal("Normal x", r.normal.x, "not equal to y", r.normal.y)
	}

	if r.normal.x*2 != r.normal.z {
		t.Fatal("normal x", r.normal.x, "was not half of z", r.normal.z)
	}
}
