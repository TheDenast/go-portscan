# go-portscan

TCP port scanner. Scans a host concurrently and prints open ports sorted.

## Install

```bash
go install github.com/TheDenast/go-portscan@latest
```

## Usage

```
go-portscan <host> <port-range>
```

## Arguments

| Argument     | Description                                 |
| ------------ | ------------------------------------------- |
| `host`       | IPv4 or IPv6 address                        |
| `port-range` | Range in `X-X` format. Valid ports: 1–65535 |

## Example

```
$ go-portscan 192.168.10.1 1-1024
Scanning 192.168.10.1 (ports 1-1024)...
[open] 22/tcp
[open] 80/tcp
[open] 443/tcp
```

## Notes

- Timeout per port: 1s
- Single port: use `N-N`

## AI Disclosure

This project is AI-**advised**:

- I've consulted AI while writing it, but didn't let it write any code directly.
- Every line of code is manually written by me
- README is AI-generated

