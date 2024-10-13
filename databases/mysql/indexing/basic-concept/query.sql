-- 1. Searching with Non-Key Conditions
SELECT * FROM students
WHERE city = 'bandung';

-- 2 Search for Students Enrolled in the Course "data structures" (2 Tables)
SELECT s.id, s.first_name, s.last_name, c.id, c.name
FROM students s
JOIN enrollments e ON s.id = e.student_id
JOIN courses c ON e.course_id = c.id
WHERE c.name = 'data structures';

-- 3 Search for Students from Bandung Enrolled in the Course with Code "csh2c3" (3 Tables)
SELECT s.id, s.first_name, s.last_name, c.id, c.name
FROM students s
JOIN enrollments e ON s.id = e.student_id
JOIN courses c ON e.course_id = c.id
WHERE s.city = 'bandung' AND c.name = 'data structures';

-- 4 Search for Students with Student ID = 1103120066 (2 Tables)
SELECT s.id, s.first_name, s.last_name, e.id, e.course_id
FROM students s
JOIN enrollments e ON s.id = e.student_id
WHERE s.id = 10;

--  5 Search for Students with Student ID Starting with '110313' Who Are Enrolled in the Course with Code "csh2c3" (3 Tables)
SELECT s.id, s.first_name, s.last_name, c.id, c.name
FROM students s
JOIN enrollments e ON s.id = e.student_id
JOIN courses c ON e.course_id = c.id
WHERE s.id LIKE '2%' AND c.id = 'c2';
