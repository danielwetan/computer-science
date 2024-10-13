Sure! Here’s the complete explanation, including the database schema and SQL queries, all in English and using lowercase for the database and field names.

### Database Schema Overview (Lowercase)

1. **students Table** (Reference Table)
2. **courses Table** (Reference Table)
3. **enrollments Table** (Non-reference Table with Foreign Keys)

### Table Structure

#### 1. students Table
- **Table Name:** `students`
- **Description:** This table contains information about the students enrolled in the university.
- **Columns:**
  - `id` (Primary Key, INT)
  - `first_name` (VARCHAR)
  - `last_name` (VARCHAR)
  - `email` (VARCHAR)
  - `city` (VARCHAR)  
- **Sample Records:**
  | id | first_name | last_name | email                 | city     |
  |----|------------|-----------|-----------------------|----------|
  | 1  | john       | doe       | john.doe@example.com   | bandung  |
  | 2  | jane       | smith     | jane.smith@example.com | jakarta  |
  | 3  | alice      | johnson   | alice.j@example.com    | bandung  |

#### 2. courses Table
- **Table Name:** `courses`
- **Description:** This table holds information about the courses offered at the university.
- **Columns:**
  - `id` (Primary Key, VARCHAR)
  - `name` (VARCHAR)
  - `credits` (INT)
- **Sample Records:**
  | id    | name                | credits |
  |-------|---------------------|---------|
  | 101   | introduction to cs   | 3       |
  | 102   | data structures      | 4       |
  | csh2c3| advanced programming | 4       |

#### 3. enrollments Table
- **Table Name:** `enrollments`
- **Description:** This table records which students are enrolled in which courses. It contains foreign keys referencing the `students` and `courses` tables.
- **Columns:**
  - `id` (Primary Key, INT)
  - `student_id` (Foreign Key, INT)
  - `course_id` (Foreign Key, VARCHAR)
  - `enrollment_date` (DATE)
- **Sample Records:**
  | id | student_id | course_id | enrollment_date |
  |----|------------|-----------|-----------------|
  | 1  | 1          | 101       | 2024-08-20      |
  | 2  | 1          | 102       | 2024-08-20      |
  | 3  | 2          | 103       | 2024-08-21      |
  | 4  | 3          | csh2c3    | 2024-08-22      |

### 2. SQL Queries with Lowercase Names

#### a. Searching with Non-Key Conditions

**Query 1: Search for Students from Bandung (1 Table)**

```sql
SELECT * FROM students
WHERE city = 'bandung';
```

**Query 2: Search for Students Enrolled in the Course "s1 informatika" (2 Tables)**

```sql
SELECT s.first_name, s.last_name, c.name
FROM students s
JOIN enrollments e ON s.id = e.student_id
JOIN courses c ON e.course_id = c.id
WHERE c.name = 's1 informatika';
```

**Query 3: Search for Students from Bandung Enrolled in the Course with Code "csh2c3" (3 Tables)**

```sql
SELECT s.first_name, s.last_name, c.name
FROM students s
JOIN enrollments e ON s.id = e.student_id
JOIN courses c ON e.course_id = c.id
WHERE s.city = 'bandung' AND c.id = 'csh2c3';
```

#### b. Searching with Key Conditions

**Query 4: Search for Students with Student ID = 1103120066 (2 Tables)**

```sql
SELECT s.first_name, s.last_name
FROM students s
JOIN enrollments e ON s.id = e.student_id
WHERE s.id = 1103120066;
```

**Query 5: Search for Students with Student ID Starting with '110313' Who Are Enrolled in the Course with Code "csh2c3" (3 Tables)**

```sql
SELECT s.first_name, s.last_name, c.name
FROM students s
JOIN enrollments e ON s.id = e.student_id
JOIN courses c ON e.course_id = c.id
WHERE s.id LIKE '110313%' AND c.id = 'csh2c3';
```

### 3. Executing the Queries in DBMS

#### Recording Execution Time
To record the execution time for each query in Oracle DBMS, you can follow these steps:

1. **Enable Timing**
   ```sql
   SET TIMING ON;
   ```

2. **Execute Each Query**
   - Execute Query 1
   - Execute Query 2
   - Execute Query 3
   - Execute Query 4
   - Execute Query 5

3. **Record Execution Time**
   After running each query, note the time displayed in the console.

#### Example Execution Time Record
Here’s an example format to record the execution times (the numbers below are illustrative; you should replace them with the actual times you get):

| Query No | Execution Time (seconds) |
|----------|--------------------------|
| Query 1  | 0.15                     |
| Query 2  | 0.25                     |
| Query 3  | 0.30                     |
| Query 4  | 0.10                     |
| Query 5  | 0.35                     |

### Conclusion
Using lowercase for table and column names, these SQL queries can be utilized to access the required information from the university database. Make sure to adjust the data and conditions according to the actual database you are working with.