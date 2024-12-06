<#
    .DESCRIPTION
    This script creates a new folder for the day of the Advent of Code and creates the necessary files to start the day.
    .INPUTS
    -year: The year of the Advent of Code. Default is the current year.
    -day: The day of the Advent of Code. Default is the next day.
    .OUTPUTS
    A new folder with the name of the day and the necessary files to start the day.
    .LINK
    https://adventofcode.com/
#>

[CmdletBinding()]
param (
    [Parameter()] [int] $year = (Get-Date).Year,
    [Parameter()] [int] $day = (Get-ChildItem -Path $year -Directory -Filter "day*").Count + 1
)

$dayPath = "$year\day$day"
$dayPage = Invoke-WebRequest -Uri "https://adventofcode.com/$year/day/$day"
$dayContent = $dayPage.Content -replace "(?s).*<main>(.*)</main>.*", '$1'

New-Item -ItemType Directory -Path $dayPath -Force
New-Item -ItemType Directory -Path "$dayPath\assets" -Force
New-Item -ItemType File -Path "$dayPath\part_one.md" -Force -Value $dayContent
New-Item -ItemType File -Path "$dayPath\part_two.md" -Force -Value $dayContent
New-Item -ItemType File -Path "$dayPath\README.md" -Force -Value "
# AOC DAY $day

---

![img]()
> Difficulty [x/x] :)  


### Take a look to the problems and my results:
- [Part 1](./part_one.md)
- [Part 2](./part_two.md)
"
