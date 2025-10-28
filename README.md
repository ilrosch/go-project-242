# Disk size analyzer

[![Actions Status](https://github.com/ilrosch/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/ilrosch/go-project-242/actions) [![main](https://github.com/ilrosch/go-project-242/actions/workflows/main.yml/badge.svg)](https://github.com/ilrosch/go-project-242/actions/workflows/main.yml)

A cli utility that determines the size of a file or directory with flexible output settings. Similar to the standard UNIX utility du.

Features:
- **Recursive traversal** – calculates the size of all nested files and folders if needed.
- **Human-readable format** – automatically selects appropriate units of measurement (bytes, kilobytes, megabytes, etc.).
- **Hidden file support** – includes files and folders starting with a period (dotfiles).

Help for use:

```
    ./bin/hexlet-path-size -h
    NAME:
        hexlet-path-size - print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)

    USAGE:
        hexlet-path-size [global options] path

    GLOBAL OPTIONS:
        --recursive, -r  recursive size of directories (default: false)
        --human, -H      human-readable sizes (auto-select unit) (default: false)
        --all, -a        include hidden files and directories (default: false)
        --help, -h       show help
```
---

## Example

With the --human (or -H) flag, a human-readable size is output:

```
    ./bin/hexlet-path-size output.dat --human
    24.0MB  output.dat
```

With the --all (or -a) option, all files are included, including hidden ones:
```
    ./bin/hexlet-path-size project/ -H -a
    27.0MB  project/
```

The --recursive (or -r) flag includes all nested files and directories:
```
    ./bin/hexlet-path-size project/ -H -a -r
    31.0MB  project/
```

Demo:

[![asciicast](https://asciinema.org/a/YTGSedsrk7oxpE6a4cXl9vmkb.svg)](https://asciinema.org/a/YTGSedsrk7oxpE6a4cXl9vmkb)

---

## Tech stack
- **Language:** Golang
- **Packages:**
    - [urfave/cli](https://cli.urfave.org/v3/getting-started/)
    - [testify](https://github.com/stretchr/testify)