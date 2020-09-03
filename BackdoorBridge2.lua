-- some Lua framework if you want it 
local function CommandHandlerStuff()

end

local function GetData()
    return HttpService:JSONDecode(game:HttpGet("YourURL/Commands.json"))
end

while wait(8) do
    local Data = GetData()
    for Name,Command in pairs(Data) do
		-- Your Handling
    end
end
