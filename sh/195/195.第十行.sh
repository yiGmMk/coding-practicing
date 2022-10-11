#
# @lc app=leetcode.cn id=195 lang=bash
#
# [195] 第十行
#
# https://leetcode.cn/problems/tenth-line/description/
#
# shell
# Easy (44.08%)
# Likes:    119
# Dislikes: 0
# Total Accepted:    46.9K
# Total Submissions: 106.4K
# Testcase Example:  'Line 1\\nLine 2\\nLine 3\\nLine 4\\nLine 5\\nLine 6\\nLine 7\\nLine 8\\nLine 9\\nLine 10'
#
# 给定一个文本文件 file.txt，请只打印这个文件中的第十行。
# 
# 示例:
# 
# 假设 file.txt 有如下内容：
# 
# Line 1
# Line 2
# Line 3
# Line 4
# Line 5
# Line 6
# Line 7
# Line 8
# Line 9
# Line 10
# 
# 
# 你的脚本应当显示第十行：
# 
# Line 10
# 
# 
# 说明:
# 1. 如果文件少于十行，你应当输出什么？
# 2. 至少有三种不同的解法，请尝试尽可能多的方法来解题。
# 
#

# @lc code=start
# Read from the file file.txt and output the tenth line to stdout.

# 1. awk
awk 'NR==10' file.txt
#或者
awk 'NR == 10{print $0}' file.txt

# 2. head & tail
tail -n +10 file.txt | head -n 1
# 或者
head -n 10 file.txt | tail -n +10 

# @lc code=end

