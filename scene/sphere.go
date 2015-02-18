package scene

import (
    "github.com/lethal-guitar/go_tracer/math32"
    "github.com/lethal-guitar/go_tracer/vecmath"
)

// "Hot" sphere data optimized for fast cache-friendly access
type SphereIntersectionInfo struct {
    Pos vecmath.Vec3d
    Radius float32
}

type Sphere struct {
    TraceableBase

    SphereIntersectionInfo
}

func MakeSimpleSphere(pos vecmath.Vec3d, radius float32) *Sphere {
    return &Sphere{
        SphereIntersectionInfo: SphereIntersectionInfo{
            Pos: pos,
            Radius: radius,
        },
    }
}

func MakeSphere(pos vecmath.Vec3d, radius float32, material Material) *Sphere {
    return &Sphere{
        TraceableBase: TraceableBase{ObjectMaterial: material},
        SphereIntersectionInfo: SphereIntersectionInfo{
            Pos: pos,
            Radius: radius,
        },
    }
}

func (self *SphereIntersectionInfo) CheckIntersection(ray *Ray) (bool, float32) {
    v := ray.Origin.Subtracted(self.Pos)
    b := v.Dot(ray.Direction)
    c := v.Dot(v) - self.Radius*self.Radius
    determinant := b*b - c

    if determinant < 0.0 {
        return false, 0.0
    }

    sqrtD := math32.Sqrtf(determinant)
    t0 := -b + sqrtD
    t1 := -b - sqrtD
    t := math32.Minf(t0, t1)

    return t > 0.0, t
}

func (self *Sphere) Normal(intersection *vecmath.Vec3d) vecmath.Vec3d {
    return vecmath.MakeDirectionVector(&self.Pos, intersection)
}

func (self *Sphere) Bounds() AABB {
    return MakeAABBV(
        self.Pos.Subtracted(vecmath.Vec3d{self.Radius, self.Radius, self.Radius}),
        self.Pos.Added(vecmath.Vec3d{self.Radius, self.Radius, self.Radius}),
    )
}
