# Auto Scaler

Automatic resource scaling for applications based on metrics and demand.

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
[![CI](https://github.com/Qyroxen/auto-scaler/actions/workflows/ci.yml/badge.svg)](https://github.com/Qyroxen/auto-scaler/actions/workflows/ci.yml)

> Automatic resource scaling for applications based on metrics and demand.

## What is it?

Auto Scaler is a command-line tool built with Go that helps developers automatic resource scaling for applications based on metrics and demand. It's designed to be fast, reliable, and easy to use.

## Why?

Every developer needs auto scaler — but existing tools are either too complex, too slow, or require cloud dependencies. We built Auto Scaler to be:
- **Fast** — Written in Go for maximum performance
- **Offline** — No cloud dependencies, your data stays on your machine
- **Simple** — Clean CLI interface with sensible defaults
- **Extensible** — Easy to customize and integrate into your workflow

## Features

- CPU/Memory-based scaling
- Custom metric scaling
- Predictive scaling
- Cost optimization
- Multi-cloud support
- CLI monitoring

## Quick Start

### Prerequisites

- Go 1.23 or later

### Install

```bash
# Install with go install
go install github.com/Qyroxen/auto-scaler@latest

# Or build from source
git clone https://github.com/Qyroxen/auto-scaler.git
cd auto-scaler
go build -o auto-scaler .
```

### Usage

```bash
# Basic usage
.auto-scaler --help

# Example
./auto-scaler configure --min 2 --max 10 --cpu-threshold 70
```

## Output

```
Auto Scaler v1.0.0

Scanning...

✓ Analysis complete
✓ Results ready

{
  "status": "success",
  "results": [...]
}
```

## Configuration

Create a `.config.yaml` file in your project root:

```yaml
# Configuration options
verbose: true
output: json
timeout: 30s
```

## CLI Flags

```
auto scaler [command]

Flags:
  --path string      Target path (default ".")
  --format string    Output format: json, text (default "text")
  --verbose          Enable verbose output
  --config string    Config file path
  --output string    Output file path
```

## Examples

### Basic Example

```bash
.auto-scaler --path ./src
```

### Advanced Example

```bash
.auto-scaler --path ./src --format json --output report.json --verbose
```

### CI/CD Integration

```yaml
# .github/workflows/ci.yml
- name: Run Auto Scaler
  run: |
    go install github.com/Qyroxen/auto-scaler@latest
    auto-scaler --path . --format json --output report.json
```

## Documentation

- [Getting Started](docs/getting-started.md)
- [Configuration](docs/configuration.md)
- [API Reference](docs/api-reference.md)
- [Examples](examples/)
- [Contributing](CONTRIBUTING.md)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Qyroxen** - [GitHub](https://github.com/Qyroxen)

---

**Found this useful?** Give it a ⭐ on GitHub!
