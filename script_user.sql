CREATE TABLE user (
    id int(11) auto_increment primary key,
    name varchar(255),
    email varchar(255),
    document_number varchar(11),
    phone_number varchar(20),
    cep varchar(8),
    street varchar(255),
    city varchar(255),
    uf varchar(2),
    created_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
)

CREATE TABLE account_info (
    id int(11) auto_increment primary key,
    id_user int(11),
    balance float,
    created_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
)

CREATE TABLE transaction_logs (
    id int(11) auto_increment primary key,
    id_user_sender int(11),
    id_user_recipient int(11),
    amount float,
    status varchar(255),
    type varchar(255),
    created_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
)