-- ROLES
INSERT INTO life_bank_v1.role (name, description, created_date, created_by, modified_date, modified_by)
VALUES ('Customer', 'Normal customer of the LifeBank services', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- TRANSACTION TYPES
INSERT INTO life_bank_v1.transaction_type (name, description, created_date, created_by, modified_date, modified_by)
VALUES ('Credit', 'Credits', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.transaction_type (name, description, created_date, created_by, modified_date, modified_by)
VALUES ('Debit', 'Debits', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- PRODUCT TYPES
INSERT INTO life_bank_v1.product_type (name, description, created_date, created_by, modified_date, modified_by)
VALUES ('Account', 'Accounts catalogue', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product_type (name, description, created_date, created_by, modified_date, modified_by)
VALUES ('Credit card', 'Credit cards catalogue', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product_type (name, description, created_date, created_by, modified_date, modified_by)
VALUES ('Loan', 'Loans catalogue', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- PRODUCTS
-- Accounts
INSERT INTO life_bank_v1.product (type, name, description, active, created_date, created_by, modified_date, modified_by)
VALUES (1, 'Personal', 'This product is intended to organize the personal money of the customer', TRUE,CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product (type, name, description, active, created_date, created_by, modified_date, modified_by)
VALUES (1, 'Christmas savings', 'This product is intended to organize the Christmas Savings category', TRUE,CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product (type, name, description, active, created_date, created_by, modified_date, modified_by)
VALUES (1, 'Student savings', 'This product is intended to organize the Student Savings category', TRUE,CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
-- Credit cards
INSERT INTO life_bank_v1.product (type, name, description, active, created_date, created_by, modified_date, modified_by)
VALUES (2, 'LifeBank', 'Basic Credit card emmited by LifeBank', TRUE,CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product (type, name, description, active, created_date, created_by, modified_date, modified_by)
VALUES (2, 'LifeBank silver', 'Level A Credit Card emmited by LifeBank', TRUE,CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product (type, name, description, active, created_date, created_by, modified_date, modified_by)
VALUES (2, 'LifeBank Gold', 'Level A+ Credit Card emmited by LifeBank', TRUE,CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product (type, name, description, active, created_date, created_by, modified_date, modified_by)
VALUES (2, 'LifeBank Platinum', 'Elite Level Credit Card emmited by LifeBank', TRUE,CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
-- Loans
INSERT INTO life_bank_v1.product (type, name, description, active, created_date, created_by, modified_date, modified_by)
VALUES (3, 'Personal', 'Personal loan for customers', TRUE,CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product (type, name, description, active, created_date, created_by, modified_date, modified_by)
VALUES (3, 'New Life', 'Loan intended for buying or building a new house', TRUE,CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product (type, name, description, active, created_date, created_by, modified_date, modified_by)
VALUES (3, 'Personal Projects', 'Loan made for the execution of personal projects', TRUE,CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- USERS
INSERT INTO life_bank_v1.user (role, username, password, active, created_date, created_by, modified_date, modified_by)
VALUES (1, 'jhondoe', 'jhonpassword', true, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.user (role, username, password, active, created_date, created_by, modified_date, modified_by)
VALUES (1, 'lifebankuser', 'lifebankuserpass', true, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- CLIENTS
INSERT INTO life_bank_v1.client (user_info, first_name, last_name, national_id, email, created_date, created_by, modified_date, modified_by)
VALUES (1, 'Jhon', 'Doe', '0123456789', 'jhondoe@mail.com', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.client (user_info, first_name, last_name, national_id, email, created_date, created_by, modified_date, modified_by)
VALUES (2, 'Juan', 'Perez', '0123456349', 'lifebankuser@mail.com', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- BENEFICIARY
--INSERT INTO life_bank_v1.beneficiary (owner, receiver, identifier, email, active, created_date, created_by, modified_date, modified_by)
--VALUES (1, 2, 'ID3NT1F13R', 'beneficiateduser@mail.com',TRUE, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- ACCOUNT
INSERT INTO life_bank_v1.account (number, name_id, created_date, created_by, modified_date, modified_by)
VALUES ('00000000000000000001', 'account1', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.account (number, name_id, created_date, created_by, modified_date, modified_by)
VALUES ('00000000000000000002', 'account2', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.account (number, name_id, created_date, created_by, modified_date, modified_by)
VALUES ('00000000000000000003', 'account3', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.account (number, name_id, created_date, created_by, modified_date, modified_by)
VALUES ('00000000000000000004', 'account4', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.account (number, name_id, created_date, created_by, modified_date, modified_by)
VALUES ('00000000000000000005', 'account5', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.account (number, name_id, created_date, created_by, modified_date, modified_by)
VALUES ('00000000000000000006', 'account6', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- LOAN
INSERT INTO life_bank_v1.loan (interest_rate, amount, due_payment, created_date, created_by, modified_date, modified_by)
VALUES (7.50, 10000.00, '2025-12-31 00:00:00', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.loan (interest_rate, amount, due_payment, created_date, created_by, modified_date, modified_by)
VALUES (7.50, 50000.00, '2025-12-31 00:00:00', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- CREDIT CARD
INSERT INTO life_bank_v1.credit_card (card_number, good_thru, cvv, card_limit, interest_rate, interest_amount, monthly_cut, created_date, created_by, modified_date, modified_by)
VALUES ('4321234567890873','2025-12-31 00:00:00','123', 30000.00, 25.90, 0.00, 25, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.credit_card (card_number, good_thru, cvv, card_limit, interest_rate, interest_amount, monthly_cut, created_date, created_by, modified_date, modified_by)
VALUES ('5678908734321234','2025-12-31 00:00:00','132', 30000.00, 25.90, 0.00, 25, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- CONTRACT
--Accounts
INSERT INTO life_bank_v1.contract (client, account, product, created_date, created_by, modified_date, modified_by)
VALUES (1, 1, 1, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.contract (client, account, product, created_date, created_by, modified_date, modified_by)
VALUES (2, 2, 1, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- Credit Card
INSERT INTO life_bank_v1.contract (client, account, credit_card,  product, created_date, created_by, modified_date, modified_by)
VALUES (1, 3, 1, 4, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.contract (client, account, credit_card,  product, created_date, created_by, modified_date, modified_by)
VALUES (2, 4, 2, 5, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
-- Loan
INSERT INTO life_bank_v1.contract (client, account, loan,  product, created_date, created_by, modified_date, modified_by)
VALUES (1, 5, 1, 8, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.contract (client, account, loan,  product, created_date, created_by, modified_date, modified_by)
VALUES (2, 6, 2, 8, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- TRANSACTIONS
-- Accounts
INSERT INTO life_bank_v1.transaction (type, contract, identifier, amount, description, created_date, created_by, modified_date, modified_by)
VALUES (1, 1, 'AC123abc123abc12', 1000.00, 'First transaction, saving some money finally', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.transaction (type, contract, identifier, amount, description, created_date, created_by, modified_date, modified_by)
VALUES (1, 1, 'AC123abc123abasd2', 500.00, 'Adding more savings', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.transaction (type, contract, identifier, amount, description, created_date, created_by, modified_date, modified_by)
VALUES (2, 1, 'ACsdfssdfasdfsdfsf', 250.00, 'I´ll pay it back I promise it', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
-- Loans
INSERT INTO life_bank_v1.transaction (type, contract, identifier, amount, description, created_date, created_by, modified_date, modified_by)
VALUES (1, 5, 'LNsdfssdfasdfsdfshghfwr', 2500.00, 'First step for a little project', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.transaction (type, contract, identifier, amount, description, created_date, created_by, modified_date, modified_by)
VALUES (1, 6, 'LNsdfssdfasdfsdfsfsdfwr', 2500.00, 'First step for a little project', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
-- Credit cards
INSERT INTO life_bank_v1.transaction (type, contract, identifier, amount, description, created_date, created_by, modified_date, modified_by)
VALUES (2, 3, 'CCsdfssdfasdfsdfsfsdfwr', 2500.00, 'Amazon items', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
