$ErrorActionPreference = 'Stop'

$exts = @('md','toml','yaml','yml','json','html','css','js')
$hits = New-Object System.Collections.Generic.List[string]

$all = Get-ChildItem -Recurse -File
foreach ($f in $all) {
    $ext = $f.Extension.TrimStart('.').ToLowerInvariant()
    if ($exts -notcontains $ext) { continue }

    $bytes = [System.IO.File]::ReadAllBytes($f.FullName)
    if ([Array]::IndexOf($bytes, 0) -ge 0) {
        $hits.Add("NULLBYTE $($f.FullName)")
    }
}

if ($hits.Count -eq 0) {
    Write-Output 'OK: keine Nullbytes in Textdateien gefunden.'
    exit 0
}

Write-Output 'Auffälligkeiten:'
$hits | ForEach-Object { Write-Output $_ }
exit 1
