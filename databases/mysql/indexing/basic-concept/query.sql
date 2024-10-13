-- 1. Searching with Non-Key Conditions
SELECT * FROM students
WHERE city = 'bandung';

-- 2 Search for students Enrolled in the Course "data structures" (2 Tables)
SELECT s.id, s.first_name, s.last_name, c.id, c.name
FROM students s
JOIN enrollments e ON s.id = e.student_id
JOIN courses c ON e.course_id = c.id
WHERE c.name = 'data structures';

-- 3 Search for students from "bandung" Enrolled in the Course with "data structures" (3 Tables)
SELECT s.id, s.first_name, s.last_name, c.id, c.name
FROM students s
JOIN enrollments e ON s.id = e.student_id
JOIN courses c ON e.course_id = c.id
WHERE s.city = 'bandung' AND c.name = 'data structures';

-- 4 Search for students with id = 10 (2 Tables)
SELECT s.id, s.first_name, s.last_name, e.id, e.course_id
FROM students s
JOIN enrollments e ON s.id = e.student_id
WHERE s.id = 10;

--  5 Search for students with id starting with '2' and enrolled in the course with code "c2" (3 Tables)
SELECT s.id, s.first_name, s.last_name, c.id, c.name
FROM students s
JOIN enrollments e ON s.id = e.student_id
JOIN courses c ON e.course_id = c.id
WHERE s.id LIKE '2%' AND c.id = 'c2';
