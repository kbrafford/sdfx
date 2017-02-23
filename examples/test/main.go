package main

import (
	"fmt"
	"math"

	"github.com/deadsy/sdfx/sdf"
)

func test1() {
	s0 := sdf.Box2D(sdf.V2{0.8, 1.2}, 0.05)
	s1 := sdf.RevolveTheta3D(s0, sdf.DtoR(225))
	sdf.RenderPNG(s1, true)
}

func test2() {
	s0 := sdf.Box2D(sdf.V2{0.8, 1.2}, 0.1)
	s1 := sdf.Extrude3D(s0, 0.3)
	sdf.RenderPNG(s1, true)
}

func test3() {
	s0 := sdf.Circle2D(0.1)
	s0 = sdf.Transform2D(s0, sdf.Translate2d(sdf.V2{1, 0}))
	s1 := sdf.Revolve3D(s0)
	sdf.RenderPNG(s1, true)
}

func test4() {
	s0 := sdf.Box2D(sdf.V2{0.2, 0.4}, 0.05)
	s0 = sdf.Transform2D(s0, sdf.Translate2d(sdf.V2{1, 0}))
	s1 := sdf.RevolveTheta3D(s0, sdf.DtoR(270))
	sdf.RenderPNG(s1, true)
}

func test5() {
	s0 := sdf.Box2D(sdf.V2{0.2, 0.4}, 0.05)
	s0 = sdf.Transform2D(s0, sdf.Rotate2d(sdf.DtoR(45)))
	s0 = sdf.Transform2D(s0, sdf.Translate2d(sdf.V2{1, 0}))
	s1 := sdf.RevolveTheta3D(s0, sdf.DtoR(315))

	sdf.RenderSTL(s1, 200, "test.stl")
	//sdf.RenderPNG(s1, true)
}

func test6() {
	s0 := sdf.Sphere3D(0.5)
	d := 0.4
	s1 := sdf.Transform3D(s0, sdf.Translate3d(sdf.V3{0, d, 0}))
	s2 := sdf.Transform3D(s0, sdf.Translate3d(sdf.V3{0, -d, 0}))
	s3 := sdf.Union3D(s1, s2)
	s3.(*sdf.UnionSDF3).SetMin(sdf.PolyMin, 0.1)
	sdf.RenderPNG(s3, true)
}

func test7() {
	s0 := sdf.Box3D(sdf.V3{0.8, 0.8, 0.05}, 0)
	s1 := sdf.Transform3D(s0, sdf.Rotate3d(sdf.V3{1, 0, 0}, sdf.DtoR(60)))
	s2 := sdf.Union3D(s0, s1)
	s2.(*sdf.UnionSDF3).SetMin(sdf.PolyMin, 0.1)
	s3 := sdf.Transform3D(s2, sdf.Rotate3d(sdf.V3{0, 0, 1}, sdf.DtoR(-30)))
	sdf.RenderPNG(s3, true)
}

func test8() {
	a := sdf.V3{0, 0, 0}
	b := sdf.V3{1, 0, 0}
	c := sdf.V3{0, 1, 0}
	d := sdf.V3{0, 0, 1}
	t1 := sdf.NewTriangle(a, b, d)
	t2 := sdf.NewTriangle(a, c, b)
	t3 := sdf.NewTriangle(a, d, c)
	t4 := sdf.NewTriangle(b, c, d)
	m := sdf.NewMesh([]*sdf.Triangle{t1, t2, t3, t4})
	err := sdf.SaveSTL("test.stl", m)
	if err != nil {
		fmt.Printf("%s", err)
	}
}

func test9() {
	s := sdf.Sphere3D(10.0)
	sdf.RenderSTL(s, 200, "test.stl")
}

func test10() {
	s0 := sdf.Box3D(sdf.V3{0.8, 0.8, 0.05}, 0)
	s1 := sdf.Transform3D(s0, sdf.Rotate3d(sdf.V3{1, 0, 0}, sdf.DtoR(60)))
	s := sdf.Union3D(s0, s1)
	s.(*sdf.UnionSDF3).SetMin(sdf.PolyMin, 0.1)
	sdf.RenderSTL(s, 200, "test.stl")
}

