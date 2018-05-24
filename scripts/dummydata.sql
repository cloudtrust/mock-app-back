CREATE TABLE IF NOT EXISTS hospital.patients (
		id INT,
		first_name STRING,
		last_name STRING,
		birth_date DATE,
		avs_number STRING,
		PRIMARY KEY (id));
DELETE FROM hospital.patients WHERE true;
INSERT INTO hospital.patients(id, first_name, last_name, birth_date, avs_number) VALUES (1, 'Mariusz', 'Wiesniewski', '1984-05-04', '756.1234.3333.55');
INSERT INTO hospital.patients(id, first_name, last_name, birth_date, avs_number) VALUES (2, 'Naimeng', 'Liu', '1993-11-11', '765.4321.0303.44');
INSERT INTO hospital.patients(id, first_name, last_name, birth_date, avs_number) VALUES (3, 'Julien', 'Roch', '1984-12-03', '333.4444.5555.66');
INSERT INTO hospital.patients(id, first_name, last_name, birth_date, avs_number) VALUES (4, 'Christophe', 'Frattino', '1985-08-08', '420.1337.1337.42');

CREATE TABLE IF NOT EXISTS hospital.hospitals (
		id INT,
		name STRING,
		city STRING,
		PRIMARY KEY (id));
DELETE FROM hospital.hospitals WHERE true;
INSERT INTO hospital.hospitals(id, name, city) VALUES (1, 'CHUV', 'Lausanne');
INSERT INTO hospital.hospitals(id, name, city) VALUES (2, 'Le Samaritain', 'Vevey');
		
CREATE TABLE IF NOT EXISTS hospital.departments (
		id INT,
		name STRING,
		hospital_id INT,
		PRIMARY KEY (id));
DELETE FROM hospital.departments WHERE true;
INSERT INTO hospital.departments(id, name, hospital_id) VALUES (1, 'Natalites', 1);
INSERT INTO hospital.departments(id, name, hospital_id) VALUES (2, 'Radiologie', 1);
INSERT INTO hospital.departments(id, name, hospital_id) VALUES (3, 'Oncologie', 2);
INSERT INTO hospital.departments(id, name, hospital_id) VALUES (4, 'Dermatologie', 2);

CREATE TABLE IF NOT EXISTS hospital.doctors (
		id INT,
		first_name STRING,
		last_name STRING,
		department_id INT,
		PRIMARY KEY (id));
DELETE FROM hospital.doctors WHERE true;
INSERT INTO hospital.doctors(id, first_name, last_name, department_id) VALUES (1, 'Jane', 'Doe', 1);
INSERT INTO hospital.doctors(id, first_name, last_name, department_id) VALUES (2, 'John', 'Doe', 2);
INSERT INTO hospital.doctors(id, first_name, last_name, department_id) VALUES (3, 'Gregory', 'House', 3);
INSERT INTO hospital.doctors(id, first_name, last_name, department_id) VALUES (4, 'James', 'Wilson', 4);