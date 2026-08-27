# Get changed files
$changes = git status --short

if (-not $changes) {
    Write-Host "No changes to commit."
    exit
}

Write-Host ""
Write-Host "Changed files:"
$changes | ForEach-Object {
    Write-Host "  $_"
}

Write-Host ""
Write-Host "What did you change?"
$description = Read-Host "Description"

Write-Host ""
Write-Host "Suggested commit message:"
Write-Host "feat: $description"
Write-Host ""

$confirm = Read-Host "Use this message? (y/n)"

if ($confirm -eq "y") {
    git add .
    git commit -m "feat: $description"
    
    Write-Host ""
    Write-Host "Commit created successfully."
}
else {
    Write-Host "Commit cancelled."
}