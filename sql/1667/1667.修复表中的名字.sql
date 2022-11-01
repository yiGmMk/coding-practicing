--
-- @lc app=leetcode.cn id=1667 lang=mysql
--
-- [1667] 修复表中的名字
--
-- https://leetcode.cn/problems/fix-names-in-a-table/description/
--
-- database
-- Easy (63.90%)
-- Likes:    53
-- Dislikes: 0
-- Total Accepted:    32.9K
-- Total Submissions: 51.4K
-- Testcase Example:  '{"headers":{"Users":["user_id","name"]},"rows":{"Users":[[1,"aLice"],[2,"bOB"]]}}'
--
-- 表： Users
-- 
-- 
-- +----------------+---------+
-- | Column Name    | Type    |
-- +----------------+---------+
-- | user_id        | int     |
-- | name           | varchar |
-- +----------------+---------+
-- user_id 是该表的主键。
-- 该表包含用户的 ID 和名字。名字仅由小写和大写字符组成。
-- 
-- 
-- 
-- 
-- 编写一个 SQL 查询来修复名字，使得只有第一个字符是大写的，其余都是小写的。
-- 
-- 返回按 user_id 排序的结果表。
-- 
-- 查询结果格式示例如下。
-- 
-- 
-- 
-- 示例 1：
-- 
-- 
-- 输入：
-- Users table:
-- +---------+-------+
-- | user_id | name  |
-- +---------+-------+
-- | 1       | aLice |
-- | 2       | bOB   |
-- +---------+-------+
-- 输出：
-- +---------+-------+
-- | user_id | name  |
-- +---------+-------+
-- | 1       | Alice |
-- | 2       | Bob   |
-- +---------+-------+
-- 
--

-- @lc code=start
# Write your MySQL query statement below
-- 拼接:第一个大写+后面的小写
select user_id,CONCAT(UPPER(SUBSTRING(name,1,1)),LOWER(SUBSTRING(name,2))) AS Name
from Users
order by user_id 


-- 作者：JSJohnsonJS
-- 链接：https://leetcode.cn/problems/fix-names-in-a-table/solution/by-clever-austinzya-lh8c/
-- 来源：力扣（LeetCode）
-- 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
select user_id, CONCAT(UPPER(left(name, 1)), LOWER(RIGHT(name, length(name) - 1))) as name
from Users
order by user_id

-- @lc code=end

