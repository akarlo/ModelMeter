# ModelMeter

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**ModelMeter** is a lightweight, cross-platform utility that calculates the total storage footprint of locally installed large language models managed through Ollama.

## 📋 Overview

Managing multiple large language models locally can quickly consume significant disk space. ModelMeter provides a simple command-line tool to instantly determine exactly how much storage your LLM collection is using, helping you manage system resources effectively.

## 🚀 Quick Start

```bash
# Python
./py/ollama_total.py

# Deno (JavaScript)
deno run --allow-run ./js/ollama_total.js

# Go
./go/ollama_total

# Example output
26.8 GB
```

## ✨ Features

- **Accurate Size Calculation:** Precisely measures total storage used by all installed Ollama models
- **Smart Unit Formatting:** Automatically selects the most appropriate unit (MB, GB, TB)
- **Multiple Language Implementations:** Choose from Python, Deno (JavaScript), or Go
- **Simple Integration:** Use as a standalone command or integrate into larger workflows
- **Minimal Dependencies:** Relies only on standard libraries in each language

## ❓ Why ModelMeter?

While you could manually add up model sizes from `ollama list` output, ModelMeter:
- Saves time with instant calculations
- Eliminates human error in size conversions and addition
- Provides consistent formatting across different model sizes
- Enables easy integration into monitoring scripts and dashboards
- Serves as a building block for more complex LLM resource management tools

## 💻 Installation

### System Requirements
- **Python:** Python 3.10+
- **JavaScript:** Deno 2.0+
- **Go:** Go 1.16+
- **All implementations:** Ollama must be installed and in your PATH

### Python Implementation

```bash
# Clone the repository
git clone https://github.com/ai-mindset/ModelMeter.git
cd ModelMeter/py

# Make the Python script executable
chmod +x ollama_total.py

# Optional: Add to your PATH
ln -s $(pwd)/ollama_total.py /usr/local/bin/ollama_total
```

### Deno Implementation

```bash
# Clone the repository
git clone https://github.com/ai-mindset/ModelMeter.git
cd ModelMeter/js

# Make the JavaScript file executable
chmod +x ollama_total.js

# Run with Deno
deno run --allow-run ollama_total.js

# Optional: Install as a command
deno install --allow-run --name ollama_total ollama_total.js
```

### Go Implementation

```bash
# Clone the repository
git clone https://github.com/ai-mindset/ModelMeter.git
cd ModelMeter/go

# Build the Go binary
go build -o ollama_total ollama_total.go

# Move to a directory in your PATH
sudo mv ollama_total /usr/local/bin/
```

## 💻 Usage

### Basic Command

```bash
ollama_total
```

### Integrating Into Scripts

#### Python

```python
from ollama_total import get_ollama_total

# Get the total storage used by Ollama models
total = get_ollama_total()
print(f"Your Ollama models are using {total}")
```

#### Deno (JavaScript)

```javascript
import { getOllamaTotal } from "./ollama_total.js";

// Get the total storage used by Ollama models
const total = await getOllamaTotal();
console.log(`Your Ollama models are using ${total}`);
```

#### Go

```go
package main

import (
    "fmt"
    "path/to/modelMeter"
)

func main() {
    total, err := getOllamaTotal()
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }
    fmt.Printf("Your Ollama models are using %s\n", total)
}
```

## 🔍 Technical Details

### How It Works

ModelMeter executes the `ollama list` command, parses the output to extract size information, converts all sizes to bytes for accurate summation, and then converts the total back to the most appropriate unit.

### Size Conversion and Parsing

- Uses SI units: 1 MB = 10^6 bytes, 1 GB = 10^9 bytes, 1 TB = 10^12 bytes (matching Ollama's output)
- Regex pattern `(\d+(?:\.\d+)?) (MB|GB|TB)` extracts size values and units
- Results are formatted with one decimal place precision for readability

## ⚠️ Troubleshooting

| Issue | Solution |
|-------|----------|
| `Command not found: ollama` | Ensure Ollama is installed and in your PATH |
| No output displayed | Check if you have any models installed with `ollama list` |
| Permission denied | Run `chmod +x` on the script file |
| Incorrect total | Ensure no other process is installing/removing models during calculation |
| Deno security errors | Ensure you've provided the `--allow-run` flag when using Deno |

## 🔄 Alternatives

| Approach | Pros | Cons |
|----------|------|------|
| ModelMeter | Automatic calculation, multiple languages, formatting | Additional tool to install |
| Manual addition from `ollama list` | No installation needed | Time-consuming, error-prone |
| Shell one-liner[^1] | Quick for one-time use | Complex syntax, limited formatting |
| Disk usage tools (du) | System-wide analysis | Includes non-model files, less precise |


## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- [Ollama](https://github.com/ollama/ollama) - For making local LLM deployment accessible

---
[^1]: Example shell one-liner: `ollama list | awk 'NR>1 {size=$3; if($4=="GB") size*=1000; if($4=="TB") size*=1000000; sum+=size} END {printf "%.1f %s\n", sum>=1000000 ? sum/1000000 : sum>=1000 ? sum/1000 : sum, sum>=1000000 ? "TB" : sum>=1000 ? "GB" : "MB"}'`
