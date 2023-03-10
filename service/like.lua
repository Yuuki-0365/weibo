-- 返回值
-- 0 -- 出错 1 -- 点赞成功 2 -- 取消点赞成功
local value = redis.call("GET", KEYS[1])
if value == false
then
    redis.call("SET", KEYS[1], 1)
    redis.call("INCR", KEYS[2])
    return 1
elseif tonumber(value) == 0
then
    redis.call("SET", KEYS[1], 1)
    redis.call("INCR", KEYS[2])
    return 1
elseif tonumber(value) == 1
then
        redis.call("SET", KEYS[1], 0)
        redis.call("DECR", KEYS[2])
        return 2
else
    return 0
end