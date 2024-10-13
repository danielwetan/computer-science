-- Create students table
CREATE TABLE students (
    id INT PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100),
    city VARCHAR(50)
);

-- Create courses table
CREATE TABLE courses (
    id VARCHAR(10) PRIMARY KEY,
    name VARCHAR(100),
    credits INT
);

-- Create enrollments table
CREATE TABLE enrollments (
    id INT PRIMARY KEY,
    student_id INT,
    course_id VARCHAR(10),
    enrollment_date DATE,
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);

-- Create composite index on first_name and last_name columns in the users table
CREATE INDEX idx_users_full_name ON users(first_name, last_name);
-- Create index on the city column in the students table
CREATE INDEX idx_students_city ON students(city);

-- Create index on the name column in the courses table
CREATE INDEX idx_courses_name ON courses(name);

-- Create composite index on the student_id and course_id columns in the enrollments table
CREATE INDEX idx_enrollments_student_course ON enrollments(student_id, course_id);

