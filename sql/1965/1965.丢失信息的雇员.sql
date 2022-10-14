--
-- @lc app=leetcode.cn id=1965 lang=mysql
--
-- [1965] 丢失信息的雇员
--
-- https://leetcode.cn/problems/employees-with-missing-information/description/
--
-- database
-- Easy (72.41%)
-- Likes:    51
-- Dislikes: 0
-- Total Accepted:    27.7K
-- Total Submissions: 38.3K
-- Testcase Example:  '{"headers":{"Employees":["employee_id","name"],"Salaries":["employee_id","salary"]},"rows":{"Employees":[[2,"Crew"],[4,"Haven"],[5,"Kristian"]],"Salaries":[[5,76071],[1,22517],[4,63539]]}}'
--
-- 表: Employees
-- 
-- 
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | employee_id | int     |
-- | name        | varchar |
-- +-------------+---------+
-- employee_id 是这个表的主键。
-- 每一行表示雇员的id 和他的姓名。
-- 
-- 
-- 表: Salaries
-- 
-- 
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | employee_id | int     |
-- | salary      | int     |
-- +-------------+---------+
-- employee_id is 这个表的主键。
-- 每一行表示雇员的id 和他的薪水。
-- 
-- 
-- 
-- 
-- 写出一个查询语句，找到所有 丢失信息 的雇员id。当满足下面一个条件时，就被认为是雇员的信息丢失：
-- 
-- 
-- 雇员的 姓名 丢失了，或者
-- 雇员的 薪水信息 丢失了，或者
-- 
-- 
-- 返回这些雇员的id  employee_id ， 从小到大排序 。
-- 
-- 查询结果格式如下面的例子所示。
-- 
-- 
-- 
-- 示例 1：
-- 
-- 
-- 输入：
-- Employees table:
-- +-------------+----------+
-- | employee_id | name     |
-- +-------------+----------+
-- | 2           | Crew     |
-- | 4           | Haven    |
-- | 5           | Kristian |
-- +-------------+----------+
-- Salaries table:
-- +-------------+--------+
-- | employee_id | salary |
-- +-------------+--------+
-- | 5           | 76071  |
-- | 1           | 22517  |
-- | 4           | 63539  |
-- +-------------+--------+
-- 输出：
-- +-------------+
-- | employee_id |
-- +-------------+
-- | 1           |
-- | 2           |
-- +-------------+
-- 解释：
-- 雇员1，2，4，5 都工作在这个公司。
-- 1号雇员的姓名丢失了。
-- 2号雇员的薪水信息丢失了。
-- 
--

-- @lc code=start
# Write your MySQL query statement below

-- 并集(缺失姓名)+缺失薪水的
select Salaries.employee_id
from Salaries
left join Employees
on employees.employee_id = Salaries.employee_id
where Employees.employee_id is null
union
select Employees.employee_id
from Employees
left join Salaries
on employees.employee_id = Salaries.employee_id
where Salaries.employee_id is null
order by employee_id

-- 作者：fqy-66
-- 链接：https://leetcode.cn/problems/employees-with-missing-information/solution/by-fqy-66-cac3/
-- 来源：力扣（LeetCode）
-- 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

-- 利用union all 计算employee_id的数量
select employee_id from (
    select employee_id from Employees
    union all
    select employee_id from Salaries
 ) as ans
 group by employee_id
 having count(employee_id) = 1
 order by employee_id;
-- 作者：serendipity-i6a
-- 链接：https://leetcode.cn/problems/employees-with-missing-information/solution/by-serendipity-i6a-e0cv/
-- 来源：力扣（LeetCode）
-- 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
-- @lc code=end

