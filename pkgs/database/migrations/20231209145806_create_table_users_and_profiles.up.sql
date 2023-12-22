CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(128) NOT NULL,
    username VARCHAR(32) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id BIGINT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS user_profiles (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    nim VARCHAR(20) NULL,
    nip VARCHAR(28) NULL,
    nidn VARCHAR(28) NULL,
    no_reg_blu VARCHAR(28) NULL,
    name VARCHAR(64) NOT NULL,
    sex VARCHAR(20) NOT NULL,
    religion VARCHAR(20) NOT NULL,
    phone VARCHAR(20) NULL,
    address VARCHAR(255) NULL,
    birth_place VARCHAR(64) NOT NULL,
    birth_date DATE NOT NULL,
    prodi_id BIGINT NULL,
    concentration_id BIGINT NULL,
    pembimbing_id BIGINT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    CONSTRAINT fk_user_profile FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_user_pembimbing FOREIGN KEY(pembimbing_id) REFERENCES users(id)
);