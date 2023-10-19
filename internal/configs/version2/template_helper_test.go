package version2

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/google/go-cmp/cmp"
)

func testTmpl(t *testing.T) *template.Template {
	t.Helper()

	funcs := template.FuncMap{
		"generateUpstreamBackup": generateUpstreamBackup,
	}
	const testTmpl = `{{- generateUpstreamBackup .Backup .BackupPort -}}`

	tmpl, err := template.New("upstreamTest").Funcs(funcs).Parse(testTmpl)
	if err != nil {
		t.Fatal(err)
	}
	return tmpl
}

func TestGenerateUpstreamBackupForValidInput(t *testing.T) {
	t.Parallel()

	tmpl := testTmpl(t)

	tt := []struct {
		name  string
		input Upstream
		want  string
	}{
		{
			name: "valid values for backup name and backup port",
			input: Upstream{
				Backup:     "backup1.example.com",
				BackupPort: 9999,
			},
			want: "server backup1.example.com:9999 backup;",
		},
		{
			name:  "no backup name nor backup port provided",
			input: Upstream{},
			want:  "",
		},
	}

	for _, tc := range tt {
		buf := bytes.Buffer{}
		err := tmpl.Execute(&buf, tc.input)
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(tc.want, buf.String()) {
			t.Error(cmp.Diff(tc.want, buf.String()))
		}
	}
}
