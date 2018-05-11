CREATE TABLE IF NOT EXISTS patients (
		id INT,
		first_name STRING,
		last_name STRING,
		birth_date DATE,
		avs_number STRING,
		PRIMARY KEY (id));
DELETE FROM patients WHERE true;
INSERT INTO patients(id, first_name, last_name, birth_date, avs_number) VALUES (0, 'Mariusz', 'Wiesniewski', '1984-05-04', '756.1234.3333.55');
INSERT INTO patients(id, first_name, last_name, birth_date, avs_number) VALUES (1, 'Naimeng', 'Liu', '1993-11-11', '765.4321.0303.44');
INSERT INTO patients(id, first_name, last_name, birth_date, avs_number) VALUES (2, 'Julien', 'Roch', '1984-12-03', '333.4444.5555.66');
INSERT INTO patients(id, first_name, last_name, birth_date, avs_number) VALUES (3, 'Christophe', 'Frattino', '1985-08-08', '420.1337.1337.42');