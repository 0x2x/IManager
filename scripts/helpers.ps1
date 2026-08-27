param(
    [Parameter(Position = 0)]
    [ValidateSet(
        "install",
        "deps",
        "build",
        "test",
        "check",
        "run",
        "port",
        "release",
        "clean",
        "help"
    )]
    [string]$Command = "help"
)

$ErrorActionPreference = "Stop"

# ============================================================
# IManager Build Helper
# ============================================================

$ProjectRoot = Split-Path -Parent $PSScriptRoot

$AppName = "imanager"

$BuildDir = Join-Path $ProjectRoot "bin"

$PortsDir = Join-Path $ProjectRoot "ports"

$ReleaseDir = Join-Path $PortsDir "release"

$GoMod = Join-Path $ProjectRoot "go.mod"

# ============================================================
# General Helpers
# ============================================================

function Write-Step {
    param(
        [string]$Message
    )

    Write-Host ""
    Write-Host "==> $Message" -ForegroundColor Cyan
}

function Write-Success {
    param(
        [string]$Message
    )

    Write-Host ""
    Write-Host $Message -ForegroundColor Green
}

function Write-ErrorMessage {
    param(
        [string]$Message
    )

    Write-Host ""
    Write-Host "ERROR: $Message" -ForegroundColor Red
}

function Check-Go {

    if (-not (Get-Command go -ErrorAction SilentlyContinue)) {

        Write-ErrorMessage "Go is not installed or is not available in PATH."

        exit 1
    }

    if (-not (Test-Path $GoMod)) {

        Write-ErrorMessage "go.mod was not found."

        Write-Host ""
        Write-Host "Expected:"
        Write-Host $GoMod

        exit 1
    }

    Write-Host "Go version:"
    go version
}

function Invoke-Go {

    param(
        [Parameter(Mandatory = $true)]
        [string[]]$Arguments
    )

    Push-Location $ProjectRoot

    try {

        & go @Arguments

        if ($LASTEXITCODE -ne 0) {

            throw "Go command failed with exit code $LASTEXITCODE"
        }

    }
    finally {

        Pop-Location
    }
}

