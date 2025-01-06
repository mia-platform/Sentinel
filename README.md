# Sentinel Agent
<div align="center">
  <img src="sentinel.png" alt="Logo" style="width: 200px; height: auto;">
</div>

## Introduction
### What is Sentinel?

Sentinel is a project that aims to provide a simple and easy-to-use tool for monitoring the status of your servers and processes. It is designed to be used with the [integration-connector-agent](https://github.com/mia-platform/integration-connector-agent).

### How does it work?


## Configuration

The configuration is done through environment variables and a configuration file. The configuration file is a JSON file that contains the configuration for the services that you want to monitor and some other settings.

### Configuration file

Here is an example of a configuration file:

```json
{
  "output": [
    {
      "type": "stdout",         // 3 type of output: stdout, file, webhook
      "file": {                 // ONLY if type is file
        "path": "./eventsFile/file.log"
      },
      {                         // ONLY if type is webhook - NOT IMPLEMENTED
        "url": "http://localhost:8080/webhook",
        "authentication": {
          // ...
        } 
      }
    },
  ],
  "monitor": {
    "interval": 30,
    "filters": {                // Optional 
      "whitelist": [
        "go"                    // filter only this processes 
      ],
      "users": [
        "<process_username>"    // filter only processes with this username
      ],
      "blacklist": [            // remove this processes from the list
        "system"
      ]
    }
  },
  "server": {
    "port": 8080                // Port for the server - NOT IMPLEMENTED
  },
  "advanced": {
    "debug": true,              // Enable debug mode - NOT IMPLEMENTED
    "omitScanning": [           // Omit output fields - NOT IMPLEMENTED
      "process.command",
      "process.user"
    ]
  }
}
```

#### Filters
Actually only this pairs of filters are supported:
- NONE
- whitelist
- blacklist
- users
- whitelist and users
- blacklist and users

NB: whitelist and blacklist are mutually exclusive

### Environment variables

```env
CONFIGURATION_PATH=<path_to_configuration_file>
LOG_LEVEL=<log_level> # Default: info
HTTP_PORT=<port> # Default: 8080
HTTP_ADDRESS=<address> # Default: 0.0.0.0
```

## Local Development

To develop the service locally you need:

Go 1.23+
To start the application locally

```shell
go run .
```

By default the service will run on port 8080
