select 
	player_id, device_id
from activity
where (player_id, event_date) in 
(
    select player_id, min(event_date)
	from activity
	group by player_id
)

-- 作者：zg104
-- 链接：https://leetcode.cn/problems/game-play-analysis-ii/solution/by-zg104-xt7d/
-- 来源：力扣（LeetCode）
-- 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。