func test11() {
	s := sdf.Capsule3D(0.3, 1.4)
	sdf.RenderSTL(s, 200, "test.stl")
}

func test12() {
	k := 0.1
	points := []sdf.V2{
		sdf.V2{0, -k},
		sdf.V2{k, k},
		sdf.V2{-k, k},
	}
	s0 := sdf.Polygon2D(points)
	s0 = sdf.Transform2D(s0, sdf.Translate2d(sdf.V2{0.8, 0}))
	s1 := sdf.RevolveTheta3D(s0, sdf.DtoR(360))
	sdf.RenderSTL(s1, 200, "test.stl")
	//sdf.RenderPNG(s1, true)
}

func test13() {
	k := 0.4
	s0 := sdf.Polygon2D([]sdf.V2{sdf.V2{k, -k}, sdf.V2{k, k}, sdf.V2{-k, k}, sdf.V2{-k, -k}})
	s0 = sdf.Transform2D(s0, sdf.Translate2d(sdf.V2{0.8, 0}))
	s1 := sdf.RevolveTheta3D(s0, sdf.DtoR(270))
	sdf.RenderSTL(s1, 200, "test.stl")
	//sdf.RenderPNG(s1, true)
}

func test14() {

	// size
	a := 0.3
	b := 0.7
	// rotation
	theta := 30.0
	c := math.Cos(sdf.DtoR(theta))
	s := math.Sin(sdf.DtoR(theta))
	// translate
	j := 10.0
	k := 2.0

	points := []sdf.V2{
		sdf.V2{j + c*a - s*b, k + s*a + c*b},
		sdf.V2{j - c*a - s*b, k - s*a + c*b},
		sdf.V2{j - c*a + s*b, k - s*a - c*b},
		sdf.V2{j + c*a + s*b, k + s*a - c*b},
	}

	s0 := sdf.Polygon2D(points)
	s1 := sdf.RevolveTheta3D(s0, sdf.DtoR(300))

	sdf.RenderSTL(s1, 200, "test.stl")
}

func test15() {
	// size
	a := 1.0
	b := 1.0
	// rotation
	theta := 0.0
	// translate
	j := 3.0
	k := 0.0

	points := []sdf.V2{
		sdf.V2{0, -b},
		sdf.V2{a, b},
		sdf.V2{-a, b},
	}

	s0 := sdf.Polygon2D(points)
	s0 = sdf.Transform2D(s0, sdf.Rotate2d(sdf.DtoR(theta)))
	s0 = sdf.Transform2D(s0, sdf.Translate2d(sdf.V2{j, k}))

	s1 := sdf.RevolveTheta3D(s0, sdf.DtoR(300))
	s1 = sdf.Transform3D(s1, sdf.Rotate3d(sdf.V3{0, 0, 1}, sdf.DtoR(30)))

	sdf.RenderSTL(s1, 200, "test.stl")
}

func test16() {
	// size
	a0 := 1.3
	b0 := 0.4
	a1 := 1.3
	b1 := 1.3
	c := 0.8
	// rotation
	theta := 20.0
	// translate
	j := 4.0
	k := 0.0

	points := []sdf.V2{
		sdf.V2{b0, -c},
		sdf.V2{a0, c},
		sdf.V2{-a1, c},
		sdf.V2{-b1, -c},
	}

	s0 := sdf.Polygon2D(points)
	s0 = sdf.Transform2D(s0, sdf.Rotate2d(sdf.DtoR(theta)))
	s0 = sdf.Transform2D(s0, sdf.Translate2d(sdf.V2{j, k}))

	s1 := sdf.RevolveTheta3D(s0, sdf.DtoR(300))
	s1 = sdf.Transform3D(s1, sdf.Rotate3d(sdf.V3{0, 0, 1}, sdf.DtoR(30)))

	sdf.RenderSTL(s1, 200, "test.stl")
}

func test17() {
	// size
	a := 1.3
	b := 0.4
	// translate
	j := 3.0
	k := 0.0

	points := []sdf.V2{
		sdf.V2{a, 0},
		sdf.V2{-a, b},
		sdf.V2{-a, -b},
	}

	s0 := sdf.Polygon2D(points)
	s0 = sdf.Transform2D(s0, sdf.Translate2d(sdf.V2{j, k}))

	s1 := sdf.RevolveTheta3D(s0, sdf.DtoR(300))
	s1 = sdf.Transform3D(s1, sdf.Rotate3d(sdf.V3{0, 0, 1}, sdf.DtoR(30)))

	sdf.RenderSTL(s1, 200, "test.stl")
}

