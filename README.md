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

I made 8 different tries, implemented in functions named "findValue", "findValue2", ..., "findValue8".

Oh, by the way, I started my indexes at 1, oops... The instructions stated coordinates should start at [0, 0]. Well, I had fun nonetheless.

## findValue

Loops that increment the value to add 1 + 2 + ..., etc. Not really optimized, but working, it is composed of two loops.

Benchmark result:

```plain
BenchmarkFindValue_1000x1000-4    	       3	 423804338 ns/op
BenchmarkFindValue_100x100-4      	    2001	    596626 ns/op
BenchmarkFindValue_20x20-4        	  276559	      4186 ns/op
BenchmarkFindValue_3x5-4          	18369159	        65.36 ns/op
```

It takes 424ms to calculate all values for a 1000×1000 table, 4 µs for a 20×20 table.

## findValue2

A very simple try to have only one loop. Really not convincing because the code in the second loop, that was executed only in some specific cases, is executed for every value.

Benchmark result:

```plain
BenchmarkFindValue2_1000x1000-4   	       2	 921937489 ns/op
BenchmarkFindValue2_100x100-4     	     844	   1412143 ns/op
BenchmarkFindValue2_20x20-4       	  145068	      8166 ns/op
BenchmarkFindValue2_3x5-4         	19991491	        58.26 ns/op
```

It takes 922ms to calculate all values for a 1000×1000 table, 4 µs for a 20×20 table.

## findValue3

Only one loop, but in a smart way!

Benchmark result:

```plain
BenchmarkFindValue3_1000x1000-4   	       3	 365511193 ns/op
BenchmarkFindValue3_100x100-4     	    2486	    488636 ns/op
BenchmarkFindValue3_20x20-4       	  240256	      4204 ns/op
BenchmarkFindValue3_3x5-4         	28127276	        42.03 ns/op
```

It takes 366ms to calculate all values for a 1000×1000 table, 4 µs for a 20×20 table.

Slightly faster on large tables, it's great.

## findValue4

Completely new approach: take the first row as the basis of calculations, add use offsets to get values for other rows, and use the fact that 1+2+...+n is equivalent to (n/2)×(n+1), completely avoiding loops.

Benchmark result:

```plain
BenchmarkFindValue4_1000x1000-4   	     214	   5590552 ns/op
BenchmarkFindValue4_100x100-4     	   21198	     56512 ns/op
BenchmarkFindValue4_20x20-4       	  519321	      2240 ns/op
BenchmarkFindValue4_3x5-4         	15263586	        77.57 ns/op
```

It takes 6ms to calculate all values for a 1000×1000 table, 2 µs for a 20×20 table.

Approximately 100 times quicker than the first try. Not bad, heh?

## findValue5

The same as findValue4, but not calling a function for the aforementioned operation. Fuck DRY, maybe we can be more performant!

Benchmark result:

```plain
BenchmarkFindValue5_1000x1000-4   	    3308	    344641 ns/op
BenchmarkFindValue5_100x100-4     	  260847	      4362 ns/op
BenchmarkFindValue5_20x20-4       	 5442354	       189.0 ns/op
BenchmarkFindValue5_3x5-4         	100000000	        10.49 ns/op
```

It takes 345µs to calculate all values for a 1000×1000 table, 189 ns for a 20×20 table.

More than 1000 times quicker than the first try! Yep, I knew calling functions is slow.

## findValue6 and findValue7

What if we use a cache for those values?

Benchmark result:

```plain
BenchmarkFindValue6_1000x1000-4   	     265	   4394929 ns/op
BenchmarkFindValue6_100x100-4     	   26583	     45017 ns/op
BenchmarkFindValue6_20x20-4       	  537818	      2159 ns/op
BenchmarkFindValue6_3x5-4         	18255672	        64.23 ns/op
BenchmarkFindValue7_1000x1000-4   	     277	   4310937 ns/op
BenchmarkFindValue7_100x100-4     	   26578	     45012 ns/op
BenchmarkFindValue7_20x20-4       	  527995	      2208 ns/op
BenchmarkFindValue7_3x5-4         	18380082	        64.15 ns/op
```

Yep, that's what I thought: storing values and interrogating this cache is way slower than recalculating every time.

# findValue8

We can simplify "(n/2)×(n+1)" in "(n×(n+1))/2", which avoids conversions to float.

Benchmark result:

```plain
BenchmarkFindValue8_1000x1000-4   	    3346	    345541 ns/op
BenchmarkFindValue8_100x100-4     	  255201	      4647 ns/op
BenchmarkFindValue8_20x20-4       	 6939322	       171.8 ns/op
BenchmarkFindValue8_3x5-4         	125461849	         9.509 ns/op
```

It takes 346µs to calculate all values for a 1000×1000 table, 171 ns for a 20×20 table.

The result is more or less the same as with findValue5, which is probably explained by the performance of the Go compiler...

How much can we compute in a reasonable amount of time?

Benchmark result:

```plain
BenchmarkFindValue8_10000x10000-4   	      33	  33543685 ns/op
BenchmarkFindValue8_100000x100000-4   	       1	3652251790 ns/op
```

100 million values take 33ms, 10 billion values take 3,7secs. Looks reasonable to me...

# findValueFJ

Another participant, FJ, has suggested a very neat and short solution. Honestly, I don't understand it, but well, I haven't tried much. It is here only for performance comparison. If you want to use it with the compiled executables, simply add the `-fj` flag.

Benchmark result:

```plain
BenchmarkFindValueFJ_1000x1000-4       	    3396	    346946 ns/op
BenchmarkFindValueFJ_100x100-4         	  255729	      4591 ns/op
BenchmarkFindValueFJ_20x20-4           	 7260736	       164.8 ns/op
BenchmarkFindValueFJ_3x5-4             	133179514	         8.939 ns/op
BenchmarkFindValueFJ_10000x10000-4     	      33	  32847506 ns/op
BenchmarkFindValueFJ_100000x100000-4   	       1	3351124175 ns/op
```

It takes 347µs to calculate all values for a 1000×1000 table, 164 ns for a 20×20 table. 100 million values take 33ms, 10 billion values take 3,4secs.
Well, I thought this solution would be better, but I realize it is comparable.

## Compiled

See compiled executables in the `compiled/` subdirectory.
