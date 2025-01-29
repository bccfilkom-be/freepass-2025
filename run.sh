#!/usr/bin/env bash

echo "Please enter your database details."

read -p "Enter the database name: " DB_NAME
read -p "Enter the database username: " DB_USERNAME
read -sp "Enter the database password: " DB_PASSWORD
echo

echo "Generating application.properties..."

cat > src/main/resources/application.properties <<EOL
spring.datasource.url=jdbc:postgresql://localhost:5432/$DB_NAME
spring.datasource.username=$DB_USERNAME
spring.datasource.password=$DB_PASSWORD
EOL

echo "application.properties file has been created."

echo "Cleaning project..."
mvn clean

echo "Building the project..."
mvn install

echo "Running the Spring Boot application..."
mvn spring-boot:run
