package project

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"

	"github.com/mooncake9527/xt/internal/base"
)

// Project is a project template.
type Project struct {
	Name string
	Path string
}

var (
	notReplace = []string{".template", "command/const.go"}
)

// New new a project from remote repo.
func (p *Project) New(ctx context.Context, dir string, layout string, branch string) error {
	to := filepath.Join(dir, p.Name)
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Printf("🚫 %s already exists\n", p.Name)
		prompt := &survey.Confirm{
			Message: "📂 Do you want to override the folder ?",
			Help:    "Delete the existing folder and create the project.",
		}
		var override bool
		e := survey.AskOne(prompt, &override)
		if e != nil {
			return e
		}
		if !override {
			return err
		}
		os.RemoveAll(to)
	}
	fmt.Printf("🚀 Creating service %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)
	repo := base.NewRepo(layout, branch)
	//if err := repo.CopyTo(ctx, to, p.Name, []string{".git", ".github"}); err != nil {
	projectName := title(p.Name)
	if err := repo.CopyToV2(ctx, to, p.Name, []string{".git", ".github"}, []string{
		"xt-layout", p.Name, "Xt-layout", projectName,
	}, notReplace); err != nil {
		return err
	}
	e := os.Rename(
		filepath.Join(to, "cmd", "xt-layout"),
		filepath.Join(to, "cmd", p.Name),
	)
	if e != nil {
		return e
	}
	base.Tree(to, dir)

	fmt.Printf("\n🍺 Project creation succeeded %s\n", color.GreenString(p.Name))
	fmt.Print("💻 Use the following command to start the project 👇:\n\n")

	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Println(color.WhiteString("$ go generate ./..."))
	fmt.Println(color.WhiteString("$ go build -o ./bin/ ./... "))
	fmt.Println(color.WhiteString("$ ./bin/%s -conf ./configs\n", p.Name))
	fmt.Println("			🤝 Thanks for using XT")
	fmt.Println("	📚 Tutorial: https://go-xt.dev/docs/getting-started/start")
	return nil
}

func title(s string) string {
	if s == "" {
		return s
	}
	// 将字符串按空格分割成单词切片
	words := strings.Fields(s)
	for i, word := range words {
		r := []rune(word)
		w := strings.ToUpper(string(r[0])) + string(r[1:])
		words[i] = w
	}
	// 将单词切片重新连接成字符串
	return strings.Join(words, " ")
}
