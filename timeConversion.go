package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
 * Complete the 'timeConversion' hhifunction below.
 * soufian 1 w
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
    // Extract the time part and AM/PM indicator
    timePart := s[:8]
    period := s[8:]

    // Split the time into hours, minutes, and seconds
    hourStr := timePart[:2]
    minuteStr := timePart[3:5]
    secondStr := timePart[6:8]

    // Convert hour string to integer
    hour := 0
    fmt.Sscanf(hourStr, "%d", &hour)

    // Perform conversion based on AM/PM
    if period == "AM" {
        // Handle special case for 12AM
        if hour == 12 {
            hour = 0
        }
    } else if period == "PM" {
        // Handle special case for 12PM
        if hour != 12 {
            hour += 12
        }
    }

    // Format the hour as two digits
    hourStr = fmt.Sprintf("%02d", hour)

    // Construct the 24-hour formatted string
    militaryTime := hourStr + ":" + minuteStr + ":" + secondStr

    return militaryTime
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
