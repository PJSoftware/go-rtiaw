package ray_test

import (
	"testing"

	"github.com/pjsoftware/go-rtiaw/lib/ray"
	"github.com/pjsoftware/go-rtiaw/lib/vec3"
)

func TestNewRayAt(t *testing.T) {
	orig := vec3.NewVec3(0.0, 0.0, 0.0)
	dir := vec3.NewVec3(2.0, 3.0, 4.0)
	ray := ray.NewRay(orig, dir)

	if !orig.Equal(ray.Origin) {
		t.Fatalf("ray has malformed origin")
	}
	if !dir.Equal(ray.Direction) {
		t.Fatalf("ray has malformed direction")
	}

	pt := ray.At(3.0)
	exp := vec3.MultBy(dir, 3.0)
	if !pt.Equal(exp) {
		t.Fatalf("At() was not calculated correctly")
	}
}
