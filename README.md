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
  completion  Generates bash/zsh completion scripts
  decrypt     Decrypt all thoughts
  remove      Remove a thought
  encrypt     Encrypt all thoughts
  help        Help about any command
  list        List all thoughts
  open        Open a thought
  rename      Rename a thought
  render      Render all markdown to beautiful html
  report      Report
  web         Server static html page

Flags:
  -h, --help   help for mt

Use "mt [command] --help" for more information about a command.
```

## Completion

### Zsh

Generate completion script

```sh
mt completion
```

Put somewhere in your `$fpath` named `_mt`

```
# example path
/usr/local/share/zsh/site-functions/_mt
```
