-- Shows the users who liked a post and 
-- are also following the poster
SELECT COUNT(*)
FROM Users_Follows uf 
	JOIN Users_Likes ul on uf.`following` = ul.username 
	JOIN Posts p on ul.post_id = p.post_id 
WHERE ul.post_id = 1

-- Counting the number of posts in each category and grouping by likes
SELECT category_name, COUNT(*) as posts, likes
FROM Posts p 
	JOIN Post_Categories pc USING (post_id)
	JOIN Categories c USING (category_name)
	JOIN (SELECT (COUNT(*)) AS likes, category_name 
	      FROM Users_Likes
		  	JOIN Post_Categories USING (post_id)
		  	GROUP BY category_name) q1 USING (category_name)
GROUP BY category_name
ORDER BY posts DESC, likes DESC

SELECT username, posts, comments
FROM Users u 
	JOIN (SELECT username, COUNT(*) AS posts
		  FROM Users_Posts up
		  GROUP BY username) q1 USING (username)
	JOIN (SELECT username, COUNT(*) AS comments 
		  FROM Users_Comments uc 
		  GROUP BY username) q2 USING (username)
WHERE u.join_date > "2021-06-01" OR posts > 1 OR comments > 1
ORDER BY posts desc, comments desc