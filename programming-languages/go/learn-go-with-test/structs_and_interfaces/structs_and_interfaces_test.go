package structandinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g but wan %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0
		checkArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})

	t.Run("table test", func(t *testing.T) {
		t.Helper()
		areaTests := []struct {
			shape Shape
			want float64
		}{
			{Rectangle{12, 6}, 72.0},
			{Circle{10}, 314.1592653589793},
			{Triangle{12, 6}, 36.0},
		}

		for _, tt := range areaTests {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %v but want %v", tt.shape, got, tt.want)
			}
		}
	})
}

