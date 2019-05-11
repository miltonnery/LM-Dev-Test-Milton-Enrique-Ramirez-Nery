-- PRODUCT TYPES
INSERT INTO life_bank_v1.product_type (name, description, created_date, created_by, modified_date, modified_by)
VALUES ('Account', 'Accounts catalogue', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product_type (name, description, created_date, created_by, modified_date, modified_by)
VALUES ('Credit card', 'Credit cards catalogue', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product_type (name, description, created_date, created_by, modified_date, modified_by)
VALUES ('Loan', 'Loans catalogue', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- ROLES
INSERT INTO life_bank_v1.role (name, description, created_date, created_by, modified_date, modified_by)
VALUES ('Customer', 'Normal customer of the LifeBank services', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- TRANSACTION TYPES
INSERT INTO life_bank_v1.transaction_type (name, description, created_date, created_by, modified_date, modified_by)
VALUES ('Credit', 'Credits', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.transaction_type (name, description, created_date, created_by, modified_date, modified_by)
VALUES ('Debit', 'Debits', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- PRODUCTS
-- Accounts
INSERT INTO life_bank_v1.product (type, name, description, created_date, created_by, modified_date, modified_by)
VALUES (1, 'Personal', 'This product is intended to organize the personal money of the customer',CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product (type, name, description, created_date, created_by, modified_date, modified_by)
VALUES (1, 'Christmas savings', 'This product is intended to organize the Christmas Savings category',CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product (type, name, description, created_date, created_by, modified_date, modified_by)
VALUES (1, 'Student savings', 'This product is intended to organize the Student Savings category',CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
-- Credit cards
INSERT INTO life_bank_v1.product (type, name, description, created_date, created_by, modified_date, modified_by)
VALUES (2, 'LifeBank', 'Basic Credit card emmited by LifeBank',CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product (type, name, description, created_date, created_by, modified_date, modified_by)
VALUES (2, 'LifeBank silver', 'Level A Credit Card emmited by LifeBank',CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product (type, name, description, created_date, created_by, modified_date, modified_by)
VALUES (2, 'LifeBank Gold', 'Level A+ Credit Card emmited by LifeBank',CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product (type, name, description, created_date, created_by, modified_date, modified_by)
VALUES (2, 'LifeBank Platinum', 'Elite Level Credit Card emmited by LifeBank',CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
-- Loans
INSERT INTO life_bank_v1.product (type, name, description, created_date, created_by, modified_date, modified_by)
VALUES (3, 'Personal', 'Personal loan for customers',CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product (type, name, description, created_date, created_by, modified_date, modified_by)
VALUES (3, 'New House', 'Loan intended for buying or building a new house',CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.product (type, name, description, created_date, created_by, modified_date, modified_by)
VALUES (3, 'Personal Projects', 'Loan made for the execution of personal projects',CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- USERS
INSERT INTO life_bank_v1.user (role, username, password, active, created_date, created_by, modified_date, modified_by)
VALUES (1, 'jhondoe', 'jhonpassword', true, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.user (role, username, password, active, created_date, created_by, modified_date, modified_by)
VALUES (1, 'lifebankuser', 'lifebankuserpass', true, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- CLIENTS
INSERT INTO life_bank_v1.client (userinfo, first_name, last_name, national_id, email, created_date, created_by, modified_date, modified_by)
VALUES (1, 'Jhon', 'Doe', '0123456789', 'jhondoe@mail.com', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.client (userinfo, first_name, last_name, national_id, email, created_date, created_by, modified_date, modified_by)
VALUES (2, 'Juan', 'Perez', '0123456349', 'lifebankuser@mail.com', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- BENEFICIARY
INSERT INTO life_bank_v1.beneficiary (owner, receiver, active, created_date, created_by, modified_date, modified_by)
VALUES (1, 2, TRUE, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- ACCOUNT
INSERT INTO life_bank_v1.account (client, product, number, name_id, created_date, created_by, modified_date, modified_by)
VALUES (1, 1, '00000000000000000001', 'jhondoe', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.account (client, product, number, name_id, created_date, created_by, modified_date, modified_by)
VALUES (2, 1, '00000000000000000002', 'juanperez', CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');

-- TRANSACTIONS
INSERT INTO life_bank_v1.transaction (type, account, identifier, amount, created_date, created_by, modified_date, modified_by)
VALUES (1, 1, '123abc123abc12', 100.00, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
INSERT INTO life_bank_v1.transaction (type, account, identifier, amount, created_date, created_by, modified_date, modified_by)
VALUES (1, 2, '1234ñk32j2k2l3', 50.00, CURRENT_TIMESTAMP, 'lb_configurator', CURRENT_TIMESTAMP, 'lb_configurator');
