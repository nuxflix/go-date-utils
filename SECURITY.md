# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 0.1.x   | :white_check_mark: |

The `main` branch is also supported during development. Security fixes are
released through the next appropriate version according to the project's
semantic-versioning policy.

## Reporting a Vulnerability

If you discover a security vulnerability in **go-date-fns**, please **do not** open a public GitHub issue.

Instead, please report it privately by:

1. Opening a [GitHub Security Advisory](https://github.com/chmenegatti/go-date-fns/security/advisories/new) (preferred), or
2. Emailing the maintainer directly (see the GitHub profile for contact info).

Please include:
- A description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if available)

Please do not include credentials, private keys, or other sensitive data in a
report. Test only against your own systems and data.

We will respond within **5 business days** and aim to release a fix within **14 days** of confirmation.

## Notes

`go-date-fns` is a pure date utility library with **zero external dependencies** and no network I/O, so the attack surface is very limited. Most security concerns would involve incorrect date/time calculations in security-critical contexts (e.g., token expiry, session management).
