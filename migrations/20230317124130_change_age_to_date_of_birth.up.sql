ALTER TABLE users ADD COLUMN date_of_birth TIMESTAMP;
ALTER TABLE users ALTER COLUMN date_of_birth SET NOT NULL;
ALTER TABLE users DROP COLUMN age;