# github-toy

Toy to operate github. 🤖

## Install

#### From source

```bash
go get github.com/yukihirop/github-toy
cd $GOPATH/src/github.com/yukihirop/github-toy
make install
```

#### From homebrew

```bash
brew tap yukihirop/github-toy
brew install gtoy
```

## Supported OSes

- Mac
- Linux

## Usage

```bash
$ gtoy
NAME:
   github-toy - Toy to operate github

USAGE:
   gtoy [global options] command [command options] [arguments...]

VERSION:
   0.1.0 (rev:HEAD)

COMMANDS:
   settings  Operate github display name
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

#### Update settings name

```bash
GITHUB_AUTH_TOKEN=<your personal token> gtoy settings --login yukihirop update --name yukihirop
GITHUB_AUTH_TOKEN=<your personal token> gtoy settings -l yukihirop update -n yukihirop@Hoge.inc
```
