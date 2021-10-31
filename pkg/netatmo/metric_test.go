package netatmo

import "testing"

func TestSanitizeName(t *testing.T) {
	type args struct {
		name string
	}

	cases := []struct {
		intention string
		args      args
		want      string
	}{
		{
			"empty",
			args{
				name: "",
			},
			"",
		},
		{
			"simple",
			args{
				name: "maison",
			},
			"maison",
		},
		{
			"detailed",
			args{
				name: "maison (indoor)",
			},
			"maison",
		},
	}

	for _, tc := range cases {
		t.Run(tc.intention, func(t *testing.T) {
			if got := sanitizeName(tc.args.name); got != tc.want {
				t.Errorf("sanitizeSource() = %s, want %s", got, tc.want)
			}
		})
	}
}
