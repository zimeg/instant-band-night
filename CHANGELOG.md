# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][changelog], and this project adheres
to [Semantic Versioning][semver].

## [Unreleased]

### Enhancement

- Created a top level 'ibn' command as the starting point for this program
- Output collected information from a joining musician back to stdout
- Included a global '--date' flag to set the event date that defaults to today
- Save information about musicians who enter musical buckets to an event file
- List all musicians with instruments using the 'musician list' command
- Join unique musicians to create a band of specified size with 'band create'
- List all performing bands and the musicians that performed with 'band list'
- Prevent musicians from finding a place in bands for a certain '--cooldown'
- Alias extended 'musician join' and 'band create' commands to 'join' and 'draw'

### Fixed

- Output a helpful message with information about common commands by default
- Hide repeated error messages and usage information from failed command output
- Exit with an error message and error code before ending failed executions
- Collect the musician IDs when loading the current event into this process

### Documentation

- Write a few basic words about the project on the project README
- Create this exact changelog file to log changes in a changelog file
- Link to the original and wonderful specifications of technical tool
- Record a short demonstration video with a few common commands

### Maintenance

- Share all source code under the open and permissive MIT license
- Setup basic build script commands with teardown using Makefile
- Instantiate development dependencies on load using a Nix flake
- Request code patterns follow the same shared linting practices

<!-- a collection of links -->
[changelog]: https://keepachangelog.com/en/1.1.0/
[semver]: https://semver.org/spec/v2.0.0.html

<!-- a collection of releases -->
[Unreleased]: https://github.com/zimeg/instant-band-night/compare/HEAD
