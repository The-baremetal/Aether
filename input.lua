-- === Variables and Assignment === --
local x = 10
local y = x + 5
local a, b = 20, 30

local const z = 42

-- === Control Structures === --


while y > 0 do
    print("y = " .. y)
    y = y - 1
end

repeat
    print("This will print at least once")
until false

for i = 1, 5 do
    print("i = " .. i)
end

for key, value in pairs({a = 1, b = 2}) do
    print(key .. " = " .. value)
end

-- === Functions === --
function greet(name)
    print("Hello, " .. name)
end

greet("World")

function add(a, b)
    return a + b
end

local sum = add(3, 4)

-- Method call syntax
local obj = {value = 5}
function obj:double()
    self.value = self.value * 2
end

obj:double()

-- === Anonymous Functions / Lambdas === --
local square = function(x)
    return x * x
end

print(square(5))

-- === Tables and Arrays === --
local person = {
    name = "Alice",
    age = 30,
    ["__metatable"] = { protected = true }
}

person.job = "Engineer"

local fruits = {"apple", "banana", "cherry"}

-- Indexed access
print(fruits[1])

-- === Expressions and Operators === --
local mathExpr = (2 + 3) * 4 / 2 - 1 % 2 ^ 3
local boolExpr = (true or false) and not false

local isEqual = (5 == 5)
local isNotEqual = (5 != 6)
local comparison = 10 >= 5 and 3 < 4

-- Concatenation
local greeting = "Hello" .. " " .. "Lua"

-- Prefix and Infix expressions
local negated = -x
local logicalNot = not true

-- Indexing
local val = person["name"]

-- Assignment
person.age = 31

-- === Varargs and Variable Arguments === --
function printAll(...)
    for _, v in ipairs({...}) do
        print(v)
    end
end

printAll(1, 2, 3, 4)

-- === Coroutines === --
coroutine.create(function()
    print("Inside coroutine")
end)

coroutine.resume(coroutine.create(function()
    print("Resuming this coroutine")
end))

-- === Async/Await === --
async function fetchData()
    local result = await somePromise()
    print("Data fetched: " .. result)
end

-- === Try-Catch-Finally === --
try
    error("Something went wrong!")
catch err
    print("Caught error: " .. tostring(err))
finally
    print("Finally block executed")
end

-- === Import Statements === --
import "utils.lua"
import "config.lua" as configModule

-- === Return Statement === --
function returnTest()
    return 1, 2, 3
end

local r1, r2, r3 = returnTest()

-- === Break Statement === --
for i = 1, 10 do
    if i > 5 then break end
    print(i)
end

-- === Nil and Literals === --
local nothing = nil
local truth = true
local falsehood = false
local str = "This is a string"
local num = 3.14159

-- === Table Fields and Metaprogramming === --
local metaTable = {
    __metatable = { hidden = true },
    [1 + 2] = "computed key",
    regularKey = "value",
    another = {
        nested = true
    }
}

-- === Function Literal with Receiver === --
local objWithMethod = {}
function objWithMethod:method()
    print("Method called on object")
end

objWithMethod:method()