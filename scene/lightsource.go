package scene

import (
    "image/color"
    "github.com/lethal-guitar/go_tracer/vecmath"
)

// A point light source
type LightSource struct {
    Position vecmath.Vec3d
    Diffuse, Specular FloatColor
}

// Create a simple colored light with default specularity
func MakeColoredLight(position vecmath.Vec3d, color color.RGBA) LightSource {
    diffuse := MakeFloatColor(color)
    return LightSource{
        Position: position,
        Diffuse: diffuse,
        Specular: FloatColor{1.0, 1.0, 1.0},
    }
}
