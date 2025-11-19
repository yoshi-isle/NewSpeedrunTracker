CREATE TABLE Activity(  
    id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    create_time DATE,
    category_id INT,
    name VARCHAR(50)

);