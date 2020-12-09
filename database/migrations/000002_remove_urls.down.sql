ALTER TABLE blogs 
  DROP COLUMN domain;

ALTER TABLE posts 
  DROP COLUMN slug;

ALTER TABLE blogs 
  ADD COLUMN url        varchar(255) NOT NULL UNIQUE;
COMMENT ON COLUMN blogs.url IS 'The URL where this blog lives on the internet.';

ALTER TABLE posts 
  ADD COLUMN url          varchar(255) NOT NULL UNIQUE;
COMMENT ON COLUMN posts.url IS 'The full URL that a user might utilize to find this post (e.g. https://ag7if.net/posts/this-is-a-post). This will typically be generated from the post''s title.';
