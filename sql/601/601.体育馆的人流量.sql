--
-- @lc app=leetcode.cn id=601 lang=mysql
--
-- [601] 体育馆的人流量
--
-- https://leetcode.cn/problems/human-traffic-of-stadium/description/
--
-- database
-- Hard (47.66%)
-- Likes:    282
-- Dislikes: 0
-- Total Accepted:    43.6K
-- Total Submissions: 91.5K
-- Testcase Example:  '{"headers": {"Stadium": ["id", "visit_date", "people"]}, "rows": {"Stadium": [[1, "2017-01-01", 10], [2, "2017-01-02", 109], [3, "2017-01-03", 150], [4, "2017-01-04", 99], [5, "2017-01-05", 145], [6, "2017-01-06", 1455], [7, "2017-01-07", 199], [8, "2017-01-09", 188]]}}'
--
-- 表：Stadium
-- 
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | id            | int     |
-- | visit_date    | date    |
-- | people        | int     |
-- +---------------+---------+
-- visit_date 是表的主键
-- 每日人流量信息被记录在这三列信息中：序号 (id)、日期 (visit_date)、 人流量 (people)
-- 每天只有一行记录，日期随着 id 的增加而增加
-- 
-- 
-- 
-- 
-- 编写一个 SQL 查询以找出每行的人数大于或等于 100 且 id 连续的三行或更多行记录。
-- 
-- 返回按 visit_date 升序排列 的结果表。
-- 
-- 查询结果格式如下所示。
-- 
-- 
-- 
-- 示例 1:
-- 
-- 
-- 输入：
-- Stadium 表:
-- +------+------------+-----------+
-- | id   | visit_date | people    |
-- +------+------------+-----------+
-- | 1    | 2017-01-01 | 10        |
-- | 2    | 2017-01-02 | 109       |
-- | 3    | 2017-01-03 | 150       |
-- | 4    | 2017-01-04 | 99        |
-- | 5    | 2017-01-05 | 145       |
-- | 6    | 2017-01-06 | 1455      |
-- | 7    | 2017-01-07 | 199       |
-- | 8    | 2017-01-09 | 188       |
-- +------+------------+-----------+
-- 输出：
-- +------+------------+-----------+
-- | id   | visit_date | people    |
-- +------+------------+-----------+
-- | 5    | 2017-01-05 | 145       |
-- | 6    | 2017-01-06 | 1455      |
-- | 7    | 2017-01-07 | 199       |
-- | 8    | 2017-01-09 | 188       |
-- +------+------------+-----------+
-- 解释：
-- id 为 5、6、7、8 的四行 id 连续，并且每行都有 >= 100 的人数记录。
-- 请注意，即使第 7 行和第 8 行的 visit_date 不是连续的，输出也应当包含第 8 行，因为我们只需要考虑 id 连续的记录。
-- 不输出 id 为 2 和 3 的行，因为至少需要三条 id 连续的记录。
-- 
--

-- @lc code=start
# Write your MySQL query statement below

-- 作者：hejy-w
-- 链接：https://leetcode.cn/problems/human-traffic-of-stadium/solution/tu-jie-lian-xu-ri-qi-ji-nan-dian-fen-xi-xnj58/
-- 来源：力扣（LeetCode）
-- 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
with t1 as(
    select *,id - row_number() over(order by id) as rk
    from stadium
    where people >= 100
)

select id,visit_date,people
from t1
where rk in(
    select rk
    from t1
    group by rk
    having count(rk) >= 3
);


-- 作者：LeetCode
-- 链接：https://leetcode.cn/problems/human-traffic-of-stadium/solution/ti-yu-guan-de-ren-liu-liang-by-leetcode/
-- 来源：力扣（LeetCode）
-- 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
-- 表 t1，t2 和 t3 相同，需要考虑添加哪些条件能够得到想要的结果。
-- 以 t1 为例，它有可能是高峰期的第 1 天，第 2 天，或第 3 天。
select distinct t1.*
from stadium t1, stadium t2, stadium t3
where t1.people >= 100 and t2.people >= 100 and t3.people >= 100
and
(
	(t1.id - t2.id = 1 and t1.id - t3.id = 2 and t2.id - t3.id =1)  -- t1, t2, t3
    or
    (t2.id - t1.id = 1 and t2.id - t3.id = 2 and t1.id - t3.id =1) -- t2, t1, t3
    or
    (t3.id - t2.id = 1 and t2.id - t1.id =1 and t3.id - t1.id = 2) -- t3, t2, t1
)
order by t1.id
;


-- @lc code=end

