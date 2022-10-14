# gh-profile

[![Release](https://github.com/gabe565/gh-profile/actions/workflows/release.yml/badge.svg)](https://github.com/gabe565/gh-profile/actions/workflows/release.yml)

Work with multiple GitHub accounts using the [gh cli](https://cli.github.com/).

## Installation

```shell
gh extension install gabe565/gh-profile
```

## Usage

- `gh profile create` - Prompts to add a new profile.
  - `gh profile create example` - Adds a profile named example.
- `gh profile switch` - Prompts to choose a profile from an interactive list.
  - `gh profile switch default` - Switches to a profile named default.
- `gh profile rename` - Prompts to rename a profile.
  - `gh profile rename example example2` - Renames a profile named example to example2.
- `gh profile list` - Lists all profiles.
- `gh profile delete` - Prompts to remove a profile.
  - `gh profile delete example` - Deletes a profile named example.

See [generated usage docs](./docs/profile.md) for more information.
