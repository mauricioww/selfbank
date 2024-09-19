CREATE TABLE IF NOT EXISTS USERS(
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  first_name char(50) NOT NULL,
  last_name char(50) NOT NULL,
  email char(50) NOT NULL,
  -- password text NOT NULL,
  password char(70) NOT NULL, -- from client
  admin BOOLEAN default false,
  active BOOLEAN default true,
  UNIQUE(email),
  INDEX email_idx(email)
);

-- CREATE INDEX email_index ON USERS(email);
