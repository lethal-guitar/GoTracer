package spatial

import (
    "github.com/lethal-guitar/go_tracer/scene"
    "github.com/lethal-guitar/go_tracer/vecmath"
)

type OctreeNode struct {
    Bounds scene.AABB
    Children [8]*OctreeNode
    isLeaf bool
    Objects *ObjectContainer
}

func MakeNode(bounds scene.AABB) *OctreeNode {
    return &OctreeNode{
        Bounds: bounds,
        isLeaf: true,
        Objects: &ObjectContainer{},
    }
}

func (self *OctreeNode) IsLeaf() bool {
    return self.isLeaf
}

func (self *OctreeNode) AddObject(object *scene.Traceable) {
    self.Objects.Insert(object)

}

func (self *OctreeNode) createChildrenLayer (
    baseIndex int,
    baseMin, baseMax vecmath.Vec3d,
    offsetX, offsetZ float32,
) {
    adjustX := vecmath.Vec3d{X: offsetX}
    adjustZ := vecmath.Vec3d{Z: offsetZ}
    adjustXZ := vecmath.Vec3d{X: offsetX, Z: offsetZ}


    nodes := [4]*OctreeNode{
        MakeNode(scene.MakeAABBV(baseMin, baseMax)),
        MakeNode(scene.MakeAABBV(baseMin.Added(adjustX), baseMax.Added(adjustX))),
        MakeNode(scene.MakeAABBV(baseMin.Added(adjustZ), baseMax.Added(adjustZ))),
        MakeNode(scene.MakeAABBV(baseMin.Added(adjustXZ), baseMax.Added(adjustXZ))),
    }

    for i, _ := range nodes {
        self.Children[baseIndex + i] = nodes[i]
    }
}

func (self *OctreeNode) Split() {
    center := self.Bounds.Center()
    offset := center.Subtracted(vecmath.ToVec3d(&self.Bounds.Min))

    self.createChildrenLayer(
        0,
        vecmath.ToVec3d(&self.Bounds.Min),
        center,
        offset.X, offset.Z,
    )
    adjustY := vecmath.Vec3d{Y: offset.Y}
    self.createChildrenLayer(
        4,
        vecmath.ToVec3d(&self.Bounds.Min).Added(adjustY),
        center.Added(adjustY),
        offset.X, offset.Z,
    )

    self.isLeaf = false
}
