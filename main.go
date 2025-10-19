package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/StackExchange/wmi"
	"golang.org/x/sys/windows"
)

type Win32_Process struct {
	ProcessId uint32
}

func WaitForAllChildProcesses(parentId int) {
	var processes []Win32_Process
	query := fmt.Sprintf("SELECT ProcessId FROM Win32_Process WHERE ParentProcessId = %d", parentId)
	err := wmi.Query(query, &processes)
	if err != nil {
		fmt.Printf("Error querying processes: %v\n", err)
		return
	}
	for _, proc := range processes {
		childPid := int(proc.ProcessId)
		fmt.Printf("> Waiting for child process with PID: %d\n", childPid)
		handle, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION|windows.SYNCHRONIZE, false, uint32(childPid))
		if err != nil {
			fmt.Printf("Error opening process %d: %v\n", childPid, err)
			continue
		}
		defer windows.CloseHandle(handle)
		_, err = windows.WaitForSingleObject(handle, windows.INFINITE)
		if err != nil {
			fmt.Printf("Error waiting for process %d: %v\n", childPid, err)
		} else {
			fmt.Printf("PID: %d already exited.\n", childPid)
		}
		WaitForAllChildProcesses(childPid)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <executable_path> [args...]")
		os.Exit(1)
	}

	exe := os.Args[1]
	var args []string

	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	cmd := exec.Command(exe, args...)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error starting process: %v\n", err)
		os.Exit(1)
	}

	parentPid := cmd.Process.Pid
	fmt.Printf("Parent process launched with PID: %d\n", parentPid)

	// Wait for the child process to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Process finished with error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Parent process exited with return code: %d\n", cmd.ProcessState.ExitCode())

	WaitForAllChildProcesses(parentPid)

	fmt.Println("All processes exited.")
}
