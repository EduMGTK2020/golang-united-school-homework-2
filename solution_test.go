package square

import "testing"

func TestCalcSquare(t *testing.T) {
	expCircle := 314.1592653589793
	expTriangle := 43.30127018922193
	expSquare := 100.0

	if expCircle != CalcSquare(10.0, SidesCircle) {
		t.Errorf("Unexpected result:\n\tExpected: %f\n\tGot: %f", expCircle, CalcSquare(10.0, SidesCircle))
	}
	if expTriangle != CalcSquare(10.0, SidesTriangle) {
		t.Errorf("Unexpected result:\n\tExpected: %f\n\tGot: %f", expCircle, CalcSquare(10.0, SidesTriangle))
	}
	if expSquare != CalcSquare(10.0, SidesSquare) {
		t.Errorf("Unexpected result:\n\tExpected: %f\n\tGot: %f", expCircle, CalcSquare(10.0, SidesSquare))
	}
}
