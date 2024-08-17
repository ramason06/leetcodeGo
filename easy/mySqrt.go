package easy

import "math"

func mySqrt(x int) int {
	return int(math.Sqrt(float64(x))) // 소수점을 버리고 정수로 변환
}
