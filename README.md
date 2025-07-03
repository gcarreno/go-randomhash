# Go RandomHash

[![Supports Windows](https://img.shields.io/badge/support-Windows-blue?logo=Windows)](https://github.com/gcarreno/go-randomhash/releases/latest)
[![Supports Linux](https://img.shields.io/badge/support-Linux-yellow?logo=Linux)](https://github.com/gcarreno/go-randomhash/releases/latest)
[![Supports macOS](https://img.shields.io/badge/support-macOS-black?logo=macOS)](https://github.com/gcarreno/go-randomhash/releases/latest)
[![License](https://img.shields.io/github/license/gcarreno/go-randomhash)](https://github.com/gcarreno/go-randomhash/blob/main/LICENSE)
[![Latest Release](https://img.shields.io/github/v/release/gcarreno/go-randomhash?label=latest%20release)](https://github.com/gcarreno/go-randomhash/releases/latest)
[![Downloads](https://img.shields.io/github/downloads/gcarreno/go-randomhash/total)](https://github.com/gcarreno/go-randomhash/releases)

A Go version of the cryptocoin's proof of work `RandomHash`.

Since there's not an official specification for `RandomHash`, in contrast to `RandomX`, this is my interpretation of it's main idea.

It has entropy measures that include random calls to `XOR`, `AND`, `RotateLeft`, `FlipBits` and `ReverseBytes`.\
Ideally, this will make `GPU`/`ASIC` sweat a little. Still.. Needs confirmation :smile: