CREATE TABLE IF NOT EXISTS articles(
    article_id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    contents text not null,
    username VARCHAR(100) NOT NULL,
    nice INTEGER NOT NULL,
    created_at DATETIME
);

CREATE TABLE IF NOT EXISTS comments(
    comment_id integer UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    article_id INTEGER UNSIGNED NOT NULL,
    message TEXT NOT NULL,
    created_at DATETIME,
    FOREIGN KEY (article_id) REFERENCES articles(article_id)
);

INSERT INTO articles (title, contents, username, nice, created_at) VALUES (
    'firstPost', 'This is my first blog', 'saki', 4, now()
);

INSERT INTO articles (title, contents, username, nice) VALUES (
    '2nd Post', 'Second Blog Post', 'saki', 9
);

INSERT INTO comments(article_id, message, created_at) VALUES(1, '1st comment yeah', now());


INSERT INTO comments(article_id, message) VALUES(1, 'welcome');
