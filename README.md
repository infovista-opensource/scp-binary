# SCP File Transfer written in golang

A single-binary SCP client written in Golang.

Original code: https://github.com/nicklasfrahm/scp-action

Created a single-binary from the upstream repository in order to avoid using a container for scp-action.

## Usage

```
Usage:
  scp-binary [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  download    downloads a file from host:SRC to DEST. SRC is comma separated
  help        Help about any command
  upload      uploads a file from SRC to host:DEST. SRC is comma separated

Flags:
      --fingerprint string             expected host key fingerprint ($SCP_KEY_FINGERPRINT)
  -h, --help                           help for scp-binary
  -H, --host string                    host to connect to ($SCP_HOST)
      --key string                     private key file ($SCP_KEY)
      --passenv string                 envvar name containing password (default "INSECURE_PASSWORD")
      --passphrase string              private key passphrase ($SCP_KEY_PASSPHRASE)
  -p, --password string                password ($INSECURE_PASSWORD)
      --port int                       port to connect to ($SCP_PORT/22) (default 22)
      --proxy string                   proxy ($SCP_PROXY)
      --proxy-key string               proxy key ($SCP_PROXY_KEY)
      --proxy-key-fingerprint string   expected proxy host key fingerprint ($SCP_PROXY_KEY_FINGERPRINT)
      --proxy-key-passphrase string    private key passphrase ($SCP_PROXY_KEY_PASSPHRASE)
      --proxy-password string          proxy password ($SCP_PROXY_PASSWORD)
      --proxy-port int                 port to connect to ($SCP_PROXY_PORT/22) (default 22)
      --proxy-skip-fingerprint         skip proxy host key verification ($SCP_IGNORE_PROXY_FINGERPRINT)
      --proxy-username string          proxy username ($SCP_PROXY_USERNAME)
      --skip-fingerprint               skip host key verification ($SCP_IGNORE_FINGERPRINT)
      --timeout string                 timeout in seconds ($SCP_TIMEOUT/30s) (default "30s")
  -u, --username string                username ($SCP_USERNAME)
```

## License

This project is licensed under the [MIT license](./LICENSE.md).
