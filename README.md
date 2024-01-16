# Select

|           |      map      |  map + mutex  | map + rwMutex  |    sync.Map     |
|:---------:|:-------------:|:-------------:|:--------------:|:---------------:|
| 1 000 000 | 535 714 ns/op | 646 762 ns/op | 637 297  ns/op | 1 167 459 ns/op |
|  100 000  | 17 125 ns/op  | 22 862 ns/op  | 21 352  ns/op  |  54 273 ns/op   |
|  10 000   |  2020 ns/op   |  2451 ns/op   |  6106  ns/op   |  11 818 ns/op   |
|   1000    |   156 ns/op   |   243 ns/op   |   225  ns/op   |    361 ns/op    |
|    100    |   34 ns/op    |   34 ns/op    |   38  ns/op    |    67 ns/op     |

# Insert

|           |      map      |  map + mutex  | map + rwMutex |    sync.Map     |
|----------:|:-------------:|:-------------:|:-------------:|:---------------:|
| 1 000 000 | 918 498 ns/op | 763 247 ns/op | 886912  ns/op | 1 996 260 ns/op |
|   100 000 | 43 260 ns/op  | 34 024 ns/op  | 60 433  ns/op |  211 956 ns/op  |
|    10 000 |  2456 ns/op   |  3153 ns/op   |  4149  ns/op  |  17 128 ns/op   |
|      1000 |   870 ns/op   |   247 ns/op   |  319  ns/op   |   1233 ns/op    |
|       100 |   80 ns/op    |   71 ns/op    |   43  ns/op   |    226 ns/op    |

# Delete

|           |     map      |  map + mutex  | map + rwMutex  |    sync.Map     |
|:---------:|:------------:|:-------------:|:---------------|:---------------:|
| 1 000 000 | 52 905 ns/op | 763 247 ns/op | 886 912  ns/op | 1 081 755 ns/op |
|  100 000  |  4018 ns/op  | 34 024 ns/op  | 35 230  ns/op  |  35 687 ns/op   |
|  10 000   |  329 ns/op   |  3153 ns/op   | 3811    ns/op  |   4287 ns/op    |
|   1000    |   57 ns/op   |  247.2 ns/op  | 403.7   ns/op  |   2664 ns/op    |
|    100    |   14 ns/op   |  70.54 ns/op  | 47.50   ns/op  |    119 ns/op    |