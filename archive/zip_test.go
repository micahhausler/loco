package archive

import (
	"bytes"
	"testing"
)

func TestCreateConfig(t *testing.T) {
	cases := []struct {
		in   LoginConfig
		want []byte
	}{
		{
			in: LoginConfig{"https://index.docker.io/v1/", "user", "password", ""},
			want: []byte(`{
    "auths": {
        "https://index.docker.io/v1/": {
            "auth": "dXNlcjpwYXNzd29yZA=="
        }
    }
}`),
		},
	}
	for _, c := range cases {
		got := c.in.CreateConfig()
		if !bytes.Equal(got, c.want) {
			t.Errorf("%q.CreateConfig() == %q, want %q", c.in, got, c.want)
		}
	}
}
