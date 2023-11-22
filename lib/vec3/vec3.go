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

// NewVec3(x,y,z) returns a new, initialised Vec3{} object
func NewVec3(x, y, z float64) *Vec3 {
	v := &Vec3{}
	v.X = x
	v.Y = y
	v.Z = z
	return v
}

// v.Equal(v3) returns true if both vectors are equal
func (v *Vec3) Equal(v3 *Vec3) bool {
	return v.X == v3.X && v.Y == v3.Y && v.Z == v3.Z
}

// v.EqualXYZ(x,y,z) returns true if both vectors are equal
func (v *Vec3) EqualXYZ(x, y, z float64) bool {
	return v.X == x && v.Y == y && v.Z == z
}

// Negative(v3) returns a new Vec3{} for which all values are the negative of
// those in the supplied Vec3{}
func Negative(v3 *Vec3) *Vec3 {
	return NewVec3(-v3.X, -v3.Y, -v3.Z)
}

// v.Negative() sets all x,y,z values to their negative value
func (v *Vec3) Negative() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

// v.Add(v1) adds the specified vector to the current vector
func (v *Vec3) Add(v1 *Vec3) {
	v.X += v1.X
	v.Y += v1.Y
	v.Z += v1.Z
}

// Add(v1,v2) adds two vectors together and returns the result
func Add(v1, v2 *Vec3) *Vec3 {
	return NewVec3(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z)
}

// v.Sub(v1) subtracts the specified vector from the current vector
func (v *Vec3) Sub(v1 *Vec3) {
	v.X -= v1.X
	v.Y -= v1.Y
	v.Z -= v1.Z
}

// Sub(v1,v2) subtracts v2 from v1 and returns the result
func Sub(v1, v2 *Vec3) *Vec3 {
	return NewVec3(v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z)
}

// v.Mult(v1) multiplies the specified vector by the current vector
func (v *Vec3) Mult(v1 *Vec3) {
	v.X *= v1.X
	v.Y *= v1.Y
	v.Z *= v1.Z
}

// Mult(v1,v2) multiplies v1 and v2 and returns the result
func Mult(v1, v2 *Vec3) *Vec3 {
	return NewVec3(v1.X*v2.X, v1.Y*v2.Y, v1.Z*v2.Z)
}

// v.MultBy(f) multiplies the current vector's x,y,z by the specified float f
func (v *Vec3) MultBy(f float64) {
	v.X *= f
	v.Y *= f
	v.Z *= f
}

// MultBy(v,f) multiplies the vector v by the float f and returns the result
func MultBy(v *Vec3, f float64) *Vec3 {
	return NewVec3(v.X*f, v.Y*f, v.Z*f)
}

// v.Div(v1) divides current vector (v) by the specified vector (v1)
func (v *Vec3) Div(v1 *Vec3) {
	v.X /= v1.X
	v.Y /= v1.Y
	v.Z /= v1.Z
}

// Div(v1,v2) divides v1 by v2 and returns the result
func Div(v1, v2 *Vec3) *Vec3 {
	return NewVec3(v1.X/v2.X, v1.Y/v2.Y, v1.Z/v2.Z)
}

// v.DivBy(f) divides the current vector's x,y,z by the specified float f
func (v *Vec3) DivBy(f float64) {
	v.MultBy(1 / f)
}

// DivBy(v,f) divides the vector v by the float f and returns the result
func DivBy(v *Vec3, f float64) *Vec3 {
	return NewVec3(v.X/f, v.Y/f, v.Z/f)
}

// v.LengthSquared() returns the squared length of the current vector
func (v *Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// v.Length() returns the length of the current vector
func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// DotProduct(v1,v2) returns the dot product of vectors v1 and v2
//
// The dot product of two vectors a and b, where the angle between a and b is θ,
// is `a · b = |a| × |b| × cos(θ)`
//
// (see https://www.mathsisfun.com/algebra/vectors-dot-product.html)
func DotProduct(v1, v2 *Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// CrossProduct(v1, v2) returns the cross product of vectors v1 and v2
//
// The Cross Product a × b of two vectors is another vector that is at
// right angles to both
//
// (see https://www.mathsisfun.com/algebra/vectors-cross-product.html)
func CrossProduct(v1, v2 *Vec3) *Vec3 {
	return NewVec3(
		v1.Y*v2.Z-v1.Z*v2.Y,
		v1.Z*v2.X-v1.X*v2.Z,
		v1.X*v2.Y-v1.Y*v2.X)
}

func UnitVector(v *Vec3) *Vec3 {
	return DivBy(v, v.Length())
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
