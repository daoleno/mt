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
NAME:
   My Thought - Rocord all my thoughts

USAGE:
   mt [global options] command [command options] [arguments...]

COMMANDS:
   open     Open a thought
   cat      View a thought
   delete   Delete a thought
   list     List all thoughts
   clean    Clean all thoughts
   encrypt  Encrypt all thoughts
   decrypt  Decrypt all thoughts
   render   Render all markdown to beautiful html
   web      Server static html page
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## Support bash/zsh autocomplete

### Run following command to enable auto-complete.

```zsh
#bash
PROG=mt source autocomplete/bash_autocomplete
#zsh
PROG=mt source autocomplete/zsh_autocomplete
```

### Distribution and Persistent Autocompletion

#### Bash Support

MacOS

Copy `autocomplete/bash_autocomplete` into `/usr/local/etc/bash_completion.d` and rename it to the name `mt`. Don't forget to source the file or restart your shell to activate the auto-completion.

Linux

Copy `autocomplete/bash_autocomplete` into `/etc/bash_completion.d/` and rename it to the name `mt`. Don't forget to source the file or restart your shell to activate the auto-completion.

#### Zsh Support

Adding the following lines to your ZSH configuration file (usually .zshrc) will allow the auto-completion to persist across new shells:

```zsh
PROG=mt
_CLI_ZSH_AUTOCOMPLETE_HACK=1

#compdef $PROG

_cli_zsh_autocomplete() {

  local -a opts
  opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} --generate-bash-completion)}")

  _describe 'values' opts

  return
}

compdef _cli_zsh_autocomplete $PROG
```
