CREATE TABLE files (
   id SERIAL PRIMARY KEY,
   user_id INT NOT NULL,
   path varchar,

   CONSTRAINT fk_user
       FOREIGN KEY(user_id)
           REFERENCES users(id)
           ON DELETE CASCADE
);
