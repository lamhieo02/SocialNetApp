-- Use the database
USE socialnetapp;

-- Create the user table
CREATE TABLE user (
  id INT AUTO_INCREMENT PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  date_of_birth VARCHAR(20) NOT NULL,
  email VARCHAR(50) NOT NULL ,
  user_name VARCHAR(50) UNIQUE NOT NULL,
  `password` VARCHAR(250) NOT NULL,
  salt VARCHAR(250) NOT NULL,
  INDEX idx_username (user_name,email)
);

-- Create the post table
CREATE TABLE post (
  id INT AUTO_INCREMENT PRIMARY KEY,
  content_text VARCHAR(500),
  content_image_path VARCHAR(255),
  user_id INT NOT NULL,
  visible BOOL NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP NOT NULL,
  FOREIGN KEY (user_id) REFERENCES user(id)
);

-- Create the friendship table
CREATE TABLE following (
    user_id INT NOT NULL ,
    friend_id INT NOT NULL ,
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (friend_id) REFERENCES user(id),
    PRIMARY KEY (user_id, friend_id)
);

-- Create the comment table
CREATE TABLE comment (
     id INT AUTO_INCREMENT PRIMARY KEY,
     content VARCHAR(255) NOT NULL,
     created_at TIMESTAMP NOT NULL,
     updated_at TIMESTAMP NOT NULL,
     deleted_at TIMESTAMP NOT NULL,
     user_id INT NOT NULL ,
     post_id INT NOT NULL ,
     FOREIGN KEY (user_id) REFERENCES user(id),
     FOREIGN KEY (post_id) REFERENCES post(id)
);

-- Create the like table
CREATE TABLE `like` (
    user_id INT NOT NULL,
    post_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (post_id) REFERENCES post(id),
    PRIMARY KEY (post_id, user_id)
);