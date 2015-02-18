package tests

import (
    "github.com/lethal-guitar/go_tracer/scene"
    "github.com/lethal-guitar/go_tracer/vecmath"
    . "gopkg.in/check.v1"
)

func (s *MySuite) TestSphereReportsFalseForNonIntersectingRay(c *C) {
    ray := &scene.Ray{}
    sphere := scene.MakeSimpleSphere(vecmath.Vec3d{0.0, 0.0, 0.0}, 1.0)

    hasIntersection, _ := sphere.CheckIntersection(ray)

    c.Assert(hasIntersection, Equals, false)
}

func (s *MySuite) TestSphereReportsTrueForIntersectingRay(c *C) {
    ray := &scene.Ray{vecmath.Vec3d{0.0, -2.0, 0.0}, vecmath.Vec3d{0.0, 1.0, 0.0}}
    sphere := scene.MakeSimpleSphere(vecmath.Vec3d{0.0, 0.0, 0.0}, 0.2)

    hasIntersection, _ := sphere.CheckIntersection(ray)

    c.Assert(hasIntersection, Equals, true)
}

func (s *MySuite) TestSphereReturnsDistanceToRayOrigin(c *C) {
    ray := &scene.Ray{vecmath.Vec3d{0.0, 0.0, 0.0}, vecmath.Vec3d{0.0, 1.0, 0.0}}
    sphere := scene.MakeSimpleSphere(vecmath.Vec3d{0.0, 10.0, 0.0}, 3.5)

    _, distance := sphere.CheckIntersection(ray)

    c.Assert(distance, Equals, float32(6.5))
}

// TODO: Implement AABB intersection tests per Traceable

//func (s *MySuite) TestSphereRejectsDisjointAABB(c *C) {
    //sphere := scene.Sphere{Pos: vecmath.Vec3d{0,0,0}, Radius: 5}
    //aabb := scene.MakeAABB(-20,-20,-20,-10,-10,-10)

    //test := sphere.CheckAABBIntersection(&aabb)

    //c.Assert(test, Equals, false)
//}

//func (s *MySuite) TestSphereAcceptsContainingAABB(c *C) {
    //sphere := scene.Sphere{Pos: vecmath.Vec3d{1,1,1}, Radius: 5}
    //aabb := scene.MakeAABB(-4, -4, -4, 6, 6, 6)

    //test := sphere.CheckAABBIntersection(&aabb)

    //c.Assert(test, Equals, true)
//}

//func (s *MySuite) TestSphereAcceptsOverlappingAABB(c *C) {
    //sphere := scene.Sphere{Pos: vecmath.Vec3d{1,1,1}, Radius: 5}
    //aabb := scene.MakeAABB(-4, -4, -4, 2, 2, 2)

    //test := sphere.CheckAABBIntersection(&aabb)

    //c.Assert(test, Equals, true)
//}

//func (s *MySuite) TestPlaneRejectsDisjointAABB(c *C) {
    //plane := scene.MakePlane(vecmath.Vec3d{Y: 5}, vecmath.Vec3d{1,0,0})
    //aabb := scene.MakeAABB(-4, -4, -4, 4, 4, 4)

    //test := plane.CheckAABBIntersection(&aabb)

    //c.Assert(test, Equals, false)
//}

//func (s *MySuite) TestPlaneAcceptsOverlappingAABB(c *C) {
    //plane := scene.MakePlane(vecmath.Vec3d{Y: 3}, vecmath.Vec3d{1,0,0})
    //aabb := scene.MakeAABB(-4, -4, -4, 4, 4, 4)

    //test := plane.CheckAABBIntersection(&aabb)

    //c.Assert(test, Equals, true)
//}
