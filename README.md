# go-wc

A Go implementation of the Unix `wc` (word count) command, built as part of the [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc) series.

## 🚀 Features

- **Line counting** (`-l`): Count the number of lines in a file or stdin
- **Word counting** (`-w`): Count the number of words in a file or stdin  
- **Byte counting** (`-c`): Count the number of bytes in a file or stdin
- **Character counting** (`-m`): Count the number of characters (UTF-8 aware) in a file or stdin
- **Default mode**: Shows lines, words, and bytes (like the original `wc`)
- **Stdin support**: Read from standard input when no file is specified
- **UTF-8 locale detection**: Automatically handles character counting based on system locale

## 📦 Installation

### Prerequisites

- Go 1.25.0 or later

### Build from source

```bash
# Clone the repository
git clone https://github.com/YPAzevedo/go-wc.git
cd go-wc

# Build the binary
make build

# Or build manually
go build -o bin/go-wc ./cmd/go-wc
```

## 🔧 Usage

### Basic Usage

```bash
# Count lines, words, and bytes (default behavior)
./bin/go-wc filename.txt

# Count from stdin
echo "Hello world" | ./bin/go-wc
```

### Options
```bash
# Count lines only
./bin/go-wc -l filename.txt

# Count words only  
./bin/go-wc -w filename.txt

# Count bytes only
./bin/go-wc -c filename.txt

# Count characters (UTF-8 aware)
./bin/go-wc -m filename.txt

# Combine multiple options
./bin/go-wc -lw filename.txt
```

### Using the Makefile

```bash
# Run with a file
make run INPUT=test.txt

# Run with options
make run OPTIONS=l INPUT=test.txt

# Run with multiple options
make run OPTIONS="l w" INPUT=test.txt

# Run from stdin
echo "Hello world" | make run
```

## 📊 Examples

```bash
# Example with a text file
$ ./bin/go-wc test2.txt
3 5 23 test2.txt

# Count only lines
$ ./bin/go-wc -l test2.txt  
3 test2.txt

# Count only words
$ ./bin/go-wc -w test2.txt
5 test2.txt

# Using with pipes
$ cat test2.txt | ./bin/go-wc
3 5 23

# UTF-8 character counting
$ echo "Hello 世界" | ./bin/go-wc -m
8
```

## 🛠️ Development

### Available Make targets
```bash
make build     # Build the application
make run       # Run the application
make test-run  # Build and run with test.txt
make clean     # Remove build artifacts  
make fmt       # Format the code
make help      # Show available targets
```

### Project Structure
```
go-wc/
├── cmd/go-wc/          # Main application entry point
│   └── main.go
├── bin/                # Build artifacts (created by make build)
├── test.txt           # Test file with sample content
├── test2.txt          # Another test file
├── Makefile           # Build and run automation
├── go.mod             # Go module definition
└── README.md          # This file
```

## 🧪 Testing

```bash
# Test with the included test files
make run INPUT=test.txt
make run INPUT=test2.txt

# Test different options
make run OPTIONS=l INPUT=test2.txt
make run OPTIONS=w INPUT=test2.txt  
make run OPTIONS=c INPUT=test2.txt
make run OPTIONS=m INPUT=test2.txt

# Test with stdin
echo "Sample text for testing" | make run
```

## 🌟 Features vs Original `wc`

| Feature | go-wc | Original wc | Notes |
|---------|--------|-------------|-------|
| Line counting (`-l`) | ✅ | ✅ | |
| Word counting (`-w`) | ✅ | ✅ | |
| Byte counting (`-c`) | ✅ | ✅ | |
| Character counting (`-m`) | ✅ | ✅ | UTF-8 aware with locale detection |
| Default output | ✅ | ✅ | Lines, words, bytes |
| Stdin support | ✅ | ✅ | |
| Multiple files | ❌ | ✅ | Single file or stdin only |

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the terms specified in the [LICENSE](LICENSE) file.

## 🎯 Challenge

This project was built as part of the [Build Your Own wc Tool](https://codingchallenges.fyi/challenges/challenge-wc) coding challenge. The goal was to implement a clone of the Unix `wc` command in Go, demonstrating:

- Command-line argument parsing
- File I/O operations  
- Text processing algorithms
- UTF-8 character handling
- Cross-platform compatibility

## 🔗 Links

- [Original Challenge](https://codingchallenges.fyi/challenges/challenge-wc)
- [Unix wc manual](https://man7.org/linux/man-pages/man1/wc.1.html)