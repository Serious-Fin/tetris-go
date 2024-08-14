package game

import (
	"math/rand"
)

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
	BlockType     int
}

var figures = []Figure{FigureI, FigureJ, FigureL, FigureO, FigureS, FigureT, FigureZ}

func GetRandomFigure() *Figure {
	return &figures[rand.Intn(len(figures))]
}

var FigureT = Figure{
	BlockType:     CellT,
	GeometryIndex: 0,
	MiddlePos: Point{
		Row: -1,
		Col: 4,
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

var FigureO = Figure{
	BlockType:     CellO,
	GeometryIndex: 0,
	MiddlePos: Point{
		Row: -1,
		Col: 4,
	},
	Geometries: []Geometry{
		{
			[]Point{
				{
					Row: 0,
					Col: 0,
				},
				{
					Row: 0,
					Col: +1,
				},
				{
					Row: +1,
					Col: 0,
				},
				{
					Row: +1,
					Col: +1,
				},
			},
		},
	},
}

var FigureS = Figure{
	BlockType:     CellS,
	GeometryIndex: 0,
	MiddlePos: Point{
		Row: -1,
		Col: 4,
	},
	Geometries: []Geometry{
		{
			[]Point{
				{
					Row: 0,
					Col: 0,
				},
				{
					Row: 0,
					Col: +1,
				},
				{
					Row: +1,
					Col: 0,
				},
				{
					Row: +1,
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
					Row: +1,
					Col: 0,
				},
				{
					Row: 0,
					Col: -1,
				},
				{
					Row: -1,
					Col: -1,
				},
			},
		},
	},
}

var FigureZ = Figure{
	BlockType:     CellZ,
	GeometryIndex: 0,
	MiddlePos: Point{
		Row: -1,
		Col: 4,
	},
	Geometries: []Geometry{
		{
			[]Point{
				{
					Row: 0,
					Col: 0,
				},
				{
					Row: 0,
					Col: -1,
				},
				{
					Row: +1,
					Col: 0,
				},
				{
					Row: +1,
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
					Row: -1,
					Col: 0,
				},
				{
					Row: 0,
					Col: -1,
				},
				{
					Row: +1,
					Col: -1,
				},
			},
		},
	},
}

var FigureI = Figure{
	BlockType:     CellI,
	GeometryIndex: 0,
	MiddlePos: Point{
		Row: -3,
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
					Row: +1,
					Col: 0,
				},
				{
					Row: +2,
					Col: 0,
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
					Row: 0,
					Col: +1,
				},
				{
					Row: 0,
					Col: -1,
				},
				{
					Row: 0,
					Col: -2,
				},
			},
		},
	},
}

var FigureL = Figure{
	BlockType:     CellL,
	GeometryIndex: 0,
	MiddlePos: Point{
		Row: -2,
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
					Row: +1,
					Col: 0,
				},
				{
					Row: +1,
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
					Row: 0,
					Col: +1,
				},
				{
					Row: 0,
					Col: -1,
				},
				{
					Row: +1,
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
					Row: +1,
					Col: 0,
				},
				{
					Row: -1,
					Col: 0,
				},
				{
					Row: -1,
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
					Row: 0,
					Col: -1,
				},
				{
					Row: 0,
					Col: +1,
				},
				{
					Row: -1,
					Col: +1,
				},
			},
		},
	},
}

var FigureJ = Figure{
	BlockType:     CellJ,
	GeometryIndex: 0,
	MiddlePos: Point{
		Row: -2,
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
					Row: +1,
					Col: 0,
				},
				{
					Row: +1,
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
					Row: 0,
					Col: +1,
				},
				{
					Row: 0,
					Col: -1,
				},
				{
					Row: -1,
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
					Row: +1,
					Col: 0,
				},
				{
					Row: -1,
					Col: 0,
				},
				{
					Row: -1,
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
					Row: 0,
					Col: -1,
				},
				{
					Row: 0,
					Col: +1,
				},
				{
					Row: +1,
					Col: +1,
				},
			},
		},
	},
}
