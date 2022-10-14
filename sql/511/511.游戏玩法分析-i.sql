--
-- @lc app=leetcode.cn id=511 lang=mysql
--
-- [511] 游戏玩法分析 I
--
-- https://leetcode.cn/problems/game-play-analysis-i/description/
--
-- database
-- Easy (71.95%)
-- Likes:    76
-- Dislikes: 0
-- Total Accepted:    60.4K
-- Total Submissions: 84K
-- Testcase Example:  '{"headers":{"Activity":["player_id","device_id","event_date","games_played"]},"rows":{"Activity":[[1,2,"2016-03-01",5],[1,2,"2016-05-02",6],[2,3,"2017-06-25",1],[3,1,"2016-03-02",0],[3,4,"2018-07-03",5]]}}'
--
-- 活动表 Activity：
-- 
-- 
-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | player_id    | int     |
-- | device_id    | int     |
-- | event_date   | date    |
-- | games_played | int     |
-- +--------------+---------+
-- 表的主键是 (player_id, event_date)。
-- 这张表展示了一些游戏玩家在游戏平台上的行为活动。
-- 每行数据记录了一名玩家在退出平台之前，当天使用同一台设备登录平台后打开的游戏的数目（可能是 0 个）。
-- 
-- 
-- 
-- 
-- 写一条 SQL 查询语句获取每位玩家 第一次登陆平台的日期。
-- 
-- 查询结果的格式如下所示：
-- 
-- 
-- Activity 表：
-- +-----------+-----------+------------+--------------+
-- | player_id | device_id | event_date | games_played |
-- +-----------+-----------+------------+--------------+
-- | 1         | 2         | 2016-03-01 | 5            |
-- | 1         | 2         | 2016-05-02 | 6            |
-- | 2         | 3         | 2017-06-25 | 1            |
-- | 3         | 1         | 2016-03-02 | 0            |
-- | 3         | 4         | 2018-07-03 | 5            |
-- +-----------+-----------+------------+--------------+
-- 
-- Result 表：
-- +-----------+-------------+
-- | player_id | first_login |
-- +-----------+-------------+
-- | 1         | 2016-03-01  |
-- | 2         | 2017-06-25  |
-- | 3         | 2016-03-02  |
-- +-----------+-------------+
-- 
-- 
--

-- @lc code=start
# Write your MySQL query statement below
-- 窗口函数求topN
-- select player_id,event_date `first_login`
-- from (
--  select player_id,event_date,dense_rank()over(partition by player_id order by event_date) as rnk
--  from Activity
-- )tmp
-- where rnk=1

-- 聚合函数
select player_id, min(event_date) `first_login`
from Activity
group by player_id
Order by event_date

-- @lc code=end

