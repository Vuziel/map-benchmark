# Select

|           |           map |   map + mutex |   map + rwMutex |        sync.Map |
|----------:|--------------:|--------------:|----------------:|----------------:|
| 1 000 000 | 535 714 ns/op | 646 762 ns/op |  637 297  ns/op | 1 564 924 ns/op |
|   100 000 |  17 125 ns/op |  22 862 ns/op |   21 352  ns/op |   110 962 ns/op |
|    10 000 |    2020 ns/op |    2451 ns/op |     6106  ns/op |    11 818 ns/op |
|      1000 |     156 ns/op |     243 ns/op |      225  ns/op |     880.0 ns/op |
|       100 |      34 ns/op |      34 ns/op |       38  ns/op |        82 ns/op |

# Insert

|           |           map |  map + mutex  | map + rwMutex |        sync.Map |
|----------:|--------------:|:-------------:|--------------:|----------------:|
| 1 000 000 | 772 332 ns/op | 763 247 ns/op | 886912  ns/op | 5 045 782 ns/op |
|   100 000 |  38 578 ns/op | 34 024 ns/op  | 60 433  ns/op |   349 506 ns/op |
|    10 000 |    3317 ns/op |  3153 ns/op   |   4149  ns/op |     30808 ns/op |
|      1000 |     628 ns/op |   247 ns/op   |  319.2  ns/op |      1810 ns/op |
|       100 |      29 ns/op |   71 ns/op    |  43.30  ns/op |     219.3 ns/op |

# Delete

|           |     map      | map + mutex  | map + rwMutex |   sync.Map    |
|:---------:|:------------:|:------------:|:--------------|:-------------:|
| 1_000_000 | 691945 ns/op | 763247 ns/op | 886912  ns/op | 1708265 ns/op |
|  100_000  | 40300 ns/op  | 34024 ns/op  | 35230  ns/op  | 155333 ns/op  |
|  10_000   |  2083 ns/op  |  3153 ns/op  | 3811    ns/op |  25738 ns/op  |
|   1000    | 190.7 ns/op  | 247.2 ns/op  | 403.7   ns/op |  1048 ns/op   |
|    100    | 38.42 ns/op  | 70.54 ns/op  | 47.50   ns/op |  172.5 ns/op  |