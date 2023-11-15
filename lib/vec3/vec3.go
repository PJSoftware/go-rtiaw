package vec3

import (
	"fmt"
	"math"
)

// Vec3{} is a struct representing a 3D vector
type Vec3 struct {
	X float64
	Y float64
	Z float64
}

// // Point{} is simply a Vec3{} -- is this how this works in Go?
// type Point Vec3

// // Color{} is simply a Vec3{} -- is this how this works in Go?
// type Color Vec3

// NewVec3(x,y,z) returns a new, initialised Vec3{} object
func NewVec3(x, y, z float64) *Vec3 {
	v := &Vec3{}
	v.X = x
	v.Y = y
	v.Z = z
	return v
}

// Negative(v3) returns a new Vec3{} for which all values are the negative of
// those in the supplied Vec3{}
func Negative(v3 Vec3) *Vec3 {
	return NewVec3(-v3.X, -v3.Y, -v3.Z)
}

// v.Negative() sets all x,y,z values to their negative value
func (v *Vec3) Negative() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

// v.AddVec(v3) adds the specified vector to the current vector
func (v *Vec3) AddVec(v3 Vec3) {
	v.X += v3.X
	v.Y += v3.Y
	v.Z += v3.Z
}

// v.MultBy(f) multiplies the current vector's x,y,z by the specified float f
func (v *Vec3) MultBy(f float64) {
	v.X *= f
	v.Y *= f
	v.Z *= f
}

// v.DivBy(f) divides the current vector's x,y,z by the specified float f
func (v *Vec3) DivBy(f float64) {
	if f == 0.0 {
		v.MultBy(f) // Not strictly correct; is this the best way to handle this?
	} else {
		v.MultBy(1/f)
	}
}

// v.LengthSquared() returns the squared length of the current vector
func (v *Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// v.Length() returns the length of the current vector
func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vec3) String() string {
	return fmt.Sprintf("%f %f %f", v.X, v.Y, v.Z)
}

func (v *Vec3) Print() {
	fmt.Printf("%f %f %f", v.X, v.Y, v.Z)
}

func (v *Vec3) Println() {
	fmt.Printf("%f %f %f\n", v.X, v.Y, v.Z)
}

func Add(v1, v2 Vec3) *Vec3 {
	return NewVec3(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z)
}
func Sub(v1, v2 Vec3) *Vec3 {
	return NewVec3(v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z)
}
func Mult(v1, v2 Vec3) *Vec3 {
	return NewVec3(v1.X*v2.X, v1.Y*v2.Y, v1.Z*v2.Z)
}

func MultBy(v Vec3, f float64) *Vec3 {
	return NewVec3(v.X*f, v.Y*f, v.Z*f)
}
func DivBy(v Vec3, f float64) *Vec3 {
	if f == 0.0 {
		return NewVec3(0,0,0)
	}

	return NewVec3(v.X/f, v.Y/f, v.Z/f)
}

func DotProduct(v1, v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func CrossProduct(v1, v2 Vec3) *Vec3 {
	return NewVec3(
		v1.Y * v2.Z - v1.Z * v2.Y,
		v1.Z * v2.X - v1.X * v2.Z,
		v1.X * v2.Y - v1.Y * v2.X )
}

func UnitVector(v Vec3) *Vec3 {
	return DivBy(v, v.Length())
}
