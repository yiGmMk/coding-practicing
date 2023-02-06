--
-- @lc app=leetcode.cn id=1148 lang=mysql
--
-- [1148] 文章浏览 I
--
-- https://leetcode.cn/problems/article-views-i/description/
--
-- database
-- Easy (71.88%)
-- Likes:    56
-- Dislikes: 0
-- Total Accepted:    41.1K
-- Total Submissions: 57.1K
-- Testcase Example:  '{"headers":{"Views":["article_id","author_id","viewer_id","view_date"]},"rows":{"Views":[[1,3,5,"2019-08-01"],[1,3,6,"2019-08-02"],[2,7,7,"2019-08-01"],[2,7,6,"2019-08-02"],[4,7,1,"2019-07-22"],[3,4,4,"2019-07-21"],[3,4,4,"2019-07-21"]]}}'
--
-- Views 表：
-- 
-- 
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | article_id    | int     |
-- | author_id     | int     |
-- | viewer_id     | int     |
-- | view_date     | date    |
-- +---------------+---------+
-- 此表无主键，因此可能会存在重复行。
-- 此表的每一行都表示某人在某天浏览了某位作者的某篇文章。
-- 请注意，同一人的 author_id 和 viewer_id 是相同的。
-- 
-- 
-- 
-- 
-- 请编写一条 SQL 查询以找出所有浏览过自己文章的作者，结果按照 id 升序排列。
-- 
-- 查询结果的格式如下所示：
-- 
-- 
-- Views 表：
-- +------------+-----------+-----------+------------+
-- | article_id | author_id | viewer_id | view_date  |
-- +------------+-----------+-----------+------------+
-- | 1          | 3         | 5         | 2019-08-01 |
-- | 1          | 3         | 6         | 2019-08-02 |
-- | 2          | 7         | 7         | 2019-08-01 |
-- | 2          | 7         | 6         | 2019-08-02 |
-- | 4          | 7         | 1         | 2019-07-22 |
-- | 3          | 4         | 4         | 2019-07-21 |
-- | 3          | 4         | 4         | 2019-07-21 |
-- +------------+-----------+-----------+------------+
-- 
-- 结果表：
-- +------+
-- | id   |
-- +------+
-- | 4    |
-- | 7    |
-- +------+
-- 
-- 
--

-- @lc code=start
# Write your MySQL query statement below
-- select distinct v1.author_id as id
-- from Views v1
-- left join Views v2 on v1.article_id=v2.article_id 
-- where v1.author_id=v2.author_id and v1.author_id=v2.viewer_id
-- order by id


select distinct v1.author_id as id
from Views v1,Views v2
where v1.author_id=v2.author_id and v1.author_id=v2.viewer_id
order by id

-- @lc code=end

