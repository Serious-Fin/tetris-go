package game

type Point struct {
	Row int
	Col int
}

type Geometry struct {
	Points []Point
}

type Figure struct {
	MiddlePos     Point
	Geometries    []Geometry
	GeometryIndex int
}

var FigureT = Figure{
	GeometryIndex: 0,
	MiddlePos: Point{
		Row: 3,
		Col: 5,
	},
	Geometries: []Geometry{
		{
			[]Point{
				{
					Row: 0,
					Col: 0,
				},
				{
					Row: -1,
					Col: 0,
				},
				{
					Row: 0,
					Col: +1,
				},
				{
					Row: 0,
					Col: -1,
				},
			},
		},
		{
			[]Point{
				{
					Row: 0,
					Col: 0,
				},
				{
					Row: -1,
					Col: 0,
				},
				{
					Row: +1,
					Col: 0,
				},
				{
					Row: 0,
					Col: +1,
				},
			},
		},
		{
			[]Point{
				{
					Row: 0,
					Col: 0,
				},
				{
					Row: +1,
					Col: 0,
				},
				{
					Row: 0,
					Col: +1,
				},
				{
					Row: 0,
					Col: -1,
				},
			},
		},
		{
			[]Point{
				{
					Row: 0,
					Col: 0,
				},
				{
					Row: -1,
					Col: 0,
				},
				{
					Row: +1,
					Col: 0,
				},
				{
					Row: 0,
					Col: -1,
				},
			},
		},
	},
}

func (g *GameBoard) MoveDown(f *Figure) {
	f.MiddlePos.Row += 1
}

func (g *GameBoard) MoveLeft(f *Figure) {
	if g.leftWallCollisionDetected(f) {
		return
	}

	// TODO: Check collision with any blocks on the left

	f.MiddlePos.Col -= 1
}

func (g *GameBoard) MoveRight(f *Figure) {
	if g.rightWallCollisionDetected(f) {
		return
	}

	// TODO: Check collision with any blocks on the right

	f.MiddlePos.Col += 1
}

func (g *GameBoard) Rotate(f *Figure) {
	// TODO: rotation collision check
	f.GeometryIndex = (f.GeometryIndex + 1) % len(f.Geometries)
}

func (f *Figure) findLeftmostOffset() int {
	leftmostOffset := 0
	for _, point := range f.Geometries[f.GeometryIndex].Points {
		if point.Col < leftmostOffset {
			leftmostOffset = point.Col
		}
	}
	return leftmostOffset
}

func (f *Figure) findRightmostOffset() int {
	rightmostOffset := 0
	for _, point := range f.Geometries[f.GeometryIndex].Points {
		if point.Col > rightmostOffset {
			rightmostOffset = point.Col
		}
	}
	return rightmostOffset
}

func (g *GameBoard) leftWallCollisionDetected(f *Figure) bool {
	leftmostOffset := f.findLeftmostOffset()
	return f.MiddlePos.Col+leftmostOffset-1 < 0
}

func (g *GameBoard) rightWallCollisionDetected(f *Figure) bool {
	rightmostOffset := f.findRightmostOffset()
	return f.MiddlePos.Col+rightmostOffset+1 >= g.Width
}
