-- 清理已经删除掉的问题在answers表中仍有记录导致的前端练习历史空白
use exam;
select * FROM answers where question_id not in (select id as question_id from questions)
