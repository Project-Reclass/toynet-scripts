$Folder = '~\documents\usb-deployment'
"Test to see if folder [$Folder]  exists"
if (Test-Path -Path $Folder) {
    "Path exists!"
} else {
    New-Item -Path "~\documents\" -Name "usb-deployment" -ItemType "directory"
    "Path doesn't exist."
}
Copy-Item -Path ".\dependencies" -Destination "$Folder" -Recurse -ErrorAction SilentlyContinue
Copy-Item -Path ".\toynet" -Destination "$Folder" -Recurse -ErrorAction SilentlyContinue