# gh-profile

[![Release](https://github.com/gabe565/gh-profile/actions/workflows/release.yml/badge.svg)](https://github.com/gabe565/gh-profile/actions/workflows/release.yml)

Work with multiple GitHub accounts using the [gh cli](https://cli.github.com/).

## Installation

```shell
gh extension install gabe565/gh-profile
```

## Usage

See the [generated usage docs](./docs/profile.md), or see a summary of each
subcommand below.

### `gh profile create [NAME]`
Creates a new profile called `name`.  
Name is optional. If not set, command will run interactively.  
Aliases: `new`, `add`

### `gh profile switch [NAME]`
Activates the profile called `name`.  
Name is optional. If not set, command will run interactively.  
Aliases: `activate`, `active`, `sw`

### `gh profile rename [NAME] [NEW_NAME]`
Renames a profile called `old` to `new`.  
Old and new name are optional. If not set, command will run interactively.  
Aliases: `mv`

### `gh profile list`
Lists all profiles. Active profile will be bold with a green check.  
Aliases: `ls`, `l`

### `gh profile remove [NAME]`
Removes a profile called `name`.  
Name is optional. If not set, command will run interactively.  
Aliases: `delete`, `rm`

### `gh profile show`
Prints the active profile name. If no profile is active, will print `none`.
