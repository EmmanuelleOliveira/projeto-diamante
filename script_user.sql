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
)