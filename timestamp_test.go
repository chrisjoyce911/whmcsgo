package whmcsgo

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimestampUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected time.Time
		wantErr  bool
	}{
		{
			name:     "RFC3339",
			input:    "2023-07-10T12:34:56Z",
			expected: time.Date(2023, 7, 10, 12, 34, 56, 0, time.UTC),
			wantErr:  false,
		},
		{
			name:     "Unix timestamp",
			input:    int64(1688992496), // 2023-07-10T12:34:56Z
			expected: time.Unix(1688992496, 0),
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := json.Marshal(tt.input)
			var ts Timestamp
			err := json.Unmarshal(b, &ts)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.True(t, ts.Time.Equal(tt.expected), "expected %v, got %v", tt.expected, ts.Time)
			}
		})
	}
}

func TestWHCMSdateUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
		wantZero bool
		wantErr  bool
	}{
		{
			name:     "Standard date",
			input:    "2023-07-10",
			expected: time.Date(2023, 7, 10, 0, 0, 0, 0, time.UTC),
			wantZero: false,
			wantErr:  false,
		},
		{
			name:     "DateTime",
			input:    "2023-07-10 15:04:05",
			expected: time.Date(2023, 7, 10, 15, 4, 5, 0, time.UTC),
			wantZero: false,
			wantErr:  false,
		},
		{
			name:     "Zero date",
			input:    "0000-00-00 00:00:00",
			expected: time.Time{},
			wantZero: true,
			wantErr:  false,
		},
		{
			name:     "Zero date",
			input:    "0000-00-00",
			expected: time.Time{},
			wantZero: true,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := json.Marshal(tt.input)
			var wd WHCMSdate
			err := json.Unmarshal(b, &wd)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.wantZero {
					assert.True(t, wd.Time.IsZero(), "expected zero time, got %v", wd.Time)
				} else {
					assert.Equal(t, tt.expected.Year(), wd.Time.Year())
					assert.Equal(t, tt.expected.Month(), wd.Time.Month())
					assert.Equal(t, tt.expected.Day(), wd.Time.Day())
					assert.Equal(t, tt.expected.Hour(), wd.Time.Hour())
					assert.Equal(t, tt.expected.Minute(), wd.Time.Minute())
					assert.Equal(t, tt.expected.Second(), wd.Time.Second())
					assert.Equal(t, tt.expected.Location(), wd.Time.Location())
				}
			}
		})
	}
}
