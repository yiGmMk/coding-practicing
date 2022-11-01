--
-- @lc app=leetcode.cn id=1729 lang=mysql
--
-- [1729] 求关注者的数量
--
-- https://leetcode.cn/problems/find-followers-count/description/
--
-- database
-- Easy (62.26%)
-- Likes:    24
-- Dislikes: 0
-- Total Accepted:    18.9K
-- Total Submissions: 30.3K
-- Testcase Example:  '{"headers":{"Followers":["user_id","follower_id"]},"rows":{"Followers":[["0","1"],["1","0"],["2","0"],["2","1"]]}}'
--
-- 表： Followers
-- 
-- 
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | user_id     | int  |
-- | follower_id | int  |
-- +-------------+------+
-- (user_id, follower_id) 是这个表的主键。
-- 该表包含一个关注关系中关注者和用户的编号，其中关注者关注用户。
-- 
-- 
-- 
-- 写出 SQL 语句，对于每一个用户，返回该用户的关注者数量。
-- 
-- 按 user_id 的顺序返回结果表。
-- 
-- 查询结果的格式如下示例所示。
-- 
-- 
-- 
-- 示例 1：
-- 
-- 
-- 输入：
-- Followers 表：
-- +---------+-------------+
-- | user_id | follower_id |
-- +---------+-------------+
-- | 0       | 1           |
-- | 1       | 0           |
-- | 2       | 0           |
-- | 2       | 1           |
-- +---------+-------------+
-- 输出：
-- +---------+----------------+
-- | user_id | followers_count|
-- +---------+----------------+
-- | 0       | 1              |
-- | 1       | 1              |
-- | 2       | 2              |
-- +---------+----------------+
-- 解释：
-- 0 的关注者有 {1}
-- 1 的关注者有 {0}
-- 2 的关注者有 {0,1}
-- 
--

-- @lc code=start
# Write your MySQL query statement below

select user_id,count(distinct follower_id) followers_count
from followers
group by user_id

-- @lc code=end

