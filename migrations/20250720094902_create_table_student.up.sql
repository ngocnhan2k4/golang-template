CREATE TABLE students(
    id VARCHAR(8) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    date_of_birth DATE NOT NULL,
    gender INT DEFAULT 0 NOT NULL,
    email VARCHAR(50) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    permanent_address TEXT NOT NULL,
    temporary_address TEXT,
    mailing_address TEXT,
    nationality VARCHAR(50) NOT NULL,
    course INT NOT NULL,
    isdeleted BOOLEAN DEFAULT FALSE NOT NULL
);

INSERT INTO students VALUES
('22120249', 'Nhan', '2004-07-05', 0, 'tranngocnhannt2004@gmail.com', '0987654321', 
 '123 Main St, Hanoi', '456 Secondary St, Hanoi', '789 Mailing St, Hanoi', 
 'Vietnamese', 2022, FALSE);