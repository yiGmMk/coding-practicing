
-- 数组长度,和包含'order'的元素个数
SELECT `oss_key_list` ,JSON_LENGTH(`oss_key_list`), JSON_LENGTH(JSON_SEARCH(`oss_key_list`,"all",'order%')) 
from `order` 
WHERE `oss_key_list` is NOT NULL ; 

-- oss_key_list内不全是包含'order'的记录
SELECT `oss_key_list` ,JSON_LENGTH(`oss_key_list`), JSON_LENGTH(JSON_SEARCH(`oss_key_list`,"all",'order%')) 
from `order` 
WHERE `oss_key_list` is NOT NULL 
and ( JSON_LENGTH(`oss_key_list`) != JSON_LENGTH(JSON_SEARCH(`oss_key_list`,"all",'order%')) -- 两边长度不相等
or JSON_LENGTH(JSON_SEARCH(`oss_key_list`,"all",'order%')) IS NULL) ; -- 或包含'order'的个数为0