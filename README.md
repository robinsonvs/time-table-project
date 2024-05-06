
## Time Table Allocation - Documentation - in progress

### Overview

The Time Table Allocation System for IT Courses is an application developed to automate the process of offering and allocating teachers for undergraduate courses in Information Technology (IT) at a university. The application aims to streamline and facilitate the generation of proposals for class offerings for each discipline, ensuring that the restrictions and assumptions established by the undergraduate directorate are met.

### Main Features

* CRUD (Create, Read, Update, Delete) of semesters, courses, disciplines, classes, professors, and availability.
* Automatic generation of class offering proposals.
* Export of class offering proposals by course.
* Export of the total offering.
*  Query of hours per professor.


## Vision Document
    
### Objective of the Document
This document aims to provide an overview of the Teacher Allocation system for Information Technology (IT) courses at a university. It highlights the necessary resources, the needs of key users, and the reasons behind these needs. Details on how the system meets these needs will be addressed in the functional and non-functional requirements documents.

### Current Situation
Currently, the teacher allocation process involves four main stages: defining the disciplines offered in each IT course, creating classes for each discipline, establishing schedules for the classes, and allocating suitable teachers for each class. This process is manual, carried out by the course coordinator using Excel spreadsheets. Upon completion of the planning, the information is passed to the university's secretariat for registration in the academic system.

### Project Objectives
Automate the teacher allocation process to streamline and facilitate the generation of class offering proposals.

### Project Scope
The system will automate the process of offering and allocating teachers for the university's IT undergraduate courses. There will be no integration with the university's academic system. The end result will be the generation of an Excel file for submission to the secretariat.

### Assumptions

- Records of all IT courses, including in-person and hybrid/online separately.
- List of disciplines for each course.
- Maximum credit goals for the upcoming semester.
- Information about classes, including credits, days of the week, and times.
- List of teachers, their eligible disciplines, and availability.
- Teacher workload goals.

### Constraints

- Compliance with credit goals per semester for each course.
- Equitable distribution of disciplines across semesters.
- Avoiding schedule overlaps for disciplines and teachers.
- Maintenance of teacher hours between semesters.

### Product Overview
### Functional Requirements:
- CRUD (Create, Read, Update, Delete) for semester credits, courses, disciplines, classes, teachers, availability, and semesters.
- Generation and updating of class offering proposals.
- Export of proposals per course and total.

### Non-Functional Requirements:
- Web-based system.
- Security.
- 24/7 availability.

### Detailed Functional Requirements

Semester (RF001)
- CRUD for academic periods.
- Information: semester.

Discipline (RF002)
- CRUD for disciplines.
- Information: name, number of credits.

Course (RF003)
- CRUD for IT courses.
- Information: name, modality, location, disciplines.

Teacher (RF004)
- CRUD for teachers.
- Information: name, hours to be allocated, eligible disciplines, available times.

Parameterization (RF005)
- CRUD for parameters for class/teacher offering planning.
- Information: semester, course, maximum credits, disciplines to be offered, number of classes.

Results (RF006)
- Automatic generation of class offering proposal.
- Considerations: disciplines, number of classes, teacher availability.

Export (RF007, RF008)
- Export of proposal per course or total to Excel spreadsheet.

Teacher Hours Inquiry (RF009)
- Inquiry of teachers with different workload than planned.
- Information: teacher name, hours to be allocated, allocated hours.

### Use Case

![](/img/usecase20240505.png)

### Class Diagram

![](/img/classDiagram20240505.png)

### ER Diagram of the Database

![](/img/database20240505.png)

### Main frameworks used:

* [PostgreSQL](https://www.postgresql.org/) - database
* [SQCL](https://sqlc.dev/) - for handling database queries
* [Golang Migrate](https://github.com/golang-migrate/migrate) - for handling database migrations
* [Go Chi](https://github.com/go-chi/chi) - for creating routes
* [Go playground validator](https://github.com/go-playground/validator) - responsible for validating input data
* [Swaggo](https://github.com/swaggo/http-swagger) - for creating OpenAPI standard documentation
* [Viper](https://github.com/spf13/viper) - for managing environment variables
* [Docker](https://www.docker.com/get-started/) - for running the database
    
    

### Structure

### Detailed description of the main folders:

* **cmd**: Here will be the main.go files, responsible for starting the application.
* **config**: Here will be some configs, such as environment variables, logs.
* **internal**: This is where the business logic will reside.
* **internal/dto**: Data types allowed to enter the application.
* **internal/entity**: Application entities.
* **internal/handler**: Routing files (controllers).
* **internal/database**: Files related to the database.
* **internal/database/migrations**: Migration files.
* **internal/database/queries**: SQL query files for database queries.
* **internal/database/sqlc**: Files automatically generated by sqlc.
* **internal/repository**: Repository layer, it wouldn't be necessary to add this layer, we could use the sqlc structures directly, but let's add this layer to make it a little more decoupled from sqlc.
* **internal/service**: Service layer, where the business logic (usecases) will reside. 


### Next Steps

### This project is constantly evolving. Some planned future improvements include:

* Implementation of a user-friendly interface.
* Addition of authentication and access control.
* Improvements in automatic generation of class offering proposals.


