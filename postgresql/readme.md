# PostgresSQL

- PostgresSQL is an advance relational database
- It is free and Opens-source
- It supports SQL and NoSQL
- It is a backend database for dynamic web applications and websites
- It supports the most important programming languages.

## Database Management

#### Create a Database

```
CREATE DATABASE [name];
```

#### List all Databases

```
\l
```

#### Switch to another Database

```
\c [database_name]
```

#### Drop a Database

```
DROP DATABASE [name];
```

## 2. Table Management

#### Create a Table

```
CREATE TABLE employees (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50),
  position VARCHAR(50),
  salary DECIMAL(10,2)
)
```

#### List all Tables

```
\dt
```

#### Describe Table Structure

```
\d employees
```

#### Alter a Table (Add Column)

```
ALTER TABLE employees ADD COLUMN email VARCHAR(50);
```

#### Alter a Table (Modify Column)

```
ALTER TABLE employees ALTER COLUMN salary SET DATA TYPE NUMERIC(10,2)
```

#### Drop a Table

```
DROP TABLE employees;
```

## 3. Data Manipulation

#### Insert Data

```
INSERT INTO employees (name, position, salary)
VALUES ('Pranshu Singh Bisht', 'Developer', 90000.00);
```

#### Select Data (Basic Query)

```
SELECT * FROM employees;
```

#### Select Data (With Conditions)

```
SELECT name, position FROM employees WHERE salary > 50000;
```

#### Update Data

```
UPDATE employees
SET salary = 95000.00
WHERE id = 1;
```

#### Delete Data

```
DELETE FROM employees where name = 'Pranshu Singh Bisht'
```

## 4. Join Operations

#### Inner Join

```
SELECT e.name, d.department_name
FROM employees e
LEFT JOIN departments d ON e.department_id = d.id;
```

#### Left Join or Left Outer Join

```
SELECT e.name, d.department_name
FROM employees e
LEFT JOIN departments d ON e.department_id = d.id;
```

#### Right Join or Right Outer Join

```
SELECT e.name, d.department_name
FROM employees e
RIGHT JOIN departments d ON e.department_id = d.id;
```

#### Full join or Full Outer Join

```
SELECT e.name, d.department_name
FROM employees e
FULL OUTER JOIN departments d ON e.department_id = d.id;
```
