CREATE TABLE
    user_base (
        user_base_id SERIAL PRIMARY KEY,
        user_base_email VARCHAR(100) UNIQUE NOT NULL, -- user email is unique
        user_base_password VARCHAR(255), -- password can be NULL for OAuth users
        user_auth_type SMALLINT NOT NULL, -- 0: credential, 1: google
        created_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    user_info (
        user_info_id INT PRIMARY KEY REFERENCES user_base (user_base_id) ON DELETE CASCADE,
        user_info_firstname VARCHAR(50),
        user_info_lastname VARCHAR(50),
        created_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    user_2fa (
        user_2fa INT PRIMARY KEY REFERENCES user_base (user_base_id) ON DELETE CASCADE,
        enabled BOOLEAN DEFAULT false,
        secret VARCHAR(255),
        created_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        WITH
            TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );