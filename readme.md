# Launcher Handler

A simple Go program designed for Windows that launches an executable and waits for both the parent process and all its child processes to finish before exiting. this is golang enhancement port version of [WaitForChildrenProcess](https://github.com/KidiXDev/WaitForChildrenProcess)

## Features

- Launches a specified executable with optional arguments
- Waits for the main process to complete
- Recursively waits for all child processes to exit
- Provides feedback on process IDs and exit status

## Usage

```bash
go run main.go <executable_path> [args...]
```

### Example

```bash
go run main.go notepad.exe myfile.txt
```

## Requirements

- Go 1.16 or later
- Windows

## Dependencies

- `github.com/StackExchange/wmi`
- `golang.org/x/sys/windows`

## License

MIT License - see LICENSE file for details
