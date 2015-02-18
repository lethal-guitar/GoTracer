package scene

import "image/color"

type Material struct {
    Ambient FloatColor
    Diffuse FloatColor
    Specular FloatColor
    Specularity float32
    Reflectivity float32
    Transparency float32

    RefractionIndex float32
}

func MakeSimpleMaterial(color color.RGBA) Material {
    return Material{
        Ambient: FloatColor{0.0, 0.0, 0.0},
        Diffuse: MakeFloatColor(color),
        Specular: FloatColor{1.0, 1.0, 1.0},
        Specularity: 90.0,
    }
}
