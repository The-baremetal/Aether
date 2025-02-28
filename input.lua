-- Basic syntax
local x = 5 * (3 + 2)
const MAX = 100
local y = "hello" .. "world"
local z = true

-- Control flow
if x > MAX then
    print("Overflow")
elseif x < 0 then
    print("Underflow")
else
    print("Within limits")
end

while x < 50 do
    x += 1
    print(x)
end

repeat
    x -= 1
until x <= 50

-- Functions
function factorial(n: number): number
    if n <= 1 then
        return 1
    end
    return n * factorial(n - 1)
end

local fib = function(n)
    if n <= 2 then
        return 1
    end
    return fib(n-1) + fib(n-2)
end

-- Tables and metatables
local player = {
    name = "Arthur",
    health = 100,
    inventory = { "sword", "shield" },
    __metatable = { __index = function(t,k) return "N/A" end }
}

-- Error handling
try
    local risky = dangerousOperation()
catch err
    print("Error occurred:", err)
finally
    cleanupResources()
end