CREATE TABLE IF NOT EXISTS "tb_student" (
  "id" varchar(5) NOT NULL,
  "name" varchar(255) NOT NULL,
  "age" int NOT NULL,
  "grade" int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


INSERT INTO "tb_student" ("id", "name", "age", "grade") VALUES
('B001', 'Jason Bourne', 29, 1),
('B002', 'James Bond', 27, 1),
('E001', 'Ethan Hunt', 27, 2),
('W001', 'John Wick', 28, 2);

ALTER TABLE "tb_student" ADD PRIMARY KEY ("id");