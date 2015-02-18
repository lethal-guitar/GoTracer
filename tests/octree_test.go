package tests

import (
    "math"
    "github.com/lethal-guitar/go_tracer/scene"
    "github.com/lethal-guitar/go_tracer/vecmath"
    . "github.com/lethal-guitar/go_tracer/spatial"
    . "gopkg.in/check.v1"
)

func (s *MySuite) TestNodeIsInitiallyALeaf(c *C) {
    node := MakeNode(scene.AABB{})
    c.Assert(node.IsLeaf(), Equals, true)
}

func (s *MySuite) TestNodeIsNoLeafAfterSplitting(c *C) {
    node := MakeNode(scene.AABB{})

    node.Split()

    c.Assert(node.IsLeaf(), Equals, false)
}

func gatherBounds(node *OctreeNode) (bounds [8]scene.AABB) {
    for i, child := range node.Children {
        bounds[i] = child.Bounds
    }

    return
}

// This shrinks both AABBs by a small amount before checking, so that
// directly adjacent boxes aren't detected as overlapping
func checkOverlap(a, b scene.AABB) bool {
    const EPS = 0.0001
    e := vecmath.Vec3d{EPS, EPS, EPS}
    minA := vecmath.ToVec3d(&a.Min).Added(e)
    minB := vecmath.ToVec3d(&b.Min).Added(e)
    maxA := vecmath.ToVec3d(&a.Max).Subtracted(e)
    maxB := vecmath.ToVec3d(&b.Max).Subtracted(e)

    adjustedA := scene.MakeAABBV(minA, maxA)
    adjustedB := scene.MakeAABBV(minB, maxB)

    return adjustedA.Overlaps(&adjustedB)
}

func extent(bounds *scene.AABB) vecmath.Vec3d {
    return vecmath.Vec3d{
        bounds.Max.V[0] - bounds.Min.V[0],
        bounds.Max.V[1] - bounds.Min.V[1],
        bounds.Max.V[2] - bounds.Min.V[2],
    }
}

func inList(item scene.Traceable, list []scene.Traceable) bool {
    for _, element := range list {
        if element == item {
            return true
        }
    }

    return false
}

func makeSplitNode() *OctreeNode {
    node := MakeNode(scene.MakeAABBV(
        vecmath.Vec3d{-2,-2,-2},
        vecmath.Vec3d{2,2,2},
    ))

    node.Split()
    return node
}

func expand(v1, v2 *vecmath.Vec3d, comp func(float32, float32) float32) {
    v1.X = comp(v1.X, v2.X)
    v1.Y = comp(v1.Y, v2.Y)
    v1.Z = comp(v1.Z, v2.Z)
}

func (s *MySuite) TestNodeSplittingLeavesOuterBoundsIntact(c *C) {
    node := makeSplitNode()

    actualBounds := gatherBounds(node)

    var totalMin, totalMax vecmath.Vec3d

    for _, bound := range actualBounds {
        min := vecmath.ToVec3d(&bound.Min)
        max := vecmath.ToVec3d(&bound.Max)
        expand(&totalMin, &min, func(a, b float32) float32 {
            return float32(math.Min(float64(a), float64(b)))
        })
        expand(&totalMax, &max, func(a, b float32) float32 {
            return float32(math.Max(float64(a), float64(b)))
        })
    }

    c.Assert(totalMin, Equals, vecmath.Vec3d{-2,-2,-2})
    c.Assert(totalMax, Equals, vecmath.Vec3d{2,2,2})
}

func (s *MySuite) TestChildNodesHaveHalfExtent(c *C) {
    node := makeSplitNode()

    actualBounds := gatherBounds(node)

    for _, bound := range actualBounds {
        boundExtent := extent(&bound)
        c.Assert(boundExtent, Equals, vecmath.Vec3d{2,2,2})
    }
}

func (s *MySuite) TestChildNodesAreDisjoint(c *C) {
    node := makeSplitNode()

    actualBounds := gatherBounds(node)

    for iOuter, bound := range actualBounds {
        for iInner, otherBound := range actualBounds {
            overlaps := iOuter != iInner && checkOverlap(bound, otherBound)
            c.Assert(overlaps, Equals, false)
        }
    }
}

func (s *MySuite) TestInsertingObjectSplitsEmptyTree(c *C) {
    tree := MakeOctree(scene.MakeAABB(0,0,0,4,4,4))
    sphere := scene.MakeSimpleSphere(vecmath.Vec3d{1.5,1.5,1.5}, 0.5)
    object := scene.Traceable(sphere)

    tree.InsertObject(&object)

    c.Assert(tree.Root.IsLeaf(), Equals, false)
}
