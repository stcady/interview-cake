package main

type Overlap struct {
	StartPoint int
	Length     int
}

type Rectangle struct {
	LeftX   int
	Width   int
	BottomY int
	Height  int
}

func findRangeOverlap(point1 int, length1 int, point2 int, length2 int) *Overlap {
	highestStartPoint := max(point1, point2)
	lowestEndPoint := min(point1+length1, point2+length2)
	if highestStartPoint > lowestEndPoint {
		return &Overlap{}
	}
	return &Overlap{
		StartPoint: highestStartPoint,
		Length:     lowestEndPoint - highestStartPoint,
	}
}

func findRectangularOverlap(rect1 *Rectangle, rect2 *Rectangle) *Rectangle {
	xOverlap := findRangeOverlap(rect1.LeftX, rect1.Width, rect2.LeftX, rect2.Width)
	yOverlap := findRangeOverlap(rect1.BottomY, rect1.Height, rect2.BottomY, rect2.Width)
	if xOverlap == (&Overlap{}) || yOverlap == (&Overlap{}) {
		return &Rectangle{}
	}
	return &Rectangle{
		LeftX:   xOverlap.StartPoint,
		Width:   xOverlap.Length,
		BottomY: yOverlap.StartPoint,
		Height:  yOverlap.Length,
	}
}

// Max returns the larger of x or y.
func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
