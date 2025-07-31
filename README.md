# Vessl - Docker Container Management CLI

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/saswatsam786/vessl)](https://goreportcard.com/report/github.com/saswatsam786/vessl)
[![Go Reference](https://pkg.go.dev/badge/github.com/saswatsam786/vessl.svg)](https://pkg.go.dev/github.com/saswatsam786/vessl)

A powerful command-line interface for managing Docker containers, built with Go and Cobra.

## 🚀 Features

- **Container Management**: Create, start, stop, and remove containers
- **Container Monitoring**: Real-time stats and resource usage
- **Image Management**: List, pull, and manage Docker images
- **Network Information**: View port mappings and network details
- **Logs & Inspection**: View logs and inspect container details
- **Interactive Commands**: Execute commands inside containers

## 📦 Installation

### Option 1: Homebrew (Recommended)

```bash
brew install saswatsam786/vessl/vessl
```

### Option 2: Go Install

```bash
go install github.com/saswatsam786/vessl@latest
```

### Option 3: Build from Source

```bash
git clone https://github.com/saswatsam786/vessl.git
cd vessl
go build -o vessl
sudo mv vessl /usr/local/bin/
```

### Option 4: Download Binary

Download the latest release from [GitHub Releases](https://github.com/saswatsam786/vessl/releases)

## 🎯 Usage

### Basic Commands

```bash
# List all containers
vessl list

# List all images
vessl images

# Create a new container
vessl create

# Start a container
vessl start <container-id>

# Stop a container
vessl stop <container-id>

# Remove a container
vessl remove <container-id>
```

### Advanced Commands

```bash
# View real-time container stats
vessl stats

# View container logs
vessl logs <container-id> [--follow] [--tail N]

# Inspect container details
vessl inspect <container-id> [--json]

# View port mappings
vessl ports <container-id>

# Pull an image
vessl pull <image-name>

# Execute command in container
vessl exec <container-id> <command>
```

### Examples

```bash
# Start a Redis container
vessl create
# Enter: redis-test
# Enter: redis:alpine

# View running containers
vessl list

# Monitor resource usage
vessl stats

# View container logs
vessl logs redis-test --tail 10
```

## 📋 Available Commands

| Command   | Description                  |
| --------- | ---------------------------- |
| `list`    | List all Docker containers   |
| `images`  | List all Docker images       |
| `create`  | Create a new container       |
| `start`   | Start a container            |
| `stop`    | Stop a container             |
| `remove`  | Remove a container           |
| `logs`    | View container logs          |
| `inspect` | Inspect container details    |
| `ports`   | View port mappings           |
| `stats`   | Real-time container stats    |
| `pull`    | Pull an image                |
| `exec`    | Execute command in container |

## 🔧 Requirements

- Go 1.19+
- Docker Engine
- Docker API access

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra)
- Uses [Docker Go SDK](https://github.com/docker/docker)

## 📞 Support

- Create an issue on GitHub
- Check the documentation
- Review existing issues

---

**Made with ❤️ using Go and Docker**
