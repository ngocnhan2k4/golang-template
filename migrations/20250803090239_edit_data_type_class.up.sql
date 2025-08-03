DROP TABLE IF EXISTS class_student;

DROP TABLE IF EXISTS classes;

CREATE TABLE classes(
	id SERIAL PRIMARY KEY,
	academic_year int NOT NULL,
	course_id int NOT NULL,
	semester int NOT NULL,
	teacher_name varchar(50) NOT NULL,
	max_students int NOT NULL,
	room varchar(10) NOT NULL,
	day_of_week int NOT NULL,
	start_time numeric NOT NULL,
	end_time numeric NOT NULL,
	deadline timestamp NOT NULL,
	is_deleted boolean NOT NULL,
	FOREIGN key (course_id) REFERENCES courses(id)
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