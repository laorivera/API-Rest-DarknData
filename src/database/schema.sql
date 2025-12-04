CREATE TABLE IF NOT EXISTS request_logs (
    id SERIAL PRIMARY KEY,
    path VARCHAR(500),
    method VARCHAR(10),
    status_code INTEGER,
    response_size INTEGER,
    duration_ms INTEGER,
    user_agent TEXT,
    ip_address VARCHAR(45),
    created_at TIMESTAMP DEFAULT NOW()
);
