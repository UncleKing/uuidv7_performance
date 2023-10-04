# What is UUIDV7 and why do you care? 

Recently I read about UUID version7 and its promised advantages over UUID v4 as database indexes. A lot has been written over the internet about issues with v4 as database indexes and I'll refer the reader to some articles below to get some understanding of the problem. I wanted to do some quick benchmarking to see the actualadvantages of v7 over v4. This is the result. 

```
(base) ➜  mongo-uuid go test -run=^$ -bench=.  example.com/m
goos: darwin
goarch: arm64
pkg: example.com/m
BenchmarkFindByUUIDV7NewRecords-8               	   14504	     84301 ns/op
BenchmarkFindByUUIDV7OldRecords-8               	   13044	     91041 ns/op
BenchmarkFindByUUIDV7RandomRecords-8            	    9464	    115489 ns/op
BenchmarkFindByUUIDV4NewRecords-8               	   12764	    152528 ns/op
BenchmarkFindByUUIDV4OldRecords-8               	   12312	    122110 ns/op
BenchmarkFindByUUIDV4RandomRecords-8            	   13550	    162564 ns/op
BenchmarkFindByMongoIDNewRecords-8              	   14280	     76537 ns/op
BenchmarkFindByMongoIDOldRecords-8              	   11852	     85678 ns/op
BenchmarkFindByMongoIDRandomRecords-8           	   12841	    117687 ns/op
BenchmarkFindByNumericIDNewRecords-8            	   11510	    102057 ns/op
BenchmarkFindByNumericIDOldRecords-8            	   11336	     90367 ns/op
BenchmarkFindByNumericIDRandomRecords-8         	   12985	    108366 ns/op
BenchmarkFindByNumericStringIDNewRecords-8      	   12171	     91524 ns/op
BenchmarkFindByNumericIDStringOldRecords-8      	   11076	     96544 ns/op
BenchmarkFindByNumericIDStringRandomRecords-8   	   10578	    124484 ns/op
PASS
ok  	example.com/m	47.439s
```

## Setup
I am running everything locally on my 2020 M1 Mac to reduce noise from network variances. 
MongoDB v7.0.2 community edition running locally
go version is go1.21.1 darwin/arm64

Indexes created manually on numericid, numericidstr, uuidv4 and uuidv7



## How to run
### Create the json files
``` go run main.go ```
This creates 100M records in the mongo collection and generates 3 files to be used for benchmarking
### Run the tests!
``` go test -run=^$ -bench=. ```

## Results summary
v7 outperforms v4 every time, however the bump might not be enough for everyone to make the switch!
v7 gives even better results when checking newer records. Think your database after 5 years, Archiving older records periodically is better if your usecase allows for it. 
Native mongoIDs give similar or slightly better performance over v7, Not everyone prefers mongo IDs 
Using dual IDs(internal & external) is probably not required given benefits of v7 as sequential IDs thus reducing code complexity

## Next steps
I want to tryout other databases to understand the impact of v7 UUIDs and I'll update this with tests to Postgres and mySQL. Create issue if you have feedback!
YMMV