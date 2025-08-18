# Clipboard Clearer

<p align="center">
  <img src="./assets/clipboard_clearer_logo.png" alt="Go Semver Release Logo" width="230">
  <br><br>
  <a href="https://img.shields.io/github/v/tag/s0ders/clipboard-clearer?label=Version&color=bb33ff"><img alt="GitHub Tag" src="https://img.shields.io/github/v/tag/s0ders/clipboard-clearer?label=Version&color=bb33ff"></a>
  <a href="https://img.shields.io/github/go-mod/go-version/s0ders/clipboard-clearer"><img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/s0ders/clipboard-clearer"></a>
  <a href="https://img.shields.io/github/actions/workflow/status/s0ders/clipboard-clearer/build.yaml?label=CI"><img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/s0ders/clipboard-clearer/build.yaml?label=CI"></a>
  <a href="https://goreportcard.com/report/github.com/s0ders/clipboard-clearer"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/s0ders/clipboard-clearer"></a>
  <a href="https://github.com/s0ders/go-semver-release/blob/main/LICENSE.md"><img alt="GitHub License" src="https://img.shields.io/github/license/s0ders/clipboard-clearer?label=License"></a>
</p>

Program that clears the content of the OS clipboard after a given amount of time.

## Configuration

The program is accessible from the system tray. From there, one can configure the clipboard expiration timeout by
increasing or decreasing the default value of one minute.
<br><br>
Available expiration delays are:
- Ten seconds
- Thirty seconds
- One minute
- Five minutes
- Ten minutes
- One hour

If the expiration delay is updated, the new timer will take into account the time that has already elapsed for the current
clipboard content. For instance, if you increase the delay from 30 seconds to 1 minute but 20 seconds already 
elapsed, the effective lifetime of the current clipboard content will be 40s. Later clipboard content will then expire 
after a full one minute.


## Installation

### Linux and Windows
Releases for Linux and Windows are available on the [release page](https://github.com/s0ders/clipboard-clearer/releases)
of the repository.

### macOS
Releases are not yet published for macOS, but you can build the program from source if [Go](https://go.dev/dl/) is 
installed on your machine.

## Backlog

The following features are either in development or being considered:
- [ ] Optional clipboard history (stored in as a log file)
- [ ] Release for macOS (requires signing and an Apple dev. account)