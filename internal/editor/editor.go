package editor

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

const (
	defaultEditor = "vi"
	defaultShell  = "/bin/bash"
	windowsEditor = "notepad"
	windowsShell  = "cmd"
)

// Editor holds the command line args to fire up the editor
type Editor struct {
	Args  []string
	Shell bool
}

func NewDefaultEditor(envs []string) Editor {
	args, shell := defaultEnvEditor(envs)
	return Editor{
		Args:  args,
		Shell: shell,
	}
}
func defaultEnvEditor(envs []string) ([]string, bool) {
	var editor string
	for _, env := range envs {
		if len(env) > 0 {
			editor = os.Getenv(env)
		}
		if len(editor) > 0 {
			break
		}
	}
	if len(editor) == 0 {
		editor = platformize(defaultEditor, windowsEditor)
	}
	if !strings.Contains(editor, "") {
		return []string{editor}, false
	}
	if !strings.ContainsAny(editor, "\"'\\") {
		return strings.Split(editor, " "), false
	}
	shell := defaultEnvShell()
	return append(shell, editor), true
}

func defaultEnvShell() []string {
	shell := os.Getenv("SHELL")
	if len(shell) == 0 {
		shell = platformize(defaultShell, windowsShell)
	}
	flag := "-c"
	if shell == windowsShell {
		flag = "/C"
	}
	return []string{shell, flag}
}

func (e Editor) LaunchTempFile(prefix string) ([]byte, string, error) {

	f, err := os.CreateTemp("", prefix+"*")
	if err != nil {
		return nil, "", err
	}
	defer f.Close()
	path := f.Name()
	// This file descriptor needs to close so the next process (Launch) can claim it.
	f.Close()
	if err = e.Launch(path); err != nil {
		return nil, path, err
	}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, path, errors.Wrap(err, "error in read file")
	}
	err = os.Remove(path)
	return bytes, path, err
}

func (e Editor) Launch(path string) error {
	if len(e.Args) == 0 {
		return fmt.Errorf("no editor is defined, can't open %s", path)
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	args := e.args(abs)
	cmd := exec.Command(args[0], args[1:]...)
	fmt.Println("args: ", args)
	fmt.Fprintf(os.Stdout, "Opening file with editor %v\n", args)
	err = cmd.Run()
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return err
}
func (e Editor) args(path string) []string {
	args := make([]string, len(e.Args))
	copy(args, e.Args)
	if e.Shell {
		last := args[len(args)-1]
		args[len(args)-1] = fmt.Sprintf("%s %q", last, path)
	} else {
		args = append(args, path)
	}
	return args
}

func platformize(linux, windows string) string {
	if runtime.GOOS == "windows" {
		return windows
	}
	return linux
}
