--
-- @lc app=leetcode.cn id=262 lang=mysql
--
-- [262] 行程和用户
--
-- https://leetcode.cn/problems/trips-and-users/description/
--
-- database
-- Hard (42.23%)
-- Likes:    323
-- Dislikes: 0
-- Total Accepted:    56.1K
-- Total Submissions: 132.8K
-- Testcase Example:  '{"headers": {"Trips": ["id", "client_id", "driver_id", "city_id", "status", "request_at"], "Users": ["users_id", "banned", "role"]}, "rows": {"Trips": [["1", "1", "10", "1", "completed", "2013-10-01"], ["2", "2", "11", "1", "cancelled_by_driver", "2013-10-01"], ["3", "3", "12", "6", "completed", "2013-10-01"], ["4", "4", "13", "6", "cancelled_by_client", "2013-10-01"], ["5", "1", "10", "1", "completed", "2013-10-02"], ["6", "2", "11", "6", "completed", "2013-10-02"], ["7", "3", "12", "6", "completed", "2013-10-02"], ["8", "2", "12", "12", "completed", "2013-10-03"], ["9", "3", "10", "12", "completed", "2013-10-03"], ["10", "4", "13", "12", "cancelled_by_driver", "2013-10-03"]], "Users": [["1", "No", "client"], ["2", "Yes", "client"], ["3", "No", "client"], ["4", "No", "client"], ["10", "No", "driver"], ["11", "No", "driver"], ["12", "No", "driver"], ["13", "No", "driver"]]}}'
--
-- 表：Trips
-- 
-- 
-- 
-- +-------------+----------+
-- | Column Name | Type     |
-- +-------------+----------+
-- | id          | int      |
-- | client_id   | int      |
-- | driver_id   | int      |
-- | city_id     | int      |
-- | status      | enum     |
-- | request_at  | date     |     
-- +-------------+----------+
-- id 是这张表的主键。
-- 这张表中存所有出租车的行程信息。每段行程有唯一 id ，其中 client_id 和 driver_id 是 Users 表中 users_id
-- 的外键。
-- status 是一个表示行程状态的枚举类型，枚举成员为(‘completed’, ‘cancelled_by_driver’,
-- ‘cancelled_by_client’) 。
-- 
-- 
-- 
-- 
-- 
-- 
-- 表：Users
-- 
-- 
-- 
-- 
-- +-------------+----------+
-- | Column Name | Type     |
-- +-------------+----------+
-- | users_id    | int      |
-- | banned      | enum     |
-- | role        | enum     |
-- +-------------+----------+
-- users_id 是这张表的主键。
-- 这张表中存所有用户，每个用户都有一个唯一的 users_id ，role 是一个表示用户身份的枚举类型，枚举成员为 (‘client’,
-- ‘driver’, ‘partner’) 。
-- banned 是一个表示用户是否被禁止的枚举类型，枚举成员为 (‘Yes’, ‘No’) 。
-- 
-- 
-- 
-- 
-- 取消率 的计算方式如下：(被司机或乘客取消的非禁止用户生成的订单数量) / (非禁止用户生成的订单总数)。
-- 
-- 写一段 SQL 语句查出 "2013-10-01" 至 "2013-10-03" 期间非禁止用户（乘客和司机都必须未被禁止）的取消率。非禁止用户即
-- banned 为 No 的用户，禁止用户即 banned 为 Yes 的用户。
-- 
-- 返回结果表中的数据可以按任意顺序组织。其中取消率 Cancellation Rate 需要四舍五入保留 两位小数 。
-- 
-- 查询结果格式如下例所示。
-- 
-- 
-- 
-- 示例：
-- 
-- 
-- 输入： 
-- Trips 表：
-- +----+-----------+-----------+---------+---------------------+------------+
-- | id | client_id | driver_id | city_id | status              | request_at |
-- +----+-----------+-----------+---------+---------------------+------------+
-- | 1  | 1         | 10        | 1       | completed           | 2013-10-01 |
-- | 2  | 2         | 11        | 1       | cancelled_by_driver | 2013-10-01 |
-- | 3  | 3         | 12        | 6       | completed           | 2013-10-01 |
-- | 4  | 4         | 13        | 6       | cancelled_by_client | 2013-10-01 |
-- | 5  | 1         | 10        | 1       | completed           | 2013-10-02 |
-- | 6  | 2         | 11        | 6       | completed           | 2013-10-02 |
-- | 7  | 3         | 12        | 6       | completed           | 2013-10-02 |
-- | 8  | 2         | 12        | 12      | completed           | 2013-10-03 |
-- | 9  | 3         | 10        | 12      | completed           | 2013-10-03 |
-- | 10 | 4         | 13        | 12      | cancelled_by_driver | 2013-10-03 |
-- +----+-----------+-----------+---------+---------------------+------------+
-- 
-- Users 表：
-- +----------+--------+--------+
-- | users_id | banned | role   |
-- +----------+--------+--------+
-- | 1        | No     | client |
-- | 2        | Yes    | client |
-- | 3        | No     | client |
-- | 4        | No     | client |
-- | 10       | No     | driver |
-- | 11       | No     | driver |
-- | 12       | No     | driver |
-- | 13       | No     | driver |
-- +----------+--------+--------+
-- 输出：
-- +------------+-------------------+
-- | Day        | Cancellation Rate |
-- +------------+-------------------+
-- | 2013-10-01 | 0.33              |
-- | 2013-10-02 | 0.00              |
-- | 2013-10-03 | 0.50              |
-- +------------+-------------------+
-- 解释：
-- 2013-10-01：
-- ⁠ - 共有 4 条请求，其中 2 条取消。
-- ⁠ - 然而，id=2 的请求是由禁止用户（user_id=2）发出的，所以计算时应当忽略它。
-- ⁠ - 因此，总共有 3 条非禁止请求参与计算，其中 1 条取消。
-- ⁠ - 取消率为 (1 / 3) = 0.33
-- 2013-10-02：
-- ⁠ - 共有 3 条请求，其中 0 条取消。
-- ⁠ - 然而，id=6 的请求是由禁止用户发出的，所以计算时应当忽略它。
-- ⁠ - 因此，总共有 2 条非禁止请求参与计算，其中 0 条取消。
-- ⁠ - 取消率为 (0 / 2) = 0.00
-- 2013-10-03：
-- ⁠ - 共有 3 条请求，其中 1 条取消。
-- ⁠ - 然而，id=8 的请求是由禁止用户发出的，所以计算时应当忽略它。
-- ⁠ - 因此，总共有 2 条非禁止请求参与计算，其中 1 条取消。
-- ⁠ - 取消率为 (1 / 2) = 0.50
-- 
-- 
-- 
-- 
--

