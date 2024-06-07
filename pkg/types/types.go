package types

type CoordinateSet2D [2]int

type Plot struct {
	Corner1 CoordinateSet2D `json:"corner1"`
	Corner2 CoordinateSet2D `json:"corner2"`
}
