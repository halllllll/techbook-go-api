CREATE TABLE IF NOT EXISTS articles(
    article_id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    contents TEXT NOT NULL,
    username VARCHAR(100) NOT NULL,
    created_at DATETIME
);

CREATE TABLE IF NOT EXISTS comments(
    comment_id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    article_id INTEGER NOT NULL,
    message text NOT NULL,
    created_at DATETIME,
    FOREING KEY (article_id) REFERENCES articles(article_id)
);

INSERT INTO articles (title, contents, username, nice, created_at) VALUES (
    'firstPoost', 'This is my first blog', 'saki', 2, now()
);

INSERT INTO articles (title, contents, username, nice) VALUES (
    '2nd Post', 'Second Blog Post', 'saki', 4
);

INSERT INTO comments(article_id, message, created_at) VALUES(1, '1st comment yeah', now());


INSERT INTO comments(article_id, message) VALUES(1, 'welcome');