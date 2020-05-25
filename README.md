# MT - My thoughts

Record your valuable thoughts on terminal.

[![asciicast](https://asciinema.org/a/325041.svg)](https://asciinema.org/a/325041)

## Build

```zsh
go build .
```

## Install

```zsh
go install .
```

## Usage

```
Rocord all my thoughts

Usage:
  mt [command]

Available Commands:
  cat         View a thought
  clean       Clean all thoughts
  decrypt     Decrypt all thoughts
  delete      Delete a thought
  encrypt     Encrypt all thoughts
  help        Help about any command
  list        List all thoughts
  open        Open a thought
  rename      Rename a thought
  render      Render all markdown to beautiful html
  web         Server static html page

Flags:
      --config string   config file (default is $HOME/.mt.yaml)
  -h, --help            help for mt
  -t, --toggle          Help message for toggle

Use "mt [command] --help" for more information about a command.
```
