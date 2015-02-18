package tracing

import (
    "image"
    "image/color"
    "image/draw"
    "math"
    "github.com/lethal-guitar/go_tracer/scene"
    //"github.com/lethal-guitar/go_tracer/spatial"
    "github.com/lethal-guitar/go_tracer/vecmath"
)

const MAX_RECURSION int = 2
const REFRACTION_INDEX_AIR float32 = 1.0

type CameraConfig struct {
    Left, Right, Top, Bottom float32
    PointOfView vecmath.Vec3d
}

type RayTracer struct {
    Width, Height int
    Camera CameraConfig
    Scene *Scene

    engine Engine
}

func iMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func (self *RayTracer) Render(tileSize int) image.Image {
    self.initEngine()

    // ALTERNATIVE: Bypass parallelization
    //return self.renderTile(image.Rectangle{
        //image.Point{0, 0},
        //image.Point{self.Width, self.Height},
    //})

    // Channel for parallel rendering
    transport := make(chan image.Image)

    // Subdivide target image into equally sized tiles and start
    // worker go routines
    var tiles []image.Image
    tilesX := int(math.Ceil(float64(self.Width) / float64(tileSize)))
    tilesY := int(math.Ceil(float64(self.Height) / float64(tileSize)))

    for y:=0; y < tilesY; y++ {
        for x:=0; x < tilesX; x++ {
            startX := x * tileSize
            startY := y * tileSize

            sizeX := iMin(tileSize, self.Width - startX)
            sizeY := iMin(tileSize, self.Height - startY)

            go self.renderTileToChannel(
                image.Rectangle{
                    image.Point{startX, startY},
                    image.Point{startX + sizeX, startY + sizeY},
                },
                transport,
            )
        }
    }

    // Now collect results from channel
    numTiles := tilesX * tilesY
    for i:=0; i<numTiles; i++ {
        tile := <-transport
        tiles = append(tiles, tile)
    }

    return self.combineTiles(tiles)
}

func (self *RayTracer) initEngine() {
    octree := self.Scene.ToOctree()
    self.engine = &octree

    // ALTERNATIVE: Use naive ray tracing
    //container := spatial.ObjectContainer{}
    //for _, object := range self.Scene.Traceables {
        //container.Insert(object)
    //}
    //self.engine = &container
}

func (self *RayTracer) combineTiles(tiles []image.Image) image.Image {
    combined := image.NewRGBA(
        image.Rectangle{
            image.Point{0, 0}, image.Point{self.Width, self.Height}})

    for _, tile := range tiles {
        draw.Draw(combined, tile.Bounds(), tile, tile.Bounds().Min, draw.Src)
    }

    return combined
}

func (self *RayTracer) renderTileToChannel(
    tile image.Rectangle,
    sink chan image.Image,
) {
    output := self.renderTile(tile)
    sink <- output
}

func (self *RayTracer) renderTile(dimensions image.Rectangle) image.Image {
    output := image.NewRGBA(dimensions)

    for y := dimensions.Min.Y; y < dimensions.Max.Y; y++ {
        for x := dimensions.Min.X; x < dimensions.Max.X; x++ {
            output.Set(x, y, self.renderPixel(x, y))
        }
    }

    return output
}

func makeRay(start, end *vecmath.Vec3d) scene.Ray {
    direction := vecmath.MakeDirectionVector(start, end)
    return scene.Ray{Origin: *start, Direction: direction}
}

func (self *RayTracer) rayForPixel(x, y int) scene.Ray {
    dx := float32(math.Abs(float64(self.Camera.Right - self.Camera.Left)))
    dy := float32(math.Abs(float64(self.Camera.Bottom - self.Camera.Top)))
    dxPixel := dx / float32(self.Width)
    dyPixel := dy / float32(self.Height)
    tx := self.Camera.Left + float32(x)*dxPixel
    ty := self.Camera.Top + float32(y)*dyPixel

    pointOnViewPlane := vecmath.Vec3d{tx, ty, 0.0}
    return makeRay(&self.Camera.PointOfView, &pointOnViewPlane)
}

func (self *RayTracer) renderPixel(x, y int) color.Color {
    ray := self.rayForPixel(x, y)

    foundObject, color := self.renderRay(&ray, 0)

    if foundObject {
        return color
    }
    return self.Scene.BackgroundColor
}

func (self *RayTracer) renderRay(ray *scene.Ray, recursionDepth int) (
    foundObject bool, color scene.FloatColor,
) {
    object, distance := self.engine.FindClosestObject(ray)

    if object != nil {
        intersection := ray.PointAt(distance)
        color =
            self.determineColor(object, ray, &intersection, recursionDepth)

    }

    return object != nil, color
}

