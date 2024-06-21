![image](https://github.com/Chinnasit/admin-manage-system/assets/76206065/30173f93-bebb-48b5-be05-10af61b6e9c8)

## Frontend
The frontend of this project is based on the tutorial from this video: [User Management UI with React JS & Material UI | React Context API for State Management](https://www.youtube.com/watch?v=KTRFoouGzvY).
- **JavaScript**
- **React**
- **MUI & MUI X Data Grid**

## Backend
The backend project struct of this project reference from this video:
[Hexagonal & Clean Architecture | GoAPI Essential EP. 7](https://www.youtube.com/watch?v=4y_JXPwDuaA)
- **Go**
- **Fiber**
- **GORM**
- **PostgresSQL**

## Database
For this project, I'm using Docker Compose to run a PostgresSQL and pgAdmin4

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
docker-compose up -d
go run main.go
```
3. Start the Frontend Server on a different port:
```bash
cd frontend
npm run dev
```

## Summary
This project is a practical exercise in building a CRUD (Create, Read, Update, Delete) application using React, Material-UI (MUI) for the frontend, and Go, Fiber, GORM, and PostgresSQL for the backend.
The primary goal of this project is to develop a user management system that allows administrators to manage user accounts, including approving new users, editing user roles, and viewing user information.
