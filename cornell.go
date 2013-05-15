package main

func cornell(w, h int) *scene {
	scene := NewScene()
	scene.camera = NewCamera(vec3{0, 0, 0.5}, vec3{0, 0, 0}, w, h, 1.5)

	// glass &sphere
	// this is the original location
	scene.Add(&body{&sphere{vec3{1, 0, -2}, 0.5}, &glass{color{0.8, 0.8, 1.0}, 1.5, 0.1}})

	// these are for fun
	scene.Add(&body{&sphere{vec3{-1, 0, -2}, 0.5}, &glass{color{1.0, 0.8, 0.8}, 1.5, 0.1}})
	scene.Add(&body{&sphere{vec3{0, 0, -2}, 0.5}, &glass{color{0.8, 1.0, 0.8}, 1.5, 0.1}})

	// chrome &sphere
	// scene.Add(&body{&sphere{vec3{-1.1, 0, -2.8}, 0.5}, &chrome{color{0.8, 0.8, 0.8}}})

	// floor
	scene.Add(&body{&sphere{vec3{0, -10e6, -3.5}, 10e6 - 0.5}, &dull{color{0.9, 0.9, 0.9}}})

	// back
	scene.Add(&body{&sphere{vec3{0, 0, -10e6}, 10e6 - 4.5}, &dull{color{0.9, 0.9, 0.9}}})

	// left
	scene.Add(&body{&sphere{vec3{-10e6, 0, -3.5}, 10e6 - 1.9}, &dull{color{0.9, 0.5, 0.5}}})

	// right
	scene.Add(&body{&sphere{vec3{10e6, 0, -3.5}, 10e6 - 1.9}, &dull{color{0.5, 0.5, 0.9}}})

	// top light
	scene.Add(&body{&sphere{vec3{0, 10e6, 0}, 10e6 - 2.5}, &glow{sunlight}})

	// front
	scene.Add(&body{&sphere{vec3{0, 0, 10e6}, 10e6 - 2.5}, &dull{color{0.9, 0.9, 0.9}}})
	return scene
}
