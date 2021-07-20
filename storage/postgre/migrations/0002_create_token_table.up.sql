CREATE TABLE IF NOT EXISTS tokens(
    id_token UUID PRIMARY KEY,
    id_account UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,

    CONSTRAINT token_owner_account_fk
        FOREIGN KEY (id_account) REFERENCES accounts
);