--
-- @lc app=leetcode.cn id=1587 lang=mysql
--
-- [1587] 银行账户概要 II
--
-- https://leetcode.cn/problems/bank-account-summary-ii/description/
--
-- database
-- Easy (83.57%)
-- Likes:    17
-- Dislikes: 0
-- Total Accepted:    18.2K
-- Total Submissions: 21.7K
-- Testcase Example:  '{"headers": {"Users": ["account", "name"], "Transactions": ["trans_id", "account", "amount", "transacted_on"]}, "rows": {"Users": [[900001, "Alice"], [900002, "Bob"], [900003, "Charlie"]], "Transactions": [[1, 900001, 7000, "2020-08-01"], [2, 900001, 7000, "2020-09-01"], [3, 900001, -3000, "2020-09-02"], [4, 900002, 1000, "2020-09-12"], [5, 900003, 6000, "2020-08-07"], [6, 900003, 6000, "2020-09-07"], [7, 900003, -4000, "2020-09-11"]]}}'
--
-- 表: Users
-- 
-- +--------------+---------+
-- | Column Name  | Type    |
-- +--------------+---------+
-- | account      | int     |
-- | name         | varchar |
-- +--------------+---------+
-- account 是该表的主键.
-- 表中的每一行包含银行里中每一个用户的账号.
-- 
-- 
-- 
-- 
-- 表: Transactions
-- 
-- +---------------+---------+
-- | Column Name   | Type    |
-- +---------------+---------+
-- | trans_id      | int     |
-- | account       | int     |
-- | amount        | int     |
-- | transacted_on | date    |
-- +---------------+---------+
-- trans_id 是该表主键.
-- 该表的每一行包含了所有账户的交易改变情况.
-- 如果用户收到了钱, 那么金额是正的; 如果用户转了钱, 那么金额是负的.
-- 所有账户的起始余额为 0.
-- 
-- 
-- 
-- 
-- 写一个 SQL,  报告余额高于 10000 的所有用户的名字和余额. 账户的余额等于包含该账户的所有交易的总和.
-- 
-- 返回结果表单没有顺序要求.
-- 
-- 查询结果格式如下例所示.
-- 
-- 
-- 
-- Users table:
-- +------------+--------------+
-- | account    | name         |
-- +------------+--------------+
-- | 900001     | Alice        |
-- | 900002     | Bob          |
-- | 900003     | Charlie      |
-- +------------+--------------+
-- 
-- Transactions table:
-- +------------+------------+------------+---------------+
-- | trans_id   | account    | amount     | transacted_on |
-- +------------+------------+------------+---------------+
-- | 1          | 900001     | 7000       |  2020-08-01   |
-- | 2          | 900001     | 7000       |  2020-09-01   |
-- | 3          | 900001     | -3000      |  2020-09-02   |
-- | 4          | 900002     | 1000       |  2020-09-12   |
-- | 5          | 900003     | 6000       |  2020-08-07   |
-- | 6          | 900003     | 6000       |  2020-09-07   |
-- | 7          | 900003     | -4000      |  2020-09-11   |
-- +------------+------------+------------+---------------+
-- 
-- Result table:
-- +------------+------------+
-- | name       | balance    |
-- +------------+------------+
-- | Alice      | 11000      |
-- +------------+------------+
-- Alice 的余额为(7000 + 7000 - 3000) = 11000.
-- Bob 的余额为1000.
-- Charlie 的余额为(6000 + 6000 - 4000) = 8000.
-- 
-- 
--

-- @lc code=start
# Write your MySQL query statement below
-- with
select Users.name,t.total as balance
from Users join (
select account,sum(amount) as total
from Transactions
group by  account
) t on Users.account=t.account and t.total>10000


-- 使用视图
-- with t as (
--     select account,sum(amount) as total
--     from Transactions
--     group by  account
--     having sum(amount)>10000
-- )
-- select Users.name,t.total as balance
-- from Users 
-- left join t on t.account=Users.account
-- where t.total is not null

-- @lc code=end

