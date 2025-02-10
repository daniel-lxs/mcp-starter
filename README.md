# MCP Starter

A lightweight Go application that parses JSON configuration files and executes commands with specified environment variables.

## Features

- Parses JSON configuration files
- Executes commands with environment variables
- Supports single server configuration
- Cross-platform (Linux, Windows, macOS)

## Installation

### Pre-built Binaries

Download the latest release for your platform from the [Releases page](https://github.com/daniel-lxs/mcp-starter/releases).

### From Source

1. Ensure you have Go installed (version 1.21 or higher)
2. Clone the repository:
   ```bash
   git clone https://github.com/daniel-lxs/mcp-starter.git
   cd mcp-starter
   ```
3. Build the application:
   ```bash
   go build -o mcp-starter
   ```

## Usage

1. Create a JSON configuration file (e.g., `config.json`):
   ```json
   {
     "mcpServers": {
       "MyServer": {
         "command": "my-command",
         "args": ["arg1", "arg2"],
         "env": {
           "MY_ENV_VAR": "value"
         }
       }
     }
   }
   ```

2. Prepare the command to add it to Cursor:
   ```bash
   path/to/mcp-starter path/to/config.json
   ```

3. Paste the command into the "Add MCP Server dialog:
![image](https://github.com/user-attachments/assets/d6d4344c-c356-41f5-84f5-427c1005a518)

Note: Use the full path to the executable and the config file.

## Configuration

The JSON configuration file must contain exactly one server definition with the following structure:

```json
{
  "mcpServers": {
    "ServerName": {
      "command": "command-to-execute",
      "args": ["argument1", "argument2"],
      "env": {
        "ENV_VAR_NAME": "value"
      }
    }
  }
}
```

- `command`: The command to execute (required)
- `args`: Array of arguments to pass to the command (optional)
- `env`: Key-value pairs of environment variables (optional)

## Development

### Building

To build the application:
```bash
go build -o mcp-starter
```

### Testing

Run the tests:
```bash
go test ./...
```

### Linting

Run the linter:
```bash
golangci-lint run
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Windows Users
The Windows executable is not code-signed. To run it:
1. Right-click the .exe file
2. Select "Properties"
3. Check "Unblock" under the General tab
4. Click "OK"
5. Run the executable
