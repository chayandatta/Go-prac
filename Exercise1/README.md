How this was made

```
❯ cd Exercise1
❯ go mod init Exercise1
go: creating new go.mod: module Exercise1
go: to add module requirements and sums:
        go mod tidy
❯ go mod tidy
❯ go run main.go
❯ go build .
❯ go run main.go
```
How to run this file
```
❯ go run main.go -h
Usage of /var/folders/yj/vwqkxq2d26v76xnxvfr77b3m0000gn/T/go-build1493658496/b001/exe/main:
  -csv string
        read the csv for question & answers (default "problems.csv")
```
or we can do this also
```
❯ go run main.go --help
Usage of /var/folders/yj/vwqkxq2d26v76xnxvfr77b3m0000gn/T/go-build4291790095/b001/exe/main:
  -csv string
        read the csv for question & answers (default "problems.csv")
```

```
❯ go run main.go -csv=abc.csv
Failed to open the file abc.csv 
exit status 1
❯ go run main.go -csv=problems.csv
```

```
❯ go run main.go -csv="abc c.csv"
Failed to open the file abc c.csv 
exit status 1
❯ go run main.go -csv="problems.csv"

```

```
❯ go build . && go run main.go -csv="problems.csv"
[[5+5 10] [1+1 2] [8+3 11] [1+2 3] [8+6 14] [3+1 4] [1+4 5] [5+1 6] [2+3 5] [3+3 6] [2+4 6] [5+2 7]]
```

```
❯ go build . && go run main.go -csv="problems.csv"
[{5+5 10} {1+1 2} {8+3 11} {1+2 3} {8+6 14} {3+1 4} {1+4 5} {5+1 6} {2+3 5} {3+3 6} {2+4 6} {5+2 7}]
```

```
❯ go build . && go run main.go -csv="problems.csv"
Problem no #1: 5+5
Problem no #2: 1+1
Problem no #3: 8+3
Problem no #4: 1+2
Problem no #5: 8+6
Problem no #6: 3+1
Problem no #7: 1+4
Problem no #8: 5+1
Problem no #9: 2+3
Problem no #10: 3+3
Problem no #11: 2+4
Problem no #12: 5+2
```

```
❯ go build . && go run main.go -csv="problems.csv"
Problem no #1: 5+5 = 
10
Correct Answer!
Problem no #2: 1+1 = 
2
Correct Answer!
Problem no #3: 8+3 = 
11
Correct Answer!
Problem no #4: 1+2 = 
3
Correct Answer!
Problem no #5: 8+6 = 
14
Correct Answer!
Problem no #6: 3+1 = 
4
Correct Answer!
Problem no #7: 1+4 = 
5
Correct Answer!
Problem no #8: 5+1 = 
6
Correct Answer!
Problem no #9: 2+3 = 
5
Correct Answer!
Problem no #10: 3+3 = 
6
Correct Answer!
Problem no #11: 2+4 = 
6
Correct Answer!
Problem no #12: 5+2 = 
7
Correct Answer!
```

```
❯ go build . && go run main.go -csv="problems.csv"
Problem no #1: 5+5 = 
10
Problem no #2: 1+1 = 
2
Problem no #3: 8+3 = 
11
Problem no #4: 1+2 = 
3
Problem no #5: 8+6 = 
14
Problem no #6: 3+1 = 
4
Problem no #7: 1+4 = 
5
Problem no #8: 5+1 = 
6
Problem no #9: 2+3 = 
5
Problem no #10: 3+3 = 
6
Problem no #11: 2+4 = 
6
Problem no #12: 5+2 = 
7
You have scored 12 out of 12
```

```
❯ go build . && go run main.go -csv="problems.csv" -limit=2
Problem no #1: 5+5 = 
10
Problem no #2: 1+1 = 
2
Time is over!!
You have scored 2 out of 12 
```