func test18() {

	r0 := 10.0
	r1 := 8.0
	r2 := 7.5
	r3 := 9.0

	h0 := 4.0
	h1 := 6.0
	h2 := 5.5
	h3 := 3.5
	h4 := 1.0

	points := []sdf.V2{
		sdf.V2{0, 0},
		sdf.V2{r0, 0},
		sdf.V2{r0, h0},
		sdf.V2{r1, h1},
		sdf.V2{r2, h2},
		sdf.V2{r3, h3},
		sdf.V2{r3, h4},
		sdf.V2{0, h4},
	}

	s0 := sdf.Polygon2D(points)
	s1 := sdf.RevolveTheta3D(s0, sdf.DtoR(300))
	s1 = sdf.Transform3D(s1, sdf.Rotate3d(sdf.V3{0, 0, 1}, sdf.DtoR(30)))

	sdf.RenderSTL(s1, 200, "test.stl")
}

func test19() {
	r := 2.0
	k := 1.9
	s0 := sdf.Circle2D(r)
	s1 := sdf.Array2D(s0, sdf.V2i{3, 7}, sdf.V2{k * r, k * r})
	s1.(*sdf.ArraySDF2).SetMin(sdf.PolyMin, 0.8)
	s2 := sdf.Extrude3D(s1, 1.0)
	sdf.RenderSTL(s2, 200, "test.stl")
}

func test20() {
	r := 4.0
	d := 20.0
	s0 := sdf.Circle2D(r)
	s0 = sdf.Transform2D(s0, sdf.Translate2d(sdf.V2{d, 0}))
	s0 = sdf.Rotate2D(s0, 5, sdf.Rotate2d(sdf.DtoR(20)))
	s0.(*sdf.RotateSDF2).SetMin(sdf.PolyMin, 1.2)
	s1 := sdf.Extrude3D(s0, 10.0)
	sdf.RenderSTL(s1, 200, "test.stl")
}

func test21() {
	r := 2.0
	k := 1.9
	s0 := sdf.Sphere3D(r)
	s1 := sdf.Array3D(s0, sdf.V3i{3, 7, 5}, sdf.V3{k * r, k * r, k * r})
	s1.(*sdf.ArraySDF3).SetMin(sdf.PolyMin, 0.8)
	sdf.RenderSTL(s1, 200, "test.stl")
}

func test22() {
	r := 4.0
	d := 20.0
	s0 := sdf.Sphere3D(r)
	s0 = sdf.Transform3D(s0, sdf.Translate3d(sdf.V3{d, 0, 0}))
	s0 = sdf.Rotate3D(s0, 5, sdf.Rotate3d(sdf.V3{0, 0, 1}, sdf.DtoR(20)))
	s0.(*sdf.RotateSDF3).SetMin(sdf.PolyMin, 1.2)
	sdf.RenderSTL(s0, 200, "test.stl")
}

func test23() {
	s0 := sdf.Circle2D(4.0)
	sdf.SDF2_RenderPNG(s0, "test.png")
}

func test24() {
	s0 := sdf.Box2D(sdf.V2{50, 70}, 4)
	sdf.SDF2_RenderPNG(s0, "test.png")
}

func test25() {
	s0 := sdf.Circle2D(0.0)
	var s1 sdf.SDF2
	for i := 0; i < 50; i++ {
		s1 = sdf.Union2D(s1, sdf.Transform2D(s0, sdf.Translate2d(sdf.RandomV2(-10, 10))))
	}
	sdf.SDF2_RenderPNG(s1, "test.png")
}

func test26() {
	s := sdf.Cylinder3D(5, 2, 1)
	sdf.RenderSTL(s, 200, "test.stl")
}

func test27() {
	r := 5.0
	posn := []sdf.V2{sdf.V2{2 * r, 2 * r}, sdf.V2{-r, r}, sdf.V2{r, -r}, sdf.V2{-r, -r}, sdf.V2{0, 0}}
	s := sdf.MultiCylinder3D(3, 1, posn)
	sdf.RenderSTL(s, 200, "test.stl")
}

