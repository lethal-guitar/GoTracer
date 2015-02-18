package tests

import (
    "testing"
    . "github.com/lethal-guitar/go_tracer/vecmath"
    . "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }
type MySuite struct {}
var _ = Suite(&MySuite{})

func (s *MySuite) TestAddingTwoVectorsAddsTheElements(c *C) {
    v1 := Vec3d{2.0, 1.0, 3.0}
    v2 := Vec3d{1.0, -1.0, 2.0}

    result := v1.Added(v2)

    c.Assert(result, Equals, Vec3d{3.0, 0.0, 5.0})
}

func (s *MySuite) TestLengthIsCalculatedCorrectly(c *C) {
    vec := Vec3d{3.0, 0.0, 0.0}

    c.Assert(vec.Length(), Equals, float32(3.0))
}

func (s *MySuite) TestLengthIsOneAfterNormalization(c *C) {
    vec := Vec3d{4.0, 1.0, 3.0}
    normalized := vec.Normalized()

    c.Assert(normalized.Length(), Equals, float32(1.0))
}
