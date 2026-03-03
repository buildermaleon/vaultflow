$r = "vaultflow"
$dir = "$env:LOCALAPPDATA\Programs"
New-Item -ItemType Directory -Path "$dir" -Force | Out-Null
Invoke-WebRequest -Uri "https://github.com/dablon/$r/releases/download/v1.0.0/vaultflow-windows-amd64.exe" -OutFile "$dir\$r.exe"
