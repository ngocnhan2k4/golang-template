DROP TABLE IF EXISTS students;

CREATE TABLE settings
(
    id SERIAL PRIMARY KEY,
    email_domain text NOT NULL,
	email_pattern text NOT NULL
);

CREATE TABLE register_cancellation_history(
	id SERIAL PRIMARY KEY,
	class_id text NOT NULL,
	course_name text NOT NULL,
	student_id text NOT NULL,
	student_name text NOT NULL,
	semester int NOT NULL,
	academic_year int NOT NULL,
	time timestamp NOT NULL
);

CREATE TABLE faculties(
	id SERIAL PRIMARY KEY,
	name varchar(50) NOT NULL,
	name_eng varchar(50) DEFAULT '' NOT NULL
);

CREATE TABLE programs(
	id SERIAL PRIMARY KEY,
	name text NOT NULL,
	name_eng text DEFAULT '' NOT NULL
);


CREATE TABLE student_statuses(
	id SERIAL PRIMARY KEY,
	name text NOT NULL,
	status_order int NOT NULL,
	name_eng text DEFAULT '' NOT NULL
);

CREATE TABLE courses(
	id SERIAL PRIMARY KEY,
	name varchar(50) NOT NULL,
	credits int NOT NULL,
	faculty_id int NOT NULL,
	description TEXT NOT NULL,
	required_course_id int NULL,
	deleted_at timestamp NULL,
	created_at timestamp NOT NULL,
	description_eng text DEFAULT '' NOT NULL,
	name_eng varchar(50) DEFAULT '' NOT NULL,
	FOREIGN KEY (required_course_id) REFERENCES courses(id),
	FOREIGN KEY (faculty_id) REFERENCES faculties(id) ON DELETE CASCADE
);

CREATE TABLE classes(
	id SERIAL PRIMARY KEY,
	academic_year int NOT NULL,
	course_id int NOT NULL,
	semester int NOT NULL,
	teacher_name varchar(50) NOT NULL,
	max_students int NOT NULL,
	room varchar(10) NOT NULL,
	day_of_week int NOT NULL,
	start_time time NOT NULL,
	end_time time NOT NULL,
	deadline timestamp NOT NULL,
	is_deleted boolean NOT NULL,
	FOREIGN key (course_id) REFERENCES courses(id)
);

CREATE TABLE students(
	id SERIAL PRIMARY KEY,
	name varchar(50) NOT NULL,
	date_of_birth date NOT NULL,
	gender varchar(10) DEFAULT 'Male' NOT NULL,
	email varchar(50) NOT NULL,
	course int NOT NULL,
	phone varchar(20) NOT NULL,
	permanent_address TEXT NOT NULL,
	temporary_address TEXT NULL,
	mailing_address TEXT NULL,
	program_id  int NOT NULL,
	status_id int NOT NULL,
	faculty_id int NOT NULL,
	nationality varchar(50) NOT NULL,
	is_deleted boolean NOT NULL,
	FOREIGN KEY (faculty_id) REFERENCES faculties(id),
	FOREIGN KEY (program_id) REFERENCES programs(id),
	FOREIGN KEY (status_id) REFERENCES student_statuses(id)
);

CREATE TABLE identity_documents(
	id SERIAL PRIMARY KEY,
	type varchar(10) DEFAULT ('CCCD') NOT NULL,
	number varchar(20) NOT NULL,
	issued_date date NOT NULL,
	expiry_date date NOT NULL,
	issue_place varchar(100) NOT NULL,
	country varchar(50) NULL,
	is_chip BOOLEAN NOT NULL,
	note varchar(100) NULL,
	student_id int NOT NULL,
	FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE
);


CREATE TABLE class_student(
	class_id int NOT NULL,
	student_id int NOT NULL,
	score numeric NOT NULL,
	final_score numeric NOT NULL,
	gpa numeric NOT NULL,
	grade varchar(1) NOT NULL,
	is_passed boolean NOT NULL,
	PRIMARY KEY(class_id, student_id),
	FOREIGN key (class_id) REFERENCES classes(id) ON DELETE CASCADE,
	FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE
);