CREATE TABLE publishers(
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE currency(
    id SERIAL PRIMARY KEY,
    rate DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    symbol VARCHAR(255) NOT NULL
);


CREATE TABLE publishers_info(
    id SERIAL PRIMARY KEY,
    publisher_id INTEGER NOT NULL,
    requests INTEGER NOT NULL DEFAULT 0,
    impressions INTEGER NOT NULL DEFAULT 0,
    clicks INTEGER NOT NULL DEFAULT 0,
    revenue DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    date_created DATE NOT NULL,
    FOREIGN KEY (publisher_id) REFERENCES publishers(id)
);

INSERT INTO `publishers`.`currency` (`rate`, `symbol`) VALUES ('1.00', 'USD');
INSERT INTO `publishers`.`currency` (`rate`, `symbol`) VALUES ('0.92', 'EUR');
INSERT INTO `publishers`.`currency` (`rate`, `symbol`) VALUES ('1.33', 'SGD');


INSERT INTO `publishers`.`publishers` (`id`, `name`) VALUES ('1', 'test_pub');
INSERT INTO `publishers`.`publishers_info` ( `publisher_id`, `requests`, `impressions`, `clicks`, `revenue`, `date_created`) VALUES ('1', '1000', '1000', '1000', '1000.00', '2018-01-03');
INSERT INTO `publishers`.`publishers_info` ( `publisher_id`, `requests`, `impressions`, `clicks`, `revenue`, `date_created`) VALUES ('1', '1000', '1000', '1000', '1000.00', '2018-01-04');
INSERT INTO `publishers`.`publishers_info` ( `publisher_id`, `requests`, `impressions`, `clicks`, `revenue`, `date_created`) VALUES ('1', '1000', '1000', '1000', '1000.00', '2018-01-05');
INSERT INTO `publishers`.`publishers_info` ( `publisher_id`, `requests`, `impressions`, `clicks`, `revenue`, `date_created`) VALUES ('1', '1000', '1000', '1000', '1000.00', '2018-01-06');
