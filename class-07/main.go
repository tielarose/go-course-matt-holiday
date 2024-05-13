package main

// // ***********************
// // Playing with Formatting
// // ***********************

// import (
// 	"io"
// )

// func main() {
// 	a, b := 12, 345
// 	c, d := 1.2, 3.45

// 	fmt.Printf("%d %d\n", a, b)
// 	fmt.Printf("%X %X\n", a, b)
// 	fmt.Printf("%#x %#x\n", a, b)
// 	fmt.Printf("%f %.2f\n", c, d)

// 	fmt.Println()

// 	fmt.Printf("|%6d|%6d|\n", a, b)
// 	fmt.Printf("|%06d|%06d|\n", a, b)
// 	fmt.Printf("|%-6d|%-6d|\n", a, b)
// 	fmt.Printf("|%6f|%6.2f|\n", c, d)

// 	s := []int{1, 2, 3}
// 	a := [3]rune{'a', 'b', 'c'}

// 	fmt.Printf("%T\n", s)
// 	fmt.Printf("%v\n", s)
// 	fmt.Printf("%#v\n", s)

// 	fmt.Println()

// 	fmt.Printf("%T\n", a)
// 	fmt.Printf("%q\n", a)
// 	fmt.Printf("%v\n", a)
// 	fmt.Printf("%#v\n", a)

// 	m := map[string]int{"and":1, "or":2}

// 	fmt.Printf("%T\n", m)
// 	fmt.Printf("%v\n", m)
// 	fmt.Printf("%#v\n", m)

// }

// // ***********************
// // Concatenating Files
// // ***********************

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	for _, fname := range os.Args[1:] {
// 		file, err := os.Open(fname)

// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			continue
// 		}

// 		if _, err := io.Copy(os.Stdout, file); err != nil {
// 			fmt.Fprint(os.Stderr, err)
// 			continue
// 		}

// 		file.Close()
// 	}
// }

// ***********************
// Printing Data About Files
// ***********************

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, fname := range os.Args[1:] {
		var lc, wc, cc int

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}

		fmt.Printf(" %7d %7d %7d %s\n", lc, wc, cc, fname)
		file.Close()
	}
}

