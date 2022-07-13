# RabbitMQ Project
 
1.) Create user account database on mysql(not docker) to store user account information
    a.) In mysql workbench, create an account name user1 with all administrative privilege

    b.) Use the following sql script to create the database and table:
        CREATE database MYSTOREDBRABBITMQ;
        USE MYSTOREDBRABBITMQ;
        CREATE TABLE Users (UserName VARCHAR(30) NOT NULL PRIMARY KEY, Password VARCHAR(256), FirstName VARCHAR(30), LastName VARCHAR(30), Language VARCHAR(30));

2.) To create MySQL Database on Docker: 5.) Create Docker Container for sql database Working: https://www.youtube.com/watch?v=X8W5Xq9e2Os 

a.) Download Docker Image docker pull mysql/mysql-server:latest

b.) Run docker container
    docker run -p 3307:3306 --name MYSTOREDBRABBITMQ -e MYSQL_ROOT_PASSWORD=password -d mysql

c.) Execute container
    docker exec -it MYSTOREDBRABBITMQ /bin/bash

d.) Login to mysql
    mysql -uroot -p -A
    enter password:password

e.) To create a table in sql database in container
    1.) CREATE DATABASE IF NOT EXISTS MYSTOREDBRABBITMQ;
    2.) USE MYSTOREDBRABBITMQ;
    3.) Press enter
    4.) CREATE TABLE Users (UserName VARCHAR(30) NOT NULL PRIMARY KEY, Password VARCHAR(256), FirstName VARCHAR(30), LastName VARCHAR(30), Language VARCHAR(30));
    5.) Press enter

f.) To check content of table:
    USE MYSTOREDBRABBITMQ; SELECT * FROM Users;

g.) To see database:
    SHOW DATABASES;

h.) To see what is in MYSTOREDBRABBITMQ User table:
    SHOW COLUMNS FROM MYSTOREDBRABBITMQ.Users;

Do either 3 or 4, not both
===================================================================================================================================================================
2.) Download RabbitMQ docker image and run
    a.) Start docker
    b.) Open Command Prompt and key in the following command:
         docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management

3.) Download RabbitMQ docker image and run using docker-compose.yml
    a.) Start docker
    b.) Open Visual Studio Code or Command Prompt cd into C:\Users\user1\Desktop\GitHub\RabbitMQ-Project
        key in the following command:
        docker-compose up
===================================================================================================================================================================
4.) To install packages used by ReactJS
    a.) open command prompt, cd into  C:\Users\Jason\Documents\GitHub\RabbitMQ-Project\ReactJS
    b.) key in the following command:
        npm install

5.) To run ReactJS
    a.) open command prompt, cd into  C:\Users\Jason\Documents\GitHub\RabbitMQ-Project\ReactJS
    b.) key in the following command:
        npm start