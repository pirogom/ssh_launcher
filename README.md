# sshl

A lightweight TUI-based host address manager and SSH client launcher written in Go.

## Features

- Manage SSH host addresses via a simple JSON configuration file.
- Interactive TUI for selecting hosts.
- Supports custom SSH users per host.
- Cross-platform support (Windows, macOS, Linux).

## Installation

Ensure you have [Go](https://go.dev/dl/) installed on your system.

### Using go install (Recommended)
```bash
go install github.com/pirogom/ssh_launcher@latest
```

After installation, the binary will be installed to your `$GOPATH/bin` directory. The executable name depends on your platform:

- **Windows**: `ssh_launch.exe`
- **macOS / Linux**: `ssh_launch`

### Building from source
```bash
git clone https://github.com/pirogom/ssh_launcher
cd ssh_launcher
go build -o sshl .
```

## Configuration

The application reads its configuration from `~/.ssh_launcher/ssh_launcher.json`.

### Configuration File Format

Create the directory and file at `~/.ssh_launcher/ssh_launcher.json` with the following JSON format:

```json
[
    { "host": "192.168.0.110", "title": "Mac Studio", "User": "pirogom" },
    { "host": "192.168.0.155", "title": "Mac Book Pro", "User": "pirogom" }
]
```

- `host`: The IP address or hostname of the remote server.
- `title`: A friendly name for the host (displayed in the TUI).
- `User`: (Optional) The SSH username. If omitted, the default system user will be used.

## Usage

Run the compiled binary:

```bash
# macOS / Linux
sshl

# Windows
sshl.exe
```

If installed via `go install`, use the platform-specific executable name:

```bash
# macOS / Linux
ssh_launch

# Windows
ssh_launch.exe
```

- Use `↑` / `↓` or `k` / `j` to navigate the host list.
- Press `Enter` to launch the SSH client for the selected host.
- Press `q` or `Ctrl+C` to quit.

## License

This project is licensed under the MIT License.
