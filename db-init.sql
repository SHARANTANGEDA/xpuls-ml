create user xpuls_client with password 'OWjWiVYtu/7dDJon';
alter user xpuls_client with password 'OWjWiVYtu7dDJon';

create database xpuls;
grant all privileges on database xpuls to xpuls_client;
grant all privileges on all tables in schema public to xpuls_client;
grant all privileges on all sequences in schema public to xpuls_client;
grant all privileges on all procedures in schema public to xpuls_client;
grant all privileges on schema public to xpuls_client;

