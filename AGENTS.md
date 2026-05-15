# AGENTS.md

This file provides guidance to coding agents (e.g. Claude Code, claude.ai/code) when working with code in this repository.

## Repository purpose

A tiny CLI that generates the URL form of SMTP relay auth credentials (hashed password into a URL-safe form) for use with `smtprelay`. Module path is `gomodules.xyz/url-generator`.

## Architecture

- `main.go` — the whole program.
- `vendor/` — checked-in deps.

## Common commands

```
go run main.go -p <password>
```

No Makefile, no tests.

## Conventions

- Single-file utility — keep it that way.
- License: see `LICENSE` if present.
- Module path is `gomodules.xyz/url-generator` (vanity URL).
