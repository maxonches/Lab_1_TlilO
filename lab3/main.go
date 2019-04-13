package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Point struct {
	X 				float64
	Y 				float64
	Z 				float64
	RadialDistance	float64
	ZenithAngle		float64
	AzimuthAngle	float64
}

type Sphere struct {
	Radius	float64
	Center	Point
}

type Game struct {
	Apoints		[]Point
	Bpoints		  Point
	PointsNum	  int
	Epsilon		  float64
	Sphere        Sphere
}

func main() {
	counterWinsA := 0
	counterWinsB := 0
	n := 100
	for i := 1; i <= n; i++ {
		center := Sphere{}.Center
		center.X = 0
		center.Y = 0
		center.Z = 0
		result := MakeGame(10, 0.1, Sphere{1, center})
		if result == 1 {
			counterWinsA++
		}
		if result == 2 {
			counterWinsB++
		}
	}
	fmt.Println("Количество побед игрока А:", counterWinsA, fmt.Sprintf("(%#.2f%%)", (float64(counterWinsA) / float64(n)) * 100))
	fmt.Println("Количество побед игрока B:", counterWinsB, fmt.Sprintf("(%#.2f%%)", (float64(counterWinsB) / float64(n)) * 100))

}

func MakeGame(pointsNum int, epsilon float64, sphere Sphere) (result int) {
	pointBSpherical := sphere.GeneratePoint()
//	pointB := 1 - rand.Float64() * (sphere.Radius - sphere.Center)
	pointB := GetPointFromSpherical(pointBSpherical.RadialDistance, pointBSpherical.AzimuthAngle, pointBSpherical.ZenithAngle)
	for i := 1; i <= pointsNum; i++ {
		pointASpherical := Point{
			RadialDistance:sphere.GeneratePoint().RadialDistance,
			AzimuthAngle:sphere.GeneratePoint().AzimuthAngle,
			ZenithAngle:sphere.GeneratePoint().ZenithAngle,
		}
		pointA := &Point{}
		pointA = GetPointFromSpherical(pointASpherical.RadialDistance, pointASpherical.AzimuthAngle, pointASpherical.ZenithAngle)
		distanceFromBtoA := pointB.GetDistance(pointA)
		if distanceFromBtoA <= epsilon {
			result = 1
			return result // игрок А выиграл
		}
		if distanceFromBtoA >= epsilon {
			result = 2
			return result // игрок B выиграл
		}
	}

	return result
}

func (p *Point) GetDistance(other *Point) (distance float64) {
	distance = math.Sqrt(math.Pow(p.X - other.X,2) + math.Pow(p.Y - other.Y,2) + math.Pow(p.Z - other.Z,2))

	return distance
}

func (s *Sphere) GeneratePoint() *Point {
	point := Point{}
	rand.Seed(time.Now().UTC().UnixNano())
	point.Z = rand.Float64() * ((s.Center.X - s.Radius) - (s.Center.X + s.Radius))
	point.AzimuthAngle = rand.Float64() * (2 * math.Pi - 0)
	point.RadialDistance = s.Radius
	point.ZenithAngle = math.Acos(point.Z/s.Radius)

	return &point
}

func GetPointFromSpherical(RadialDistance float64, AzimuthAngle float64, ZenithAngle float64) *Point {
	point := Point{}
	point.RadialDistance = RadialDistance
	point.AzimuthAngle = AzimuthAngle
	point.ZenithAngle = ZenithAngle
	point.X = RadialDistance * math.Sin(ZenithAngle) * math.Cos(AzimuthAngle)
	point.Y = RadialDistance * math.Sin(ZenithAngle) * math.Sin(AzimuthAngle)
	point.Z = RadialDistance * math.Cos(ZenithAngle)

	return &point
}