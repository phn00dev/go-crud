CREATE TABLE posts(
    id SERIAL PRIMARY KEY,
    post_title VARCHAR(500) NOT NULL,
    post_slug VARCHAR(520) NOT NULL,
    post_desc TEXT NOT NULL,
    post_image VARCHAR(255) NOT NULL,
    view_count INTEGER DEFAULT 0,
    user_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);