package tests

import (
    "github.com/lethal-guitar/go_tracer/scene"
    "github.com/lethal-guitar/go_tracer/vecmath"
    . "gopkg.in/check.v1"
)

func (s *MySuite) TestAABBAcceptsIntersectingRay(c *C) {
    aabb := scene.MakeAABB(1.001,1.001,1.001, 1.124,1.124,1.124)
    center := aabb.Center()

    start := vecmath.Vec3d{}
    ray := scene.Ray{start, vecmath.MakeDirectionVector(&start, &center)}

    test := aabb.IntersectsBasic(&ray)
    c.Assert(test, Equals, true)
}

func (s *MySuite) TestAABBRejectsNonIntersectingRay(c *C) {
    aabb := scene.MakeAABB(1, 1, 1, 3, 3, 3)

    start := vecmath.Vec3d{8, 8, -20}
    p := vecmath.Vec3d{8, 9, 180}
    ray := scene.Ray{start, vecmath.MakeDirectionVector(&start, &p)}

    test := aabb.IntersectsBasic(&ray)
    c.Assert(test, Equals, false)
}

func (s *MySuite) TestAABBCorrectlyCalculatesCenter(c *C) {
    aabb := scene.MakeAABB(3, 3, 3, 6, 6, 6)
    center := aabb.Center()

    c.Assert(center, Equals, vecmath.Vec3d{4.5, 4.5, 4.5})
}

func (s *MySuite) TestAABBIntersectionWorksProperly(c *C) {
    aabb := scene.MakeAABB(-24, -5, 187, 25, 17, 232)
    ray := scene.Ray{
        vecmath.Vec3d{0, 0, -10},
        vecmath.Vec3d{0.09208211, -0.023144966, 0.9954823},
    }

    c.Assert(aabb.IntersectsBasic(&ray), Equals, true)
}
