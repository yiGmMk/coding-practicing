#
# @lc app=leetcode.cn id=194 lang=bash
#
# [194] 转置文件
#
# https://leetcode.cn/problems/transpose-file/description/
#
# shell
# Medium (34.46%)
# Likes:    68
# Dislikes: 0
# Total Accepted:    13.7K
# Total Submissions: 39.7K
# Testcase Example:  'a'
#
# 给定一个文件 file.txt，转置它的内容。
# 
# 你可以假设每行列数相同，并且每个字段由 ' ' 分隔。
# 
# 
# 
# 示例：
# 
# 假设 file.txt 文件内容如下：
# 
# 
# name age
# alice 21
# ryan 30
# 
# 
# 应当输出：
# 
# 
# name alice ryan
# age 21 30
# 
# 
#

# @lc code=start
# Read from the file file.txt and print its transposed content to stdout.

# 作者：LuciferLiu
# 链接：https://leetcode.cn/problems/transpose-file/solution/leetcodejie-ti-xi-lie-194-zhuan-zhi-wen-u4r72/
# 来源：力扣（LeetCode）
# 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。 

# 先算有多少列,按列取数据,再用xargs转成行
# awk '{print $1}' file 取第一列的数据打印(print $i可以取指定列的数据)
file=file.txt
columns=$(cat $file  | head -n 1 | wc -w)
# seq start end(1为步长,循环)
for i in $(seq 1 $columns)
do
awk '{print $'"$i"'}' $file | xargs #这里$'''$i''' 或$'$i' 或$'"$i"'都可以,实质上是字符串拼接
done

# @lc code=end

