
# MEAN Stack with MySQL

MEAN stack is collection of JavaScript framework MongoDB, ExpressJS, Angular and NodeJs. Using MEAN, you can develop fast and robust web application.

But in this tutorial we will create an web application using MEAN (MySQL, ExpressJs, Angular, NodeJS). We are using M for MySQL instead of MongoDB. In this tutorial include features introduction, Angular with NodeJs, User Registration, User Login using JWT Token, CRUD Operations, Pagination using Server Side, Change Password etc.

All features available on this link MEAN (MySQL, ExpressJs, Angular, NodeJS). Next features will be available soon.

MySQL the most popular Open Source SQL database management system is developed, distributed and supported by Oracle Corporation. MySQL is the most popular language for adding, accessing and managing content in a database. In this MEAN adding and accesing data.

ExpressJS is a NodeJS web application framework that provides a robust set of features for web and mobile applications. In MEAN it is use for creating NodeJS app with set of features.

Angular is framework for creating single page application. Angular for building mobile and desktop applictions. In MEAN it use for client side.

NodeJS is a plateform thats run JavaScript on server side. In MEAN NodeJs use for creating API.

In this lesson we will learn setup mysql, nodejs, create nodejs app using express framework and connect mysql with nodejs.

## Setup MySQL

Before use mysql in Node.js must be ensured MySQL Database should be installed in your machine. If MySQL is not installed in your machine then you can install from Installing MySQL documentation.

https://dev.mysql.com/doc/refman/8.0/en/installing.html

After setup MySQL then create database meanapp and create two tables users and products

```
CREATE TABLE `products` (
  `id` int(11) PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `sku` varchar(20) DEFAULT NULL,
  `price` decimal(8,2) NOT NULL,
  `is_active` tinyint(1) NOT NULL,
  `created_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `users` (
  `id` int(11) PRIMARY KEY AUTO_INCREMENT NOT NULL,
  `username` varchar(20) DEFAULT NULL,
  `password` text DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `first_name` varchar(30) DEFAULT NULL,
  `last_name` varchar(30) DEFAULT NULL,
  `is_active` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

NodeJS should be installed in your system. If not installed then you can watch this video [Node.js Installation](https://www.youtube.com/watch?v=E9ECq9qDwMo).

## Create NodeJS App using ExpressJS

### Intallation express-generator tool

We will install express-generator tools globaly. If you have already installed express-generator tool then no need to install.

```
npm install express-generator -g
```

### Create a Application
Create a application using express-generator tool
```
express --view=hbs meanapi
```

### Install Dependencies
```
cd meanapi
npm install
```
### Install mysql node modules in node.js
Install mysql node modules in your application.
```
npm install mysql
npm install -g @angular/cli
```

The file `db.js` shall be moved to the root folder for database connection and use this connection in your entire project. Connect your database using your MySQL database credential host, database user, database password and database name.

### Check Database Connection
Now `var db = require("../db");` file in `routes/users.js`

### nodemon
nodemon is a tool that helps run node.js based applications by automatically restarting the node application when file changes in the directory are detected.
```
npm install -g nodemon
```

### Run NodeJs Application
You can run nodejs application using below commands:

```
nodemon start
```
Access nodejs application using url `http://localhost:3000/`
