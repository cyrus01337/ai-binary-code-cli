package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func textToBinary(text string) string {
	var builder strings.Builder

	for i, c := range text {
		if i > 0 {
			builder.WriteByte(' ')
		}

		fmt.Fprintf(&builder, "%08b", c)
	}

	return builder.String()
}

func binaryToText(binary string) (string, error) {
	segments := strings.Fields(binary)

	if len(segments) == 0 {
		return "", errors.New("no binary data provided")
	}

	var builder strings.Builder

	for _, segment := range segments {
		if len(segment) != 8 {
			return "", fmt.Errorf("binary segment of non-8 characters '%s' provided", segment)
		}

		asDecimal, err := strconv.ParseInt(segment, 2, 64)

		if err != nil {
			return "", fmt.Errorf("invalid binary segment '%s': %w", segment, err)
		}

		builder.WriteRune(rune(asDecimal))
	}

	return builder.String(), nil
}

func main() {
	encode := flag.Bool("encode", false, "Encode text to binary")
	decode := flag.Bool("decode", false, "Decode binary to text")

	flag.Parse()

	arguments := flag.Args()

	if len(arguments) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	content := strings.Join(arguments, " ")

	if *encode {
		fmt.Println(textToBinary(content))
	} else if *decode {
		output, err := binaryToText(content)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(output)
	} else {
		flag.Usage()
	}
}
