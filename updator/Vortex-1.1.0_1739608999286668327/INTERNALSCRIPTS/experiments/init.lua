-- AFTER THE PARSER IS DONE, PLEASE EMBED THIS CODE (WHEN COMPILED) INTO THE PROJECT SO WHEN A USER REQUESTS AN EXPERIMENT, IT IS INSTANTLY AVAILABLE. AND YES, I KNOW THIS IS A LITTLE BIT OF A HACK, BUT IT IS THE ONLY WAY TO GET THE JOB DONE. MODULES IN THIS PROJECT ARE STRICTLY WRITTEN IN THE LANGUAGE THIS PROJECT IS.
local experiments = {
    safe_nav = false,
    pipe = false,
    decorators = false
}

return setmetatable(experiments, {
    __index = function(_, name)
        error(("Experiment %q not registered"):format(name))
    end,
    __call = function(self, name)
        local ok, mod = pcall(require, "experiments."..name)
        if ok then
            self[name] = true
            return mod
        end
        error(("Experiment %q not found"):format(name))
    end
}) 