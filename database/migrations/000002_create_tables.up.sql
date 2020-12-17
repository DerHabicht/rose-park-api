CREATE TABLE blogs (
  id         SERIAL NOT NULL, 
  created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL, 
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL, 
  deleted_at timestamp, 
  name       varchar(255) NOT NULL, 
  url        varchar(255) NOT NULL UNIQUE, 
  PRIMARY KEY (id));
COMMENT ON TABLE blogs IS 'The blogs supported by this backend.';
COMMENT ON COLUMN blogs.name IS 'The human-readable name of the blog.';
COMMENT ON COLUMN blogs.url IS 'The URL where this blog lives on the internet.';

CREATE TABLE authors (
  id         SERIAL NOT NULL, 
  public_id  uuid DEFAULT UUID_GENERATE_V4() NOT NULL UNIQUE, 
  created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL, 
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL, 
  deleted_at timestamp, 
  name       varchar(255) NOT NULL, 
  email      varchar(255) NOT NULL, 
  bio        text NOT NULL, 
  PRIMARY KEY (id));
COMMENT ON TABLE authors IS 'The authors that write for one or more blogs.';
COMMENT ON COLUMN authors.public_id IS 'This author''s lookup ID. Primarily used for looking up the author''s bio on the blog''s website.';
COMMENT ON COLUMN authors.name IS 'The name that is displayed on each blog post that this author writes.';
COMMENT ON COLUMN authors.email IS 'Primary contact email for this author (not public).';
COMMENT ON COLUMN authors.bio IS 'Markdown-formatted author bio.';

CREATE TABLE posts (
  id           SERIAL NOT NULL, 
  created_at   timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL, 
  updated_at   timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL, 
  deleted_at   timestamp, 
  url          varchar(255) NOT NULL UNIQUE, 
  title        int4 NOT NULL, 
  publish_date date NOT NULL, 
  body         text NOT NULL, 
  blog_id      int4 NOT NULL, 
  PRIMARY KEY (id));
COMMENT ON TABLE posts IS 'Individual posts that have been published on one of the blogs.';
COMMENT ON COLUMN posts.url IS 'The full URL that a user might utilize to find this post (e.g. https://ag7if.net/posts/this-is-a-post). This will typically be generated from the post''s title.';
COMMENT ON COLUMN posts.title IS 'The display title of this post.';
COMMENT ON COLUMN posts.publish_date IS 'The date when this post should become visible to the public.';
COMMENT ON COLUMN posts.body IS 'The markdown-formatted body of the post.';
COMMENT ON COLUMN posts.blog_id IS 'References the blog on which this post was published.';

CREATE TABLE blog_authors (
  blog_id   int4 NOT NULL, 
  author_id int4 NOT NULL, 
  PRIMARY KEY (blog_id, 
  author_id));

CREATE TABLE post_authors (
  post_id   int4 NOT NULL, 
  author_id int4 NOT NULL, 
  PRIMARY KEY (post_id, 
  author_id));
ALTER TABLE posts ADD CONSTRAINT blogs_posts_post_id_fk FOREIGN KEY (blog_id) REFERENCES blogs (id) ON DELETE Cascade;
ALTER TABLE blog_authors ADD CONSTRAINT authors_blog_authors_author_id_fk FOREIGN KEY (author_id) REFERENCES authors (id);
ALTER TABLE blog_authors ADD CONSTRAINT blog_blog_authors_blog_id_fk FOREIGN KEY (blog_id) REFERENCES blogs (id);
ALTER TABLE post_authors ADD CONSTRAINT authors_post_authors_author_id_fk FOREIGN KEY (author_id) REFERENCES authors (id);
ALTER TABLE post_authors ADD CONSTRAINT posts_post_authors_post_id_fk FOREIGN KEY (post_id) REFERENCES posts (id);
