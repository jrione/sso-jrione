CREATE TABLE IF NOT EXISTS "User" (
  username CHAR(20) NOT NULL PRIMARY KEY,
  full_name VARCHAR(50) NOT NULL,
  email VARCHAR(50) NOT NULL,
  password VARCHAR(60) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);


CREATE TABLE IF NOT EXISTS "refresh_token" (
  id_refresh SERIAL PRIMARY KEY,
  username CHAR(20) NOT NULL,
  refresh_token TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  FOREIGN KEY (username) REFERENCES "User"(username) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE INDEX refresh_token_index ON refresh_token(username)

CREATE OR REPLACE FUNCTION updated_auto_func() RETURNS TRIGGER
LANGUAGE 'plpgsql'
AS $$
BEGIN
	NEW.updated_at = CURRENT_TIMESTAMP;
	RETURN NEW;
END;
$$;


CREATE OR REPLACE TRIGGER updated_auto 
  BEFORE UPDATE 
  ON 
    "User"
  FOR EACH ROW
EXECUTE PROCEDURE updated_auto_func();

CREATE OR REPLACE TRIGGER updated_auto 
  BEFORE UPDATE 
  ON 
   "refresh_token"
  FOR EACH ROW
EXECUTE PROCEDURE updated_auto_func();

INSERT INTO "User" (username,full_name,email,password) VALUES
('njir','njirlah coeg','njircoeg@tahoo.com','thispassword'),
('njir2','njirlah coeg two','njircoeg2@tahoo.com','thispassword');

