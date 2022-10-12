--
-- @lc app=leetcode.cn id=178 lang=mysql
--
-- [178] 分数排名
--
-- https://leetcode.cn/problems/rank-scores/description/
--
-- database
-- Medium (60.92%)
-- Likes:    1009
-- Dislikes: 0
-- Total Accepted:    172.5K
-- Total Submissions: 283.1K
-- Testcase Example:  '{"headers": {"Scores": ["id", "score"]}, "rows": {"Scores": [[1, 3.50], [2, 3.65], [3, 4.00], [4, 3.85], [5, 4.00], [6, 3.65]]}}'
--
-- 表: Scores
-- 
-- 
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | score       | decimal |
-- +-------------+---------+
-- Id是该表的主键。
-- 该表的每一行都包含了一场比赛的分数。Score是一个有两位小数点的浮点值。
-- 
-- 
-- 
-- 
-- 编写 SQL 查询对分数进行排序。排名按以下规则计算:
-- 
-- 
-- 分数应按从高到低排列。
-- 如果两个分数相等，那么两个分数的排名应该相同。
-- 在排名相同的分数后，排名数应该是下一个连续的整数。换句话说，排名之间不应该有空缺的数字。
-- 
-- 
-- 按 score 降序返回结果表。
-- 
-- 查询结果格式如下所示。
-- 
-- 
-- 
-- 示例 1:
-- 
-- 
-- 输入: 
-- Scores 表:
-- +----+-------+
-- | id | score |
-- +----+-------+
-- | 1  | 3.50  |
-- | 2  | 3.65  |
-- | 3  | 4.00  |
-- | 4  | 3.85  |
-- | 5  | 4.00  |
-- | 6  | 3.65  |
-- +----+-------+
-- 输出: 
-- +-------+------+
-- | score | rank |
-- +-------+------+
-- | 4.00  | 1    |
-- | 4.00  | 1    |
-- | 3.85  | 2    |
-- | 3.65  | 3    |
-- | 3.65  | 3    |
-- | 3.50  | 4    |
-- +-------+------+
-- 
--

-- @lc code=start
# Write your MySQL query statement below

# 同分同名,不升级问题,使用窗口函数
-- select score,dense_rank() over(order by score desc)as `rank`
-- from Scores

# 使用自定义变量解决,rank是mysql内置函数,要加引号
# 发现这样写,rank返回的是字符串类型,必须cast转换一下
-- select score,cast(`rank` as SIGNED INTEGER) as `rank`
-- from (
-- select score,@r:=IF(@p=score,@r,@r+1) as `rank`,@p:=score
-- from scores,(select @r:=0,@p:=NULL) init
-- order by score desc
-- )as tmp;

#TODO 如何在声明时设定mysql variable的类型,而不是在返回时cast转换
# 多条语句在外部能正常执行,提交到leetcode无法执行,可能不支持
-- SET @r:=0;
-- SET @p:=1;
-- select score,`rank` -- cast(`rank` as SIGNED INTEGER)
-- from (
-- select score,@r:=IF(@p=score,@r,@r+1) as `rank`,@p:=score
-- from scores -- ,(select @r:=0,@p:=NULL) init
-- order by score desc
-- )as tmp;


# if条件语句里加0,就不会使r变成字符类型了
select score,cast(`rank` as SIGNED INTEGER) as `rank`
from (
select score,@r:=IF(@p=score,@r+0,@r+1) as `rank`,@p:=score
from scores,(select @r:=0,@p:=NULL) init
order by score desc
)as tmp;
-- @lc code=end

