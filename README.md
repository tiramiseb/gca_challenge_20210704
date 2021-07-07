# GCA Challenge 04/07/2021

This is my participation to a coding challenge. This code is Open Source (MIT licence).

The goal was:

1. to write a function that returns a 2 dimensions table filled as in the following example table, with any given dimensions.
2. as a bonus, to write a method which returns the value corresponding to coordinates in the table.

|     |     |     |     |     |
| --- | --- | --- | --- | --- |
| 1   | 3   | 6   | 9   | 12  |
| 2   | 5   | 8   | 11  | 14  |
| 4   | 7   | 10  | 13  | 15  |

## My choice

I chose to not return the 2 dimensions table, but to focus on the second question, because answering it virtually also answers the first question.

I tried to write code as performant as possible, in Go. I chose not to rely on values from previous cells to find value of a specific cell.

I made 8 different tries, implemented in functions named `findValue`, `findValue2`, ..., `findValue8`.

Oh, by the way, I started my indexes at 1, oops... The instructions stated coordinates should start at `[0, 0]`. Well, I had fun nonetheless.

Well, after reading other comments on the challenge (I purposedly didn't read them before completing my participation), I learnt that the solution I used for findValue4 and successive answers is based on [triangular numbers](https://en.wikipedia.org/wiki/Triangular_number).

## Other implementations

I have written `findValueFJ` and `findValueDB`, which are implementations from other participants.

FJ has correctly implemented what has been requested (with index starting at `[0, 0]` and first value being `1`), DB had correctly used index `[0, 0]` but has started values at `0`.

## findValue

Loops that increment the value to add 1 + 2 + ..., etc. Not really optimized, but working, it is composed of two loops.

Benchmark result:

```plain
BenchmarkFindValue_1000x1000-4                         3         422101940 ns/op
BenchmarkFindValue_100x100-4                        1868            545759 ns/op
BenchmarkFindValue_20x20-4                        221552              4603 ns/op
BenchmarkFindValue_3x5-4                        17055450                63.12 ns/op
```

It takes 422ms to calculate all values for a 1000×1000 table, 5µs for a 20×20 table.

## findValue2

A very simple try to have only one loop. Really not convincing because the code in the second loop, that was executed only in some specific cases, is executed for every value.

Benchmark result:

```plain
BenchmarkFindValue2_1000x1000-4                        2         919721678 ns/op
BenchmarkFindValue2_100x100-4                        804           1388507 ns/op
BenchmarkFindValue2_20x20-4                       136218              7799 ns/op
BenchmarkFindValue2_3x5-4                       23145858                51.87 ns/op
```

It takes 920ms to calculate all values for a 1000×1000 table, 8µs for a 20×20 table.

## findValue3

Only one loop, but in a smart way!

Benchmark result:

```plain
BenchmarkFindValue3_1000x1000-4                        3         362642040 ns/op
BenchmarkFindValue3_100x100-4                       2655            512495 ns/op
BenchmarkFindValue3_20x20-4                       274640              4329 ns/op
BenchmarkFindValue3_3x5-4                       20750596                48.84 ns/op
```

It takes 363ms to calculate all values for a 1000×1000 table, 4µs for a 20×20 table.

Slightly faster on large tables, it's great.

## findValue4

Completely new approach: take the first row as the basis of calculations, add use offsets to get values for other rows, and use the fact that 1+2+...+n is equivalent to (n/2)×(n+1), completely avoiding loops.

Benchmark result:

```plain
BenchmarkFindValue4_1000x1000-4                      207           5599216 ns/op
BenchmarkFindValue4_100x100-4                      20478             57617 ns/op
BenchmarkFindValue4_20x20-4                       536888              2224 ns/op
BenchmarkFindValue4_3x5-4                       15338473                77.17 ns/op
```

It takes 6ms to calculate all values for a 1000×1000 table, 2µs for a 20×20 table.

Approximately 100 times quicker than the first try. Not bad, heh?

## findValue5

The same as findValue4, but not calling a function for the aforementioned operation. Fuck DRY, maybe we can be more performant!

Benchmark result:

```plain
BenchmarkFindValue5_1000x1000-4                     3553            334879 ns/op
BenchmarkFindValue5_100x100-4                     262231              4590 ns/op
BenchmarkFindValue5_20x20-4                      6480378               165.7 ns/op
BenchmarkFindValue5_3x5-4                       128176962                9.267 ns/op
```

It takes 335µs to calculate all values for a 1000×1000 table, 166ns for a 20×20 table.

More than 1000 times quicker than the first try! Yep, I knew calling functions is slow.

## findValue6 and findValue7

What if we use a cache for those values?

Benchmark result:

```plain
BenchmarkFindValue6_1000x1000-4                      244           4725972 ns/op
BenchmarkFindValue6_100x100-4                      25052             46808 ns/op
BenchmarkFindValue6_20x20-4                       556388              2146 ns/op
BenchmarkFindValue6_3x5-4                       17464030                69.17 ns/op
BenchmarkFindValue7_1000x1000-4                      261           4346229 ns/op
BenchmarkFindValue7_100x100-4                      24811             46901 ns/op
BenchmarkFindValue7_20x20-4                       552018              2156 ns/op
BenchmarkFindValue7_3x5-4                       17447212                68.86 ns/op
```

Yep, that's what I thought: storing values and interrogating this cache is way slower than recalculating every time.

# findValue8

We can simplify "(n/2)×(n+1)" in "(n×(n+1))/2", which avoids conversions to float.

Benchmark result:

```plain
BenchmarkFindValue8_1000x1000-4                     3555            335002 ns/op
BenchmarkFindValue8_100x100-4                     297558              3999 ns/op
BenchmarkFindValue8_20x20-4                      6660996               166.0 ns/op
BenchmarkFindValue8_3x5-4                       120974073                9.793 ns/op
```

It takes 335µs to calculate all values for a 1000×1000 table, 166ns for a 20×20 table.

The result is more or less the same as with findValue5, which is probably explained by the performance of the Go compiler...

How much can we compute in a reasonable amount of time?

Benchmark result:

```plain
BenchmarkFindValue8_10000x10000-4                     34          32972852 ns/op
BenchmarkFindValue8_100000x100000-4                    1        3327526386 ns/op
```

100 million values take 33ms, 10 billion values take 3secs. Looks reasonable to me...

## findValueFJ

Another participant, FJ, has suggested a very neat and short solution, initially written in C++. After reading it carefully, it looks very similar to `findvalue8`. It is here for performance comparison. If you want to use it with the compiled executables, simply add the `-fj` flag.

Benchmark result:

```plain
BenchmarkFindValueFJ_1000x1000-4                    3532            336833 ns/op
BenchmarkFindValueFJ_100x100-4                    258498              4601 ns/op
BenchmarkFindValueFJ_20x20-4                     6646764               166.2 ns/op
BenchmarkFindValueFJ_3x5-4                      131364908                8.905 ns/op
BenchmarkFindValueFJ_10000x10000-4                    34          33028187 ns/op
BenchmarkFindValueFJ_100000x100000-4                   1        3315161415 ns/op
```

It takes 337µs to calculate all values for a 1000×1000 table, 166ns for a 20×20 table. 100 million values take 33ms, 10 billion values take 3secs.

Well, I thought this solution would be better, but I realize it is comparable.

## findValueDB

Another participant, DB, has suggested another solution, initially written in Python. It is here for performance comparison.

Benchmark result:

```plain
BenchmarkFindValueDB_1000x1000-4                     151           7739120 ns/op
BenchmarkFindValueDB_100x100-4                     14912             80003 ns/op
BenchmarkFindValueDB_20x20-4                      334302              3549 ns/op
BenchmarkFindValueDB_3x5-4                       8814028               125.2 ns/op
BenchmarkFindValueDB_10000x10000-4                     2         784594814 ns/op
BenchmarkFindValueDB_100000x100000-4                   1        77396716443 ns/op
```

It takes 8ms to calculate all values for a 1000×1000 table, 4µs for a 20×20 table. 100 million values take 785ms, 10 billion values take 77s.

## findValueDBoptim

I tried to optimize DB's answer, removing the function call and avoiding variables initialization when they are unneeded, while keeping the variables, in order to avoid getting too far from the implementation idea.

Benchmark result:

```plain
BenchmarkFindValueDBoptim_1000x1000-4                164           7188232 ns/op
BenchmarkFindValueDBoptim_100x100-4                15918             74141 ns/op
BenchmarkFindValueDBoptim_20x20-4                 352962              3375 ns/op
BenchmarkFindValueDBoptim_3x5-4                  8932149               121.5 ns/op
BenchmarkFindValueDBoptim_10000x10000-4                2         722839825 ns/op
BenchmarkFindValueDBoptim_100000x100000-4              1        71728461858 ns/op
```

It takes 7ms to calculate all values for a 1000×1000 table, 3µs for a 20×20 table. 100 million values take 723ms, 10 billion values take 72s. It is slightly better than original suggestion, but not by as much as I imagined...

## findValueDBoptim2

After talking with DB, we tried to improve his implementation, which results in a huge improvement. However, it is still slower than FJ's and mine, sorry bro :sweat_smile:. If you want to use it with the compiled executables, simply add the `-db` flag.

Benchmark result:

```plain
BenchmarkFindValueDBoptim2_1000x1000-4               324           3457180 ns/op
BenchmarkFindValueDBoptim2_100x100-4               31789             36537 ns/op
BenchmarkFindValueDBoptim2_20x20-4                715617              1655 ns/op
BenchmarkFindValueDBoptim2_3x5-4                20649829                57.84 ns/op
BenchmarkFindValueDBoptim2_10000x10000-4               3         339722170 ns/op
BenchmarkFindValueDBoptim2_100000x100000-4             1        36047289752 ns/op
```

## The whole benchmark values

You love tables? Here they are!

The comparison table:

| impl.             |    1000×1000 |      100×100 |        20×20 |          3×5 |   10000×10000 | 100000×100000 |
| ----------------- | -----------: | -----------: | -----------: | -----------: | ------------: | ------------: |
| findValue         | 422 101,9 µs |   545 759 ns |   4 603,0 ns |    63,120 ns |             / |             / |
| findValue2        | 919 721,7 µs | 1 388 507 ns |   7 799,0 ns |    51,870 ns |             / |             / |
| findValue3        | 362 642,0 µs |   512 495 ns |   4 329,0 ns |    48,840 ns |             / |             / |
| findValue4        |   5 599,2 µs |    57 617 ns |   2 224,0 ns |    77,170 ns |             / |             / |
| findValue5        | **334,8 µs** |     4 590 ns | **165,7 ns** |     9,267 ns |             / |             / |
| findValue6        |   4 726,0 µs |    46 808 ns |   2 146,0 ns |    69,170 ns |             / |             / |
| findValue7        |   4 346,2 µs |    46 901 ns |   2 156,0 ns |    68,860 ns |             / |             / |
| findValue8        | **335,0 µs** | **3 999 ns** | **166,0 ns** |     9,793 ns | **32,972 ms** |   **3,327 s** |
| findValueFJ       | **336,8 µs** |     4 601 ns | **166,2 ns** | **8,905 ns** | **33,028 ms** |   **3,315 s** |
| findValueDB       |   7 739,1 µs |    80 003 ns |   3 549,0 ns |   125,200 ns |    784,594 ms |      77,396 s |
| findValueDBoptim  |   7 188,2 µs |    74 141 ns |   3 375,0 ns |   121,500 ns |    722,839 ms |      71,728 s |
| findValueDBoptim2 |   3 457,2 µs |    36 537 ns |   1 655,0 ns |    57,840 ns |    339,722 ms |      36,047 s |

Here are some lessons I can think of:

- Go does not make a difference between bitwise shifting and division by 2 (some other languages treat them differently and can make the shifting a single CPU operation, whereas a division is more complex)
- with an efficient compiler (like Go), there is no point in spending days seeking to improve an algorithm that is already short and simple (let's say "mostly maths" simple)
- mixing two efficient algorithms may make the mix slower
- loops are expensive
- storing multiple variables is expensive
- function calls are expensive

Regarding optimization, it is something that I learnt long ago, but it's always good to remember: try to improve input/outputs and data storage first, because it is way more expensive than your algorithms.

The benchmarking command complete return values:

```plain
goos: linux
goarch: amd64
pkg: github.com/tiramiseb/gca_challenge_20210704
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
BenchmarkFindValue_1000x1000-4                         3         422101940 ns/op
BenchmarkFindValue_100x100-4                        1868            545759 ns/op
BenchmarkFindValue_20x20-4                        221552              4603 ns/op
BenchmarkFindValue_3x5-4                        17055450                63.12 ns/op
BenchmarkFindValue2_1000x1000-4                        2         919721678 ns/op
BenchmarkFindValue2_100x100-4                        804           1388507 ns/op
BenchmarkFindValue2_20x20-4                       136218              7799 ns/op
BenchmarkFindValue2_3x5-4                       23145858                51.87 ns/op
BenchmarkFindValue3_1000x1000-4                        3         362642040 ns/op
BenchmarkFindValue3_100x100-4                       2655            512495 ns/op
BenchmarkFindValue3_20x20-4                       274640              4329 ns/op
BenchmarkFindValue3_3x5-4                       20750596                48.84 ns/op
BenchmarkFindValue4_1000x1000-4                      207           5599216 ns/op
BenchmarkFindValue4_100x100-4                      20478             57617 ns/op
BenchmarkFindValue4_20x20-4                       536888              2224 ns/op
BenchmarkFindValue4_3x5-4                       15338473                77.17 ns/op
BenchmarkFindValue5_1000x1000-4                     3553            334879 ns/op
BenchmarkFindValue5_100x100-4                     262231              4590 ns/op
BenchmarkFindValue5_20x20-4                      6480378               165.7 ns/op
BenchmarkFindValue5_3x5-4                       128176962                9.267 ns/op
BenchmarkFindValue6_1000x1000-4                      244           4725972 ns/op
BenchmarkFindValue6_100x100-4                      25052             46808 ns/op
BenchmarkFindValue6_20x20-4                       556388              2146 ns/op
BenchmarkFindValue6_3x5-4                       17464030                69.17 ns/op
BenchmarkFindValue7_1000x1000-4                      261           4346229 ns/op
BenchmarkFindValue7_100x100-4                      24811             46901 ns/op
BenchmarkFindValue7_20x20-4                       552018              2156 ns/op
BenchmarkFindValue7_3x5-4                       17447212                68.86 ns/op
BenchmarkFindValue8_1000x1000-4                     3555            335002 ns/op
BenchmarkFindValue8_100x100-4                     297558              3999 ns/op
BenchmarkFindValue8_20x20-4                      6660996               166.0 ns/op
BenchmarkFindValue8_3x5-4                       120974073                9.793 ns/op
BenchmarkFindValue8_10000x10000-4                     34          32972852 ns/op
BenchmarkFindValue8_100000x100000-4                    1        3327526386 ns/op
BenchmarkFindValueFJ_1000x1000-4                    3532            336833 ns/op
BenchmarkFindValueFJ_100x100-4                    258498              4601 ns/op
BenchmarkFindValueFJ_20x20-4                     6646764               166.2 ns/op
BenchmarkFindValueFJ_3x5-4                      131364908                8.905 ns/op
BenchmarkFindValueFJ_10000x10000-4                    34          33028187 ns/op
BenchmarkFindValueFJ_100000x100000-4                   1        3315161415 ns/op
BenchmarkFindValueDB_1000x1000-4                     151           7739120 ns/op
BenchmarkFindValueDB_100x100-4                     14912             80003 ns/op
BenchmarkFindValueDB_20x20-4                      334302              3549 ns/op
BenchmarkFindValueDB_3x5-4                       8814028               125.2 ns/op
BenchmarkFindValueDB_10000x10000-4                     2         784594814 ns/op
BenchmarkFindValueDB_100000x100000-4                   1        77396716443 ns/op
BenchmarkFindValueDBoptim_1000x1000-4                164           7188232 ns/op
BenchmarkFindValueDBoptim_100x100-4                15918             74141 ns/op
BenchmarkFindValueDBoptim_20x20-4                 352962              3375 ns/op
BenchmarkFindValueDBoptim_3x5-4                  8932149               121.5 ns/op
BenchmarkFindValueDBoptim_10000x10000-4                2         722839825 ns/op
BenchmarkFindValueDBoptim_100000x100000-4              1        71728461858 ns/op
BenchmarkFindValueDBoptim2_1000x1000-4               324           3457180 ns/op
BenchmarkFindValueDBoptim2_100x100-4               31789             36537 ns/op
BenchmarkFindValueDBoptim2_20x20-4                715617              1655 ns/op
BenchmarkFindValueDBoptim2_3x5-4                20649829                57.84 ns/op
BenchmarkFindValueDBoptim2_10000x10000-4               3         339722170 ns/op
BenchmarkFindValueDBoptim2_100000x100000-4             1        36047289752 ns/op
```

## Compiled

See compiled executables in the `compiled/` subdirectory.
