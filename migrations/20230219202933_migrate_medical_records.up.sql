CREATE TABLE IF NOT EXISTS medical_records (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    uuid UUID NOT NULL DEFAULT gen_random_uuid(), 
    health_care_provider_type BOOLEAN NOT NULL,
    health_care_provider_name VARCHAR(100) NOT NULL,
    emergency_medical_service BOOLEAN NOT NULL,
    multiple_sclerosis_type VARCHAR(4) NOT NULL,
    laboral_condition VARCHAR(100) NOT NULL,
    conmorbidity BOOLEAN NOT NULL,
    treating_neurologist VARCHAR(100) NOT NULL,
    support_network BOOLEAN NOT NULL,
    is_disabled BOOLEAN NOT NULL,
    educational_level VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT FK_user FOREIGN KEY(user_id)
    REFERENCES users(id)
);