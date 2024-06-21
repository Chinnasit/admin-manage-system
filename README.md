![image](https://github.com/Chinnasit/admin-manage-system/assets/76206065/30173f93-bebb-48b5-be05-10af61b6e9c8)

## Frontend
The frontend of this project is based on the tutorial from this video: [User Management UI with React JS & Material UI | React Context API for State Management](https://www.youtube.com/watch?v=KTRFoouGzvY).

## Backend
The backend is built from scratch using the following technologies:
- **Go**: A statically typed, compiled programming language for building fast and efficient applications.
- **Fiber**: An Express-inspired web framework built on top of Fasthttp, a high-performance HTTP engine for Go.
- **GORM**: An Object-Relational Mapping (ORM) library for Go, designed to work with relational databases.
- **MySQL**: A popular open-source relational database management system.

## Database
For this project, I'm using Docker to run a MySQL database instance and connecting to it using the DBeaver client. 
You can use any database server of your choice and configure the connection details accordingly.

## Features
- User registration
- Admin user management
   - Approve new users
   - Edit user roles, email
   - View and manage user information

# Installation
1. Clone the project:
```bash
  git clone https://github.com/Chinnasit/admin-manage-system.git
```
2. Install Backend Dependencies:
```bash
cd backend
go get .
   ```
3. Install Frontend Dependencies:
```bash
cd frontend
npm install
```

# Running the Application
1. Start the Database Server:
Use any database server of your choice and configure the connection details accordingly.
3. Start the Backend Server:
```bash
cd backend
go run main.go
```
3. Start the Frontend Server on a different port:
```bash
cd frontend
npm run dev
```

## Summary
This project is a practical exercise in building a CRUD (Create, Read, Update, Delete) application using React.js, Material-UI (MUI) for the frontend, and Go, Fiber, GORM, and MySQL for the backend.
The primary goal of this project is to develop a user management system that allows administrators to manage user accounts, including approving new users, editing user roles, and viewing user information.
