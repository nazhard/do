package do

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const VERSION = "v0.0.1"

func RunScript() (*exec.Cmd, error) {
	args := os.Args

	switch len(args) {
	case 1:
		fmt.Println("do " + VERSION)
		fmt.Println("Automatically run your .sh file in the ./scripts/ directory.")
		return nil, nil
	case 2:
		return processScript("scripts/" + args[1])
	case 3:
		if args[1] == "." {
			return processScript("scripts/" + args[2])
		}
	}

	return nil, fmt.Errorf("invalid argument")
}

func processScript(scriptPath string) (*exec.Cmd, error) {
	scriptPath = resolveScriptPath(scriptPath)

	if err := validateFile(scriptPath); err != nil {
		return nil, err
	}

	interpreter, err := detectInterpreter(scriptPath)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(interpreter, scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd, nil
}

func resolveScriptPath(script string) string {
	exts := []string{"", ".mjs", ".js", ".py", ".sh"}

	for _, ext := range exts {
		path := "./" + script + ext

		if fileExists(path) {
			return path
		}
	}

	return script
}

func detectInterpreter(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("Failed to open %s: %v", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		firstLine := scanner.Text()
		return getInterpreterFromLine(firstLine), nil
	}

	return "", errors.New("unable to read script's first line")
}

func getInterpreterFromLine(line string) string {
	switch {
	case strings.Contains(line, "bash"):
		return "bash"
	case strings.Contains(line, "fish"):
		return "fish"
	case strings.Contains(line, "python"):
		return "python"
	case strings.Contains(line, "node"):
		return "node"
	default:
		return ""
	}
}

func fileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func validateFile(filePath string) error {
	if !fileExists(filePath) {
		return fmt.Errorf("file %s doesn't exists", filePath)
	}
	return nil
}
