package main

import "testing"

func Test_NewCamera(t *testing.T) {
	c := NewCamera(V0, Z, 2, 1)
	if c.aspect != 2 {
		t.Fail()
	}
	if c.normal != Z {
		t.Fail()
	}
}

func Test_CameraShoot(t *testing.T) {
	c := NewCamera(V0, Z, 3, 3)
	r := c.Shoot(1, 1)
	expected := ray{V0, Z}
	if r != expected {
		t.Fail()
	}

	r = c.Shoot(0, 0)
	if r.normal.x != r.normal.y {
		t.Fail()
	}

	if r.normal.x*2 != r.normal.z {
		t.Fail()
	}
}
