CREATE TABLE send_messages(
    serial_id       BIGSERIAL PRIMARY KEY,
    idempotency_key UUID        NOT NULL,
    external_id     VARCHAR(64) NOT NULL,
    message         BYTEA       NOT NULL,
    created_at      BIGINT      NOT NULL
);

create table send_message_status(
    serial_id        BIGSERIAL PRIMARY KEY,
    message_id BIGINT    NOT NULL,
    status           SMALLINT  NOT NULL, -- 1. CREATED, 2. SENDED, 3. SYNC, 4. INVALID_MESSAGE, 5. ERROR, 6. MAX_RETRY,  
    created_at       BIGINT    NOT NULL 
);