#!/usr/bin/env bash

dbname="freepassbccbe2025"

read -p "Enter the database username: " dbuser
read -sp "Enter the database password: " dbpassword
echo

echo "Generating application.properties..."

cat <<EOL > src/main/resources/application.properties
spring.datasource.url=jdbc:postgresql://localhost:5432/$dbname
spring.datasource.username=$dbuser
spring.datasource.password=$dbpassword
EOL

echo "application.properties generated."

echo "Checking if database '$dbname' exists..."

DB_EXISTS=$(psql -U "$dbuser" -d postgres -tAc "SELECT 1 FROM pg_database WHERE datname='$dbname'")

if [ "$DB_EXISTS" != "1" ]; then
    echo "Database '$dbname' does not exist. Creating database..."
    PGPASSWORD=$dbpassword psql -U "$dbuser" -d postgres -c "CREATE DATABASE $dbname;"

    echo "Running SQL to create tables and insert initial data..."

    SQL="CREATE TABLE roles (
             role_id SERIAL PRIMARY KEY,
             role_name VARCHAR(50) NOT NULL UNIQUE
         );

         CREATE TABLE users (
             user_id SERIAL PRIMARY KEY,
             username VARCHAR(100) NOT NULL UNIQUE,
             password VARCHAR(255) NOT NULL,
             email VARCHAR(100) NOT NULL UNIQUE,
             full_name VARCHAR(100),
             role_id INT DEFAULT 3 REFERENCES roles(role_id) ON DELETE CASCADE
         );

         CREATE TABLE sessions (
             session_id SERIAL PRIMARY KEY,
             title VARCHAR(255) NOT NULL,
             description TEXT,
             start_time TIMESTAMP NOT NULL,
             end_time TIMESTAMP NOT NULL,
             max_seats INT NOT NULL,
             available_seats INT NOT NULL,
             created_by INT REFERENCES users(user_id) ON DELETE CASCADE
         );

         CREATE TABLE session_registrations (
             registration_id SERIAL PRIMARY KEY,
             user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
             session_id INT REFERENCES sessions(session_id) ON DELETE CASCADE,
             registration_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
             UNIQUE (user_id, session_id)
         );

         CREATE TABLE feedback (
             feedback_id SERIAL PRIMARY KEY,
             session_id INT REFERENCES sessions(session_id) ON DELETE CASCADE,
             user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
             feedback_text TEXT NOT NULL,
             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
         );

         CREATE TABLE session_proposals (
             proposal_id SERIAL PRIMARY KEY,
             title VARCHAR(255) NOT NULL,
             description TEXT,
             start_time TIMESTAMP NOT NULL,
             end_time TIMESTAMP NOT NULL,
             created_by INT REFERENCES users(user_id) ON DELETE CASCADE,
             status VARCHAR(50) DEFAULT 'Pending',
             UNIQUE (created_by, start_time, end_time)
         );

         INSERT INTO roles (role_name) VALUES
         ('Admin'),
         ('Coordinator'),
         ('User');

         INSERT INTO users (username, password, email, full_name, role_id) VALUES
         ('admin_user', 'adminpass123', 'admin@example.com', 'Admin User', 1),
         ('instructor_01', 'instrpass456', 'instructor@example.com', 'Instructor One', 2),
         ('student_01', 'studpass789', 'student01@example.com', 'Student One', 3),
         ('student_02', 'studpass000', 'student02@example.com', 'Student Two', 3);

         INSERT INTO sessions (title, description, start_time, end_time, max_seats, available_seats, created_by) VALUES
         ('Intro to Programming', 'A beginner session for learning programming basics.', '2025-02-01 09:00:00', '2025-02-01 12:00:00', 30, 30, 2),
         ('Advanced Java', 'Deep dive into advanced Java concepts and frameworks.', '2025-02-05 14:00:00', '2025-02-05 17:00:00', 20, 20, 2),
         ('Database Design', 'Learn best practices in designing databases.', '2025-02-10 10:00:00', '2025-02-10 13:00:00', 25, 25, 2);

         INSERT INTO session_registrations (user_id, session_id) VALUES
         (3, 1),
         (4, 1),
         (3, 2);

         INSERT INTO feedback (session_id, user_id, feedback_text) VALUES
         (1, 3, 'Great introductory session, very informative!'),
         (1, 4, 'Good session, but the pacing was a bit slow.'),
         (2, 3, 'The session was challenging but rewarding.');

         INSERT INTO session_proposals (title, description, start_time, end_time, created_by, status) VALUES
         ('Python Programming', 'A session to introduce Python for data analysis.', '2025-03-01 10:00:00', '2025-03-01 13:00:00', 2, 'Pending'),
         ('Web Development', 'Learn how to build a modern web application.', '2025-03-05 14:00:00', '2025-03-05 17:00:00', 2, 'Pending');"

    PGPASSWORD=$dbpassword psql -U "$dbuser" -d "$dbname" -c "$SQL"

else
    echo "Database '$dbname' already exists."
fi

echo "Cleaning project..."
mvn clean

echo "Building the project..."
mvn install

echo "Running the Spring Boot application..."
mvn spring-boot:run
