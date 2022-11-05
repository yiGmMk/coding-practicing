--
-- @lc app=leetcode.cn id=1141 lang=mysql
--
-- [1141] 查询近30天活跃用户数
--
-- https://leetcode.cn/problems/user-activity-for-the-past-30-days-i/description/
--
-- database
-- Easy (45.75%)
-- Likes:    65
-- Dislikes: 0
-- Total Accepted:    35.1K
-- Total Submissions: 76.6K
-- Testcase Example:  '{"headers":{"Activity":["user_id","session_id","activity_date","activity_type"]},"rows":{"Activity":[[1,1,"2019-07-20","open_session"],[1,1,"2019-07-20","scroll_down"],[1,1,"2019-07-20","end_session"],[2,4,"2019-07-20","open_session"],[2,4,"2019-07-21","send_message"],[2,4,"2019-07-21","end_session"],[3,2,"2019-07-21","open_session"],[3,2,"2019-07-21","send_message"],[3,2,"2019-07-21","end_session"],[4,3,"2019-06-25","open_session"],[4,3,"2019-06-25","end_session"]]}}'
--
-- 活动记录表：Activity
-- 
-- 
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | user_id       | int     |
-- | session_id    | int     |
-- | activity_date | date    |
-- | activity_type | enum    |
-- +---------------+---------+
-- 该表是用户在社交网站的活动记录。
-- 该表没有主键，可能包含重复数据。
-- activity_type 字段为以下四种值 ('open_session', 'end_session', 'scroll_down',
-- 'send_message')。
-- 每个 session_id 只属于一个用户。
-- 
-- 
-- 
-- 
-- 请写SQL查询出截至 2019-07-27（包含2019-07-27），近 30 天的每日活跃用户数（当天只要有一条活动记录，即为活跃用户）。
-- 
-- 以 任意顺序 返回结果表。
-- 
-- 查询结果示例如下。
-- 
-- 
-- 
-- 示例 1:
-- 
-- 
-- 输入：
-- Activity table:
-- +---------+------------+---------------+---------------+
-- | user_id | session_id | activity_date | activity_type |
-- +---------+------------+---------------+---------------+
-- | 1       | 1          | 2019-07-20    | open_session  |
-- | 1       | 1          | 2019-07-20    | scroll_down   |
-- | 1       | 1          | 2019-07-20    | end_session   |
-- | 2       | 4          | 2019-07-20    | open_session  |
-- | 2       | 4          | 2019-07-21    | send_message  |
-- | 2       | 4          | 2019-07-21    | end_session   |
-- | 3       | 2          | 2019-07-21    | open_session  |
-- | 3       | 2          | 2019-07-21    | send_message  |
-- | 3       | 2          | 2019-07-21    | end_session   |
-- | 4       | 3          | 2019-06-25    | open_session  |
-- | 4       | 3          | 2019-06-25    | end_session   |
-- +---------+------------+---------------+---------------+
-- 输出：
-- +------------+--------------+ 
-- | day        | active_users |
-- +------------+--------------+ 
-- | 2019-07-20 | 2            |
-- | 2019-07-21 | 2            |
-- +------------+--------------+ 
-- 解释：注意非活跃用户的记录不需要展示。
-- 
--

-- @lc code=start
# Write your MySQL query statement below

-- select activity_date as day,count(distinct user_id) as active_users
-- from activity
-- where activity_date between '2019-06-28' and '2019-07-28'
-- group by day
-- having count(distinct user_id)>0


# Write your MySQL query statement below
select activity_date day, count(distinct user_id) active_users
from activity
where datediff('2019-07-27', activity_date) >= 0 AND datediff('2019-07-27', activity_date) <30
group by activity_date
-- 作者：mythicmyuu
-- 链接：https://leetcode.cn/problems/user-activity-for-the-past-30-days-i/solution/by-mythicmyuu-mlj9/
-- 来源：力扣（LeetCode）
-- 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
-- @lc code=end

