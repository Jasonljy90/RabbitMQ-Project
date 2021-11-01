# RabbitMQ Project
 
1.) Create user account database on mysql(not docker) to store user account information
    a.) In mysql workbench, create an account name user1 with all administrative privilege

    b.) Use the following sql script to create the database and table:
        CREATE database MYSTOREDBRABBITMQ;
        USE MYSTOREDBRABBITMQ;
        CREATE TABLE Users (UserName VARCHAR(30) NOT NULL PRIMARY KEY, Password VARCHAR(256), FirstName VARCHAR(30), LastName VARCHAR(30), Language VARCHAR(30));

Do either 2 or 3, not both
===================================================================================================================================================================
2.) Download RabbitMQ docker image and run
    a.) Start docker
    b.) Open Command Prompt and key in the following command:
        $ docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management

3.) Download RabbitMQ docker image and run using docker-compose.yml
    a.) Start docker
    b.) Open Visual Studio Code or Command Prompt cd into C:\Users\user1\Desktop\GitHub\RabbitMQ-Project
        key in the following command:
        docker-compose up
===================================================================================================================================================================
