package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartTime(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "2022-01-01 00:19:00",
			args: "2022-01-01 00:19:00",
			want: "2022-01-01 00:00:00",
		},
		{
			name: "2022-01-01 23:59:59",
			args: "2022-01-01 23:59:59",
			want: "2022-01-01 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm, _ := time.Parse(DateTimeFmt, tt.args)
			assert.Equalf(t, tt.want, StartTime(tm).Format(DateTimeFmt), "StartTime(%v)", tt.args)
		})
	}
}

func TestEndTime(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "2022-01-01 00:19:00",
			args: "2022-01-01 00:19:00",
			want: "2022-01-01 23:59:59",
		},
		{
			name: "2022-01-01 00:00:00",
			args: "2022-01-01 00:00:00",
			want: "2022-01-01 23:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm, _ := time.Parse(DateTimeFmt, tt.args)
			assert.Equalf(t, tt.want, EndTime(tm).Format(DateTimeFmt), "EndTime(%v)", tt.args)
		})
	}
}

func TestDiffDays(t *testing.T) {
	// today := time.Now()
	tests := []struct {
		name  string
		start string
		end   string
		want  int
	}{
		{
			name:  "1",
			start: "2022-01-01 23:00:00",
			end:   "2022-01-02 12:00:00",
			want:  1,
		},
		{
			name:  "2",
			start: "2023-01-14 08:00:00",
			end:   "2024-01-10 07:59:59",
			want:  361,
		},
		{
			name:  "2023-01-14 08:00:00",
			start: "2023-01-14 08:00:00",
			end:   "2023-01-30 07:59:59",
			want:  16,
		},
		{
			name:  "2023-01-14 08:00:00",
			start: "2023-01-15 08:00:00",
			end:   "2023-01-14 08:00:00",
			want:  -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts, _ := time.Parse(DateTimeFmt, tt.start)
			te, _ := time.Parse(DateTimeFmt, tt.end)
			assert.Equalf(t, tt.want, DiffDays(ts, te), "DiffDays(%v, %v)", tt.start, tt.end)
		})
	}
}
