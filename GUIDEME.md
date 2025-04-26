# AETHER Code Quality Guide  
*Best practices for writing clean, maintainable AETHER code*

---

### 📛 1. Meaningful Names  
**DO THIS**              | **AVOID THIS**  
-------------------------|-----------------
`local total_price = 100`|`local a = 100`  
`local function calculate_discount()`|`local function do_stuff()`  

**Key rules:**  
- Use snake_case for variables/functions  
- Reserve ALL_CAPS for constants  
- Make names self-documenting:  
  ```lua
  -- Good
  local player_health = 100
  local MAX_INVENTORY_SLOTS = 10
  
  -- Bad
  local h = 100
  local mis = 10
  ```

---

### 🧩 2. Modular Code Structure  
**File: items/weapon.lua**  
```lua
local Weapon = {}

function Weapon.new(damage, attack_speed)
  return {
    damage = damage,
    attack_speed = attack_speed,
    last_used = 0
  }
end

return Weapon
```

**Usage:**  
```lua
import items.weapon as Weapon
local sword = Weapon.new(15, 1.2)
```

---

### ✂️ 3. Focused Functions  
**Good (Single Responsibility):**  
```lua
function calculate_circle_area(radius)
  return math.pi * radius^2
end

function format_area_output(area)
  return string.format("Area: %.2f m²", area)
end
```

**Bad (Multi-purpose):**  
```lua
function handle_circle_stuff(radius)
  local area = math.pi * radius^2
  print(string.format("Area: %.2f m²", area))
  -- Also handles 10 other unrelated operations
end
```

---

### 📝 4. Effective Documentation  
**Function Documentation:**  
```lua
--- Calculates combat rating based on player stats
-- @param attack number  Player's attack power
-- @param defense number  Player's defense rating
-- @return number  Combat effectiveness score
function calculate_combat_rating(attack, defense)
  return (attack * 0.7) + (defense * 0.3)
end
```

**Contextual Comments:**  
```lua
-- Using memoization to cache NPC dialogue results
-- because dialogue trees are parsed multiple times
local dialogue_cache = setmetatable({}, { __mode = "v" })
```

---

### 🎩 5. Magic Number Prevention  
**config.lua**  
```lua
local Config = {
  MAX_PARTY_MEMBERS = 4,
  BASE_MOVEMENT_SPEED = 5.2,
  CRITICAL_HIT_CHANCE = 0.15
}

return Config
```

**Usage:**  
```lua
if party_size > Config.MAX_PARTY_MEMBERS then
  error("Party size exceeds maximum allowed")
end
```

---

### 🧹 6. Consistent Formatting  
**Vertical Alignment:**  
```lua
local character = {
  name          = "Arthas",
  class         = "Death Knight",
  level         = 80,
  inventory     = {},
  equipped_gear = { "Frostmourne", "Plate Armor" }
}

if is_active_player and 
   not is_npc and 
   has_quest_flag then
  start_cutscene()
end
```

---

### 🚨 7. Error Handling  
**Safe Execution:**  
```lua
local success, err = pcall(function()
  load_external_plugin("new_quest_pack.dll")
end)

if not success then
  log_error("Plugin load failed: " .. err)
  show_player_message("Content failed to load")
end
```

---

### 🧪 8. Test Practices  
**tests/combat_test.lua**  
```lua
describe("Damage calculation", function()
  it("should apply critical hit multiplier", function()
    local base_damage = 100
    local result = calculate_damage(base_damage, true)
    assert.equal(150, result)
  end)

  it("should handle zero damage edge case", function()
    local result = calculate_damage(0, false)
    assert.equal(0, result)
  end)
end)
```

---

### 🔄 9. Version Control Tips  
- Commit messages format:  
  `[system] Brief description`  
  Examples:  
  `[combat] Fix damage calculation overflow`  
  `[ui] Add inventory sorting buttons`

---

### ♻️ 10. Regular Refactoring  
**Before:**  
```lua
if quest.status == "active" then
  if #quest.objectives > 0 then
    for i, obj in ipairs(quest.objectives) do
      if not obj.complete then
        -- 15 lines of nested logic
      end
    end
  end
end
```

**After:**  
```lua
function has_incomplete_objectives(quest)
  for _, objective in ipairs(quest.objectives) do
    if not objective.complete then
      return true
    end
  end
  return false
end

if quest.status == "active" and has_incomplete_objectives(quest) then
  handle_active_quest(quest)
end
```

---

### 🚫 11. Dependency Management  
**Recommended Stack:**  
| Purpose       | Library     |
|---------------|-------------|
| Testing       | None yet    |
| Utility       | None yet    |
| JSON          | None yet    |
| Networking    | None yet    |

---

### 🚫12. Scopes
**Avoid this:**
```lua
-- AVOID THIS, Even though the compiler may free it when it is not needed, it slows down compile times and makes your code a bit unmaintainable
onetimecompute = "hello"
print(hello..hello)
```
**Do this:**
```lua
-- Do this
local onetimecomputer = "hello"
print(hello..hello)
```

---

### ✅ 13. Code Review Checklist  
1. All functions under 40 lines?  
2. No magic numbers/strings?  
3. Error handling present?  
4. Tests cover edge cases?  
5. Documentation clear?  
6. No duplicated logic?
