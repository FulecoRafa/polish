package polish

import java.util.Stack;
import kotlin.math.pow;

fun main(args: Array<String>) {
    val stack = Stack<Float>();
    for (operand in args) {
      var n_operand = operand.toFloatOrNull();
      if (n_operand != null) {
        stack.push(n_operand);
      } else {
        var right = stack.pop();
        var left = stack.pop();
        val resul = when (operand) {
          "+" -> left + right
          "-" -> left - right
          "*" -> left * right
          "/" -> left / right
          "^" -> left.pow(right)
          else -> throw Exception("Operand $operand is invalid")
        }
        stack.push(resul);
      }
    }
    println(stack.pop())
}
