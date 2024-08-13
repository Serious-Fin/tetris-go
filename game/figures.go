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
