
drop index if exists idx_course_name;
drop index if exists idx_discipline_name;
drop index if exists idx_professor_name;
drop index if exists idx_availability_professor_id;
drop index if exists idx_parameterization_semester_id;
drop index if exists idx_parameterization_course_id;
drop index if exists idx_class_discipline_id;
drop index if exists idx_class_professor_id;
drop index if exists idx_proposal_semester_id;
drop index if exists idx_proposal_course_id;

drop table if exists users;
drop table if exists class;
drop table if exists proposal;
drop table if exists parameterization;
drop table if exists availability;
drop table if exists discipline;
drop table if exists professor;
drop table if exists semester;
drop table if exists course;
drop table if exists eligible_disciplines;

drop sequence if exists course_id_seq;
drop sequence if exists semester_id_seq;
drop sequence if exists discipline_id_seq;
drop sequence if exists professor_id_seq;
drop sequence if exists availability_id_seq;
drop sequence if exists parameterization_id_seq;
drop sequence if exists proposal_id_seq;
drop sequence if exists class_id_seq;
drop sequence if exists user_id_seq;
drop sequence if exists eligible_disciplines_id_seq;



