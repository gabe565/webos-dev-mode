# LG Dev Mode CLI

A command-line tool to extend the LG dev mode session timer.

## Installation

### Docker

<details>
  <summary>Click to expand</summary>

A Docker image is available at [ghcr.io/gabe565/lg-dev-mode](https://ghcr.io/gabe565/lg-dev-mode)

```shell
sudo docker run --rm -it ghcr.io/gabe565/lg-dev-mode cron --token SESSION_TOKEN
```
</details>

### Homebrew (macOS, Linux)

<details>
  <summary>Click to expand</summary>

Install lg-dev-mode from [gabe565/homebrew-tap](https://github.com/gabe565/homebrew-tap):
```shell
brew install gabe565/tap/lg-dev-mode
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

4. Install lg-dev-mode
   ```shell
   sudo apt install lg-dev-mode
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

3. Install lg-dev-mode
   ```shell
   sudo dnf install lg-dev-mode
   ```
</details>

### AUR (Arch Linux)

<details>
  <summary>Click to expand</summary>

Install [lg-dev-mode-bin](https://aur.archlinux.org/packages/lg-dev-mode-bin) with your [AUR helper](https://wiki.archlinux.org/index.php/AUR_helpers) of choice.
</details>

### Manual Installation

<details>
  <summary>Click to expand</summary>

Download and run the [latest release binary](https://github.com/gabe565/lg-dev-mode/releases/latest) for your system and architecture.
</details>

## Usage

- To extend the LG dev mode session:
  ```shell
  lg-dev-mode extend --token SESSION_TOKEN
  ```
- To run a task which will extend the LG dev mode session once per day:
  ```shell
  lg-dev-mode cron --token SESSION_TOKEN
  ```
- To check the current session expiration:
  ```shell
  lg-dev-mode check --token SESSION_TOKEN
  ```

For full command-line reference, see [docs](docs/lg-dev-mode.md).
