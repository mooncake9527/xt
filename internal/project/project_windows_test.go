//go:build windows
// +build windows

package project

import (
	"testing"
)

func Test_processProjectParams(t *testing.T) {
	type args struct {
		projectName      string
		fallbackPlaceDir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"absWindows", args{projectName: "c:\\xt\\awesome\\go\\demo", fallbackPlaceDir: ""}, "c:\\xt\\awesome\\go"},
		//{"relativeWindows", args{projectName: "/home/xt/awesome/go/demo", fallbackPlaceDir: ""}, "/home/xt/awesome/go"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := processProjectParams(tt.args.projectName, tt.args.fallbackPlaceDir); got != tt.want {
				t.Errorf("getProjectPlaceDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
