
CREATE SEQUENCE if not exists course_id_seq START 1;
CREATE SEQUENCE if not exists semester_id_seq START 1;
CREATE SEQUENCE if not exists discipline_id_seq START 1;
CREATE SEQUENCE if not exists professor_id_seq START 1;
CREATE SEQUENCE if not exists availability_id_seq START 1;
CREATE SEQUENCE if not exists parameterization_id_seq START 1;
CREATE SEQUENCE if not exists proposal_id_seq START 1;
CREATE SEQUENCE if not exists class_id_seq START 1;
CREATE SEQUENCE if not exists user_id_seq START 1;
CREATE SEQUENCE if not exists eligible_disciplines_id_seq START 1;


CREATE TABLE if not exists course (
    id BIGINT PRIMARY KEY DEFAULT nextval('course_id_seq'),
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    modality VARCHAR(50) NOT NULL,
    location VARCHAR(255) NOT NULL
    );

CREATE TABLE if not exists semester (
    id BIGINT PRIMARY KEY DEFAULT nextval('semester_id_seq'),
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    semester VARCHAR(50) NOT NULL,
    CONSTRAINT semester_unique UNIQUE (semester)
    );

CREATE TABLE if not exists professor (
    id BIGINT PRIMARY KEY DEFAULT nextval('professor_id_seq'),
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    hoursToAllocate INT NOT NULL
    );

CREATE TABLE if not exists discipline (
    id BIGINT PRIMARY KEY DEFAULT nextval('discipline_id_seq'),
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    ----code VARCHAR(20) NOT NULL,
    name VARCHAR(255) NOT NULL,
    credits INT NOT NULL,
    course_id BIGINT not null,
    constraint discipline_course_id_fk foreign key(course_id) references course(id)
    );

CREATE TABLE if not exists availability (
    id BIGINT PRIMARY KEY DEFAULT nextval('availability_id_seq'),
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    dayOfWeek VARCHAR(50) NOT NULL,
    shift VARCHAR(50) NOT NULL,
    professor_id BIGINT not null,
    constraint availability_professor_id_fk foreign key(professor_id) references professor(id)
    );

CREATE TABLE parameterization (
      id BIGINT PRIMARY KEY DEFAULT nextval('parameterization_id_seq'),
      uuid UUID NOT NULL DEFAULT gen_random_uuid(),
      maxCreditsToOffer INT NOT NULL,
      numClassesPerDiscipline INT NOT NULL,
      semester_id BIGINT not null,
      course_id BIGINT not null,
      constraint parameterization_semester_id_fk foreign key(semester_id) references semester(id),
      constraint parameterization_course_id_fk foreign key(course_id) references course(id)
);

CREATE TABLE proposal (
      id BIGINT PRIMARY KEY DEFAULT nextval('proposal_id_seq'),
      uuid UUID NOT NULL DEFAULT gen_random_uuid(),
      semester_id BIGINT not null,
      course_id BIGINT not null,
      constraint proposal_semester_id_fk foreign key(semester_id) references semester(id),
      constraint proposal_course_id_fk foreign key(course_id) references course(id)
);

CREATE TABLE class (
       id BIGINT PRIMARY KEY DEFAULT nextval('class_id_seq'),
       uuid UUID NOT NULL DEFAULT gen_random_uuid(),
       dayOfWeek VARCHAR(50) NOT NULL,
       shift VARCHAR(50) NOT NULL,
       startTime TIMESTAMP not null,
       endTime TIMESTAMP not null,
       discipline_id BIGINT not null,
       professor_id BIGINT not null,
       proposal_id BIGINT NOT NULL,
       constraint class_discipline_id_fk foreign key(discipline_id) references discipline(id),
       constraint class_professor_id_fk foreign key(professor_id) references professor(id),
       CONSTRAINT class_proposal_id_fk FOREIGN KEY (proposal_id) REFERENCES proposal(id)
);



CREATE TABLE if not exists eligible_disciplines (
    id BIGINT PRIMARY KEY DEFAULT nextval('eligible_disciplines_id_seq'),
    professor_id BIGINT NOT NULL,
    discipline_id BIGINT NOT NULL,
    CONSTRAINT eligible_disciplines_professor_id_fk FOREIGN KEY (professor_id) REFERENCES professor(id),
    CONSTRAINT eligible_disciplines_discipline_id_fk FOREIGN KEY (discipline_id) REFERENCES discipline(id),
    CONSTRAINT eligible_disciplines_unique UNIQUE (professor_id, discipline_id)
    );


CREATE TABLE users (
   id BIGINT PRIMARY KEY DEFAULT nextval('user_id_seq'),
   uuid UUID NOT NULL DEFAULT gen_random_uuid(),
   name VARCHAR(255) NOT NULL,
   email VARCHAR(255) NOT NULL UNIQUE,
   password VARCHAR(255) NOT NULL
);

CREATE INDEX idx_course_name ON course(name);
CREATE INDEX idx_discipline_name ON discipline(name);
CREATE INDEX idx_professor_name ON professor(name);
CREATE INDEX idx_availability_professor_id ON availability(professor_id);
CREATE INDEX idx_parameterization_semester_id ON parameterization(semester_id);
CREATE INDEX idx_parameterization_course_id ON parameterization(course_id);
CREATE INDEX idx_proposal_semester_id ON proposal(semester_id);
CREATE INDEX idx_proposal_course_id ON proposal(course_id);
CREATE INDEX idx_class_discipline_id ON class(discipline_id);
CREATE INDEX idx_class_professor_id ON class(professor_id);


alter table public.professor
    add constraint professor_pk
        unique (name);

alter table public.discipline
    add constraint discipline_pk
        unique (name);

alter table public.availability
    add constraint availability_pk
        unique (dayofweek, professor_id, shift);

