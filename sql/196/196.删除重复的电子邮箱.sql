--
-- @lc app=leetcode.cn id=196 lang=mysql
--
-- [196] 删除重复的电子邮箱
--
-- https://leetcode.cn/problems/delete-duplicate-emails/description/
--
-- database
-- Easy (68.62%)
-- Likes:    680
-- Dislikes: 0
-- Total Accepted:    190.7K
-- Total Submissions: 277.8K
-- Testcase Example:  '{"headers": {"Person": ["id", "email"]}, "rows": {"Person": [[1, "john@example.com"], [2, "bob@example.com"], [3, "john@example.com"]]}}'
--
-- 表: Person
-- 
-- 
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | email       | varchar |
-- +-------------+---------+
-- id是该表的主键列。
-- 该表的每一行包含一封电子邮件。电子邮件将不包含大写字母。
-- 
-- 
-- 
-- 
-- 编写一个 SQL 删除语句来 删除 所有重复的电子邮件，只保留一个id最小的唯一电子邮件。
-- 
-- 以 任意顺序 返回结果表。 （注意： 仅需要写删除语句，将自动对剩余结果进行查询）
-- 
-- 查询结果格式如下所示。
-- 
-- 
-- 
-- 
-- 
-- 示例 1:
-- 
-- 
-- 输入: 
-- Person 表:
-- +----+------------------+
-- | id | email            |
-- +----+------------------+
-- | 1  | john@example.com |
-- | 2  | bob@example.com  |
-- | 3  | john@example.com |
-- +----+------------------+
-- 输出: 
-- +----+------------------+
-- | id | email            |
-- +----+------------------+
-- | 1  | john@example.com |
-- | 2  | bob@example.com  |
-- +----+------------------+
-- 解释: john@example.com重复两次。我们保留最小的Id = 1。
-- 
--

-- @lc code=start
# Please write a DELETE statement and DO NOT write a SELECT statement.
# Write your MySQL query statement below

-- 找出需要删除的记录
-- select p1.*
-- from Person p1,Person p2
-- where p1.email=p2.email and p2.id>p1.id;



-- TODO
-- 删除
-- 作者：LeetCode
-- 链接：https://leetcode.cn/problems/delete-duplicate-emails/solution/shan-chu-zhong-fu-de-dian-zi-you-xiang-by-leetcode/
-- 来源：力扣（LeetCode）
-- 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
DELETE p1 FROM Person p1,
    Person p2
WHERE
    p1.Email = p2.Email AND p1.Id > p2.Id

--  
-- @lc code=end

