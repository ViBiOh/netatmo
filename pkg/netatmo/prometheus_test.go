package netatmo

import "testing"

func TestSanitizeSource(t *testing.T) {
	type args struct {
		source string
	}

	var cases = []struct {
		intention string
		args      args
		want      string
	}{
		{
			"simple",
			args{
				source: "maison",
			},
			"maison",
		},
		{
			"detailed",
			args{
				source: "maison (indoor)",
			},
			"maison",
		},
	}

	for _, tc := range cases {
		t.Run(tc.intention, func(t *testing.T) {
			if got := sanitizeSource(tc.args.source); got != tc.want {
				t.Errorf("sanitizeSource() = %s, want %s", got, tc.want)
			}
		})
	}
}
