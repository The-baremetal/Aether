local yes = "yes"
local no = "no"
local a = 1

local b = 2
local c = a + b
local d = a - b
local e = a * b
local f = a / b
local g = a // b
local h = a % b
local i = a ^ b
local j = a & b
local k = a | b
local l = a << b
local m = a >> b
local n = a .. b
local o = a ?? b
local p = a +_ b
local q = a -_ b
local r = a <<= b
local s = a >>= b
local t = a ..= b
local u = a ??= b

-- Test table method definitions and constructor
local counter = {
    value = 0,
    increment = function(self)  -- Traditional method
        self.value = self.value + 1
    end,
    reset = :function()  -- Shorthand method syntax
        self.value = 0
    end
}

-- Test recursive variable access
local data = {
    nested = {
        [1] = { value = "deep" },
        info = { version = 2.1 }
    }
}
assert(data.nested[1].value == "deep")
assert(data["nested"].info.version == 2.1)

-- Test metatable functionality
local secureTable = setmetatable({}, {
    __metatable = "protected",
    __index = function(t, k)
        return "restricted access"
    end
})
assert(getmetatable(secureTable) == "protected")
assert(secureTable.secretData == "restricted access")

-- Test typeof and bitwise operators
local typeCheck = typeof("hello") .. typeof(42)
assert(typeCheck == "stringnumber")
assert(~15 == -16)  -- Bitwise NOT

-- Test local function with parameters
local function factorial(n)
    return n <= 1 and 1 or n * factorial(n-1)
end
assert(factorial(5) == 120)

-- Test vararg handling
local function sum(...)
    local total = 0
    for _, v in ipairs({...}) do
        total = total + v
    end
    return total
end
assert(sum(1,2,3,4) == 10)

-- Test method chaining
local builder = {
    value = "",
    append = :function(s)
        self.value = self.value .. s
        return self
    end,
    reset = :function()
        self.value = ""
        return self
    end
}
builder:append("hello"):append(" "):append("world")
assert(builder.value == "hello world")

-- Test enhanced table constructor
local complexTable = {
    id = 1,
    ["computed".."Key"] = 2,
    { "array-style" },
    method = :function() return "works" end,
    __metatable = { __index = function() return "fallback" end }
}
assert(complexTable.method() == "works")
assert(complexTable[1][1] == "array-style")

local user = { profile = { name = "Alex" } }
print(user?.missing?.name ?? "anonymous")  -- Should print "anonymous"