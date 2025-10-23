package main

import "os"

const (
	Max = 9223372036854775807
	Min = -9223372036854775808
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := atoi(os.Args[1])
	b, ok2 := atoi(os.Args[3])
	op := os.Args[2]
	if !ok1 || !ok2 {
		return
	}

	if (op == "/" || op == "%") && b == 0 {
		if op == "/" {
			printStr("No division by 0\n")
		} else {
			printStr("No modulo by 0\n")
		}
		return
	}

	res, ok := calc(a, b, op)
	if !ok {
		return
	}
	printInt(res)
	printStr("\n")
}

func calc(a, b int64, op string) (int64, bool) {
	switch op {
	case "+":
		if (b > 0 && a > Max-b) || (b < 0 && a < Min-b) {
			return 0, false
		}
		return a + b, true
	case "-":
		if (b < 0 && a > Max+b) || (b > 0 && a < Min+b) {
			return 0, false
		}
		return a - b, true
	case "*":
		if a != 0 && (a == Min && b == -1 || b == Min && a == -1 || abs(a) > Max/abs(b)) {
			return 0, false
		}
		return a * b, true
	case "/":
		if a == Min && b == -1 {
			return 0, false
		}
		return a / b, true
	case "%":
		return a % b, true
	default:
		return 0, false
	}
}

func atoi(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}
	sign := int64(1)
	i := 0
	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}
	var n int64
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		d := int64(c - '0')
		if n > (Max-d)/10 {
			return 0, false
		}
		n = n*10 + d
	}
	n *= sign
	if n > Max || n < Min {
		return 0, false
	}
	return n, true
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func printInt(n int64) {
	if n == Min {
		printStr("-9223372036854775808")
		return
	}
	if n < 0 {
		os.Stdout.Write([]byte{'-'})
		n = -n
	}
	var buf [20]byte
	i := 0
	for {
		buf[i] = byte(n%10) + '0'
		n /= 10
		i++
		if n == 0 {
			break
		}
	}
	for j := i - 1; j >= 0; j-- {
		os.Stdout.Write([]byte{buf[j]})
	}
}

func printStr(s string) {
	os.Stdout.Write([]byte(s))
}
