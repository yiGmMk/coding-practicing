--
-- @lc app=leetcode.cn id=197 lang=mysql
--
-- [197] 上升的温度
--
-- https://leetcode.cn/problems/rising-temperature/description/
--
-- database
-- Easy (53.91%)
-- Likes:    360
-- Dislikes: 0
-- Total Accepted:    148K
-- Total Submissions: 274.4K
-- Testcase Example:  '{"headers": {"Weather": ["Id", "RecordDate", "Temperature"]}, "rows": {"Weather": [[1, "2015-01-01", 10], [2, "2015-01-02", 25], [3, "2015-01-03", 20], [4, "2015-01-04", 30]]}}'
--
-- 
-- 
-- 表： Weather
-- 
-- 
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | recordDate    | date    |
-- | temperature   | int     |
-- +---------------+---------+
-- id 是这个表的主键
-- 该表包含特定日期的温度信息
-- 
-- 
-- 
-- 编写一个 SQL 查询，来查找与之前（昨天的）日期相比温度更高的所有日期的 id 。
-- 
-- 返回结果 不要求顺序 。
-- 
-- 查询结果格式如下例。
-- 
-- 
-- 
-- 示例 1：
-- 
-- 
-- 输入：
-- Weather 表：
-- +----+------------+-------------+
-- | id | recordDate | Temperature |
-- +----+------------+-------------+
-- | 1  | 2015-01-01 | 10          |
-- | 2  | 2015-01-02 | 25          |
-- | 3  | 2015-01-03 | 20          |
-- | 4  | 2015-01-04 | 30          |
-- +----+------------+-------------+
-- 输出：
-- +----+
-- | id |
-- +----+
-- | 2  |
-- | 4  |
-- +----+
-- 解释：
-- 2015-01-02 的温度比前一天高（10 -> 25）
-- 2015-01-04 的温度比前一天高（20 -> 30）
-- 
-- 
-- 
--

-- @lc code=start
# Write your MySQL query statement below
select Weather.id 
from Weather left join Weather as t
on Weather.recordDate=date_add(t.recordDate, INTERVAL 1 day)
where Weather.Temperature > t.Temperature

-- 下面几个是不同写法,查询速度差不多
-- select a.ID
-- from weather as a cross join weather as b 
--      on datediff(a.recordDate , b.recordDate ) = 1
-- where a.Temperature  > b.Temperature ;

-- select a.ID
-- from weather as a cross join weather as b 
--      on timestampdiff(day,a.recordDate , b.recordDate ) = -1
-- where a.Temperature  > b.Temperature ;


# 作者：houzidata
# 链接：https://leetcode.cn/problems/rising-temperature/solution/tu-jie-sqlmian-shi-ti-ru-he-bi-jiao-ri-qi-shu-ju-b/
# 来源：力扣（LeetCode）
# 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
# Write your MySQL query statement below
-- select b.id 
-- from (
--     select id, date_add(recordDate,INTERVAL 1 DAY)as rd,Temperature from Weather 
--     ) a 
--     inner join Weather b 
--     on a.rd=b.recordDate
-- where a.Temperature < b.Temperature;
-- @lc code=end

