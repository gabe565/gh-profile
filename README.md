# gh-profile

[![Release](https://github.com/gabe565/gh-profile/actions/workflows/release.yml/badge.svg)](https://github.com/gabe565/gh-profile/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gabe565/gh-profile?v=1)](https://goreportcard.com/report/github.com/gabe565/gh-profile)

Work with multiple GitHub accounts using the [gh cli](https://cli.github.com/).

## Installation

```shell
gh extension install gabe565/gh-profile
```

## Usage

See the [generated usage docs](./docs/profile.md), or see a summary of each
subcommand below.


### `gh profile create [NAME]`
Creates a new profile.

#### Params
- `NAME` is optional. If not set, command will run interactively.


### `gh profile switch [NAME] [--local-dir]`
Activates a profile.

#### Params
- `NAME` is optional. If not set, command will run interactively.
- `--local-dir`/`-l` activates the profile only for the current directory.
  - For this to work, you must install a per-directory env tool like
  [direnv](https://direnv.net).


### `gh profile rename [NAME] [NEW_NAME]`
Renames a profile.

#### Params
- `NAME` and `NEW_NAME` are optional. If not set, command will run interactively.


### `gh profile list`
Lists all profiles. Active profile will be bold with a green check.  


### `gh profile remove [NAME]`
Removes a profile.

#### Params
- `NAME` is optional. If not set, command will run interactively.


### `gh profile show`
Prints the active profile name. If no profile is active, will print `none`.
