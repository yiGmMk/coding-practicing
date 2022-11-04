--
-- @lc app=leetcode.cn id=1179 lang=mysql
--
-- [1179] 重新格式化部门表
--
-- https://leetcode.cn/problems/reformat-department-table/description/
--
-- database
-- Easy (64.35%)
-- Likes:    194
-- Dislikes: 0
-- Total Accepted:    42.1K
-- Total Submissions: 65.5K
-- Testcase Example:  '{"headers":{"Department":["id","revenue","month"]},"rows":{"Department":[[1,8000,"Jan"],[2,9000,"Jan"],[3,10000,"Feb"],[1,7000,"Feb"],[1,6000,"Mar"]]}}'
--
-- 部门表 Department：
-- 
-- 
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | revenue       | int     |
-- | month         | varchar |
-- +---------------+---------+
-- (id, month) 是表的联合主键。
-- 这个表格有关于每个部门每月收入的信息。
-- 月份（month）可以取下列值
-- ["Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"]。
-- 
-- 
-- 
-- 
-- 编写一个 SQL 查询来重新格式化表，使得新的表中有一个部门 id 列和一些对应 每个月 的收入（revenue）列。
-- 
-- 查询结果格式如下面的示例所示：
-- 
-- 
-- Department 表：
-- +------+---------+-------+
-- | id   | revenue | month |
-- +------+---------+-------+
-- | 1    | 8000    | Jan   |
-- | 2    | 9000    | Jan   |
-- | 3    | 10000   | Feb   |
-- | 1    | 7000    | Feb   |
-- | 1    | 6000    | Mar   |
-- +------+---------+-------+
-- 
-- 查询得到的结果表：
-- +------+-------------+-------------+-------------+-----+-------------+
-- | id   | Jan_Revenue | Feb_Revenue | Mar_Revenue | ... | Dec_Revenue |
-- +------+-------------+-------------+-------------+-----+-------------+
-- | 1    | 8000        | 7000        | 6000        | ... | null        |
-- | 2    | 9000        | null        | null        | ... | null        |
-- | 3    | null        | 10000       | null        | ... | null        |
-- +------+-------------+-------------+-------------+-----+-------------+
-- 
-- 注意，结果表有 13 列 (1个部门 id 列 + 12个月份的收入列)。
-- 
-- 
--

-- @lc code=start
# Write your MySQL query statement below
select id,
sum(if(month="Jan", revenue,null)) Jan_Revenue,
sum(if(month="Feb", revenue,null)) Feb_Revenue,
sum(if(month="Mar", revenue,null)) Mar_Revenue,
sum(if(month="Apr", revenue,null)) Apr_Revenue,
sum(if(month="May", revenue,null)) May_Revenue,
sum(if(month="Jun", revenue,null)) Jun_Revenue,
sum(if(month="Jul", revenue,null)) Jul_Revenue,
sum(if(month="Aug", revenue,null)) Aug_Revenue,
sum(if(month="Sep", revenue,null)) Sep_Revenue,
sum(if(month="Oct", revenue,null)) Oct_Revenue,
sum(if(month="Nov", revenue,null)) Nov_Revenue,
sum(if(month="Dec", revenue,null)) Dec_Revenue
from department
group by id
-- @lc code=end

