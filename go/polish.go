package main

import (
	"fmt"
	"os"
	"strconv"
  "math"
)

func main() {
  var stack Stack;
  var args = os.Args[:]
  if len(args) <= 1 {fmt.Println("Provide arguments!"); return}
  for _, arg := range args[1:] {
    if val, err := strconv.ParseFloat(arg, 64); err == nil {
      stack.push(val)
    } else {
      vals := stack.pop(2);
      right, left := vals[0], vals[1];
      var resp float64;
      switch arg {
      case "+":
        resp = left + right
      case "-":
        resp = left - right
      case "*":
        resp = left * right
      case "/":
        resp = left / right
      case "^":
        resp = math.Pow(left, right)
      default:
        fmt.Println(fmt.Errorf("Invalid operand %s", arg));
        return
      }
      stack.push(resp);
    }
  }
  fmt.Println(stack.pop(1)[0]);
}
