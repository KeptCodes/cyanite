param (
    [string]$exePath
)

# Check if the executable exists
if (Test-Path $exePath) {
    try {
        # Open registry key for startup
        $regKey = "HKCU:\Software\Microsoft\Windows\CurrentVersion\Run"
        Set-ItemProperty -Path $regKey -Name "Cyanite" -Value $exePath
        Write-Output "Successfully added Cyanite to startup."
    } catch {
        Write-Error "Failed to add Cyanite to startup. Error: $_"
    }
} else {
    Write-Error "Executable not found: $exePath"
}
