# In that project you can visualize your local git contributions by Golang

### A CLI command that generates a graph similar to:
>go run main.go scan.go stats.go  -add "Folder Dir" -email "Your email"


### How To Run Project?
1. create the `.gogitlocalstats` in home directory
2. >go run main.go scan.go stats.go  -add "Folder Dir" -email "Your email"
   * After running this command `scan.go` responsible to scan your directory and put the path in `.gogitlocalstats`
3. Run >go run main.go scan.go stats.go


### Package for connect to git:
> go get "github.com/go-git/go-git/v5"



### Your Output lookalike this:
![](https://raw.githubusercontent.com/YoussefAshraf397/Track-Git-Locally-Golang/main/images/Screenshot%20from%202022-05-18%2012-11-54.png)


## Get all commits in the last six months only, but you can modify it:
```Language
const outOfRange = 99999
const daysInLastSixMonths = 183 // This is the last day  in six months
const weeksInLastSixMonths = 26 // This is the number of the last week in the last month 183/7
```






