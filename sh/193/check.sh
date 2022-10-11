#!/usr/bin bash
if cmp res.txt correct.txt
then
  echo '---' test: PASS
else
  echo '---' test Failed
  cat res.txt
  echo 'not correct'
fi