func test28() {
	s := sdf.Cone3D(20, 12, 8, 2)
	sdf.RenderSTL(s, 200, "test.stl")
}

func test29() {
	s0 := sdf.Line2D(10, 3)
	s1 := sdf.Extrude3D(s0, 4)
	sdf.RenderSTL(s1, 200, "test.stl")
}

func test30() {
	s0 := sdf.Line2D(10, 3)
	s0 = sdf.Cut2D(s0, sdf.V2{4, 0}, sdf.V2{1, 1})
	s1 := sdf.Extrude3D(s0, 4)
	sdf.RenderSTL(s1, 200, "test.stl")
}

func test31() {
	s := sdf.CounterSunk_Hole3d(30, 2)
	sdf.RenderSTL(s, 200, "test.stl")
}

func test32() {
	s0, err := sdf.MakeFlatFlankCam(0.094, sdf.DtoR(2.0*57.5), 0.625)
	if err != nil {
		panic(err)
	}
	s1 := sdf.Extrude3D(s0, 0.1)
	sdf.RenderSTL(s1, 200, "test.stl")
}

func test33() {
	s0 := sdf.ThreeArcCam2D(30, 20, 5, 50000)
	fmt.Printf("%+v\n", s0)
	s1 := sdf.Extrude3D(s0, 4)
	sdf.RenderSTL(s1, 200, "test.stl")
}

func test34() {
	s0, err := sdf.MakeThreeArcCam(0.1, sdf.DtoR(2.0*80), 0.7, 1.1)
	if err != nil {
		panic(err)
	}
	s1 := sdf.Extrude3D(s0, 0.1)
	sdf.RenderSTL(s1, 200, "test.stl")
}

func test35() {
	r := 7.0
	d := 20.0
	s0 := sdf.Line2D(r, 1.0)
	s0 = sdf.Transform2D(s0, sdf.Translate2d(sdf.V2{d, 0}))
	s0 = sdf.RotateCopy2D(s0, 15)
	s1 := sdf.Extrude3D(s0, 10.0)
	sdf.RenderSTL(s1, 200, "test.stl")
}

func test36() {
	s_driver, s_driven, err := sdf.MakeGenevaCam(6, 100, 40, 80, 5, 0.5)
	if err != nil {
		panic(err)
	}
	sdf.RenderSTL(sdf.Extrude3D(s_driver, 10), 200, "driver.stl")
	sdf.RenderSTL(sdf.Extrude3D(s_driven, 10), 200, "driven.stl")
}

func test37() {
	r := 5.0
	p := 2.0
	s := sdf.Screw3D(sdf.ISOThread(r, p), 50, p, 1)
	sdf.RenderSTL(s, 400, "screw.stl")
}

func test38() {
	fmt.Printf("%+v\n", sdf.ThreadLookup("unc_1/4"))
	fmt.Printf("%+v\n", sdf.ThreadLookup("blah"))
}

func test39() {
	s0 := sdf.NewFlange1(30, 20, 10)
	fmt.Printf("%+v\n", s0)
	s1 := sdf.Extrude3D(s0, 5)
	sdf.RenderSTL(s1, 200, "test.stl")
}

func test40() {
	d := 30.0
	wall := 5.0
	s0 := sdf.Box3D(sdf.V3{d, d, d}, wall/2)
	s1 := sdf.Box3D(sdf.V3{d - wall, d - wall, d}, wall/2)
	s1 = sdf.Transform3D(s1, sdf.Translate3d(sdf.V3{0, 0, wall / 2}))
	s := sdf.Difference3D(s0, s1)
	s.(*sdf.DifferenceSDF3).SetMax(sdf.PolyMax, 2)
	sdf.RenderSTL(s, 200, "test.stl")
}

func test41() {
	s0 := sdf.Cylinder3D(20.0, 5.0, 0)
	s1 := sdf.Slice2D(s0, sdf.V3{0, 0, 0}, sdf.V3{0, 1, 1})
	s2 := sdf.Revolve3D(s1)
	sdf.RenderSTL(s2, 200, "test.stl")
}

func main() {
	test41()
}
