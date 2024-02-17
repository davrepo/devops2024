SELECT messages.*, users.*
FROM messages
JOIN users ON messages.author = users.username
WHERE (username = 'david' OR id IN (
    SELECT following FROM follows WHERE follower = 1
)) AND messages.flagged = false;