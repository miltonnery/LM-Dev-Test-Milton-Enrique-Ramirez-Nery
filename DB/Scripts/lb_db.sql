-- SCHEMA CREATION
create schema life_bank_v1;
grant usage on schema life_bank_v1 to postgres;

-- TABLES CREATION
DROP TABLE IF EXISTS life_bank_v1.product_type;
CREATE TABLE life_bank_v1.product_type
(
    id            SERIAL UNIQUE NOT NULL,
    name          VARCHAR(50) NOT NULL,
    description   VARCHAR(100) NOT NULL,
    created_date  TIMESTAMP,
    created_by    VARCHAR(50),
    modified_date TIMESTAMP,
    modified_by   VARCHAR(50),
    CONSTRAINT product_type_pk PRIMARY KEY (id)
);

DROP TABLE IF EXISTS life_bank_v1.role;
CREATE TABLE life_bank_v1.role
(
    id            SERIAL UNIQUE NOT NULL,
    name          VARCHAR(50) NOT NULL,
    description   VARCHAR(100) NOT NULL,
    created_date  TIMESTAMP,
    created_by    VARCHAR(50),
    modified_date TIMESTAMP,
    modified_by   VARCHAR(50),
    CONSTRAINT role_pk PRIMARY KEY (id)
);

DROP TABLE IF EXISTS life_bank_v1.transaction_type;
CREATE TABLE life_bank_v1.transaction_type
(
    id            SERIAL UNIQUE NOT NULL,
    name          VARCHAR(50) NOT NULL,
    description   VARCHAR(50) NOT NULL,
    created_date  TIMESTAMP,
    created_by    VARCHAR(50),
    modified_date TIMESTAMP,
    modified_by   VARCHAR(50),
    CONSTRAINT transaction_type_pk PRIMARY KEY (id)
);

DROP TABLE IF EXISTS life_bank_v1.product;
CREATE TABLE life_bank_v1.product
(
    id            SERIAL UNIQUE NOT NULL,
    type          INT,
    name          VARCHAR(100) NOT NULL,
    description   VARCHAR(100) NOT NULL,
    active        BOOLEAN,
    created_date  TIMESTAMP,
    created_by    VARCHAR(50),
    modified_date TIMESTAMP,
    modified_by   VARCHAR(50),
    CONSTRAINT product_pk PRIMARY KEY (id),
    CONSTRAINT product_product_type_fk FOREIGN KEY (type) REFERENCES life_bank_v1.product_type (id)
);

DROP TABLE IF EXISTS life_bank_v1.user;
CREATE TABLE life_bank_v1.user
(
    id            SERIAL UNIQUE NOT NULL,
    role          INT NOT NULL,
    username      VARCHAR(20) NOT NULL,
    password      VARCHAR(100) NOT NULL,
    active        BOOLEAN DEFAULT TRUE,
    created_date  TIMESTAMP,
    created_by    VARCHAR(50),
    modified_date TIMESTAMP,
    modified_by   VARCHAR(50),
    CONSTRAINT user_pk PRIMARY KEY (id),
    CONSTRAINT user_role_fk FOREIGN KEY (role) REFERENCES life_bank_v1.role (id)
);

DROP TABLE IF EXISTS life_bank_v1.client;
CREATE TABLE life_bank_v1.client
(
    id            SERIAL UNIQUE NOT NULL,
    userinfo      INT,
    first_name    VARCHAR(50),
    last_name     VARCHAR(50),
    national_id   VARCHAR(20),
    email         VARCHAR(50),
    created_date  TIMESTAMP,
    created_by    VARCHAR(50),
    modified_date TIMESTAMP,
    modified_by   VARCHAR(50),
    CONSTRAINT client_pk PRIMARY KEY (id),
    CONSTRAINT client_user_fk FOREIGN KEY (userinfo) REFERENCES life_bank_v1.user (id)
);

