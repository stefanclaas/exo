// e(moji)xo - you may use it with my xo program and crockford-base32
// encoder/decoder
// Usage: [encode] crockford-base32 < msg.txt | xo | exo
//        [decode] exo -d | xo -d | crockford-base32 -d
// Pretty much overhead, the whole procedure, but fun.

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

func encodeMessage(message string) string {
    // Replace 'x' with 🔴, 'o' with 🔵, and ' ' with 🔲
    message = strings.ReplaceAll(message, "x", "🔴")
    message = strings.ReplaceAll(message, "o", "🔵")
    message = strings.ReplaceAll(message, " ", "🔲")
    return message
}

func decodeMessage(encodedMessage string) string {
    // Replace 🔴 with 'x', 🔵 with 'o', and 🔲 with ' '
    encodedMessage = strings.ReplaceAll(encodedMessage, "🔴", "x")
    encodedMessage = strings.ReplaceAll(encodedMessage, "🔵", "o")
    encodedMessage = strings.ReplaceAll(encodedMessage, "🔲", " ")
    return encodedMessage
}

func main() {
    args := os.Args[1:]

    if len(args) == 1 && args[0] == "-d" {
        // Decoding mode, read from stdin and decode
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            decoded := decodeMessage(scanner.Text())
            fmt.Println(decoded)
        }
        if scanner.Err() != nil {
            fmt.Fprintln(os.Stderr, "Error reading input:", scanner.Err())
            os.Exit(1)
        }
    } else {
        // Encoding mode, read from stdin and encode
        var inputBuilder strings.Builder
        _, err := io.Copy(&inputBuilder, os.Stdin)
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error reading input:", err)
            os.Exit(1)
        }

        encoded := encodeMessage(inputBuilder.String())
        fmt.Println(encoded)
    }
}

