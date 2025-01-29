local count = 3
repeat
    print("Repeat count:", count)
    count = count - 1
until count == 0

-- Testing bitwise operators
local a = 0b1010 & 0b1100  -- 8 (1000)
local b = 0b1010 | 0b1100   -- 14 (1110)
print("Bitwise AND:", a, "OR:", b)

-- Testing floor division
print("Floor division:", 5 // 2)  -- 2

-- Testing anonymous function
local square = function(x) return x * x end
print("Square of 5:", square(5))

-- Testing method call
local person = {
    name = "Alice",
    greet = function(self)
        return "Hello, " .. self.name
    end
}
print(person:greet())

-- Testing metatable
local meta = {
    __index = {
        default = "unknown"
    }
}
local obj = setmetatable({}, meta)
print("Metatable access:", obj.default)

-- Testing goto and labels
::start::
print("Testing goto")
goto finish

-- This won't execute due to goto
print("This should be skipped")

::finish::
print("Jumped to finish")

-- Testing return with multiple values
function getValues()
    return 1, "two", true
end
local x, y, z = getValues()
print("Multiple returns:", x, y, z)

-- Testing break in loop
while true do
    print("Breaking loop")
    break
end