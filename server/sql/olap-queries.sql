-- Shows the users who liked a post and 
-- are also following the poster
SELECT COUNT(*)
FROM Users_Follows uf 
	JOIN Users_Likes ul on uf.`following` = ul.username 
	JOIN Posts p on ul.post_id = p.post_id 
WHERE ul.post_id = 1
