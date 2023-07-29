# Compozify

Compozify is a simple (yet complicated) tool to generate a `docker-compose.yml` file from a `docker run` command.

# Usage

## Screenshot

![image](https://github.com/profclems/compozify/assets/41906128/bcd27512-8692-44f3-9113-63bfb112e38e)


## Installation
Download a binary suitable for your OS at the [releases page](https://github.com/profclems/compozify/releases/latest).

### Quick install
#### Linux and macOS
```sh
curl -sfL https://raw.githubusercontent.com/profclems/compozify/main/install.sh | sh
```

#### Windows (PowerShell)
Open a PowerShell terminal (version 5.1 or later) and run:

```powershell
Set-ExecutionPolicy RemoteSigned -Scope CurrentUser # Optional: Needed to run a remote script the first time

irm https://raw.githubusercontent.com/profclems/compozify/main/install.ps1 | iex
```

### From source

#### Prerequisites for building from source
- `make`
- Go 1.18+

1. Verify that you have Go 1.18+ installed

   ```sh
   go version
   ```

   If `go` is not installed, follow instructions on [the Go website](https://golang.org/doc/install).

2. Clone this repository

   ```sh
   git clone https://github.com/profclems/compozify.git
   cd compozify
   ```
   If you have `$GOPATH/bin` or `$GOBIN` in your `$PATH`, you can just install with `make install` (install compozify in `$GOPATH/bin`) and **skip steps 3 and 4**.

3. Build the project
   ```sh
   make build
   ```

4. Change PATH to find newly compiled `compozify`

   ```sh
   export PATH=$PWD/bin:$PATH
   ```

4. Run `compozify --version` to confirm that it worked

## License
Copyright © [Clement Sam](https://twitter.com/clems_dev)

`compozify` is open-sourced software licensed under the [MIT](LICENSE) license.
