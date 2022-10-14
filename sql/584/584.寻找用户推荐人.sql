--
-- @lc app=leetcode.cn id=584 lang=mysql
--
-- [584] 寻找用户推荐人
--
-- https://leetcode.cn/problems/find-customer-referee/description/
--
-- database
-- Easy (68.56%)
-- Likes:    106
-- Dislikes: 0
-- Total Accepted:    93K
-- Total Submissions: 135.8K
-- Testcase Example:  '{"headers":{"Customer":["id","name","referee_id"]},"rows":{"Customer":[[1,"Will",null],[2,"Jane",null],[3,"Alex",2],[4,"Bill",null],[5,"Zack",1],[6,"Mark",2]]}}'
--
-- 给定表 customer ，里面保存了所有客户信息和他们的推荐人。
-- 
-- 
-- +------+------+-----------+
-- | id   | name | referee_id|
-- +------+------+-----------+
-- |    1 | Will |      NULL |
-- |    2 | Jane |      NULL |
-- |    3 | Alex |         2 |
-- |    4 | Bill |      NULL |
-- |    5 | Zack |         1 |
-- |    6 | Mark |         2 |
-- +------+------+-----------+
-- 
-- 
-- 写一个查询语句，返回一个客户列表，列表中客户的推荐人的编号都 不是 2。
-- 
-- 对于上面的示例数据，结果为：
-- 
-- 
-- +------+
-- | name |
-- +------+
-- | Will |
-- | Jane |
-- | Bill |
-- | Zack |
-- +------+
-- 
-- 
--

-- @lc code=start
# Write your MySQL query statement below
-- select name 
-- from
-- (select distinct id,name
-- from customer
-- where referee_id!=2 or referee_id is NULL) t


-- 
select name from customer where referee_id != 2 or referee_id is null

-- @lc code=end

