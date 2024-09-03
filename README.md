<div align="center">

# Akin

A fuzzy comparison algorithm for Go values.

[![Documentation](https://img.shields.io/badge/go.dev-documentation-007d9c?&style=for-the-badge)](https://pkg.go.dev/github.com/dogmatiq/akin)
[![Latest Version](https://img.shields.io/github/tag/dogmatiq/akin.svg?&style=for-the-badge&label=semver)](https://github.com/dogmatiq/akin/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/dogmatiq/akin/ci.yml?style=for-the-badge&branch=main)](https://github.com/dogmatiq/akin/actions/workflows/ci.yml)
[![Code Coverage](https://img.shields.io/codecov/c/github/dogmatiq/akin/main.svg?style=for-the-badge)](https://codecov.io/github/dogmatiq/akin)

</div>

## Notation and terminology

- `𝑷` denotes some `predicate`
- `𝒙` is a `value` against which `𝑷` can be `evaluated`
- `𝐐` or `𝐐ₙ` is a `constituent` predicate of `𝑷`
- The `≔` symbol defines some truth as a `given`, for example `𝒙 ≔ 7`
- `=` and `≠` show that a value is (or is not) equal to another value,
  respectively.

- `≍` and `≭` show that a `value` is (or is not) equivalent to some abstract
  value, which is not necessarily representable as a Go value.
