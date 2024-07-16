package metadata

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestReflectivityCalibrationStatus_Age(t *testing.T) {
	type TestCase struct {
		name string
		s    ReflectivityCalibrationStatus
		want time.Duration
	}

	now := time.Now()

	cases := []TestCase{
		{
			"valid with timestamp",
			ReflectivityCalibrationStatus{Valid: true, Timestamp: now.Add(-1 * time.Minute)},
			1 * time.Minute,
		},
		{
			"valid with timestamp in past",
			ReflectivityCalibrationStatus{Valid: true, Timestamp: now.Add(-1 * time.Hour)},
			1 * time.Hour,
		},
		{
			"valid with no timestamp",
			ReflectivityCalibrationStatus{Valid: true, Timestamp: time.Time{}},
			time.Since(time.Time{}),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.s.Age()
			assert.Equal(t, c.want.Round(time.Second), got.Round(time.Second), "it should provide an accurate delta")
		})
	}
}
