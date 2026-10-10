# osc7

`osc7` prints an OSC 7 escape sequence for the current working directory:

```text
ESC ] 7 ; file://<hostname><encoded-path> BEL
```

Terminals may use this sequence to update their current working directory. That lets tools which inspect the terminal process, such as launchers or file-open helpers, easily find the shell's current directory.

Correct URI path escaping is awkward in shell. This helper keeps prompt integration small and efficient.

The local hostname is emitted as the URI host. The directory path is encoded as a URI path, so `/` remains `/`, spaces become `%20`, and bytes such as `%`, `?`, `#`, and control characters are escaped.

## Build

```sh
make
```

## Usage

Run `osc7` from a prompt hook. It writes no trailing newline.

Bash:

```sh
PROMPT_COMMAND='osc7'
```

Zsh (untested):

```sh
precmd() { osc7 }
```

An absolute directory can be supplied explicitly:

```sh
osc7 /tmp
```

