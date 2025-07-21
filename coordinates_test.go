package dutchprovinces

import (
	"testing"
)

func TestLookupLatitudeLongitude(t *testing.T) {
	tests := []struct {
		lat, lon float64
		want     string
	}{
		{
			lat:  51.8395096,
			lon:  5.8499342,
			want: "NL-GE",
		},
		{
			lat:  51.575868,
			lon:  4.2928775,
			want: "NL-NB",
		},
		{
			lat:  52.3695322,
			lon:  5.1678144,
			want: "NL-FL",
		},
	}
	for _, tc := range tests {
		p, _ := LookupLatitudeLongitude(tc.lat, tc.lon)
		if p != tc.want {
			t.Errorf("LookupLatitudeLongitude(%f, %f) returned %q, want %q", tc.lat, tc.lon, p, tc.want)
		}
	}
}

func BenchmarkLookupLatitudeLongitude(b *testing.B) {
	for b.Loop() {
		_, _ = LookupLatitudeLongitude(51.8395096, 5.8499342)
	}
}
