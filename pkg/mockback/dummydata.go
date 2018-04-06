package mockback

import "time"

var dummyData []Hospital

var sequence int32

func next() int32 {
	var value = sequence
	sequence++
	return value
}

// GetDummyData returns some dummy data
func GetDummyData() []Hospital {
	if len(dummyData) == 0 {
		// We create hospitals
		var chuv = Hospital{ID: next(), Name: "CHUV", City: "Lausanne"}
		var samaritain = Hospital{ID: next(), Name: "Le Samaritain", City: "Vevey"}

		// We create departments
		var natalitesChuv = Department{ID: next(), Name: "Natalités", Hospital: chuv}
		var radiologieChuv = Department{ID: next(), Name: "Radiologie", Hospital: chuv}
		var oncologieSamaritain = Department{ID: next(), Name: "Oncologie", Hospital: samaritain}
		var dermatologieSamaritain = Department{ID: next(), Name: "Dermatologie", Hospital: samaritain}
		chuv.Departments = []Department{natalitesChuv, radiologieChuv}
		samaritain.Departments = []Department{oncologieSamaritain, dermatologieSamaritain}

		// We create doctors
		var janeDoe = Doctor{ID: next(), FirstName: "Jane", LastName: "Doe", Departments: []Department{natalitesChuv}}
		var johnDoe = Doctor{ID: next(), FirstName: "John", LastName: "Doe", Departments: []Department{radiologieChuv}}
		var gregoryHouse = Doctor{ID: next(), FirstName: "Gregory", LastName: "House", Departments: []Department{oncologieSamaritain, dermatologieSamaritain}}
		var jamesWilson = Doctor{ID: next(), FirstName: "James", LastName: "Wilson", Departments: []Department{oncologieSamaritain}}
		natalitesChuv.Doctors = []Doctor{janeDoe}
		radiologieChuv.Doctors = []Doctor{johnDoe}
		oncologieSamaritain.Doctors = []Doctor{gregoryHouse, jamesWilson}
		dermatologieSamaritain.Doctors = []Doctor{gregoryHouse}

		// We create patients
		var mariuszWiesniewski = Patient{ID: next(), FirstName: "Mariusz", LastName: "Wiesniwski", AVSNumber: "756.1234.3333.55",
			BirthDate: time.Date(1984, time.May, 4, 0, 0, 0, 0, time.UTC), Doctors: []Doctor{janeDoe}}
		var naimengLiu = Patient{ID: next(), FirstName: "Naimeng", LastName: "Liu", AVSNumber: "765.4321.0303.44",
			BirthDate: time.Date(1993, time.November, 11, 0, 0, 0, 0, time.UTC), Doctors: []Doctor{janeDoe}}
		var julienRoch = Patient{ID: next(), FirstName: "Julien", LastName: "Roch", AVSNumber: "333.4444.5555.66",
			BirthDate: time.Date(1984, time.December, 3, 0, 0, 0, 0, time.UTC), Doctors: []Doctor{johnDoe}}
		var christopheFrattino = Patient{ID: next(), FirstName: "Christophe", LastName: "Frattino", AVSNumber: "420.1337.1337.42",
			BirthDate: time.Date(1985, time.August, 8, 0, 0, 0, 0, time.UTC), Doctors: []Doctor{jamesWilson, gregoryHouse}}
		janeDoe.Patients = []Patient{mariuszWiesniewski, naimengLiu}
		johnDoe.Patients = []Patient{julienRoch}
		gregoryHouse.Patients = []Patient{christopheFrattino}
		jamesWilson.Patients = []Patient{christopheFrattino}

		// We "persist" it for next call
		dummyData = []Hospital{chuv, samaritain}
	}
	return dummyData
}
