ALTER TABLE blogs 
  DROP COLUMN url;

ALTER TABLE posts 
  DROP COLUMN url;

ALTER TABLE blogs 
  ADD COLUMN domain varchar(255) NOT NULL UNIQUE;
COMMENT ON COLUMN blogs.domain IS 'The domain name where this blog lives on the internet.';

ALTER TABLE posts 
  ADD COLUMN slug varchar(255) NOT NULL UNIQUE;
COMMENT ON COLUMN posts.slug IS 'The part of the URL that, in connection with the blog''s domain property, would be used to specify this post. E.g. if the post lives at https://ag7if.net/posts/how-computers-work-pt-1, the slug would be ''how-computers-work-pt-1''.';
