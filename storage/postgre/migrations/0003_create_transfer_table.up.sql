CREATE TABLE IF NOT EXISTS transfers(
    id UUID PRIMARY KEY,
    origin_account_id UUID NOT NULL,
    destination_account_id UUID NOT NULL,
    amount INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,

    CONSTRAINT  transfer_origin_account_id_fk
        FOREIGN KEY (origin_account_id) REFERENCES accounts,

    CONSTRAINT transfer_destination_account_id_fk
        FOREIGN KEY (destination_account_id) REFERENCES accounts
)