ALTER TABLE movies
ADD CONSTRAINT users_id_movies_user_id_fk
FOREIGN KEY (user_id)
REFERENCES users(id)