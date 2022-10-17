--
-- @lc app=leetcode.cn id=1873 lang=mysql
--
-- [1873] 计算特殊奖金
--
-- https://leetcode.cn/problems/calculate-special-bonus/description/
--
-- database
-- Easy (64.80%)
-- Likes:    95
-- Dislikes: 0
-- Total Accepted:    49.6K
-- Total Submissions: 76.6K
-- Testcase Example:  '{"headers":{"Employees":["employee_id","name","salary"]},"rows":{"Employees":[[2,"Meir",3000],[3,"Michael",3800],[7,"Addilyn",7400],[8,"Juan",6100],[9,"Kannon",7700]]}}'
--
-- 表: Employees
-- 
-- 
-- +-------------+---------+
-- | 列名        | 类型     |
-- +-------------+---------+
-- | employee_id | int     |
-- | name        | varchar |
-- | salary      | int     |
-- +-------------+---------+
-- employee_id 是这个表的主键。
-- 此表的每一行给出了雇员id ，名字和薪水。
-- 
-- 
-- 
-- 
-- 写出一个SQL 查询语句，计算每个雇员的奖金。如果一个雇员的id是奇数并且他的名字不是以'M'开头，那么他的奖金是他工资的100%，否则奖金为0。
-- 
-- Return the result table ordered by employee_id.
-- 
-- 返回的结果集请按照employee_id排序。
-- 
-- 查询结果格式如下面的例子所示。
-- 
-- 
-- 
-- 示例 1:
-- 
-- 
-- 输入：
-- Employees 表:
-- +-------------+---------+--------+
-- | employee_id | name    | salary |
-- +-------------+---------+--------+
-- | 2           | Meir    | 3000   |
-- | 3           | Michael | 3800   |
-- | 7           | Addilyn | 7400   |
-- | 8           | Juan    | 6100   |
-- | 9           | Kannon  | 7700   |
-- +-------------+---------+--------+
-- 输出：
-- +-------------+-------+
-- | employee_id | bonus |
-- +-------------+-------+
-- | 2           | 0     |
-- | 3           | 0     |
-- | 7           | 7400  |
-- | 8           | 0     |
-- | 9           | 7700  |
-- +-------------+-------+
-- 解释：
-- 因为雇员id是偶数，所以雇员id 是2和8的两个雇员得到的奖金是0。
-- 雇员id为3的因为他的名字以'M'开头，所以，奖金是0。
-- 其他的雇员得到了百分之百的奖金。
-- 
--

-- @lc code=start
# Write your MySQL query statement below
-- select employee_id,(
--     CASE WHEN MOD(employee_id,2)!=0 and left(name,1)!='M' THEN salary
--     ELSE 0
--     END) as bonus
-- from Employees
-- order by employee_id


# Write your MySQL query statement below
# employee_id bonus

select employee_id,if(employee_id % 2 != 0 and left(name,1) != 'M',salary,0) as bonus
from Employees
order by employee_id



-- 作者：lBEI5874xF
-- 链接：https://leetcode.cn/problems/calculate-special-bonus/solution/kao-cha-by-lbei5874xf-ahpo/
-- 来源：力扣（LeetCode）
-- 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
-- @lc code=end

