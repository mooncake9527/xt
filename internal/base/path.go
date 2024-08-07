package base

import (
	"bytes"
	"fmt"
	"github.com/mooncake9527/x/xerrors/xerror"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func xtHome() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	home := filepath.Join(dir, ".xt")
	if _, err := os.Stat(home); os.IsNotExist(err) {
		if err := os.MkdirAll(home, 0o700); err != nil {
			log.Fatal(err)
		}
	}
	return home
}

func xtHomeWithDir(dir string) string {
	home := filepath.Join(xtHome(), dir)
	if _, err := os.Stat(home); os.IsNotExist(err) {
		if err := os.MkdirAll(home, 0o700); err != nil {
			log.Fatal(err)
		}
	}
	return home
}

func copyFile(src, dst string, replaces []string, notReplace []string) error {
	srcinfo, err := os.Stat(src)
	if err != nil {
		return xerror.New(err.Error())
	}
	buf, err := os.ReadFile(src)
	if err != nil {
		return xerror.New(err.Error())
	}

	if !isIn(src, notReplace) {
		var old string
		for i, next := range replaces {
			if i%2 == 0 {
				old = next
				continue
			}
			buf = bytes.ReplaceAll(buf, []byte(old), []byte(next))
		}
	}

	return os.WriteFile(dst, buf, srcinfo.Mode())
}

func isIn(r string, rs []string) bool {
	is := false
	for _, v := range rs {
		if strings.Contains(r, v) {
			is = true
		}
	}
	return is
}

func copyDir(src, dst string, replaces, ignores []string, notReplace []string) error {
	srcinfo, err := os.Stat(src)
	if err != nil {
		return xerror.New(err.Error())
	}

	err = os.MkdirAll(dst, srcinfo.Mode())
	if err != nil {
		return xerror.New(err.Error())
	}

	fds, err := os.ReadDir(src)
	if err != nil {
		return xerror.New(err.Error())
	}
	for _, fd := range fds {
		if hasSets(fd.Name(), ignores) {
			continue
		}
		srcfp := filepath.Join(src, fd.Name())
		dstfp := filepath.Join(dst, fd.Name())
		var e error
		if fd.IsDir() {
			e = copyDir(srcfp, dstfp, replaces, ignores, notReplace)
		} else {
			e = copyFile(srcfp, dstfp, replaces, notReplace)
		}
		if e != nil {
			return xerror.New(e.Error())
		}
	}
	return nil
}

func hasSets(name string, sets []string) bool {
	for _, ig := range sets {
		if ig == name {
			return true
		}
	}
	return false
}

func Tree(path string, dir string) {
	_ = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && info != nil && !info.IsDir() {
			fmt.Printf("%s %s (%v bytes)\n", color.GreenString("CREATED"), strings.Replace(path, dir+"/", "", -1), info.Size())
		}
		return nil
	})
}
