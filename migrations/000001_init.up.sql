CREATE TABLE IF NOT EXISTS carrier_quotes (
    id SERIAL PRIMARY KEY,
    carrier_name VARCHAR(255) NOT NULL,
    carrier_service VARCHAR(255) NOT NULL,
    deadline_days INTEGER NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
