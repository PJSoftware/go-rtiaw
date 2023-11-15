package vec3_test

import (
	"testing"

	"github.com/pjsoftware/go-rtiaw/lib/vec3"
)

var testV1 vec3.Vec3 = vec3.Vec3{1.0, 2.0, 3.0}
var testV2 vec3.Vec3 = vec3.Vec3{4.0, 5.0, 6.0}
var testV3 vec3.Vec3 = vec3.Vec3{4.0, 3.0, 2.0}

func TestNewVec(t *testing.T) {
	wantX := 1.0
	wantY := 2.0
	wantZ := 3.0

	vec := vec3.NewVec3(wantX, wantY, wantZ)

	if vec.X != wantX {
		t.Fatalf("Returned value of X (%f) does not match expected (%f)", vec.X, wantX)
	}
	if vec.Y != wantY {
		t.Fatalf("Returned value of Y (%f) does not match expected (%f)", vec.Y, wantY)
	}
	if vec.Z != wantZ {
		t.Fatalf("Returned value of Z (%f) does not match expected (%f)", vec.Z, wantZ)
	}

	if !vec.EqualXYZ(wantX, wantY, wantZ) {
		t.Fatalf("EqualXYZ did not return true")
	}
}

func TestNegatives(t *testing.T) {
	vn := vec3.Negative(testV1)
	if !vn.EqualXYZ(-1.0, -2.0, -3.0) {
		t.Fatalf("Negative of 'testV1' (%s) did not EqualXYZ() expected -1,-2,-3", vn.String())
	}

	vn.Negative()
	if !vn.Equal(testV1) {
		t.Fatalf("vn.Negative() (%s) did not Equal() expected 1,2,3", vn.String())
	}
}

func TestAdd(t *testing.T) {
	v := vec3.Add(testV1, testV2)
	if !v.EqualXYZ(5.0, 7.0, 9.0) {
		t.Fatalf("Sum of two vectors (%s) did not equal expected value", v.String())
	}
	v.Add(testV3)
	if !v.EqualXYZ(9.0, 10.0, 11.0) {
		t.Fatalf("Adding another vector gave (%s), not the expected value", v.String())
	}
}

func TestSub(t *testing.T) {
	v := vec3.Sub(testV1, testV2)
	if !v.EqualXYZ(-3.0, -3.0, -3.0) {
		t.Fatalf("Sum of two vectors (%s) did not equal expected value", v.String())
	}
	v.Sub(testV3)
	if !v.EqualXYZ(-7.0, -6.0, -5.0) {
		t.Fatalf("Adding another vector gave (%s), not the expected value", v.String())
	}
}

func TestMult(t *testing.T) {
	v := vec3.Mult(testV1, testV2)
	if !v.EqualXYZ(4.0, 10.0, 18.0) {
		t.Fatalf("Multiplication of two vectors (%s) did not equal expected value", v.String())
	}
	v.Mult(testV3)
	if !v.EqualXYZ(16.0, 30.0, 36.0) {
		t.Fatalf("Multiplying another vector gave (%s), not the expected value", v.String())
	}
	v = vec3.MultBy(testV1, 4.0)
	if !v.EqualXYZ(4.0, 8.0, 12.0) {
		t.Fatalf("Multiplication by float (%s) did not equal expected value", v.String())
	}
	v.MultBy(0.5)
	if !v.EqualXYZ(2.0, 4.0, 6.0) {
		t.Fatalf("Multiplying another vector gave (%s), not the expected value", v.String())
	}
}

func TestDiv(t *testing.T) {
	v := vec3.Div(testV1, testV2)
	if !v.EqualXYZ(0.25, 0.4, 0.5) {
		t.Fatalf("Division of two vectors (%s) did not equal expected value", v.String())
	}
	v.Div(testV3)
	if !v.EqualXYZ(0.25/4, 0.4/3, 0.5/2) {
		t.Fatalf("Dividing another vector gave (%s), not the expected value", v.String())
	}
	v = vec3.DivBy(testV1, 4.0)
	if !v.EqualXYZ(0.25, 0.5, 0.75) {
		t.Fatalf("Division by float (%s) did not equal expected value", v.String())
	}
	v.DivBy(0.5)
	if !v.EqualXYZ(0.5, 1.0, 1.5) {
		t.Fatalf("Dividing another vector gave (%s), not the expected value", v.String())
	}
}

func TestLength(t *testing.T) {
	ls1 := testV1.LengthSquared()
	ls2 := testV2.LengthSquared()
	ls3 := testV3.LengthSquared()

	exp1 := float64(1+4+9)
	exp2 := float64(16+25+36)
	exp3 := float64(16+9+4)

	if ls1 != exp1 {
		t.Fatalf("testV1 LengthSquared (%f) is not expected %f", ls1, exp1)
	}
	if ls2 != exp2 {
		t.Fatalf("testV2 LengthSquared (%f) is not expected %f", ls2, exp2)
	}
	if ls3 != exp3 {
		t.Fatalf("testV3 LengthSquared (%f) is not expected %f", ls3, exp3)
	}

	v := vec3.NewVec3(3.0, 4.0, 12.0)
	e := 13.0
	l := v.Length()
	if l != e {
		t.Fatalf("vector %s Length (%f) is not expected %f", v.String(), l, e)
	}
}

func TestDotProduct(t *testing.T) {
	dp := vec3.DotProduct(testV1, testV2)
	exp := 32.0
	if dp != exp {
		t.Fatalf("Dot product (%f) did not equal expected value %f", dp, exp)
	}
}

func TestCrossProduct(t *testing.T) {
	cp := vec3.CrossProduct(testV1, testV2)
	exp := vec3.Vec3{-3.0, 6.0, -3.0}
	if !cp.Equal(exp) {
		t.Fatalf("Cross product (%s) did not equal expected value %s", cp.String(), exp.String())
	}
}

func TestUnitVector(t *testing.T) {
	uv := vec3.UnitVector(vec3.Vec3{3.0, 4.0, 0.0})
	exp := vec3.Vec3{0.6, 0.8, 0.0}
	if !uv.Equal(exp) {
		t.Fatalf("Unit vector (%s) did not match expected vector (%s)", uv.String(), exp.String())
	}
}