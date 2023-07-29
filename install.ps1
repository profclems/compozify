<#
.SYNOPSIS
    Installs Compozify to the system or user level.
.DESCRIPTION
    This script downloads the latest release of Compozify from GitHub and installs it to the system or user level.
    The installation directory is added to the PATH environment variable.
.PARAMETER Global
    Installs Compozify at the system level. This requires administrator privileges.
.PARAMETER Usage
    Shows the usage message.
.LINK
    https://github.com/profclems/compozify
#>
param (
    [switch]$Global,
    [switch]$Usage
)

# Disable StrictMode in this script
Set-StrictMode -Off

# Show usage message if the -Usage flag is specified
if ($Usage) {
    Write-Host "Compozify Installer"
    Write-Host "Usage: install.ps1 [-Global] [-Usage]"
    Write-Host "Options:"
    Write-Host "  -Global   : Install Compozify at the system level (requires administrator privileges)."
    Write-Host "  -Usage    : Show this usage message."
    exit
}
# Variables
$RepoOwner = "profclems"
$RepoName = "compozify"
$ReleaseTag = "" # We will fetch the actual release tag from the GitHub API response
$DownloadFileName = ""
$InstallDirUser = Join-Path -Path $env:USERPROFILE -ChildPath "Programs\compozify"
$InstallDirSystem = Join-Path -Path $env:ProgramFiles -ChildPath "compozify"
$ExtractedDirUser = Join-Path -Path $InstallDirUser -ChildPath "bin\"
$ExtractedDirSystem = Join-Path -Path $InstallDirSystem -ChildPath "bin\"

# Determine the architecture (amd64, arm64, or 386)
if ([Environment]::Is64BitOperatingSystem) {
    if ([System.Runtime.InteropServices.RuntimeInformation]::ProcessArchitecture -eq [System.Runtime.InteropServices.Architecture]::Arm64) {
        $Arch = "arm64"
    } else {
        $Arch = "amd64"
    }
} else {
    $Arch = "386"
}

# GitHub release API endpoint
$ReleaseUrl = "https://api.github.com/repos/$RepoOwner/$RepoName/releases/latest"

# Function to download a file using Invoke-WebRequest
function Download-File {
    param (
        [string]$Url,
        [string]$OutputPath
    )
    Invoke-WebRequest -Uri $Url -OutFile $OutputPath
}

# Function to check if the directory exists in the PATH environment variable
function DirectoryExistsInPath {
    param (
        [string]$Directory
    )
    $existingPaths = $Env:Path -split ";"
    return $existingPaths -contains $Directory
}

# Fetch the latest release information from GitHub API
try {
    Write-Host "Fetching the latest release information from GitHub API..."
    $releaseInfo = Invoke-RestMethod -Uri $ReleaseUrl
    $ReleaseTag = $releaseInfo.tag_name -replace '^v'
    $DownloadFileName = "compozify_${ReleaseTag}_windows_${Arch}.zip"
    Write-Host "Latest release tag: $ReleaseTag"
    Write-Host "Download file name: $DownloadFileName"
    Write-Host "Fetching the download URL for the release asset..."
    $AssetUrl = ($releaseInfo.assets | Where-Object { $_.name -eq $DownloadFileName }).browser_download_url
    Write-Host "Download URL: $AssetUrl"
} catch {
    Write-Host "Error: Failed to fetch release information from GitHub API. Installation failed."
    exit 1
}

# Select the appropriate installation directory based on the --global flag
if ($Global) {
    $InstallDir = $InstallDirSystem
    $ExtractedDir = $ExtractedDirSystem
} else {
    $InstallDir = $InstallDirUser
    $ExtractedDir = $ExtractedDirUser
}

# Create the installation directory if it doesn't exist
if (-Not (Test-Path -Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
}

# Download the release asset to the installation directory
$DownloadFilePath = Join-Path -Path $InstallDir -ChildPath $DownloadFileName
try {
    Download-File -Url $AssetUrl -OutputPath $DownloadFilePath
} catch {
    Write-Host "Error: Failed to download the release asset. Installation failed."
    exit 1
}

# Extract the downloaded zip file to the installation directory
Expand-Archive -Path $DownloadFilePath -DestinationPath $InstallDir -Force

# Clean up the downloaded zip file
Remove-Item -Path $DownloadFilePath -Force

# Check if the binary is inside /bin/compozify directory
if (-Not (Test-Path -Path $ExtractedDir)) {
    Write-Host "Error: The expected binary was not found in the zip file. Installation failed."
    exit 1
}

# Display installation completed message
Write-Host "Compozify has been installed to: $InstallDir"

# Add Compozify binary directory to the PATH environment variable if not already added
Write-Host "Adding Compozify binary directory to the PATH environment variable..."
if ($Global) {
    if (-Not (DirectoryExistsInPath $ExtractedDir)) {
        $existingPath = [System.Environment]::GetEnvironmentVariable("PATH", "Machine")
        $newPath = "$existingPath;$ExtractedDir"
        [System.Environment]::SetEnvironmentVariable("PATH", $newPath, "Machine")

        # Display path added message
        Write-Host "Compozify binary directory added to the PATH environment variable at the system level."
    } else {
        Write-Host "Compozify binary directory already exists in the PATH environment variable. Skipping update."
    }
} else {
    if (-Not (DirectoryExistsInPath $ExtractedDir)) {
        $existingPath = [System.Environment]::GetEnvironmentVariable("PATH", "User")
        $newPath = "$existingPath;$ExtractedDir"
        [System.Environment]::SetEnvironmentVariable("PATH", $newPath, "User")

        # Display path added message
        Write-Host "Compozify binary directory added to the PATH environment variable at the user level."
    } else {
        Write-Host "Compozify binary directory already exists in the PATH environment variable. Skipping update."
    }
}