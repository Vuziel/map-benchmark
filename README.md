go test -run ^$ -bench=. -benchtime=100x

* 12 threads

# Select

|           |      map      |  map + mutex  | map + rwMutex |    sync.Map     |   hashmap    |
|:---------:|:-------------:|:-------------:|:-------------:|:---------------:|:------------:|
| 1 000 000 | 535 714 ns/op | 756 899 ns/op | 687 297 ns/op | 1 167 459 ns/op |              |
|  100 000  | 17 125 ns/op  | 22 862 ns/op  | 21 352 ns/op  |  54 273 ns/op   | 34 861 ns/op |
|  10 000   |  2020 ns/op   |  4336 ns/op   |  3160 ns/op   |  11 818 ns/op   |  1921 ns/op  |
|   1000    |   156 ns/op   |   243 ns/op   |   301 ns/op   |    361 ns/op    |  400 ns/op   |
|    100    |   34 ns/op    |   34 ns/op    |   38 ns/op    |    67 ns/op     |   65 ns/op   |

# Insert

|           |      map      |  map + mutex  | map + rwMutex  |    sync.Map     |   hashmap    |
|----------:|:-------------:|:-------------:|:--------------:|:---------------:|:------------:|
| 1 000 000 | 918 498 ns/op | 754 723 ns/op | 838 765  ns/op | 1 996 260 ns/op |              |
|   100 000 | 43 260 ns/op  | 36 024 ns/op  | 29 526  ns/op  |  147 351 ns/op  | 69 183 ns/op |
|    10 000 |  2456 ns/op   |  2757 ns/op   |  2608  ns/op   |  17 018 ns/op   |  3211 ns/op  |
|      1000 |   870 ns/op   |   253 ns/op   |   319  ns/op   |   1196 ns/op    |  1074 ns/op  |
|       100 |   80 ns/op    |   37 ns/op    |   43  ns/op    |    105 ns/op    |  115 ns/op   |

# Delete

|           |     map      | map + mutex  | map + rwMutex  |    sync.Map     |   hashmap    |
|:---------:|:------------:|:------------:|:---------------|:---------------:|:------------:|
| 1 000 000 | 52 905 ns/op | 92 331 ns/op | 139 578  ns/op | 1 081 755 ns/op |              |
|  100 000  |  4018 ns/op  |  8572 ns/op  | 9445  ns/op    |  35 687 ns/op   | 15 928 ns/op |
|  10 000   |  329 ns/op   |  1300 ns/op  | 1643    ns/op  |   4287 ns/op    |  973 ns/op   |
|   1000    |   57 ns/op   |   95 ns/op   | 158   ns/op    |   2664 ns/op    |   90 ns/op   |
|    100    |   14 ns/op   |   24 ns/op   | 29   ns/op     |    119 ns/op    |   38 ns/op   |