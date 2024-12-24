CREATE TABLE unit (--1
    id SERIAL PRIMARY KEY,
    unit_name varchar(100) UNIQUE,
    short_name varchar(50) UNIQUE
);

CREATE TABLE enum_classifier(--2
    id SERIAL PRIMARY KEY,
    name varchar(100),
    parent_id int,
    FOREIGN KEY (parent_id) REFERENCES enum_classifier(id) ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED,
    unit_id int,
    FOREIGN KEY (unit_id) REFERENCES unit(id) ON DELETE CASCADE
);

CREATE TABLE prod_class (--1
    id SERIAL PRIMARY KEY,
    class_name varchar(100) UNIQUE,
    unit_id int,
    FOREIGN KEY (unit_id) REFERENCES unit(id) ON DELETE CASCADE,
    parent_id int,
    FOREIGN KEY (parent_id) REFERENCES prod_class(id) ON DELETE CASCADE
);

CREATE TABLE param (--3
id SERIAL PRIMARY KEY,
name varchar(100) UNIQUE,
short_name varchar(50),
unit_id int,
FOREIGN KEY (unit_id) REFERENCES unit(id) ON DELETE CASCADE,
enum_classifier_id int,
FOREIGN KEY (enum_classifier_id) REFERENCES enum_classifier(id) ON DELETE CASCADE
);

CREATE TABLE enum_position (--2
    id SERIAL PRIMARY KEY,
    name varchar(255),
    short_name varchar(100),
    integer_value int,
    real_value float,
    string_value varchar(255),
    classifier_id int,
    FOREIGN KEY (classifier_id) REFERENCES enum_classifier(id) ON DELETE CASCADE
);

CREATE TABLE product (--1
    id_product SERIAL PRIMARY KEY,
    product_name varchar(100) UNIQUE,
    prod_class_id int,
    FOREIGN KEY (prod_class_id) REFERENCES prod_class(id) ON DELETE CASCADE,
    enum_classifier_id int,
    FOREIGN KEY (enum_classifier_id) REFERENCES enum_classifier(id) ON DELETE CASCADE
);

CREATE TABLE param_class (--3
id SERIAL PRIMARY KEY,
min_value int,
max_value int,
param_id int,
FOREIGN KEY (param_id) REFERENCES param(id) ON DELETE CASCADE,
prodclass_id int,
FOREIGN KEY (prodclass_id) REFERENCES prod_class(id) ON DELETE CASCADE
);

CREATE TABLE param_product (--3
id SERIAL PRIMARY KEY,
value int,
product_id int,
FOREIGN KEY (product_id) REFERENCES product(id_product) ON DELETE 
CASCADE,
param_class_id int,
FOREIGN KEY (param_class_id) REFERENCES param_class(id) ON DELETE 
CASCADE
);