function Ensure-Directory {

    param(
        [string]$Path
    )

    if (-not (Test-Path $Path)) {

        New-Item `
            -ItemType Directory `
            -Path $Path `
            -Force | Out-Null
    }
}

# ============================================================
# Install Dependencies
# ============================================================

function Install-Dependencies {

    Write-Step "Checking Go"

    Check-Go

    Write-Step "Downloading Go dependencies"

    Invoke-Go @(
        "mod",
        "download"
    )

    Write-Step "Tidying Go modules"

    Invoke-Go @(
        "mod",
        "tidy"
    )

    Write-Step "Installing Cobra CLI"

    go install github.com/spf13/cobra-cli@latest

    if ($LASTEXITCODE -ne 0) {

        throw "Failed to install Cobra CLI."
    }

    Write-Success "Dependencies installed successfully."
}

# ============================================================
# Update Dependencies
# ============================================================

function Update-Dependencies {

    Write-Step "Checking Go"

    Check-Go

    Write-Step "Downloading dependencies"

    Invoke-Go @(
        "mod",
        "download"
    )

    Write-Step "Tidying dependencies"

    Invoke-Go @(
        "mod",
        "tidy"
    )

    Write-Success "Dependencies updated successfully."
}

# ============================================================
# Windows Build
# ============================================================

function Build-App {

    Write-Step "Checking Go"

    Check-Go

    Ensure-Directory $BuildDir

    $Output = Join-Path `
        $BuildDir `
        "$AppName.exe"

    Write-Step "Building Windows application"

    Invoke-Go @(
        "build",
        "-o",
        $Output,
        "."
    )

    Write-Success "Windows build completed."

    Write-Host ""
    Write-Host "Output:"
    Write-Host $Output
}

# ============================================================
# Tests
# ============================================================

function Run-Tests {

    Write-Step "Checking Go"

    Check-Go

    Write-Step "Running Go tests"

    Invoke-Go @(
        "test",
        "./..."
    )

    Write-Success "All tests passed."
}

# ============================================================
# Test + Build
# ============================================================

function Check-App {

    Write-Step "Running tests"

    Run-Tests

    Write-Step "Building application"

    Build-App

    Write-Success "Application passed tests and built successfully."
}

# ============================================================
# Run Application
# ============================================================

function Run-App {

    Write-Step "Checking Go"

    Check-Go

    Write-Step "Starting $AppName"

    Push-Location $ProjectRoot

    try {

        & go run .

        if ($LASTEXITCODE -ne 0) {

            throw "Application exited with code $LASTEXITCODE"
        }

    }
    finally {

        Pop-Location
    }
}

# ============================================================
# Build Darwin
# ============================================================

function Build-Darwin {

    $DarwinDir = Join-Path `
        $PortsDir `
        "darwin"

    Write-Step "Preparing Darwin directory"

    Ensure-Directory $DarwinDir

    # --------------------------------------------------------
    # macOS ARM64
    # --------------------------------------------------------

    $Arm64Dir = Join-Path `
        $DarwinDir `
        "arm64"

    Ensure-Directory $Arm64Dir

    $Arm64Output = Join-Path `
        $Arm64Dir `
        $AppName

    Write-Step "Building macOS ARM64"

    Push-Location $ProjectRoot

    try {

        $env:GOOS = "darwin"
        $env:GOARCH = "arm64"
        $env:CGO_ENABLED = "0"

        & go build `
            -o $Arm64Output `
            .

        if ($LASTEXITCODE -ne 0) {

            throw "macOS ARM64 build failed."
        }

    }
    finally {

        Remove-Item Env:GOOS `
            -ErrorAction SilentlyContinue

        Remove-Item Env:GOARCH `
            -ErrorAction SilentlyContinue

        Remove-Item Env:CGO_ENABLED `
            -ErrorAction SilentlyContinue

        Pop-Location
    }

    # --------------------------------------------------------
    # macOS Intel AMD64
    # --------------------------------------------------------

    $Amd64Dir = Join-Path `
        $DarwinDir `
        "amd64"

    Ensure-Directory $Amd64Dir

    $Amd64Output = Join-Path `
        $Amd64Dir `
        $AppName

    Write-Step "Building macOS Intel AMD64"

    Push-Location $ProjectRoot

    try {

        $env:GOOS = "darwin"
        $env:GOARCH = "amd64"
        $env:CGO_ENABLED = "0"

        & go build `
            -o $Amd64Output `
            .

        if ($LASTEXITCODE -ne 0) {

            throw "macOS AMD64 build failed."
        }

    }
    finally {

        Remove-Item Env:GOOS `
            -ErrorAction SilentlyContinue

        Remove-Item Env:GOARCH `
            -ErrorAction SilentlyContinue

        Remove-Item Env:CGO_ENABLED `
            -ErrorAction SilentlyContinue

        Pop-Location
    }

    Write-Success "Darwin builds completed."

    Write-Host ""
    Write-Host "macOS ARM64:"
    Write-Host $Arm64Output

    Write-Host ""
    Write-Host "macOS Intel:"
    Write-Host $Amd64Output
}

# ============================================================
# Build Linux
# ============================================================

function Build-Linux {

    $LinuxDir = Join-Path `
        $PortsDir `
        "linux"

    Write-Step "Preparing Linux directory"

    Ensure-Directory $LinuxDir

    # --------------------------------------------------------
    # Linux AMD64
    # --------------------------------------------------------

    $Amd64Dir = Join-Path `
        $LinuxDir `
        "amd64"

    Ensure-Directory $Amd64Dir

    $Amd64Output = Join-Path `
        $Amd64Dir `
        $AppName

    Write-Step "Building Linux AMD64"

    Push-Location $ProjectRoot

    try {

        $env:GOOS = "linux"
        $env:GOARCH = "amd64"
        $env:CGO_ENABLED = "0"

        & go build `
            -o $Amd64Output `
            .

        if ($LASTEXITCODE -ne 0) {

            throw "Linux AMD64 build failed."
        }

    }
    finally {

        Remove-Item Env:GOOS `
            -ErrorAction SilentlyContinue

        Remove-Item Env:GOARCH `
            -ErrorAction SilentlyContinue

        Remove-Item Env:CGO_ENABLED `
            -ErrorAction SilentlyContinue

        Pop-Location
    }

    # --------------------------------------------------------
    # Linux ARM64
    # --------------------------------------------------------

    $Arm64Dir = Join-Path `
        $LinuxDir `
        "arm64"

    Ensure-Directory $Arm64Dir

    $Arm64Output = Join-Path `
        $Arm64Dir `
        $AppName

    Write-Step "Building Linux ARM64"

    Push-Location $ProjectRoot

    try {

        $env:GOOS = "linux"
        $env:GOARCH = "arm64"
        $env:CGO_ENABLED = "0"

        & go build `
            -o $Arm64Output `
            .

        if ($LASTEXITCODE -ne 0) {

            throw "Linux ARM64 build failed."
        }

    }
    finally {

        Remove-Item Env:GOOS `
            -ErrorAction SilentlyContinue

        Remove-Item Env:GOARCH `
            -ErrorAction SilentlyContinue

        Remove-Item Env:CGO_ENABLED `
            -ErrorAction SilentlyContinue

        Pop-Location
    }

    Write-Success "Linux builds completed."
}

# ============================================================
# Build All Ports
# ============================================================

function Port-All {

    Write-Step "Preparing ports"

    Check-Go

    if (Test-Path $PortsDir) {

        Remove-Item `
            $PortsDir `
            -Recurse `
            -Force
    }

    Ensure-Directory $PortsDir

    Build-Darwin

    Build-Linux

    Write-Success "All platform ports completed."

    Write-Host ""
    Write-Host "Ports directory:"
    Write-Host $PortsDir
}

