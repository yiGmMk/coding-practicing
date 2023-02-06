--
-- @lc app=leetcode.cn id=176 lang=mysql
--
-- [176] 第二高的薪水
--
-- https://leetcode.cn/problems/second-highest-salary/description/
--
-- database
-- Medium (36.27%)
-- Likes:    1251
-- Dislikes: 0
-- Total Accepted:    379.5K
-- Total Submissions: 1M
-- Testcase Example:  '{"headers":{"Employee":["id","salary"]},"rows":{"Employee":[[1,100],[2,200],[3,300]]}}'
--
-- Employee 表：
-- 
-- 
-- 
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | id          | int  |
-- | salary      | int  |
-- +-------------+------+
-- id 是这个表的主键。
-- 表的每一行包含员工的工资信息。
-- 
-- 
-- 
-- 
-- 编写一个 SQL 查询，获取并返回 Employee 表中第二高的薪水 。如果不存在第二高的薪水，查询应该返回 null 。
-- 
-- 查询结果如下例所示。
-- 
-- 
-- 
-- 示例 1：
-- 
-- 
-- 输入：
-- Employee 表：
-- +----+--------+
-- | id | salary |
-- +----+--------+
-- | 1  | 100    |
-- | 2  | 200    |
-- | 3  | 300    |
-- +----+--------+
-- 输出：
-- +---------------------+
-- | SecondHighestSalary |
-- +---------------------+
-- | 200                 |
-- +---------------------+
-- 
-- 
-- 示例 2：
-- 
-- 
-- 输入：
-- Employee 表：
-- +----+--------+
-- | id | salary |
-- +----+--------+
-- | 1  | 100    |
-- +----+--------+
-- 输出：
-- +---------------------+
-- | SecondHighestSalary |
-- +---------------------+
-- | null                |
-- +---------------------+
-- 
-- 
-- 
-- 
--

-- @lc code=start
# Write your MySQL query statement below
-- select (
--     select salary 
--     from ( 
--         select distinct(salary), dense_rank() over(ORDER BY salary DESC) AS rnk
--         from Employee
--     ) t
--     where t.rnk=2
-- ) as SecondHighestSalary 




-- 作者：LeetCode
-- 链接：https://leetcode.cn/problems/second-highest-salary/solution/di-er-gao-de-xin-shui-by-leetcode/
-- 来源：力扣（LeetCode）
-- 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
-- 将不同的薪资按降序排序，然后使用 LIMIT 子句获得第二高的薪资。
-- SELECT DISTINCT
--     Salary AS SecondHighestSalary
-- FROM
--     Employee
-- ORDER BY Salary DESC
-- LIMIT 1 OFFSET 1
-- 然而，如果没有这样的第二最高工资，这个解决方案将被判断为 “错误答案”，
-- 因为本表可能只有一项记录。为了克服这个问题，我们可以将其作为临时表。
SELECT
    (SELECT DISTINCT
            Salary
        FROM
            Employee
        ORDER BY Salary DESC
        LIMIT 1 OFFSET 1) AS SecondHighestSalary

-- 

-- @lc code=end

