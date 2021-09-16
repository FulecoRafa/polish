require'stack'

local stack = Stack:create()

local op_table = {
  ['+'] = function(a, b) return a+b end,
  ['-'] = function(a, b) return a-b end,
  ['/'] = function(a, b) return a/b end,
  ['*'] = function(a, b) return a*b end,
  ['^'] = function(a, b) return a^b end,
}

for i=1,#arg do
  local operand = arg[i]
  local num_operand = tonumber(operand, 10)
  if num_operand then
    stack:push(num_operand)
  else
    local right, left = stack:pop(2)
    if not (tonumber(left) and tonumber(right)) then
      error( string.format( "Invalid operation: [%s %s %s] is not valid", left, operand, right))
    end
    local op = op_table[operand]
    if (op) then
      local result = op(left, right)
      stack:push(result)
    else
      error( string.format( "Invalid operand: [%s] is not a valid operation", operand))
    end
  end
end
print(stack:pop())
