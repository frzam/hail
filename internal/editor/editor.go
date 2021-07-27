package editor

import (
	"bufio"
	"fmt"
	"io"
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

func (e Editor) LaunchTempFile(prefix string, edit bool, r io.Reader) ([]byte, string, error) {
	path, err := CreateTempFile(prefix, edit, r)
	if err != nil {
		return nil, path, err
	}

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

func CreateTempFile(prefix string, edit bool, r io.Reader) (string, error) {

	f, err := os.CreateTemp("", prefix+"*")
	if err != nil {
		return "", err
	}
	defer f.Close()
	path := f.Name()
	if edit {
		if _, err := io.Copy(f, r); err != nil {
			os.Remove(path)
			return "", err
		}
	}
	// This file descriptor needs to close so the next process (Launch) can claim it.
	f.Close()
	return path, err
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
	fmt.Fprintf(os.Stdout, "Opening file with editor %v\n", args)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Run()
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

func (e Editor) RunScript(filename string, command string) ([]byte, error) {

	f, _ := os.Open(filename)
	defer f.Close()

	_, err := io.Copy(f, strings.NewReader(command))
	if err != nil {
		fmt.Println("error in copy", err)

	}
	scanner := bufio.NewScanner(f)
	var shebangList []string
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if strings.HasPrefix(scanner.Text(), "#!") {
			shebangList = strings.Split(scanner.Text(), " ")
			break
		}
	}
	var interpreter string
	if len(shebangList) > 1 {
		interpreter = shebangList[1]
	} else if len(shebangList) == 1 {
		s := strings.Split(shebangList[0], "/")
		interpreter = s[len(s)-1]
	}
	if interpreter == "" {
		interpreter = "bash"
	}
	fmt.Println("interpreter: ", interpreter)

	err = os.Chmod(filename, 0500)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(2)
	}
	path, err := exec.LookPath(interpreter)
	if err != nil {

	}
	fmt.Println("path: ", path)
	fmt.Println("command: ", command)
	cmd, err := exec.Command(path, "echo ").Output()
	return cmd, err
}
