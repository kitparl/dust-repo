# PostgresSQL

- PostgresSQL is an advance relational database
- It is free and Opens-source
- It supports SQL and NoSQL
- It is a backend database for dynamic web applications and websites
- It supports the most important programming languages.

## 1. Database Management

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

## 5. Aggregates and Grouping

#### Group by (with aggregate functions)

```
SELECT position, COUNT(*) as count
FROM employees
GROUP BY position;
```

#### Common Aggregate Functions

- Count Rows

```
SELECT COUNT(*) FROM employees;
```

- Sum of Salaries

```
SELECT SUM(salary) FROM employees;
```

- Average Salary

```
SELECT AVG(salary) FROM employees;
```

- Maximum Salary

```
SELECT MAX(salary) FROM employees;
```

- Minimum Salary

```
SELECT MIN(salary) FROM employees;
```

- Having Clause (Filter Grouped Data)

```
SELECT position, COUNT(*) as count
FROM employees
GROUP BY position
HAVING COUNT(*) > 6;
```

## 6. Transaction

#### Start Transaction

```
BEGIN;
```

#### Commit Transaction

```
COMMIT;
```

#### Rollback Transaction

```
ROLLBACK;
```

#### Set Transaction Isolation Level

```
SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
```

## 7. Constraints

#### Add a Primary Key Constraint

```
ALTER TABLE employees
ADD CONSTRAINT pk_id PRIMARY KEY (id);
```

#### Add a Foreign Key Constraint

```
ALTER TABLE employees
ADD CONSTRAINT fk_department_id
FOREIGN KEY (department_id) REFERENCES departments(id);
```

#### Add a Unique Constraint

```
ALTER TABLE employees
ADD CONSTRAINT uq_emplyee_email UNIQUE (email);
```

#### Add a Check Constraint

```
ALTER TABLE employees
ADD CONSTRAINT ck_salary CHECK (salary > 0);
```

## 8. User and Role Management

#### Create a User

```
CREATE USER pranshu WITH PASSWORD 'password';
```

#### Create a Role

```
CREATE ROLE admin;
```

#### Grant Role to User

```
GRANT <role_name> TO <username>;

GRANT admin TO pranshu;
```

#### List all Users/Roles

```
\du
```

#### Drop a User

```
DROP USER pranshu;
```

#### Grant Permission to User

```
GRANT <privilege> ON <object> TO <username>;

GRANT SELECT ON employees TO pranshu;

-----------------------------
Common Privileges
-----------------------------

- SELECT
- INSERT
- UPDATE
- DELETE
- ALL PRIVILEGES
```

## 9. Indexes

#### Create an Index

```
CREATE INDEX <index_name> ON <table_name> (<column_name>);

CREATE INDEX idx_name ON employees(name);
```

#### List all Indexes

```
\di
```

#### Drop an Index

```
DROP INDEX <index_name>;
```

## Advance Features

#### Create a View

```
CREATE VIEW <view_name> AS SELECT * FROM <table_name> WHERE condition;
```

#### Materialized View (stores results for faster access)

```
CREATE MATERIALIZED VIEW <view_name> AS SELECT * FROM <table_name>;
```

#### Trigger Creation (for automatic actions)

```
CREATE TRIGGER <trigger_name> 
BEFORE INSERT ON <table_name>
FOR EACH ROW EXECUTE FUNCTION <function_name>();
```

## Other Useful Commands

#### List all Relations (Tables, Views, Sequences)

```
\d
```

#### List all Views

```
\dv
```

#### List all Sequences

```
\ds
```

#### Explain Query Plan

```
EXPLAIN <query>;
```

#### Describe a table

```
\d table_name

\d employees
```

#### Addtional Describe a table

```
\d+ table_name

\d+ employees
```

#### Show Current Database Connection

```
\conninfo
```

#### Show Server Version

```
SELECT version();
```

#### Get Current User

```
SELECT current_user;
```
