[Setup]
AppName=Syra
AppVersion={#AppVersion}
DefaultDirName={pf}\Syra
DefaultGroupName=Syra
OutputDir=.\Output
OutputBaseFilename=Syra-Win64-Setup
Compression=lzma
SolidCompression=yes
DisableProgramGroupPage=yes

[Files]
; The main executable file
Source: "syra.exe"; DestDir: "{app}"; Flags: ignoreversion

; The PowerShell script to add Syra to startup
Source: "post_installation.ps1"; DestDir: "{app}"; Flags: ignoreversion

[Run]
; Run the PowerShell script after installation to add Syra to startup
Filename: "powershell.exe"; Parameters: "-ExecutionPolicy Bypass -File ""{app}\post_installation.ps1"" -exePath ""{app}\syra.exe"""; StatusMsg: "Adding Syra to startup..."; Flags: runhidden

[Icons]
Name: "{group}\Syra"; Filename: "{app}\syra.exe"
