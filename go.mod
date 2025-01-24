module main

go 1.19

require info/cpuInfo v0.0.0
replace info/cpuInfo => ./info/cpuInfo

require info/memInfo v0.0.0
replace info/memInfo => ./info/memInfo

require info/diskInfo v0.0.0
replace info/diskInfo => ./info/diskInfo

require info/netInfo v0.0.0
replace info/netInfo => ./info/netInfo