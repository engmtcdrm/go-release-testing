# go-release-testing

[![Build and Test](https://github.com/engmtcdrm/go-release-testing/actions/workflows/build.yml/badge.svg)](https://github.com/engmtcdrm/go-release-testing/actions/workflows/build.yml)
[![Release](https://img.shields.io/github/v/release/engmtcdrm/go-release-testing.svg?label=Release)](https://github.com/engmtcdrm/go-release-testing/releases/latest)

This repository exists for testing CI/CD pipelines for Go releases.

## Usage

There are four files that are the core components of this repository for CI/CD functionality.

- `.github/workflow/build.yml` - This action builds and tests your Go program everytime a push to main occurs. It will ignore tagged commits because assumingly you are doing a release when a commit is tagged.
- `.github/workflow/release.yml` - This action does the actual release of your Go program by compiling it for multiple OSs. It also reads the file `.github/latest_release_notes.md` for use as the release notes for the release. It will only create a release if it matches semantic versioning (vX.Y.Z). At the moment it does not support any additional tags.
- `.github/latest_release_notes.md` - This is a Markdown file that contains the release notes for the latest release. This is used in `.github/workflow/release.yml` when creating a release.
- `.goreleaser.yml` - This file contains information when building a release. This includes any build variables, and what OSs to build for.
