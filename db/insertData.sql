INSERT INTO articles (title, contents, username, nice, created_at) VALUES (
    'firstPost', 'This is my first blog', 'saki', 4, now()
);

INSERT INTO articles (title, contents, username, nice) VALUES (
    '2nd Post', 'Second Blog Post', 'saki', 9
);

INSERT INTO comments(article_id, message, created_at) VALUES(1, '1st comment yeah', now());


INSERT INTO comments(article_id, message) VALUES(1, 'welcome');
