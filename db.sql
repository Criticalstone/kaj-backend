# CREATE DATABASE products;
USE products;
CREATE TABLE products
(
  id              INT unsigned NOT NULL AUTO_INCREMENT, # Unique ID for the record
  name            VARCHAR(150) NOT NULL,                # Name of the cat
  PRIMARY KEY     (id)                                  # Make the id the primary key
);