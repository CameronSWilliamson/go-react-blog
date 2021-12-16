INSERT INTO Users (username, email, join_date, bio) VALUES
    ('johndoe', 'johndoe@gmail.com', '2017-01-01', 'I am a user'),
    ('janedoe', 'janedoe@gmail.com', '2017-01-01', 'I am a user'),
    ('jimdoe', 'jimdoe@gmail.com', '2017-01-01', 'I am not a user'),
    ('joey', 'joey@gmail.com', '2021-09-09', 'I am joey'),
    ('king', 'kingbrother@mail.com', '2021-09-09', 'I am the king'),
    ('queen', 'queen@email.com', '2022-09-09', 'I am the queen'),
    ('aceofspades', 'aces@gmail.com', '2021-10-10', 'I go first');

INSERT INTO Posts (title, content, post_date) VALUES
    ('John Doe Joins the platform', 'This is my very first post on this platform I am so excited to see how it does in the long term', now()),
    ('Jane Doe Eats Pasta', 'I tried pasta for the first time today. It was one of the most amazing foods I had ever had. 10/10 would try again', now()),
    ('King renounces King account on This blog', 'Someone has created the username "King" in order to appear as though they are the king. The king of what though? Definitely not England.', now()),
    ('Queen Joins the platform', 'The Queen is so excited to see how this platform works', now());
    ('Jimmy checking in', 'Hello? Is anyone inside of this box? The cake is a lie by the way', now());

INSERT INTO Comments(post_id, content, comment_date) VALUES
    (1, 'John! Welcome to the platform my friend. There is lots to do here... I dont know why you are so excited to join though.', now()),
    (1, 'John! Dont listen to those others we are happy to have you here', now()),
    (1, 'John! I am so excited to see how this platform works too!', now()),
    (2, 'Jane, have you seriously never had pasta before? I dont believe it', now()),
    (2, "I completely agree that is proposterous"),
    (2, "No I am serious this was my first time"),
    (3, "Uhh King why did you post this yourself? You're going to expose both of us"),
    (4, "This is probably a fake like the king was");

INSERT INTO Categories (category_name) VALUES
    ('Food'),
    ('Fun'),
    ('Politics'),
    ('Science'),
    ('Sports');

INSERT INTO Post_Categories (post_id, category_name) VALUES
    (1, 'Fun'),
    (2, 'Food'),
    (3, 'Politics'),
    (4, 'Politics'),
    (5, 'Fun');

INSERT INTO Users_Comments (username, comment_id) VALUES
    ('janedoe', 1),
    ('jimdoe', 2),
    ('king', 3),
    ('queen', 4),
    ('aceofspades', 5),
    ('king', 6),
    ('queen', 7),
    ('king', 8);

INSERT INTO Users_Posts (username, post_id) VALUES
    ('johndoe', 1),
    ('janedoe', 2),
    ('king', 3),
    ('queen', 4),
    ('jimdoe', 5);

INSERT INTO Users_Likes (username, post_id) VALUES
    ('johndoe', 1),
    ('johndoe', 2),
    ('johndoe', 3),
    ('johndoe', 4),
    ('johndoe', 5),
    ('jimdoe', 1),
    ('janedoe', 6),
    ('king', 4),
    ('queen', 3);

INSERT INTO Users_Follows (username, following) VALUES
    ('johndoe', 'king'),
    ('johndoe', 'queen'),
    ('johndoe', 'aceofspades'),
    ('janedoe', 'king'),
    ('janedoe', 'queen'),
    ('jimdoe', 'king'),
    ('jimdoe', 'queen'),
    ('jimdoe', 'aceofspades'),
    ('king', 'queen'),
    ('king', 'aceofspades'),
    ('queen', 'aceofspades');