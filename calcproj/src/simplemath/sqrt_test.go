package simplemath
import "testing"

func TestSqrt(t *testing.T) {
	v := int(Sqrt(4))
	if v != 2 {
		t.Errorf("Sqrt(4) failed, Got %d, expected 2", v)
	}
}
