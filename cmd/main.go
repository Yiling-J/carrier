package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
	"golang.org/x/tools/go/packages"
)

var schemaCode = `
package schema
import (
	"github.com/Yiling-J/carrier"
)
var (
	Schemas = []carrier.Schema{}
)
`

var generateCode = `
package main

import "os"
import "strings"
import "github.com/Yiling-J/carrier"
import "{{.pkg}}"


func main() {
    // replace schema package
    pkg := strings.TrimSuffix("{{.pkg}}", "schema") + "factory"
    err := carrier.SchemaToMetaFactory(pkg, "{{.path}}", schema.Schemas)
    if err != nil {
        os.Exit(1)
    }
}

`

func run(target string) (string, error) {
	cmd := exec.Command("go", "run", target)
	stderr := bytes.NewBuffer(nil)
	stdout := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	if err := cmd.Run(); err != nil {
		fmt.Println(stdout.String())
		fmt.Println(stderr.String())
		return "", fmt.Errorf("generate error: %s", stderr)
	}
	fmt.Println(stdout.String())
	return stdout.String(), nil
}

func initCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "init carrier package",
		Run: func(cmd *cobra.Command, args []string) {
			target := "carrier/schema/schema.go"
			if err := os.MkdirAll("carrier/schema", os.ModePerm); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			tmpl, err := template.New("init").Parse(schemaCode)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			b := &bytes.Buffer{}
			err = tmpl.Execute(b, nil)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			var buf []byte
			if buf, err = format.Source(b.Bytes()); err != nil {
				fmt.Println("formatting output:", err)
				os.Exit(1)
			}
			// nolint: gosec
			if err := ioutil.WriteFile(target, buf, 0644); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	return cmd
}

func generateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "generate carrier.go",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := &packages.Config{Mode: packages.NeedName}
			path := "./carrier/schema"
			if len(args) > 0 {
				path = args[0]
			}
			abs, err := filepath.Abs(path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			dir := filepath.Dir(abs)
			if err := os.MkdirAll(dir+"/.gen", os.ModePerm); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer os.RemoveAll(dir + "/.gen")
			target := dir + "/.gen/main.go"
			pkgs, err := packages.Load(cfg, path)
			if err != nil {
				fmt.Println("Can't load package: ", err)
				os.Exit(1)
			}
			pkg := pkgs[0].PkgPath
			tmpl, err := template.New("generare").Parse(generateCode)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			b := &bytes.Buffer{}
			err = tmpl.Execute(b, map[string]string{"pkg": pkg, "path": path})

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			var buf []byte
			if buf, err = format.Source(b.Bytes()); err != nil {
				fmt.Println("formatting output:", err)
				os.Exit(1)
			}
			// nolint: gosec
			if err := ioutil.WriteFile(target, buf, 0644); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			_, err = run(target)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		},
	}
	return cmd
}

func main() {
	cmd := &cobra.Command{Use: "carrier"}
	cmd.AddCommand(
		initCmd(),
		generateCmd(),
	)
	_ = cmd.Execute()
}
