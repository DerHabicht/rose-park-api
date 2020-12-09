ALTER TABLE post_authors DROP CONSTRAINT posts_post_authors_post_id_fk;
ALTER TABLE post_authors DROP CONSTRAINT authors_post_authors_author_id_fk;
ALTER TABLE blog_authors DROP CONSTRAINT blog_blog_authors_blog_id_fk;
ALTER TABLE blog_authors DROP CONSTRAINT authors_blog_authors_author_id_fk;
ALTER TABLE posts DROP CONSTRAINT blogs_posts_post_id_fk;

DROP TABLE post_authors;
DROP TABLE blog_authors;
DROP TABLE posts;
DROP TABLE authors;
DROP TABLE blogs;
