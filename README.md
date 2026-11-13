# 🗺️ marauder

> Maps a Linux host and reveals privilege-escalation paths — a lean, single-binary post-exploitation enumeration tool in Go.

![status](https://img.shields.io/badge/status-early%20development-orange)
![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![license](https://img.shields.io/badge/license-MIT-blue)

**marauder** is a lightweight, single-binary enumeration tool for Linux post-exploitation. Once you have a foothold on a host (in an authorized engagement or a CTF), it maps the system — like a Marauder's Map for a box — and reveals local privilege-escalation paths and useful secrets.

It is an **assessment tool**: it enumerates and reports — it does **not** exploit anything.

> ⚠️ **Early development / work in progress.**

## ⚖️ Scope & authorized use

This tool is intended **only** for systems you own or have explicit written permission to test (CTFs, labs, authorized penetration tests). Using it against systems without authorization is illegal. You are responsible for how you use it.

## Commands

marauder is organized into three modules, each named for what it does:

| Command | What it does |
|---|---|
| `marauder` (no args) | Full map — run every module |
| `marauder revelio` | **Reveal** — system enumeration (OS, kernel, users, groups, mounts) |
| `marauder alohomora` | **Unlock** — privilege-escalation paths (SUID, sudo, cron, writable services) |
| `marauder niffler` | **Dig** — hunt for secrets (SSH keys, shell history, creds in configs) |

## What it checks (planned)

**revelio** — system
- OS and kernel version (for known-vuln lookup)
- Users and groups, high-UID accounts
- Environment variables and `PATH` (writable entries?)
- Mounted / unmounted drives, `fstab` creds

**alohomora** — privilege-escalation paths
- SUID / SGID binaries
- Sudo rules and misconfigurations
- Cron jobs and writable cron/service paths
- World-writable files and directories

**niffler** — secrets
- SSH private keys
- Shell history files
- Credentials in config files
- Other interesting files

Checks are added incrementally — see the [roadmap](#roadmap).

## Design goals

- Single static binary, no runtime dependencies on the target
- Standard library only, where possible
- Native Go data collection over shelling out, to keep a small footprint
- Clear output (human-readable now; machine-readable / JSON planned)

## Build

```bash
# Standard build
go build -o marauder .

# Static, stripped build (smaller, no libc dependency)
CGO_ENABLED=0 go build -ldflags="-s -w" -o marauder .
```

## Usage

```bash
./marauder              # full map — every module
./marauder revelio      # system enumeration only
./marauder alohomora    # privilege-escalation paths only
./marauder niffler      # secret-hunting only
```

## Roadmap

- [x] Project skeleton
- [ ] First check: SUID binaries (`alohomora`)
- [ ] `revelio` core system checks
- [ ] `alohomora` core privesc checks
- [ ] `niffler` core secret checks
- [ ] Structured (JSON) output
- [ ] Concurrent check execution
- [ ] Windows support

## Why this exists

A learning project — built to go deep on Linux privilege-escalation enumeration, and to learn Go by writing it from the ground up.

## License

Released under the MIT License — see [`LICENSE`](LICENSE).