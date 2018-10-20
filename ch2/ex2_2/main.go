// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
	"gopl.io/ch2/weightconv"
)

func main() {
	scanForNumbers(func(arg string) {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatalf("gf: %v\n", err)
		}
		f := tempconv.Fahrenheit(v)
		c := tempconv.Celsius(v)
		kg := weightconv.Kg(v)
		lb := weightconv.Pound(v)
		fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c), kg, kg.ToPound(), lb, lb.ToKg())
	})
}

func scanForNumbers(wordFunc func(string)) {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			wordFunc(arg)
		}
	} else {
		fmt.Println("Waiting for input:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			wordFunc(scanner.Text())
		}
	}
}

//!-
