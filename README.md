# namecheap-cli

This CLI was created to speed up the process for my engineers in procurement/renewal/management of domains, SSL certificates
in Namecheap.

As this CLI was created with the goal of using it in a company scenario,
the registrant information is constant for all products purchased in the account.
As such, the registrant information is placed in a consts.go file.

Please note that due to the limitation with Namecheap API operations, this
CLI will only work with account balance and not credit card payment methods.

Please refer to Namecheap documentations for more information.

## Configure

Configure the constants in consts/consts.go before you build the executable.

Note: Remember to change the URL to production before usage. Defaults to sandbox.

## Build

```bash
go build namecheap.go
```

## Build with Earthly

Building with Earthly will generate an executable in the bin/ directory

```bash
earthly +build
```

## Sample CLI Output

```bash

        A CLI tool built from Go
        that interact with NameCheap API

Usage:
  namecheap [command]

Available Commands:
  domains     command to interact with domains in namecheap
  help        Help about any command
  ssl         command to interact with ssl certificates in namecheap
  users       get information and perform actions on your namecheap account

Flags:
  -h, --help   help for namecheap

Use "namecheap [command] --help" for more information about a command.
```