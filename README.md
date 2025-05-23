# webOS Dev Mode CLI
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/gabe565/webos-dev-mode)](https://github.com/gabe565/webos-dev-mode/releases)
[![Build](https://github.com/gabe565/webos-dev-mode/actions/workflows/build.yml/badge.svg)](https://github.com/gabe565/webos-dev-mode/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gabe565/webos-dev-mode)](https://goreportcard.com/report/github.com/gabe565/webos-dev-mode)

A command-line tool to extend the webOS dev mode session timer.

## Installation

### Docker

<details>
  <summary>Click to expand</summary>

A Docker image is available at [ghcr.io/gabe565/webos-dev-mode](https://ghcr.io/gabe565/webos-dev-mode)

```shell
sudo docker run --rm -it ghcr.io/gabe565/webos-dev-mode cron --token SESSION_TOKEN
```
</details>

### Homebrew (macOS, Linux)

<details>
  <summary>Click to expand</summary>

Install webos-dev-mode from [gabe565/homebrew-tap](https://github.com/gabe565/homebrew-tap):
```shell
brew install gabe565/tap/webos-dev-mode
```
</details>

### APT (Ubuntu, Debian)

<details>
  <summary>Click to expand</summary>

1. If you don't have it already, install the `ca-certificates` package
   ```shell
   sudo apt install ca-certificates
   ```

2. Add gabe565 apt repository
   ```
   echo 'deb [trusted=yes] https://apt.gabe565.com /' | sudo tee /etc/apt/sources.list.d/gabe565.list
   ```

3. Update apt repositories
   ```shell
   sudo apt update
   ```

4. Install webos-dev-mode
   ```shell
   sudo apt install webos-dev-mode
   ```
</details>

### RPM (CentOS, RHEL)

<details>
  <summary>Click to expand</summary>

1. If you don't have it already, install the `ca-certificates` package
   ```shell
   sudo dnf install ca-certificates
   ```

2. Add gabe565 rpm repository to `/etc/yum.repos.d/gabe565.repo`
   ```ini
   [gabe565]
   name=gabe565
   baseurl=https://rpm.gabe565.com
   enabled=1
   gpgcheck=0
   ```

3. Install webos-dev-mode
   ```shell
   sudo dnf install webos-dev-mode
   ```
</details>

### AUR (Arch Linux)

<details>
  <summary>Click to expand</summary>

Install [webos-dev-mode-bin](https://aur.archlinux.org/packages/webos-dev-mode-bin) with your [AUR helper](https://wiki.archlinux.org/index.php/AUR_helpers) of choice.
</details>

### Manual Installation

<details>
  <summary>Click to expand</summary>

Download and run the [latest release binary](https://github.com/gabe565/webos-dev-mode/releases/latest) for your system and architecture.
</details>

## Usage

1. Set up the webOS dev mode by following the [webosbrew](https://www.webosbrew.org/devmode/) guide.
3. Fetch the session token:
   ```shell
   ares-novacom --run 'cat /var/luna/preferences/devmode_enabled; echo'
   ```

Now you are ready to run `webos-dev-mode` commands.
- To extend the webOS dev mode session:
  ```shell
  webos-dev-mode extend --token SESSION_TOKEN
  ```
- To run a task which will extend the webOS dev mode session once per day:
  ```shell
  webos-dev-mode cron --token SESSION_TOKEN
  ```
- To check the current session expiration:
  ```shell
  webos-dev-mode check --token SESSION_TOKEN
  ```

Flag values can also be set using environment variables. To do this, capitalize all characters, replace `-` with `_`, and prefix with `WEBOS_`. For example, `--token=example` would become `WEBOS_TOKEN=example`, and `--request-timeout=10m` would become `WEBOS_REQUEST_TIMEOUT=10m`.

For full command-line reference, see [docs](docs/webos-dev-mode.md).
