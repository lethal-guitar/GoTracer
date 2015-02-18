package scene

import (
    "math"
    "github.com/lethal-guitar/go_tracer/math32"
    "github.com/lethal-guitar/go_tracer/vecmath"
)

const EPSILON = 0.001

// Axis aligned bounding box
type AABB struct {
    Min, Max vecmath.Vec3dA
}

func MakeAABBV(v1, v2 vecmath.Vec3d) AABB {
    return MakeAABB(v1.X, v1.Y, v1.Z, v2.X, v2.Y, v2.Z)
}


func MakeAABB(x, y, z, xm, ym, zm float32) AABB {
    return AABB{
        Min: vecmath.ToVec3dA(&vecmath.Vec3d{x, y, z}),
        Max: vecmath.ToVec3dA(&vecmath.Vec3d{xm, ym, zm}),
    }
}

func (self *AABB) Overlaps(other *AABB) bool {
    return !(
        self.Min.V[0] > other.Max.V[0] || other.Min.V[0] > self.Max.V[0] ||
        self.Min.V[1] > other.Max.V[1] || other.Min.V[1] > self.Max.V[1] ||
        self.Min.V[2] > other.Max.V[2] || other.Min.V[2] > self.Max.V[2])
}

func (self *AABB) Contains(other *AABB) bool {
    return (
        self.Min.V[0] <= other.Min.V[0] && self.Max.V[0] >= other.Max.V[0] &&
        self.Min.V[1] <= other.Min.V[1] && self.Max.V[1] >= other.Max.V[1] &&
        self.Min.V[2] <= other.Min.V[2] && self.Max.V[2] >= other.Max.V[2])
}

func (self *AABB) Center() vecmath.Vec3d {
    min := vecmath.ToVec3d(&self.Min)
    max := vecmath.ToVec3d(&self.Max)
    return min.Added(max.Subtracted(min).Scaled(0.5))
}

func (self *AABB) IntersectsBasic(ray *Ray) bool {
    return self.Intersects(ray, float32(math.Inf(1)))
}

func (self *AABB) Expand(other *AABB) {
    self.Min.V[0] = math32.Minf(self.Min.V[0], other.Min.V[0])
    self.Min.V[1] = math32.Minf(self.Min.V[1], other.Min.V[1])
    self.Min.V[2] = math32.Minf(self.Min.V[2], other.Min.V[2])
    self.Max.V[0] = math32.Maxf(self.Max.V[0], other.Max.V[0])
    self.Max.V[1] = math32.Maxf(self.Max.V[1], other.Max.V[1])
    self.Max.V[2] = math32.Maxf(self.Max.V[2], other.Max.V[2])
}

func (self *AABB) Intersects(ray *Ray, distanceRequired float32) bool {
    invDir := vecmath.Vec3d{
        1/ray.Direction.X,
        1/ray.Direction.Y,
        1/ray.Direction.Z,
    }

    tx1 := (self.Min.V[0] - ray.Origin.X) * invDir.X
    tx2 := (self.Max.V[0] - ray.Origin.X) * invDir.X

    tmin := math32.Minf(tx1, tx2)
    tmax := math32.Maxf(tx1, tx2)

    ty1 := (self.Min.V[1] - ray.Origin.Y) * invDir.Y
    ty2 := (self.Max.V[1] - ray.Origin.Y) * invDir.Y

    tmin = math32.Maxf(tmin, math32.Minf(ty1, ty2))
    tmax = math32.Minf(tmax, math32.Maxf(ty1, ty2))

    tz1 := (self.Min.V[2] - ray.Origin.Z) * invDir.Z
    tz2 := (self.Max.V[2] - ray.Origin.Z) * invDir.Z

    tmin = math32.Maxf(tmin, math32.Minf(tz1, tz2))
    tmax = math32.Minf(tmax, math32.Maxf(tz1, tz2))

    return tmax >= math32.Maxf(0, tmin) && tmin < distanceRequired
}

//func (self *AABB) IntersectsSSE(ray *Ray, distanceRequired float32) bool {
    //invDir := vecmath.Vec3d{
        //1/ray.Direction.X,
        //1/ray.Direction.Y,
        //1/ray.Direction.Z,
    //}

    //t1 := vecmath.MulPerElem()
    //t2 := vecmath.MulPerElem()
    //Aos::Vector3 t1(Aos::mulPerElem(m_min - ray.m_pos, invDir));
    //Aos::Vector3 t2(Aos::mulPerElem(m_max - ray.m_pos, invDir));

    //tmin1 := vecmath.MinPerElem(&t1, &t2)
    //tmax1 := vecmath.MaxPerElem(&t1, &t2)

    //tmin := vecmath.MaxElem(tmin1);
    //tmax := vecmath.MinElem(tmax1);

    //return tmax >= math32.Maxf(0, tmin) && tmin < distanceRequired
//}
