package vec3

// Declare as an alias to avoid excessive typecasting between the two
type Point3 = Vec3

func NewPoint3(x, y, z float64) *Point3 {
	return NewVec3(x, y, z)
}