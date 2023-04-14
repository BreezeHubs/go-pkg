package cmdpkg

import (
	"fmt"
	"github.com/BreezeHubs/go-pkg/filepkg"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func GetPidOf(pidName string) []int {
	var ret []int
	var dirs []string
	dirs, err := filepkg.DirList("/proc")
	if err != nil {
		return ret
	}

	count := len(dirs)
	for i := 0; i < count; i++ {
		pid, err := strconv.Atoi(dirs[i])
		if err != nil {
			continue
		}

		cmdlineFile := fmt.Sprintf("/proc/%d/cmdline", pid)
		if !filepkg.IsExist(cmdlineFile) {
			continue
		}

		cmdlineBytes, err := filepkg.ReadBytes(cmdlineFile)
		if err != nil {
			continue
		}

		cmdlineBytesLen := len(cmdlineBytes)
		if cmdlineBytesLen == 0 {
			continue
		}

		noNut := make([]byte, 0, cmdlineBytesLen)
		for j := 0; j < cmdlineBytesLen; j++ {
			if cmdlineBytes[j] != 0 {
				noNut = append(noNut, cmdlineBytes[j])
			}
		}

		if strings.Contains(string(noNut), pidName) {
			ret = append(ret, pid)
		}
	}

	return ret
}

func KillPidOf(pidName string) error {
	pidName = strings.TrimSpace(pidName)
	if pidName == "" {
		return errors.New("pidName is empty")
	}

	pids := GetPidOf(pidName)
	for _, pid := range pids {
		if out, err := KillPid(pid); err != nil {
			return errors.Wrapf(err, "kill -9 %d fail, output: %s", pid, out)
		}
	}
	return nil
}

func KillPid(pid int) (string, error) {
	out, err := CmdOutput("kill", "-9", strconv.Itoa(pid))
	if err != nil {
		return out, errors.Wrapf(err, "kill -9 %d fail, output: %s", pid, out)
	}
	return out, nil
}
