DROP TABLE IF EXISTS Post_Categories;
DROP TABLE IF EXISTS Post_Comments;
DROP TABLE IF EXISTS Users_Comments;
DROP TABLE IF EXISTS Users_Posts;
DROP TABLE IF EXISTS Users_Likes;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Posts;
DROP TABLE IF EXISTS Comments;
DROP TABLE IF EXISTS Categories;


CREATE TABLE Users
(
    username VARCHAR(255),
    email VARCHAR(255),
    country VARCHAR(255),
    PRIMARY KEY (username)
);

CREATE TABLE Posts
(
    post_id INTEGER,
    title VARCHAR(255),
    author VARCHAR(255),
    content VARCHAR(255),
    post_date DATETIME,
    PRIMARY KEY (post_id)
);

CREATE TABLE Comments
(
    comment_id INTEGER,
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

CREATE TABLE Post_Comments
(
    post_id INTEGER,
    comment_id INTEGER,
    PRIMARY KEY (post_id, comment_id),
    FOREIGN KEY (post_id) REFERENCES Posts(post_id),
    FOREIGN KEY (comment_id) REFERENCES Comments(comment_id)
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


INSERT INTO Users
    (username, email, country)
VALUES
    ('jdoe', 'jdoe@example.com' , 'USA'),
    ('xyz', 'xyz@example.com', 'CA'),
    ('abc', 'abc@example.com', 'MX'),
    ('pqr', 'pqr@example.com', 'UK'),
    ('mno', 'mno@example.com', 'US'),
    ('stu', 'stu@example.com', 'US'),
    ('wxy', 'wxy@example.com', 'US'),
    ('zxc', 'zxc@example.com', 'US'),
    ('qwe', 'qwe@example.com', 'US');

INSERT INTO Posts
    (post_id, title, author, content, post_date)
VALUES
    (1, 'Post 1', 'jdoe', 'This is the first post', '2018-01-01 00:00:00'),
    (2, 'Post 2', 'xyz', 'This is the second post', '2018-01-02 00:00:00'),
    (3, 'Post 3', 'qwe', 'This is the third post', '2018-01-09 00:00:00'),
    (4, 'Post 4', 'abc', 'This is the fourth post', '2018-01-03 00:00:00'),
    (5, 'Post 5', 'pqr', 'This is the fifth post', '2018-01-04 00:00:00'),
    (6, 'Post 6', 'mno', 'This is the sixth post', '2018-01-05 00:00:00'),
    (7, 'Post 7', 'stu', 'This is the seventh post', '2018-01-06 00:00:00'),
    (8, 'Post 8', 'wxy', 'This is the eighth post', '2018-01-07 00:00:00'),
    (9, 'Post 9', 'zxc', 'This is the ninth post', '2018-01-08 00:00:00');

INSERT INTO Comments
    (comment_id, post_id, content, comment_date)
VALUES
    (1, 1, 'This is the first comment', '2018-01-01 00:00:00'),
    (2, 1, 'This is the second comment', '2019-01-02 00:00:00'),
    (3, 1, 'This is the third comment', '2020-01-03 00:00:00'),
    (4, 1, 'This is the fourth comment', '2021-01-04 00:00:00'),
    (5, 1, 'This is the fifth comment', '2018-01-05 00:00:00'),
    (6, 1, 'This is the sixth comment', '2010-01-06 00:00:00'),
    (7, 1, 'This is the seventh comment', '2017-01-07 00:00:00'),
    (8, 1, 'This is the eighth comment', '2013-01-08 00:00:00'),
    (9, 1, 'This is the ninth comment', '2019-01-09 00:00:00'),
    (10, 2, 'This is the first comment', '2018-01-01 00:00:00'),
    (11, 2, 'This is the second comment', '2018-01-02 00:00:00'),
    (12, 2, 'This is the third comment', '2018-01-03 00:00:00'),
    (13, 2, 'This is the fourth comment', '2018-01-04 00:00:00'),
    (14, 2, 'This is the fifth comment', '2018-01-05 00:00:00');

INSERT INTO Categories
    (category_name)
VALUES
    ('Cooking'),
    ('Sports'),
    ('Technology');

INSERT INTO Post_Categories
    (post_id, category_name)
VALUES
    (1, 'Cooking'),
    (2, 'Sports'),
    (3, 'Technology'),
    (4, 'Cooking'),
    (5, 'Sports'),
    (6, 'Technology'),
    (7, 'Cooking'),
    (8, 'Sports'),
    (9, 'Technology');

INSERT INTO Post_Comments
    (post_id, comment_id)
VALUES
    (1, 1),
    (2, 2),
    (3, 3),
    (4, 4),
    (5, 5),
    (6, 6),
    (7, 7),
    (8, 8),
    (9, 9),
    (1, 10),
    (2, 11),
    (3, 12),
    (4, 13),
    (5, 14);

INSERT INTO Users_Comments
    (username, comment_id)
VALUES
    ('jdoe', 1),
    ('xyz', 2),
    ('abc', 3),
    ('pqr', 4),
    ('mno', 5),
    ('stu', 6),
    ('wxy', 7),
    ('zxc', 8),
    ('qwe', 9),
    ('jdoe', 10),
    ('xyz', 11),
    ('abc', 12),
    ('pqr', 13),
    ('mno', 14);

INSERT INTO Users_Posts
    (username, post_id)
VALUES
    ('jdoe', 1),
    ('xyz', 2),
    ('abc', 3),
    ('pqr', 4),
    ('mno', 5),
    ('stu', 6),
    ('wxy', 7),
    ('zxc', 8),
    ('qwe', 9);

INSERT INTO Users_Likes
    (username, post_id)
VALUES
    ('jdoe', 1),
    ('xyz', 2),
    ('abc', 3),
    ('pqr', 4),
    ('mno', 5),
    ('stu', 6),
    ('wxy', 7),
    ('zxc', 8),
    ('qwe', 9),
    ('jdoe', 9),
    ('xyz', 8),
    ('abc', 7),
    ('pqr', 6),
    ('stu', 4),
    ('wxy', 3),
    ('zxc', 2),
    ('qwe', 1);


SELECT *
FROM Users;
SELECT *
FROM Posts;
SELECT *
FROM Comments;
SELECT *
FROM Categories;
SELECT *
FROM Post_Categories;
SELECT *
FROM Post_Comments;
SELECT *
FROM Users_Comments;
SELECT *
FROM Users_Posts;
SELECT *
FROM Users_Likes;
