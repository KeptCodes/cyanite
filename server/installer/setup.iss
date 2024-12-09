[Setup]
AppName=Cyanite
AppVersion={#AppVersion}
DefaultDirName={pf}\Cyanite
DefaultGroupName=Cyanite
OutputDir=.\Output
OutputBaseFilename=Cyanite-Win64-Setup
Compression=lzma
SolidCompression=yes
DisableProgramGroupPage=yes

[Files]
; The main executable file
Source: "cyanite.exe"; DestDir: "{app}"; Flags: ignoreversion

; The PowerShell script to add Cyanite to startup
Source: "post_installation.ps1"; DestDir: "{app}"; Flags: ignoreversion

[Run]
; Run the PowerShell script after installation to add Cyanite to startup
Filename: "powershell.exe"; Parameters: "-ExecutionPolicy Bypass -File ""{app}\post_installation.ps1"" -exePath ""{app}\cyanite.exe"""; StatusMsg: "Adding Cyanite to startup..."; Flags: runhidden

[Icons]
Name: "{group}\Cyanite"; Filename: "{app}\cyanite.exe"
