<#
.SYNOPSIS
  Smart wrapper for `go test` in PowerShell

.DESCRIPTION
  Runs Go tests or benchmarks with optional filters for test names, directories, verbosity, and memory stats.

.PARAMETER Directory
  Folder path to run `go test` on (default: ./...)

.PARAMETER RunPattern
  Regex pattern to match test functions (default: .)

.PARAMETER BenchPattern
  Regex pattern to match benchmark functions (default: none)

.PARAMETER Verbose
  Switch to enable verbose test output

.PARAMETER BenchMem
  Switch to include memory stats in benchmarks

.EXAMPLE
  ./gtest.ps1 -Directory ./tests -BenchPattern . -BenchMem
#>

param(
  [string]$Directory = "./...",
  [string]$RunPattern = ".",
  [string]$BenchPattern = "",
  [switch]$Verbose,
  [switch]$BenchMem
)

# Build base command
$cmd = "go test $Directory -run=`"$RunPattern`""

# Add benchmark flags
if ($BenchPattern) {
    $cmd += " -bench=`"$BenchPattern`""
    if ($BenchMem) {
        $cmd += " -benchmem"
    }
}

# Verbose output
if ($Verbose) {
    $cmd += " -v"
}

Write-Host "Running: $cmd" -ForegroundColor Cyan
Invoke-Expression $cmd