func makeLightVectorOptimized(intersection, lightPos *vecmath.Vec3d) (
    lightVector vecmath.Vec3d,
    lightDistance float32,
) {
    vector := lightPos.Subtracted(*intersection)
    lightDistance = vector.Length()

    // Conceptually, this is just normalizing the vector, but by reusing the
    // already calculated length, we save one (expensive) length calculation.
    // This actually makes a difference due to the huge number of occlusion
    // check rays that are cast.
    lightVector = vector.Scaled(1.0 / lightDistance)
    return
}

func (self *RayTracer) determineColor(
    pObject *scene.Traceable,
    ray *scene.Ray,
    intersectionPoint *vecmath.Vec3d,
    recursionDepth int,
) (finalColor scene.FloatColor) {
    object := *pObject
    material := object.Material()
    normal := object.Normal(intersectionPoint)

    // Primary color
    //////////////////////////////////////////////////////////////////////////
    viewVector := ray.Direction.Scaled(-1)
    var totalAmbient, totalDiffuse, totalSpecular scene.FloatColor

    for _, light := range self.Scene.Lights {
        lightVector, lightDistance :=
            makeLightVectorOptimized(intersectionPoint, &light.Position)
        lightRay := scene.Ray{*intersectionPoint, lightVector}

        inShadow := self.engine.FindOccluder(&lightRay, lightDistance)

        if !inShadow {
            diffuseFactor, specularFactor := self.calculatePhongFactors(
                intersectionPoint,
                &viewVector,
                &lightVector,
                &normal,
                material.Specularity,
            )
            // TODO: Add ambient element
            //totalAmbient.Add()
            totalDiffuse.Add(
                material.Diffuse.Multiplied(light.Diffuse.Scaled(diffuseFactor)))
            totalSpecular.Add(
                material.Specular.Multiplied(light.Specular.Scaled(specularFactor)))
        }
    }
    step := totalAmbient.Added(totalDiffuse)
    primary := step.Added(totalSpecular)
    finalColor.Add(primary)


    // Reflection
    //////////////////////////////////////////////////////////////////////////

    if material.Reflectivity > 0.0 && recursionDepth < MAX_RECURSION {
        reflection := vecmath.MakeReflect(&ray.Direction, &normal).Scaled(-1)
        reflectionRay := scene.Ray{*intersectionPoint, reflection}

        foundObject, reflectionColor :=
            self.renderRay(&reflectionRay, recursionDepth + 1)

        if foundObject {
            finalColor.Blend(material.Reflectivity, reflectionColor)
        }
    }

    // Refraction
    //////////////////////////////////////////////////////////////////////////

    // TBD.
    // Refraction requires being able to determine whether a ray 

    //if material.Transparency > 0.0 && recursionDepth < MAX_RECURSION {
        //index := material.RefractionIndex
        //n := REFRACTION_INDEX_AIR / index

        //if (*object).IsInside(&ray.Origin) {
            //normal = normal.Scaled(-1)
            //n = index / REFRACTION_INDEX_AIR
        //}

        //cosI := -normal.Dot(ray.Direction)
        //cosT2 := 1.0 - n * n * (1.0 - cosI*cosI)

        //if cosT2 > 0.0 {
            //refractionVector :=
                //ray.Direction.Scaled(n).Added(
                    //normal.Scaled(
                        //n * cosI - float32(math.Sqrt(float64(cosT2)))))

            //start := intersectionPoint.Added(refractionVector.Scaled(EPSILON))
            //refractionRay := scene.Ray{start, refractionVector}

            //foundObject, refractionColor :=
                //self.renderRay(&refractionRay, recursionDepth + 1)

            //if foundObject {
                //// Beer's law (C++ version, from 3rd party source)
                ////Color absorbance = prim->GetMaterial()->GetColor() * 0.15f * -dist;
                ////Color transparency =
                ////    Color( expf( absorbance.r ), expf( absorbance.g ), expf( absorbance.b ) );
                //finalColor.Blend(material.Transparency, refractionColor)
            //}
        //}
    //}

    return finalColor
}

func (self *RayTracer) calculatePhongFactors(
    intersectionPoint *vecmath.Vec3d,
    viewVector *vecmath.Vec3d,
    lightVector *vecmath.Vec3d,
    normal *vecmath.Vec3d,
    specularity float32,
) (
    diffuse, specular float32,
) {
    diffuse = float32(math.Max(0, float64(lightVector.Dot(*normal))))
    if diffuse > 0.0 {
        lightReflectVector := vecmath.MakeReflect(lightVector, normal)
        specularDot := lightReflectVector.Dot(*viewVector)
        if specularDot > 0.0 {
            specular = float32(
                math.Pow(float64(specularDot), float64(specularity)))
        }
    }
    return
}
