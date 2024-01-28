-- Project name: iTask
-- Description: A task management system application
-- Author: Tran Khanh

-- *Functional Requirements:

-- User
-- 1.	Register an account
-- 2.	Login to the system
-- 3.	Logout from the system
-- 4.	Reset password if forgotten
-- 5.	View the calendar
-- 6.	Use the chatbox
-- 7.	Access dashboard
-- 8.	View project details
-- 9.	Attach files to a project 
-- 10.	Remove own files from a project
-- 11.	View tasks within a project
-- 12.	Create sub-tasks within a task
-- 13.	Assign tags to tasks
-- 14.	Comment on tasks within a project
-- 15.	Edit own tasks
-- 16.	Delete own tasks
-- 17.	Change assignment for own tasks
-- 18.	Add comments to tasks
-- 19.	Edit own comments
-- Admin
-- 1.	All permissions of a User
-- 2.	Edit other users' tasks
-- 3.	Delete other users' tasks
-- 4.	Change assignment for other users' tasks
-- 5.	Manage users (add, edit, delete)
-- 6.	Manage projects (add, edit, delete)
-- 7.	Manage members of any project (add member, edit roles, delete member)
-- 8.	Manage tasks within a project (add, edit, delete)
-- 9.	Manage sub-tasks within a task (add, edit, delete)
-- 10.	Manage tags within a task (add, edit, delete)
-- 11.	Manage comments within a task (add, edit, delete)
-- Project Manager
-- 1.	All permissions of a User
-- 2.	Edit tasks within own projects
-- 3.	Delete tasks within own projects
-- 4.	Manage members of own projects (add member, change owners, delete member) 
-- 5.	Manage attachments within own projects (add, edit, delete)
-- 6.	Change assignment for tasks within own projects
-- 7.	Add comments to tasks within own projects
-- 8.	Edit own comments within own projects
-- 9.	Manage own projects (add, edit, delete)
-- 10.	Manage tasks within a project (add, edit, delete)
-- 11.	Manage sub-tasks within a task (add, edit, delete)
-- 12.	Manage tags within a task (add, edit, delete)

-- *Roles: User, Admin, Project Manager

-- *Features: 
-- 1.	Login/Logout
-- 2.	Reset password
-- 3.	Chatbox
-- 4.	Calendar
-- 5.	Dashboard
-- 6.	Project management
-- 7.	Task management
-- 8.	Sub-task management
-- 9.	Tag management
-- 10.	Comment management
-- 11.	Attachment management
-- 12.	User management
-- 13.	Role management
-- 14.	Project member management
-- 15.	Project owner management

-- *Entities: Accounts, Projects, Tasks, Tags, Comments, Attachments, Permissionsk, Project_Users

-- *Relationships:
-- 1.	User - Project: many-to-many
-- 2.	User - Task: many-to-many
-- 3.	User - Sub-task: many-to-many
-- 4.	User - Tag: many-to-many
-- 5.	User - Comment: many-to-many
-- 6.	User - Attachment: many-to-many
-- 7.	Project - Task: one-to-many
-- 8.	Project - Sub-task: one-to-many
-- 9.	Project - Tag: one-to-many
-- 10.	Project - Comment: one-to-many
-- 11.	Project - Attachment: one-to-many
-- 12.	Task - Sub-task: one-to-many
-- 13.	Task - Tag: one-to-many
-- 14.	Task - Comment: one-to-many
-- 15.	Task - Attachment: one-to-many

-- Database Schema (dbdiagram.io)

-- ? How can we save information of Project Memebers w/ roles. Task assignment, etc.?
--> Create a new table called ProjectMember, which contains project_id, user_id, role_id, created_at, updated_at
-- ? Project description, task description,... can save as bold, italic, inserted image, to do list,... 
--> Save as HTML format, datatype = text 
-- ? Should we add draft, published status for project ?
--> Column status in Project table
-- ? Should we allow a task to have multiple assginee, how can we save this information ?
--> Create a new table called TaskAssignment, which contains task_id, user_id, created_at, updated_at
-- ? Should we allow a task to have multiple tags, how can we save this information ?
--> Create a new table called TaskTag, which contains task_id, tag_id, created_at, updated_at
-- ? Can we store Tasks, Sub-tasks in the same table ?
--> Yes, we can. We can add a column called parent_task_id to store the parent task of a sub-task
-- ? How can we save permission for each role to perform Role-Based Access Control (RBAC) ?
--> Create a new table called RolePermission, which contains role_id, permission_id, created_at, updated_at
-- ? How can we use the priority of a task, a project ?
--> We can use it to sort tasks, projects in the dashboard
-- ? How can we use the privacy of a project ?
--> We can use it to determine who can view the project
-- ? How can we use the status of a project ?  
--> We can use it to determine the status of a project (draft, published, archived)


-- type User struct {
-- 	common.SQLModel
-- 	Username        string `json:"username" gorm:"column:username"`
-- 	Email           string `json:"email" gorm:"column:email"`
-- 	FullName        string `json:"full_name" gorm:"column:full_name"`
-- 	Role            int    `json:"role" gorm:"role"`
-- 	Status          int    `json:"status" gorm:"column:status"`
-- 	Password        string `json:"password" gorm:"column:password"`
-- 	Address         string `json:"address" gorm:"column:address"`
-- 	Phone           string `json:"phone" gorm:"column:phone"`
-- 	Dob             string `json:"dob" gorm:"column:dob"`
-- 	Avatar          string `json:"avatar" gorm:"avatar"`
-- 	IsEmailVerified int    `json:"is_email_verified" gorm:"is_email_verified"`
-- 	Bio             string `json:"bio" gorm:"bio"`
-- }

Table Accounts {
    id int [pk, increment] // auto-increment
    username varchar
    email varchar
    full_name varchar
    role int
    status int
    password varchar
    address varchar
    phone varchar
    dob varchar
    avatar varchar
    is_email_verified int
    bio varchar
    created_at timestamp
    updated_at timestamp
}

Table Projects {
    id int [pk, increment] // auto-increment
    title varchar
    description varchar
    status int
    priority int
    privacy int
    owner_id int
    parent_project_id int
    deadline timestamp
    started_at timestamp
    created_at timestamp
    updated_at timestamp
    thumbnail varchar
}

Table Tasks {
    id int [pk, increment] // auto-increment
    project_id int
    status int
    owner_id int
    title varchar
    priority int
    due_date timestamp
    started_at timestamp
    completed_at timestamp
    created_at timestamp
    updated_at timestamp
}

-- [Activities]