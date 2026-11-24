# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2026-11-23

### Added
- `Execute` method on `Template` — callers no longer need to reach into the struct
- Nix and `go install` instructions in README

### Changed
- Go version bumped to 1.27
- Library: `Make` renamed to `Compile`
- Library: sentinel errors renamed to idiomatic `ErrInvalidExtends` and `ErrUnresolvableExtension`
- Library: `Template.Inner` is now private; use `Template.Execute` instead
- CLI: `--std-in` flag renamed to `--stdin`
- CLI: error output now uses `stderr` directly instead of `log`
- Filesystem path handling switched from `path` to `path/filepath`

### Fixed
- `template.Must(Clone())` replaced with proper error handling — no more panics on Clone failure
- `extends` line parsing now uses `SplitN` to correctly handle paths containing spaces

## [0.0.1] - 2024-01-01

### Added
- Initial release: `extends` directive for Go `html/template`
- CLI with `--std-in` and `--print-dependencies` flags
- Nix flake build

[Unreleased]: https://github.com/sdedovic/htpp/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/sdedovic/htpp/compare/v0.0.1...v0.1.0
[0.0.1]: https://github.com/sdedovic/htpp/releases/tag/v0.0.1
