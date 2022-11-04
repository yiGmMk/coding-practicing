--
-- @lc app=leetcode.cn id=1407 lang=mysql
--
-- [1407] 排名靠前的旅行者
--
-- https://leetcode.cn/problems/top-travellers/description/
--
-- database
-- Easy (59.63%)
-- Likes:    36
-- Dislikes: 0
-- Total Accepted:    23.4K
-- Total Submissions: 39.3K
-- Testcase Example:  '{"headers":{"Users":["id","name"],"Rides":["id","user_id","distance"]},"rows":{"Users":[[1,"Alice"],[2,"Bob"],[3,"Alex"],[4,"Donald"],[7,"Lee"],[13,"Jonathan"],[19,"Elvis"]],"Rides":[[1,1,120],[2,2,317],[3,3,222],[4,7,100],[5,13,312],[6,19,50],[7,7,120],[8,19,400],[9,7,230]]}}'
--
-- 表：Users
-- 
-- 
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | name          | varchar |
-- +---------------+---------+
-- id 是该表单主键。
-- name 是用户名字。
-- 
-- 
-- 
-- 表：Rides
-- 
-- 
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | user_id       | int     |
-- | distance      | int     |
-- +---------------+---------+
-- id 是该表单主键。
-- user_id 是本次行程的用户的 id, 而该用户此次行程距离为 distance 。
-- 
-- 
-- 
-- 
-- 写一段 SQL , 报告每个用户的旅行距离。
-- 
-- 返回的结果表单，以 travelled_distance 降序排列 ，如果有两个或者更多的用户旅行了相同的距离, 那么再以 name 升序排列 。
-- 
-- 查询结果格式如下例所示。
-- 
-- 
-- Users 表：
-- +------+-----------+
-- | id   | name      |
-- +------+-----------+
-- | 1    | Alice     |
-- | 2    | Bob       |
-- | 3    | Alex      |
-- | 4    | Donald    |
-- | 7    | Lee       |
-- | 13   | Jonathan  |
-- | 19   | Elvis     |
-- +------+-----------+
-- 
-- Rides 表：
-- +------+----------+----------+
-- | id   | user_id  | distance |
-- +------+----------+----------+
-- | 1    | 1        | 120      |
-- | 2    | 2        | 317      |
-- | 3    | 3        | 222      |
-- | 4    | 7        | 100      |
-- | 5    | 13       | 312      |
-- | 6    | 19       | 50       |
-- | 7    | 7        | 120      |
-- | 8    | 19       | 400      |
-- | 9    | 7        | 230      |
-- +------+----------+----------+
-- 
-- Result 表：
-- +----------+--------------------+
-- | name     | travelled_distance |
-- +----------+--------------------+
-- | Elvis    | 450                |
-- | Lee      | 450                |
-- | Bob      | 317                |
-- | Jonathan | 312                |
-- | Alex     | 222                |
-- | Alice    | 120                |
-- | Donald   | 0                  |
-- +----------+--------------------+
-- Elvis 和 Lee 旅行了 450 英里，Elvis 是排名靠前的旅行者，因为他的名字在字母表上的排序比 Lee 更小。
-- Bob, Jonathan, Alex 和 Alice 只有一次行程，我们只按此次行程的全部距离对他们排序。
-- Donald 没有任何行程, 他的旅行距离为 0。
-- 
-- 
--

-- @lc code=start

-- with t as (
-- select user_id,sum(distance) as distance
-- from Rides
-- group by user_id
-- )
-- select Users.name,if( distance>0 ,distance,0) as travelled_distance
-- from Users left join t on Users.id=t.user_id
-- order by travelled_distance desc,Users.name



-- 作者：int_64
-- 链接：https://leetcode.cn/problems/top-travellers/solution/by-jam007-htxi/
-- 来源：力扣（LeetCode）
-- 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
-- https://dev.mysql.com/doc/refman/8.0/en/join.html
-- COALESCE(x, y) = (CASE WHEN x IS NOT NULL THEN x ELSE y END)
select 
    name, coalesce(sum(distance), 0) travelled_distance
from 
    users u
left join rides r on u.id=r.user_id
group by u.id
order by travelled_distance desc, name;

-- @lc code=end

