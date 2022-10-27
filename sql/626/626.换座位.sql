--
-- @lc app=leetcode.cn id=626 lang=mysql
--
-- [626] 换座位
--
-- https://leetcode.cn/problems/exchange-seats/description/
--
-- database
-- Medium (68.52%)
-- Likes:    344
-- Dislikes: 0
-- Total Accepted:    63.6K
-- Total Submissions: 92.8K
-- Testcase Example:  '{"headers": {"Seat": ["id","student"]}, "rows": {"Seat": [[1,"Abbot"],[2,"Doris"],[3,"Emerson"],[4,"Green"],[5,"Jeames"]]}}'
--
-- 表: Seat
-- 
-- 
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | name        | varchar |
-- +-------------+---------+
-- Id是该表的主键列。
-- 该表的每一行都表示学生的姓名和ID。
-- Id是一个连续的增量。
-- 
-- 
-- 
-- 
-- 编写SQL查询来交换每两个连续的学生的座位号。如果学生的数量是奇数，则最后一个学生的id不交换。
-- 
-- 按 id 升序 返回结果表。
-- 
-- 查询结果格式如下所示。
-- 
-- 
-- 
-- 示例 1:
-- 
-- 
-- 输入: 
-- Seat 表:
-- +----+---------+
-- | id | student |
-- +----+---------+
-- | 1  | Abbot   |
-- | 2  | Doris   |
-- | 3  | Emerson |
-- | 4  | Green   |
-- | 5  | Jeames  |
-- +----+---------+
-- 输出: 
-- +----+---------+
-- | id | student |
-- +----+---------+
-- | 1  | Doris   |
-- | 2  | Abbot   |
-- | 3  | Green   |
-- | 4  | Emerson |
-- | 5  | Jeames  |
-- +----+---------+
-- 解释:
-- 请注意，如果学生人数为奇数，则不需要更换最后一名学生的座位。
-- 
--

-- @lc code=start
# Write your MySQL query statement below

-- 作者：int_64
-- 链接：https://leetcode.cn/problems/exchange-seats/solution/by-jam007-2sy3/
-- 来源：力扣（LeetCode）
-- 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


-- SELECT 
--     id,
--     IF(id % 2 = 0, last, next) student
-- FROM (
--     SELECT 
--     id,student,
--     lag(student,1,student) over(order by id) last,
--     lead(student,1,student) over(order by id) next
--     FROM seat
-- ) t;

SELECT (
    case when id%2!=0  and counts!=id then id+1 -- 奇数
         when id%2!=0 and counts=id then id -- 奇数
         else id-1 -- 偶数
    end
)as id,student
FROM seat,(
    SELECT count(*) as counts FROM seat 
)as seatnum
order by id asc

-- @lc code=end

