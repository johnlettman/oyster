package util

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"testing"
)

func gitProjectDir() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	gitOutput, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return NormalizePath(string(gitOutput)), nil
}

func TestProjectRootDir(t *testing.T) {
	want, err := gitProjectDir()
	if err != nil {
		t.Fatalf("failed to determine git project dir: %v", err)
	}

	got := ProjectRootDir()
	assert.Equalf(t, got, want, "it should match the git project root %s", want)
}

func TestProjectDir(t *testing.T) {
	git, err := gitProjectDir()
	if err != nil {
		t.Fatalf("failed to determine git project dir: %v", err)
	}

	sub := "util"
	got := ProjectDir(sub)
	want := filepath.Join(git, sub)
	assert.Equalf(t, got, want, "it should match the git project directory for %s", sub)
}

func TestNormalizePath(t *testing.T) {
	type TestCase struct {
		name string
		path string
		want string
	}

	cases := []TestCase{
		{
			name: "spaces around the path",
			path: "       /etc/os-release     ",
			want: "/etc/os-release",
		},
		{
			name: "new lines",
			path: "/etc/\nlegal",
			want: "/etc/legal",
		},
		{
			name: "carriage returns",
			path: "/etc/\rlegal",
			want: "/etc/legal",
		},
		{
			name: "tilde user home directory shorthand",
			path: "~/Documents",
			want: (func() string {
				home, err := os.UserHomeDir()
				if err != nil {
					t.Fatalf("failed to determine user home: %v", err)
				}

				return filepath.Join(home, "Documents")
			})(),
		},
		{
			name: "dots",
			path: "/root/../etc/../etc/././././os-release",
			want: "/etc/os-release",
		},
		{
			name: "directory separators",
			path: "///////etc/////os-release",
			want: "/etc/os-release",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := NormalizePath(c.path)
			assert.Equal(t, c.want, got, "it should normalize the path")
		})
	}

	t.Run("filepath.Abs error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			patches := gomonkey.ApplyFuncReturn(filepath.Abs, "", syscall.ENOTDIR)
			defer patches.Reset()

			path := "/etc/os-release"
			want := path
			got := NormalizePath(path)
			assert.Equal(t, want, got, "it should silently handle filepath.Abs error")
		}, "it should not panic")
	})
}
