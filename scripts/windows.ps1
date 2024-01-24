# Get the directory of the current PowerShell script
$currentScriptDirectory = $PSScriptRoot

# Construct the path to file.exe
$fileExePath = Join-Path -Path $currentScriptDirectory -ChildPath "..\bin\file.exe"

# Get the current user's PATH environment variable
$currentPath = [System.Environment]::GetEnvironmentVariable("PATH", [System.EnvironmentVariableTarget]::User)

# Check if the directory is already in the PATH
if ($currentPath -notlike "*$currentScriptDirectory*") {
    # Add the directory to the PATH
    [System.Environment]::SetEnvironmentVariable("PATH", "$currentPath;$currentScriptDirectory", [System.EnvironmentVariableTarget]::User)

    # Display a message
    Write-Host "Directory added to the user's PATH. Changes will take effect in new sessions."
} else {
    # Display a message if the directory is already in the PATH
    Write-Host "Directory is already in the user's PATH."
}

# Refresh the environment variables in the current PowerShell session
$env:PATH = [System.Environment]::GetEnvironmentVariable("PATH", [System.EnvironmentVariableTarget]::User)

# Display the full path to file.exe
Write-Host "Path to file.exe: $fileExePath"
