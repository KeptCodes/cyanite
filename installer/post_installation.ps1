param (
    [string]$exePath
)

# Check if the executable exists
if (Test-Path $exePath) {
    try {
        # Open registry key for startup
        $regKey = "HKCU:\Software\Microsoft\Windows\CurrentVersion\Run"
        Set-ItemProperty -Path $regKey -Name "Syra" -Value $exePath
        Write-Output "Successfully added Syra to startup."
    } catch {
        Write-Error "Failed to add Syra to startup. Error: $_"
    }
} else {
    Write-Error "Executable not found: $exePath"
}
