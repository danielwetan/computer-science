
-- create database
create database database_system_security;
use database_system_security;

-- create table
create table payments (id int primary key);
create table users (id int primary key);
create table products (id int primary key);
create table transactions (id int primary key);
create table carts (id int primary key);
create table wishlists (id int primary key);
create table stores (id int primary key);
create table promotions (id int primary key);
create table shippings (id int primary key);
create table addresses (id int primary key);

-- create user
CREATE user daniel;
CREATE user amry;
CREATE user septa;
CREATE user yuri;

-- grant access
GRANT SELECT ON carts TO daniel;
GRANT SELECT ON wishlists TO daniel;
GRANT INSERT ON stores TO daniel;
GRANT SELECT ON promotions TO amry;
GRANT UPDATE ON shippings TO amry;
GRANT UPDATE ON addresses TO amry;
GRANT SELECT ON users TO septa;
GRANT DELETE ON users TO septa;
GRANT SELECT ON products TO yuri;
GRANT DELETE ON products TO yuri;

-- revoke
REVOKE DELETE ON products FROM yuri;
