$Folder = '~\documents\usb-deployment'
Get-ChildItem * -Path $Folder  -Recurse | Remove-Item