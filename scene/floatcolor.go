package scene

import (
    "image/color"
    "github.com/lethal-guitar/go_tracer/math32"
)

// Color whose components are represented as float value in the range [0; 1]
type FloatColor struct {
    R, G, B float32
}

func MakeFloatColor(color color.RGBA) FloatColor {
    return FloatColor{
        float32(color.R) / 255,
        float32(color.G) / 255,
        float32(color.B) / 255,
    }
}

//
// Pure members

func convert(value float32) uint8 {
    return uint8(math32.Minf(1.0, value) * 255)
}

// Satisfy color interface
func (self FloatColor) RGBA() (r, g, b, a uint32) {
    nrgba := color.NRGBA{convert(self.R), convert(self.G), convert(self.B), 255}
    return nrgba.RGBA()
}

func (self *FloatColor) Added(other FloatColor) FloatColor {
    return FloatColor{self.R + other.R, self.G + other.G, self.B + other.B}
}

func (self *FloatColor) Multiplied(other FloatColor) FloatColor {
    return FloatColor{self.R * other.R, self.G * other.G, self.B * other.B}
}

func (self *FloatColor) Scaled(scalar float32) FloatColor {
    return FloatColor{
        self.R * scalar,
        self.G * scalar,
        self.B * scalar,
    }
}

//
// Mutating members

func (self *FloatColor) Add(other FloatColor) {
    self.R += other.R
    self.G += other.G
    self.B += other.B
}

func (self *FloatColor) Blend(factor float32, otherColor FloatColor) {
    result := self.Scaled(1.0 - factor)
    *self = result.Added(otherColor.Scaled(factor))
}
