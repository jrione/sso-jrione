CREATE TABLE "User" (
    username CHAR(20) NOT NULL PRIMARY KEY,
    full_name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(16) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
)

INSERT INTO "User" (username,full_name,email,password) VALUES
('njir','njirlah coeg','njircoeg@tahoo.com','thispassword'),
('njir2','njirlah coeg two','njircoeg2@tahoo.com','thispassword');