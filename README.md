# Clipboard Clearer

<p align="center">
  <a href="https://img.shields.io/github/v/tag/s0ders/clipboard-clearer?label=Version&color=bb33ff"><img alt="GitHub Tag" src="https://img.shields.io/github/v/tag/s0ders/clipboard-clearer?label=Version&color=bb33ff"></a>
  <a href="https://img.shields.io/github/go-mod/go-version/s0ders/clipboard-clearer"><img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/s0ders/clipboard-clearer"></a>
  <a href="https://img.shields.io/github/actions/workflow/status/s0ders/clipboard-clearer/main.yaml?label=CI"><img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/s0ders/clipboard-clearer/main.yaml?label=CI"></a>
  <a href="https://goreportcard.com/report/github.com/s0ders/clipboard-clearer"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/s0ders/clipboard-clearer"></a>
  <a href="https://github.com/s0ders/go-semver-release/blob/main/LICENSE.md"><img alt="GitHub License" src="https://img.shields.io/github/license/s0ders/clipboard-clearer?label=License"></a>
</p>

Clipboard Clearer is a program that clears the content of your OS clipboard after a given amount of time.

## TODOs

- [x] GitHub Actions pipeline to lint, tests and release
- [x] Handle image clipboard
- [ ] Make clipboard liveness time configurable (cfg file with a default of 10s ?) via systray maybe?
- [ ] Create App bundle for macOS in CI/CD workflow