-- @lc code=start
# Write your MySQL query statement below
select trips.request_at as "Day",round(
    SUM(
        IF(trips.status='completed',0,1)
    )
    /
    count(trips.status),2) as "Cancellation Rate" 
from trips
join users as d on trips.driver_id=d.users_id and d.banned="No"
join users as u on trips.client_id=u.users_id and u.banned="No"
where trips.request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP by trips.request_at;

-- SELECT T.request_at AS `Day`, 
-- 	ROUND(
-- 			SUM(
-- 				IF(T.STATUS = 'completed',0,1)
-- 			)
-- 			/ 
-- 			COUNT(T.STATUS),
-- 			2
-- 	) AS `Cancellation Rate`
-- FROM Trips AS T
-- JOIN Users AS U1 ON (T.client_id = U1.users_id AND U1.banned ='No')
-- JOIN Users AS U2 ON (T.driver_id = U2.users_id AND U2.banned ='No')
-- WHERE T.request_at BETWEEN '2013-10-01' AND '2013-10-03'
-- GROUP BY T.request_at


#剔除掉被禁止的用户
SELECT T.request_at AS `Day`, 
	ROUND(
			SUM(
				IF(T.STATUS = 'completed',0,1)
			)
			/ 
			COUNT(T.STATUS),
			2
	) AS `Cancellation Rate`
FROM trips AS T
WHERE 
T.Client_Id NOT IN (
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
)
AND
T.Driver_Id NOT IN (
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
)
AND T.request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY T.request_at

-- 作者：jason-2
-- 链接：https://leetcode.cn/problems/trips-and-users/solution/san-chong-jie-fa-cong-nan-dao-yi-zong-you-gua-he-n/
-- 来源：力扣（LeetCode）
-- 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
-- @lc code=end