DROP TABLE IF EXISTS life_bank_v1.beneficiary;
CREATE TABLE life_bank_v1.beneficiary
(
    id            SERIAL UNIQUE NOT NULL,
    owner         INT,
    receiver      INT,
    active        BOOLEAN DEFAULT TRUE,
    created_date  TIMESTAMP,
    created_by    VARCHAR(50),
    modified_date TIMESTAMP,
    modified_by   VARCHAR(50),
    CONSTRAINT beneficiary_pk PRIMARY KEY (id),
    CONSTRAINT beneficiary_owner_fk FOREIGN KEY (owner) REFERENCES life_bank_v1.client (id),
    CONSTRAINT beneficiary_receiver_fk FOREIGN KEY (receiver) REFERENCES life_bank_v1.client (id)
);

DROP TABLE IF EXISTS life_bank_v1.account;
CREATE TABLE life_bank_v1.account
(
    id            SERIAL UNIQUE NOT NULL,
    product       INT,
    number        VARCHAR(20),
    name_id       VARCHAR(20),
    created_date  TIMESTAMP,
    created_by    VARCHAR(50),
    modified_date TIMESTAMP,
    modified_by   VARCHAR(50),
    CONSTRAINT account_pk PRIMARY KEY (id),
    CONSTRAINT account_product_fk FOREIGN KEY (product) REFERENCES life_bank_v1.product (id)
);

DROP TABLE IF EXISTS life_bank_v1.loan;
CREATE TABLE life_bank_v1.loan
(
    id            SERIAL UNIQUE NOT NULL,
    product       INT,
    interest_rate NUMERIC(5, 2),
    amount        NUMERIC(8, 2),
    due_payment   TIMESTAMP,
    created_date  TIMESTAMP,
    created_by    VARCHAR(50),
    modified_date TIMESTAMP,
    modified_by   VARCHAR(50),
    CONSTRAINT loan_pk PRIMARY KEY (id),
    CONSTRAINT loan_product_fk FOREIGN KEY (product) REFERENCES life_bank_v1.product (id)
);

DROP TABLE IF EXISTS life_bank_v1.credit_card;
CREATE TABLE life_bank_v1.credit_card
(
    id              SERIAL UNIQUE NOT NULL,
    product         INT,
    card_number     varchar(16),
    good_thru       TIMESTAMP,
    cvv             varchar(3),
    card_limit      NUMERIC(8, 2),
    interest_rate   NUMERIC(5, 2),
    interest_amount NUMERIC(8, 2),
    montly_cut      INT,
    created_date    TIMESTAMP,
    created_by      VARCHAR(50),
    modified_date   TIMESTAMP,
    modified_by     VARCHAR(50),
    CONSTRAINT credit_card_pk PRIMARY KEY (id),
    CONSTRAINT credit_card_product_fk FOREIGN KEY (product) REFERENCES life_bank_v1.product (id)
);

DROP TABLE IF EXISTS life_bank_v1.transaction;
CREATE TABLE life_bank_v1.transaction
(
    id            SERIAL UNIQUE NOT NULL,
    client        INT,
    type          INT,
    account       INT,
    credit_card   INT,
    loan          INT,
    identifier    VARCHAR(100),
    amount        NUMERIC(8, 2),
    description   VARCHAR(100),
    created_date  TIMESTAMP,
    created_by    VARCHAR(50),
    modified_date TIMESTAMP,
    modified_by   VARCHAR(50),
    CONSTRAINT transaction_pk PRIMARY KEY (id),
    CONSTRAINT transaction_client_fk FOREIGN KEY (client) REFERENCES life_bank_v1.client (id),
    CONSTRAINT transaction_credit_card_fk FOREIGN KEY (credit_card) REFERENCES life_bank_v1.credit_card (id),
    CONSTRAINT transaction_loan_fk FOREIGN KEY (loan) REFERENCES life_bank_v1.loan (id),
    CONSTRAINT transaction_transaction_type_fk FOREIGN KEY (type) REFERENCES life_bank_v1.transaction_type (id),
    CONSTRAINT transaction_account_fk FOREIGN KEY (account) REFERENCES life_bank_v1.account (id)
);

DROP INDEX IF EXISTS life_bank_v1.transaction_identifier_idx;
CREATE INDEX transaction_identifier_idx ON life_bank_v1.transaction (identifier);
