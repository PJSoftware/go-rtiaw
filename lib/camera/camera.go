package camera

import (
	"strconv"
	"strings"

	"github.com/pjsoftware/go-rtiaw/lib/ray"
	sb "github.com/pjsoftware/go-rtiaw/lib/screenBuffer"
	"github.com/pjsoftware/go-rtiaw/lib/vec3"
)

type Camera struct {
	Image       *sb.ScreenBuffer

	ImgWidth  int
	ImgHeight int

	aspectRatio	float64
	focalLength float64
	vpHeight 		float64
	vpWidth 		float64

	position *vec3.Point3

	viewportU *vec3.Vec3
	viewportV *vec3.Vec3

	pixelDeltaU *vec3.Vec3
	pixelDeltaV *vec3.Vec3

	vpUpperLeft *vec3.Vec3
	px00Loc *vec3.Vec3
}

func NewCamera(aspectRatio string, imgWidth int) *Camera {
	cam := &Camera{}
	cam.aspectRatio = calcAspectRatio(aspectRatio)
	cam.ImgWidth = imgWidth
	cam.ImgHeight = int(float64(cam.ImgWidth) / cam.aspectRatio)
	cam.Image = sb.NewScreenBuffer(cam.ImgWidth, cam.ImgHeight)

	cam.focalLength = 1.0
	cam.vpHeight = 2.0
	cam.vpWidth = cam.vpHeight * float64(cam.ImgWidth)/float64(cam.ImgHeight)

	cam.position = vec3.NewPoint3(0.0, 0.0, 0.0)
	cam.viewportU = vec3.NewVec3(cam.vpWidth, 0.0, 0.0)
	cam.viewportV = vec3.NewVec3(0.0, -cam.vpHeight, 0.0)

	cam.pixelDeltaU = vec3.DivBy(cam.viewportU, float64(cam.ImgWidth))
	cam.pixelDeltaV = vec3.DivBy(cam.viewportV, float64(cam.ImgHeight))

	cam.vpUpperLeft = cam.position
	cam.vpUpperLeft.Sub(vec3.NewVec3(0, 0, cam.focalLength))
	cam.vpUpperLeft.Sub(vec3.DivBy(cam.viewportU, 2.0))
	cam.vpUpperLeft.Sub(vec3.DivBy(cam.viewportV, 2.0))

	cam.px00Loc = vec3.Add(cam.vpUpperLeft,vec3.MultBy(vec3.Add(cam.pixelDeltaU, cam.pixelDeltaV),0.5))

	return cam
}

func calcAspectRatio(ar string) float64 {
	arVal := strings.Split(ar, ":")
	num, _ := strconv.ParseFloat(arVal[0], 64)
	denom, _ := strconv.ParseFloat(arVal[1], 64)
	return num/denom
}

func (cam *Camera) CalcPixel(i, j int) {
	pxI := vec3.MultBy(cam.pixelDeltaU, float64(i))
	pxJ := vec3.MultBy(cam.pixelDeltaV, float64(j))
	pxCenter := vec3.Add(cam.px00Loc,vec3.Add(pxI,pxJ))

	rayDirection := vec3.Sub(pxCenter, cam.position)
	r := ray.NewRay(cam.position, rayDirection)

	pixel := RayColour(r)

	cam.Image.SetPixel(i, j, pixel)
}

func RayColour(r *ray.Ray) *sb.Pixel {
	ud := vec3.UnitVector(r.Direction)
	a := 0.5 * (ud.Y + 1.0)
	px3 := vec3.Add(
		vec3.MultBy(vec3.NewVec3(1.0,1.0,1.0),(1.0-a)),
		vec3.MultBy(vec3.NewVec3(0.5,0.7,1.0),a))
	px := sb.NewPixelByColour(float32(px3.X), float32(px3.Y), float32(px3.Z))
	return px
}