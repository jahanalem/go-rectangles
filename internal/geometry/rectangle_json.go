package geometry

type RectangleJSONModel struct {
	Point1 Point `json:"point1"`
	Point2 Point `json:"point2"`
	Point3 Point `json:"point3"`
	Point4 Point `json:"point4"`
}

func NewRectangleJSONModel(r *Rectangle) RectangleJSONModel {
	orderedPoints := r.GetOrderedPoints()

	return RectangleJSONModel{
		Point1: orderedPoints[0],
		Point2: orderedPoints[1],
		Point3: orderedPoints[2],
		Point4: orderedPoints[3],
	}
}