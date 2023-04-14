package cmdpkg

import (
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

func CmdOutput(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	b, err := cmd.CombinedOutput()
	return *(*string)(unsafe.Pointer(&b)), err
}

func CmdOutputWithCmd(execCmd *exec.Cmd, name string, arg ...string) *exec.Cmd {
	cmd := exec.Command(name, arg...)
	renderCmd(cmd, execCmd)

	return cmd
}

func renderCmd(dst, src *exec.Cmd) {
	if src.Path != "" {
		dst.Path = src.Path
	}

	if len(src.Args) > 0 {
		dst.Args = append(src.Args, dst.Args...)
	}

	if len(src.Env) > 0 {
		dst.Env = append(os.Environ(), dst.Env...)
	}

	if src.Dir != "" {
		dst.Dir = src.Dir
	}

	if src.Stdin != nil {
		dst.Stdin = src.Stdin
	}

	if src.Stdout != nil {
		dst.Stdout = src.Stdout
	}

	if src.Stderr != nil {
		dst.Stderr = src.Stderr
	}

	if len(src.ExtraFiles) > 0 {
		dst.ExtraFiles = append(src.ExtraFiles, dst.ExtraFiles...)
	}

	if src.SysProcAttr != nil {
		dst.SysProcAttr = &syscall.SysProcAttr{}
		*dst.SysProcAttr = *src.SysProcAttr
	}

	if src.Process != nil {
		dst.Process = &os.Process{}
		*dst.Process = *src.Process
	}

	if src.ProcessState != nil {
		dst.ProcessState = &os.ProcessState{}
		*dst.ProcessState = *src.ProcessState
	}

	if src.Err != nil {
		dst.Err = src.Err
	}

	if src.Cancel != nil {
		dst.Cancel = src.Cancel
	}

	dst.WaitDelay = src.WaitDelay
}
