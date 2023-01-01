package challenges

import "testing"

func TestMinMaxSum(t *testing.T) {
	assertSolution := func(t testing.TB, gotMin, gotMax, expectedMin, expectedMax int64) {
		t.Helper()
		if gotMin != expectedMin || gotMax != expectedMax {
			t.Errorf("min max: expected [%d, %d] but got [%d, %d]", expectedMin, expectedMax, gotMin, gotMax)
		}
	}

	t.Run("test case 1", func (t *testing.T) {
		arr := []int32{1, 3, 5, 7, 9}	
		gotMin, gotMax := minMaxSum(arr)
		expectedMin, expectedMax := int64(16), int64(24)
	
		assertSolution(t, gotMin, gotMax, expectedMin, expectedMax)
	})

	t.Run("test case 2", func (t *testing.T) {
		arr := []int32{7, 69, 2, 221, 8974}
		gotMin, gotMax := minMaxSum(arr)
		expectedMin, expectedMax := int64(299), int64(9271)

		assertSolution(t, gotMin, gotMax, expectedMin, expectedMax)
	})

	t.Run("test case 3", func (t *testing.T) {
		arr := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
		gotMin, gotMax := minMaxSum(arr)
		expectedMin, expectedMax := int64(2063136757), int64(2744467344)

		assertSolution(t, gotMin, gotMax, expectedMin, expectedMax)
	})
}

