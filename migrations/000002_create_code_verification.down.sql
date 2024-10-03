CREATE TABLE codeVerification (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    already_used BOOLEAN DEFAULT FALSE,
    used_at TIMESTAMP,
    expire_date TIMESTAMP,
    expired BOOLEAN DEFAULT FALSE,
    used_by INTEGER,
    CONSTRAINT fk_user FOREIGN KEY (used_by) REFERENCES users(id) ON DELETE CASCADE
)