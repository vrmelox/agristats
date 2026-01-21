package main
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
    end := time.Date(2030, 12, 31, 0, 0, 0, 0, time.UTC).Unix()

    randomTimestamp := rand.Int63n(end-start) + start
    randomDate := time.Unix(randomTimestamp, 0).UTC()

    // ne garder que la date
    fmt.Println("Date aléatoire :", randomDate.Format("2006-01-02"))
}
