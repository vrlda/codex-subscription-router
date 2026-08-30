# Changelog

All notable changes follow [Keep a Changelog](https://keepachangelog.com/) and
this project uses [Semantic Versioning](https://semver.org/).

## [Unreleased]

## [0.2.0] - 2026-08-30

### Added

- One-command installer with safe source updates, prerequisite checks, signed
  rebuilds, recoverable upgrades, and automatic launch.
- Reset-aware routing that prioritizes weekly quota at risk of expiring and
  gives a bounded boost to subscriptions with banked usage resets.
- Shared rollout storage and thread indexing so router-created chats remain
  discoverable by every subscription and the official app.
- Manual per-chat subscription switching with active-turn and quota guards.
- Automatic continuation on another account after a terminal usage-limit
  failure.
- Per-account five-hour and weekly usage, reset timestamps, and a native-style
  selector in the pinned thread summary.
- Fail-closed support for ChatGPT desktop `26.825.32147` (build `7303`).

### Changed

- New-turn routing now excludes accounts whose five-hour or weekly window is
  depleted.
- Apple silicon is selected explicitly for the launcher and multiplexer build.
- Machines without a team-backed certificate automatically use ad-hoc signing.

### Fixed

- First-run installation under macOS Bash 3.2 no longer expands an empty array
  with `set -u` enabled.
- Restricted push entitlements are removed from independently signed copies so
  macOS does not terminate ad-hoc builds at launch.
- Legacy isolated thread indexes are backed up and imported into the shared
  primary index during upgrade.

## [0.1.0] - 2026-08-15

### Added

- Multi-subscription routing with quota-aware balancing and sticky threads.
- Account isolation, device-code sign-in, pooled usage, and quota failover.
- Native account menu, masked emails, plan labels, and profile photos.
- Combined Profile statistics with per-account selection.
- Account-scoped Apps and MCP connection state in Settings → Plugins.
- Per-account rate-limit reset selection and pooled depletion handling.
- Independently signed Appshots and Computer Use support.
- Fail-closed upstream compatibility checks and deepest-first nested helper signing.
- Loopback-only, token-authenticated diagnostic UI states.
- Source-only CI, draft release automation, security documentation, and smoke tests.

[Unreleased]: https://github.com/vrlda/codex-subscription-router/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/vrlda/codex-subscription-router/releases/tag/v0.2.0
[0.1.0]: https://github.com/b-nnett/codex-subscription-router/releases/tag/v0.1.0
