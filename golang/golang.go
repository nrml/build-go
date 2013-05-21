package golang

import (
	"github.com/nrml/build-go"
	"os/exec"
)

func Build(prg build.Program) (string, error) {
	cmd := exec.Command("go", "build", prg.Src)

	out, err := cmd.Output()

	return string(out), err
}
func Install(prg build.Program) (string, error) {
	out, err := Build(prg)

	cmd := exec.Command("go", "install", prg.Src)

	res, err := cmd.Output()

	out = string(res)

	return out, err
}
