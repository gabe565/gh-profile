# gh-profile

[![Build](https://github.com/gabe565/gh-profile/actions/workflows/build.yml/badge.svg)](https://github.com/gabe565/gh-profile/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gabe565/gh-profile?v=1)](https://goreportcard.com/report/github.com/gabe565/gh-profile)

Work with multiple GitHub accounts using the [gh cli](https://cli.github.com/).

<img alt="gh-profile demo" src="../demo/demo.gif" width="640">

## Installation

```shell
gh extension install gabe565/gh-profile
```

## Usage

See the [generated usage docs](./docs/gh-profile.md), or see a summary of each
subcommand below.

> **Note**
> As of v2.26.0, the gh cli now uses secure auth tokens by default.
> Secure auth tokens are not yet supported by gh-profile, so when logging into GitHub, make sure to run `gh auth login --insecure-storage`.

- **`gh profile create [NAME]`:** Creates a new profile.
  <details>
    <summary>Details</summary>

  **Params**
  - `NAME` is optional. If not set, command will run interactively.

  **Example**
  ```shell
  $ gh profile create example
  ✨ Creating profile: example
  🔧 Activating global profile: example
  ```
  </details>
- **`gh profile switch [NAME] [--local-dir]`:** Activates a profile.
  <details>
    <summary>Details</summary>

  **Params**
  - `NAME` is optional. If not set, command will run interactively.
    - If set to `-`, gh-profile will switch back to the previous profile.
  - `--local-dir`/`-l` activates the profile only for the current directory.
    - For this to work, you must install a per-directory env tool like [direnv](https://direnv.net).

  **Example**
  ```shell
  $ gh profile switch example
  🔧 Activating global profile: example
  ```
  </details>
- **`gh profile rename [NAME] [NEW_NAME]`:** Renames a profile.
  <details>
    <summary>Details</summary>

  **Params**
  - `NAME` and `NEW_NAME` are optional. If not set, command will run interactively.

  **Example**
  ```shell
  $ gh profile rename example example2
  🚚 Renaming profile: example to example2
  🔧 Activating global profile: example2
  ```
  </details>
- **`gh profile list`:** Lists all profiles. Active profile will be bold with a green check.
  <details>
    <summary>Details</summary>

  **Example**
  ```shell
  $ gh profile list
  ✓ example
    gabe565
  ```
  </details>
- **`gh profile remove [NAME]`:** Removes a profile.
  <details>
    <summary>Details</summary>

  **Params**
  - `NAME` is optional. If not set, command will run interactively.

  **Example**
  ```shell
  $ gh profile remove example2
  🔥 Removing profile: example2
  ```
  </details>

- **`gh profile show`:** Prints the active profile name. If no profile is active, nothing will be printed. Useful as a [prompt element](#prompt-element).
  <details>
    <summary>Details</summary>

  **Example**:
  ```shell
  $ gh profile show
  example
  ```
  </details>

## Prompt Element

`gh profile show` is useful for displaying the current profile in your
shell's prompt. This command will work for any prompt, but configuration
with [Powerlevel10k](https://github.com/romkatv/powerlevel10k) is provided
below.

### Powerlevel10k

Powerlevel10k ships with a custom formatter for `git` repositories. This
formatter can be easily modified to show the current profile.

1. Edit `~/.p10k.zsh`.
2. Find the `my_git_formatter` function
3. Find the line `local res`
4. Add the following below that line:
    ```shell
        local profile="$(gh profile show 2>/dev/null)"
        [[ -n "$profile" ]] && res+="$profile "
    ```

Now, the current profile will be shown when you are in a git repo!

#### Example

| Before | After |
|--------|-------|
| <img width="280" alt="Before" src="https://user-images.githubusercontent.com/114527278/199317857-876031b4-ac6f-45e5-84c5-304eadcbf5e6.png"> | <img width="345" alt="After" src="https://user-images.githubusercontent.com/114527278/199317888-7901518a-2a9c-40f8-8416-5c95cb62d60a.png"> |