# ============================================================
# Create Release Archives
# ============================================================

function Create-Releases {

    Write-Step "Creating release directory"

    $ReleaseDir = Join-Path `
        $PortsDir `
        "release"

    if (Test-Path $ReleaseDir) {

        Remove-Item `
            $ReleaseDir `
            -Recurse `
            -Force
    }

    Ensure-Directory $ReleaseDir

    # --------------------------------------------------------
    # Darwin ARM64
    # --------------------------------------------------------

    $Source = Join-Path `
        $PortsDir `
        "darwin\arm64\$AppName"

    $Archive = Join-Path `
        $ReleaseDir `
        "$AppName-darwin-arm64.zip"

    if (Test-Path $Source) {

        Write-Step "Packaging Darwin ARM64"

        Compress-Archive `
            -Path $Source `
            -DestinationPath $Archive `
            -Force
    }

    # --------------------------------------------------------
    # Darwin AMD64
    # --------------------------------------------------------

    $Source = Join-Path `
        $PortsDir `
        "darwin\amd64\$AppName"

    $Archive = Join-Path `
        $ReleaseDir `
        "$AppName-darwin-amd64.zip"

    if (Test-Path $Source) {

        Write-Step "Packaging Darwin AMD64"

        Compress-Archive `
            -Path $Source `
            -DestinationPath $Archive `
            -Force
    }

    # --------------------------------------------------------
    # Linux AMD64
    # --------------------------------------------------------

    $Source = Join-Path `
        $PortsDir `
        "linux\amd64\$AppName"

    $Archive = Join-Path `
        $ReleaseDir `
        "$AppName-linux-amd64.zip"

    if (Test-Path $Source) {

        Write-Step "Packaging Linux AMD64"

        Compress-Archive `
            -Path $Source `
            -DestinationPath $Archive `
            -Force
    }

    # --------------------------------------------------------
    # Linux ARM64
    # --------------------------------------------------------

    $Source = Join-Path `
        $PortsDir `
        "linux\arm64\$AppName"

    $Archive = Join-Path `
        $ReleaseDir `
        "$AppName-linux-arm64.zip"

    if (Test-Path $Source) {

        Write-Step "Packaging Linux ARM64"

        Compress-Archive `
            -Path $Source `
            -DestinationPath $Archive `
            -Force
    }

    # --------------------------------------------------------
    # Windows AMD64
    # --------------------------------------------------------

    $WindowsDir = Join-Path `
        $PortsDir `
        "windows\amd64"

    Ensure-Directory $WindowsDir

    $WindowsOutput = Join-Path `
        $WindowsDir `
        "$AppName.exe"

    Write-Step "Building Windows AMD64 release"

    Push-Location $ProjectRoot

    try {

        $env:GOOS = "windows"
        $env:GOARCH = "amd64"
        $env:CGO_ENABLED = "0"

        & go build `
            -o $WindowsOutput `
            .

        if ($LASTEXITCODE -ne 0) {

            throw "Windows AMD64 release build failed."
        }

    }
    finally {

        Remove-Item Env:GOOS `
            -ErrorAction SilentlyContinue

        Remove-Item Env:GOARCH `
            -ErrorAction SilentlyContinue

        Remove-Item Env:CGO_ENABLED `
            -ErrorAction SilentlyContinue

        Pop-Location
    }

    $Archive = Join-Path `
        $ReleaseDir `
        "$AppName-windows-amd64.zip"

    Write-Step "Packaging Windows AMD64"

    Compress-Archive `
        -Path $WindowsOutput `
        -DestinationPath $Archive `
        -Force

    # --------------------------------------------------------
    # SHA256 Checksums
    # --------------------------------------------------------

    Write-Step "Generating SHA256 checksums"

    $ChecksumFile = Join-Path `
        $ReleaseDir `
        "SHA256SUMS.txt"

    if (Test-Path $ChecksumFile) {

        Remove-Item `
            $ChecksumFile `
            -Force
    }

    Get-ChildItem `
        -Path $ReleaseDir `
        -Filter "*.zip" |
    ForEach-Object {

        $Hash = Get-FileHash `
            $_.FullName `
            -Algorithm SHA256

        "$($Hash.Hash)  $($_.Name)" |
            Out-File `
                -FilePath $ChecksumFile `
                -Append `
                -Encoding utf8
    }

    Write-Success "Release packages created."

    Write-Host ""
    Write-Host "Release files:"
    Write-Host ""

    Get-ChildItem `
        $ReleaseDir |
    ForEach-Object {

        Write-Host "  $($_.Name)"
    }
}

# ============================================================
# Full Release
# ============================================================

function Build-Release {

    Write-Step "Starting full release build"

    Check-Go

    Write-Step "Running tests"

    Run-Tests

    Write-Step "Building platform ports"

    Port-All

    Write-Step "Creating release archives"

    Create-Releases

    Write-Host ""
    Write-Host "========================================" -ForegroundColor Green
    Write-Host " RELEASE BUILD COMPLETE " -ForegroundColor Green
    Write-Host "========================================" -ForegroundColor Green
    Write-Host ""
    Write-Host "Public release directory:"
    Write-Host $ReleaseDir
    Write-Host ""
}

# ============================================================
# Clean
# ============================================================

function Clean-App {

    Write-Step "Cleaning project"

    if (Test-Path $BuildDir) {

        Remove-Item `
            $BuildDir `
            -Recurse `
            -Force

        Write-Host "Removed bin/"
    }

    if (Test-Path $PortsDir) {

        Remove-Item `
            $PortsDir `
            -Recurse `
            -Force

        Write-Host "Removed ports/"
    }

    Check-Go

    Write-Step "Running go clean"

    Invoke-Go @(
        "clean"
    )

    Write-Success "Clean complete."
}

# ============================================================
# Help
# ============================================================

function Show-Help {

    Write-Host ""
    Write-Host "IManager Build Helper" -ForegroundColor Cyan
    Write-Host "========================================"
    Write-Host ""

    Write-Host "Usage:"
    Write-Host ""
    Write-Host "  .\scripts\helpers.ps1 <command>"
    Write-Host ""

    Write-Host "Development:"
    Write-Host ""
    Write-Host "  install   Install Go dependencies + Cobra CLI"
    Write-Host "  deps      Update Go dependencies"
    Write-Host "  test      Run all tests"
    Write-Host "  run       Run the application"
    Write-Host "  build     Build Windows application"
    Write-Host "  check     Test + build"
    Write-Host ""

    Write-Host "Releases:"
    Write-Host ""
    Write-Host "  port      Build macOS + Linux ports"
    Write-Host "  release   Test + build all platforms + package releases"
    Write-Host ""

    Write-Host "Maintenance:"
    Write-Host ""
    Write-Host "  clean     Remove build and release files"
    Write-Host "  help      Show this help message"
    Write-Host ""
}

# ============================================================
# Command Router
# ============================================================

switch ($Command) {

    "install" {
        Install-Dependencies
    }

    "deps" {
        Update-Dependencies
    }

    "build" {
        Build-App
    }

    "test" {
        Run-Tests
    }

    "check" {
        Check-App
    }

    "run" {
        Run-App
    }

    "port" {
        Port-All
    }

    "release" {
        Build-Release
    }

    "clean" {
        Clean-App
    }

    "help" {
        Show-Help
    }
}