ALTER TABLE peers
    ADD COLUMN agent_version VARCHAR(255),
    ADD COLUMN protocol      VARCHAR(255) ARRAY;
