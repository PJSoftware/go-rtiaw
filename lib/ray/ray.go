package ray

import "github.com/pjsoftware/go-rtiaw/lib/vec3"

type Ray struct {
	Origin    *vec3.Point3
	Direction *vec3.Vec3
}

func NewRay(origin *vec3.Point3, dir *vec3.Vec3) *Ray {
	r := &Ray{origin, vec3.UnitVector(dir)}
	return r
}

func (r *Ray) At(t float64) *vec3.Point3 {
	pt := *vec3.Add(r.Origin, vec3.MultBy(r.Direction, t))
	return &pt
}
