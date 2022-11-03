--
-- @lc app=leetcode.cn id=1484 lang=mysql
--
-- [1484] 按日期分组销售产品
--
-- https://leetcode.cn/problems/group-sold-products-by-the-date/description/
--
-- database
-- Easy (67.70%)
-- Likes:    110
-- Dislikes: 0
-- Total Accepted:    34.2K
-- Total Submissions: 50.6K
-- Testcase Example:  '{"headers":{"Activities":["sell_date","product"]},"rows":{"Activities":[["2020-05-30","Headphone"],["2020-06-01","Pencil"],["2020-06-02","Mask"],["2020-05-30","Basketball"],["2020-06-01","Bible"],["2020-06-02","Mask"],["2020-05-30","T-Shirt"]]}}'
--
-- 表 Activities：
-- 
-- 
-- +-------------+---------+
-- | 列名         | 类型    |
-- +-------------+---------+
-- | sell_date   | date    |
-- | product     | varchar |
-- +-------------+---------+
-- 此表没有主键，它可能包含重复项。
-- 此表的每一行都包含产品名称和在市场上销售的日期。
-- 
-- 
-- 
-- 
-- 编写一个 SQL 查询来查找每个日期、销售的不同产品的数量及其名称。
-- 每个日期的销售产品名称应按词典序排列。
-- 返回按 sell_date 排序的结果表。
-- 查询结果格式如下例所示。
-- 
-- 
-- 
-- 示例 1:
-- 
-- 
-- 输入：
-- Activities 表：
-- +------------+-------------+
-- | sell_date  | product     |
-- +------------+-------------+
-- | 2020-05-30 | Headphone   |
-- | 2020-06-01 | Pencil      |
-- | 2020-06-02 | Mask        |
-- | 2020-05-30 | Basketball  |
-- | 2020-06-01 | Bible       |
-- | 2020-06-02 | Mask        |
-- | 2020-05-30 | T-Shirt     |
-- +------------+-------------+
-- 输出：
-- +------------+----------+------------------------------+
-- | sell_date  | num_sold | products                     |
-- +------------+----------+------------------------------+
-- | 2020-05-30 | 3        | Basketball,Headphone,T-shirt |
-- | 2020-06-01 | 2        | Bible,Pencil                 |
-- | 2020-06-02 | 1        | Mask                         |
-- +------------+----------+------------------------------+
-- 解释：
-- 对于2020-05-30，出售的物品是 (Headphone, Basketball, T-shirt)，按词典序排列，并用逗号 ',' 分隔。
-- 对于2020-06-01，出售的物品是 (Pencil, Bible)，按词典序排列，并用逗号分隔。
-- 对于2020-06-02，出售的物品是 (Mask)，只需返回该物品名。
-- 
--

-- @lc code=start
# Write your MySQL query statement below

select sell_date,count(distinct product) as num_sold,group_concat(distinct product order by product asc) as products
from  Activities
group by sell_date
order by sell_date
-- @lc code=end

