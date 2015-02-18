package vecmath

import "math"

type Vec3d struct {
    X, Y, Z float32
}

type Vec3dA struct {
    V [3]float32
}

func ToVec3dA(vec *Vec3d) Vec3dA {
    return Vec3dA{
        V: [3]float32{vec.X, vec.Y, vec.Z},
    }
}

func ToVec3d(vec *Vec3dA) Vec3d {
    return Vec3d{vec.V[0], vec.V[1], vec.V[2]}
}

func MakeVec3d(value float32) Vec3d {
    return Vec3d{value, value, value}
}

func (self Vec3d) Added(other Vec3d) Vec3d {
    return Vec3d{self.X + other.X, self.Y + other.Y, self.Z + other.Z}
}

func (self *Vec3d) Length() float32 {
    x2, y2, z2 := self.X*self.X, self.Y*self.Y, self.Z*self.Z
    return float32(math.Sqrt(float64(x2 + y2 + z2)))
}

func (self Vec3d) Normalized() Vec3d {
    l := float32(1.0) / self.Length()
    return Vec3d{self.X * l, self.Y * l, self.Z * l}
}

func (self Vec3d) Dot(other Vec3d) float32 {
    return self.X*other.X + self.Y*other.Y + self.Z*other.Z
}

func (self Vec3d) Subtracted(other Vec3d) Vec3d {
    return Vec3d{self.X - other.X, self.Y - other.Y, self.Z - other.Z}
}

func (self Vec3d) Scaled(scalar float32) Vec3d {
    return Vec3d{self.X * scalar, self.Y * scalar, self.Z * scalar}
}

func (self Vec3d) DistanceTo(other *Vec3d) float32 {
    direction := other.Subtracted(self)
    return direction.Length()
}

func (self Vec3d) Cross(other Vec3d) Vec3d {
    return Vec3d{
        X: self.Y*other.Z - self.Z*other.Y,
        Y: self.Z*other.X - self.X*other.Z,
        Z: self.X*other.Y - self.Y*other.X,
    }
}

func MakeDirectionVector(from, to *Vec3d) Vec3d {
    return to.Subtracted(*from).Normalized()
}

func MakeReflect(vec, normal *Vec3d) Vec3d {
    dot := normal.Dot(*vec)
    return normal.Scaled(float32(2.0) * dot).Subtracted(*vec).Normalized()
}
