#
# @lc app=leetcode.cn id=192 lang=bash
#
# [192] 统计词频
#
# https://leetcode.cn/problems/word-frequency/description/
#
# shell
# Medium (35.88%)
# Likes:    200
# Dislikes: 0
# Total Accepted:    28.1K
# Total Submissions: 78.4K
# Testcase Example:  'a'
#
# 写一个 bash 脚本以统计一个文本文件 words.txt 中每个单词出现的频率。
# 
# 为了简单起见，你可以假设：
# 
# 
# words.txt只包括小写字母和 ' ' 。
# 每个单词只由小写字母组成。
# 单词间由一个或多个空格字符分隔。
# 
# 
# 示例:
# 
# 假设 words.txt 内容如下：
# 
# the day is sunny the the
# the sunny is is
# 
# 
# 你的脚本应当输出（以词频降序排列）：
# 
# the 4
# is 3
# sunny 2
# day 1
# 
# 
# 说明:
# 
# 
# 不要担心词频相同的单词的排序问题，每个单词出现的频率都是唯一的。
# 你可以使用一行 Unix pipes 实现吗？
# 
# 
#

# @lc code=start
# Read from the file words.txt and output the word frequency list to stdout.
#!/bin/bash
cat words.txt | tr -s ' ' '\n'|sort|uniq -c|sort -nr|awk '{print $2" "$1}'

# 1.切割,使用tr命令将空格转成换行符(-s：缩减连续重复的字符成指定的单个字符)
# 2.排序
# 3.统计单词出现次数,uniq 命令用于检查及删除文本文件中重复出现的行列，一般与 sort 命令结合使用(-c：在每列旁边显示该行重复出现的次数)
# 4.以出现次数排序,-r倒序排序,-n以数值排序
# 5.输出
# 当单词的出现次数大于10时,sort需要考虑按数字排序,而非默认的按 ascii 码排序=>第二个sort加个参数-nr

# @lc code=end

