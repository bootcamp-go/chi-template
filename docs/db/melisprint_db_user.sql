-- DCL (Data Control Language)
-- drop user
DROP USER IF EXISTS 'melisprint_user'@'localhost';

-- create user 
CREATE USER 'melisprint_user'@'localhost' IDENTIFIED BY 'MeLiSprint_Pass_123!';

-- grant privileges
GRANT ALL PRIVILEGES ON melisprint.* TO 'melisprint_user'@'localhost';

-- flush privileges
FLUSH PRIVILEGES;