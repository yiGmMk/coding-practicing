--
-- @lc app=leetcode.cn id=177 lang=mysql
--
-- [177] 第N高的薪水
--
-- https://leetcode.cn/problems/nth-highest-salary/description/
--
-- database
-- Medium (46.81%)
-- Likes:    644
-- Dislikes: 0
-- Total Accepted:    181.2K
-- Total Submissions: 387K
-- Testcase Example:  '{"headers": {"Employee": ["id", "salary"]}, "argument": 2, "rows": {"Employee": [[1, 100], [2, 200], [3, 300]]}}'
--
-- 表: Employee
-- 
-- 
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | id          | int  |
-- | salary      | int  |
-- +-------------+------+
-- Id是该表的主键列。
-- 该表的每一行都包含有关员工工资的信息。
-- 
-- 
-- 
-- 
-- 编写一个SQL查询来报告 Employee 表中第 n 高的工资。如果没有第 n 个最高工资，查询应该报告为 null 。
-- 
-- 查询结果格式如下所示。
-- 
-- 
-- 
-- 示例 1:
-- 
-- 
-- 输入: 
-- Employee table:
-- +----+--------+
-- | id | salary |
-- +----+--------+
-- | 1  | 100    |
-- | 2  | 200    |
-- | 3  | 300    |
-- +----+--------+
-- n = 2
-- 输出: 
-- +------------------------+
-- | getNthHighestSalary(2) |
-- +------------------------+
-- | 200                    |
-- +------------------------+
-- 
-- 
-- 示例 2:
-- 
-- 
-- 输入: 
-- Employee 表:
-- +----+--------+
-- | id | salary |
-- +----+--------+
-- | 1  | 100    |
-- +----+--------+
-- n = 2
-- 输出: 
-- +------------------------+
-- | getNthHighestSalary(2) |
-- +------------------------+
-- | null                   |
-- +------------------------+
-- 
--

-- @lc code=start

-- 作者：luanhz
-- 链接：https://leetcode.cn/problems/nth-highest-salary/solution/mysql-zi-ding-yi-bian-liang-by-luanz/
-- 来源：力扣（LeetCode）
-- 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
-- CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
-- BEGIN
--     SET N := N-1;
--   RETURN (
--       # Write your MySQL query statement below.
--       SELECT 
--             salary
--       FROM 
--             employee
--       GROUP BY 
--             salary
--       ORDER BY 
--             salary DESC
--       LIMIT N, 1
--       -- 排名第N高意味着要跳过N-1个薪水，由于无法直接用limit N-1(N-1可能是负数,0)，所以需先在函数开头处理N为N=N-1。
--   );
-- END


-- mysql8开始支持的窗口函数
-- row_number(): 同薪不同名，相当于行号，例如3000、2000、2000、1000排名后为1、2、3、4
-- rank(): 同薪同名，有跳级，例如3000、2000、2000、1000排名后为1、2、2、4
-- dense_rank(): 同薪同名，无跳级，例如3000、2000、2000、1000排名后为1、2、2、3
-- ntile(): 分桶排名，即首先按桶的个数分出第一二三桶，然后各桶内从1排名，实际不是很常用
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
      # Write your MySQL query statement below.
      SELECT 
          DISTINCT salary
      FROM 
          (SELECT 
              salary, dense_rank() over(ORDER BY salary DESC) AS rnk
           FROM 
              employee) tmp
      WHERE rnk = N
  );
END

-- 自定义函数更一般化和常用的写法应该是分三步：
-- 定义变量接收返回值
-- 执行查询条件，并赋值给相应变量
-- 返回结果
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    # i 定义变量接收返回值
    DECLARE ans INT DEFAULT NULL;  
    # ii 执行查询语句，并赋值给相应变量
    SELECT 
        DISTINCT salary INTO ans
    FROM 
        (SELECT 
            salary, @r:=IF(@p=salary, @r, @r+1) AS rnk,  @p:= salary 
        FROM  
            employee, (SELECT @r:=0, @p:=NULL)init 
        ORDER BY 
            salary DESC) tmp
    WHERE rnk = N;
    # iii 返回查询结果，注意函数名中是 returns，而函数体中是 return
    RETURN ans;
END

-- @lc code=end

