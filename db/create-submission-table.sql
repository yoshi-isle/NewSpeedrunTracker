CREATE TABLE Submission(  
    id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    create_time DATE,
    activity_id INT,
    submitters VARCHAR(50)[],
    imgur_url VARCHAR(30),
    is_approved BIT

);