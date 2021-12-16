DROP TABLE IF EXISTS Post_Categories;
DROP TABLE IF EXISTS Post_Comments;
DROP TABLE IF EXISTS Users_Comments;
DROP TABLE IF EXISTS Users_Posts;
DROP TABLE IF EXISTS Users_Likes;
DROP TABLE IF EXISTS Users_Follows;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Posts;
DROP TABLE IF EXISTS Comments;
DROP TABLE IF EXISTS Categories;


CREATE TABLE Users
(
    username VARCHAR(255),
    email VARCHAR(255),
    join_date DATE,
    bio VARCHAR(255),
    PRIMARY KEY (username)
);

CREATE TABLE Posts
(
    post_id INTEGER AUTO_INCREMENT,
    title VARCHAR(255),
    content VARCHAR(255),
    post_date DATETIME,
    PRIMARY KEY (post_id)
);

CREATE TABLE Comments
(
    comment_id INTEGER AUTO_INCREMENT,
    post_id INTEGER,
    content VARCHAR(255),
    comment_date DATETIME,
    PRIMARY KEY (comment_id)
);

CREATE TABLE Categories
(
    category_name VARCHAR(255),
    PRIMARY KEY (category_name)
);

CREATE TABLE Post_Categories
(
    post_id INTEGER,
    category_name VARCHAR(255),
    PRIMARY KEY (post_id, category_name),
    FOREIGN KEY (post_id) REFERENCES Posts(post_id),
    FOREIGN KEY (category_name) REFERENCES Categories(category_name)
);

CREATE TABLE Users_Comments
(
    username VARCHAR(255),
    comment_id INTEGER,
    PRIMARY KEY (username, comment_id),
    FOREIGN KEY (username) REFERENCES Users(username),
    FOREIGN KEY (comment_id) REFERENCES Comments(comment_id)
);

CREATE TABLE Users_Posts
(
    username VARCHAR(255),
    post_id INTEGER,
    PRIMARY KEY (username, post_id),
    FOREIGN KEY (username) REFERENCES Users(username),
    FOREIGN KEY (post_id) REFERENCES Posts(post_id)
);

CREATE TABLE Users_Likes
(
    username VARCHAR(255),
    post_id INTEGER,
    PRIMARY KEY (username, post_id),
    FOREIGN KEY (username) REFERENCES Users(username),
    FOREIGN KEY (post_id) REFERENCES Posts(post_id)
);

CREATE TABLE Users_Follows
(
    username VARCHAR(255),
    following VARCHAR(255),
    PRIMARY KEY (username, following),
    FOREIGN KEY (username) REFERENCES Users(username),
    FOREIGN KEY (following) REFERENCES Users(username)
);