# timerounding

Rounds time.Time to a given time interval (e.g. 9:37 rounded to 5 minutes is 9:35) for all sorts of statistical, analytical and rate limiting purposes.

See [docs on godoc.org](https://godoc.org/github.com/andreyvit/timerounding).


## Installation

    go get -u github.com/andreyvit/timerounding


## Example

```go
func main() {
    t := time.Date(2017, 1, 7, 9, 37, 12, 0, time.UTC)

    fmt.Println(timerounding.Round(t, 5*time.Minute))
    // 2017-01-07 09:35:00 +0000 UTC

    fmt.Println(timerounding.Round(t, 24*time.Hour))
    // 2017-01-07 00:00:00 +0000 UTC

    fmt.Println(timerounding.FormatRounded(t, 5*time.Minute, timerounding.Concise))
    // 20170107-0935

    fmt.Println(timerounding.FormatRounded(t, 48*time.Hour, timerounding.Concise))
    // 20170106

    fmt.Println(timerounding.MakeInterval(20*time.Minute).Round(t))
    // 2017-01-07 09:20:00 +0000 UTC
    
    fmt.Println(timerounding.MakeInterval(20*time.Minute).Next(t, 1))
    // 2017-01-07 09:40:00 +0000 UTC
    
    fmt.Println(timerounding.MakeInterval(20*time.Minute).Prev(t, 1))
    // 2017-01-07 09:00:00 +0000 UTC
}
```
