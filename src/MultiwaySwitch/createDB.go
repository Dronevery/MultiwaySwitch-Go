package MultiwaySwitch

//package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

var dbmap *gorp.DbMap

func createDatabase() {
	// initialize the DbMap
	dbmap = initDb()
	defer dbmap.Db.Close()

	// delete any existing rows
	err := dbmap.TruncateTables()
	checkErr(err, "TruncateTables failed")

	// create two drones
	p1 := newDrone("2MDSlOW8qd/kkqYjCdV/HBDq5bfpaeXApPZe9kgneKs+vjCK8Bjav9s549ghXXkzvOPvriHbQmgPG/5yL43bag==", "4khSEpIm9p8Ds+Q+GJXMY+x2kKzWpwUiyZHPGKQcLkNPy8TgyxkvzR60+eH5Jobyw9I3q8BUky6a69JTp/siyORK4kY9Pxew/xT5je5F/+XMO3z9Y/5UxBnvBujECtLApDqb/DPTKlHiwUIu58GmbSIiDBr/o9KpH5bUdYr2dOIaT0XFZrMopy1YOM1eI1SEk9GyaP2jB+z4FSerm2r3DhF4E/ZNZmtJ3GtrMTDZTBbATVLnvyGg70m+9kbeajpAJDlvyjvrBzCxLGzK7c2F+FEsBgGrro6b5hKj4k3dhj4mSXd1diJDmJCVQQijjpN+mz3ddBYd2vi0NjIuNngSmDXfAZ5T7cNj1NHOIZjinJdWx4r3RadXdQUN2/XvgKJJcFc487wvx9XYCVIJP7QyFxeZOMWEi/lT95m6aEwUfeO6+sTI1M37kTNdicZShAaMzBRNQhPQdAd0a4sa0cdt4QOMiQITVZE+4vHQmHlgLr7M5qsuWprw8UX+ueRDSp+iDbWwpbIrU9sSe/Oequ83357vF9RVj0xcbi9HDOuIL3eQfvniy5rWPNgqMvJ8vpXm8ix03spppmBnxSbrT6uVtAXVb7NzndWRSONY0QFYkT5fF3V1W+fPt1rMuTHAawJ7f9uj7BGqtsF1eSvlwG56+qij/66tpXcu1UB1MAdQ4UbuoVj9xEsPsZaOEn+qKEt/uTD0AleUMtlgFwInSjIu6PIXvzzrCuWKZR0kZ/KVe50crLACm80xI6ksY5L331KJLL9eLfuG7mRq8hCxkdxgHP+s+y7ItRUbTFfT+AcK4cL4qhpJjzMPZ2bLZuxCdZ6SD5cSgmb2eJ7qNf9opqiyfi5wzbgUHTEujE/Vj1NdNq3XhRlq5+NkSkSWwNmlAEsGnc6+y49phs6qQlVr4guMbCKOaZ6dYzB2xxh3TC9yU7Dp+AT97HTIAUGFIMj6I6auBHC4FmsNCC/vfAHZmndeBnkRUNAXUcXzWoLSgAByAVaIW9SVUdOc9fU4G5pUNvL+4qOF4pZEm9GY0zmeJ5e+ftakIVYMlaej0sF+f8Cv3RQRtBhSEYHv4L+vyetowCaevn3fJKbK//ZoxyzM/UH9JKFQ49kc+ausuZN8q59C1y5iy6yLR3ca4bwf7rjzlFLYEC9x5m3k6KWWriJ5GfDi9cAf/Ll11d/HPC32hqwwQ1iSOMWYfBz4EPnqjm6qh1+WrPlGbM4j+qIXieMFF3hEgZJU1yMhqezX63kyRCdXG5BZBn8reTTkmoaRQN7ET/X77kCIgDYNHxGVhErHAcBNB9aolpMIKk0742b6Kd1annv7cUbCD7DuCGDWmtNiYvOcYfy/Aypy2H5s3dquIIsy5Q==")
	p2 := newDrone("4p2QMVEIy+zzsLpmFFEtkWFDapSk4SwZb43D7OpwvpN83x/U6T0R6B+vKDc4IRdzlDyaSwTJrkwk6hCqPY64kA==", "Ne2QpVaZLFUqpRNezeHlH2gzHpDQFeIHxCymL7Io5VEDwana/Nc/iFvSKTp4jGyTZ7rLbPZ64GzgJ4QmMji6pSw8JL5qB1GLjcF1e77Uao754+bnye4DKyeLK+hMItdgoUKzF+seOpqZa9wNlYcHZbyWFCqNUmT7igx8C2N23Mht6jjk40M3GIlW0g9in36A/PTWMCsk0pt0k30D7VFTUMzBgDFsSuuM1N2N9qqnhIfh5pkGJnKmcBTzUwezuNLX5gdnK6vlBUOehRxIALce0zS0oX03E5HJlSpJRLdBJ9oU0g8aMGr+ufav6jCVrDeKdghaZPzFAU2n7rPtwOpWlbrUk6HdpXquk4hewZFwdbyIuEZ0I+EM1xN3Jsy2RRGrYX8ANNk0Jb3cRSY/E+ab76RgAWT8DVaGe6iv++NPa9gJGgfNfqiAAoJx9X0lvEduxDwvDPrzDXSbBu/40gs+Bkol6YD8e8c+J3JNTDOMLcD3YA0LeBDayLuYm3Y8zaDNd7nD+WKIn2C0eiD/L6QGD98XN3zjK6CgyTtUoTcw6nKUjkJHLmsOSLaqIg+Fka323bNjGSMyXzocT/KtGVkhA4dz+AsDfj3rcfVt2YRF7uuCwgtsulIlzjLBN/WwdCoAAx6hKA5bTdZq53XQuY8LnLmrMLKAH0Uk0z2ubtuVhcRsu0xdkFPE3PE0V9BwaOVEx5dVvlrJ80Hnm9Vqce4XK4DukDtDOjJgQpa91QwRF5gVwLrPaokgHkiBeE+FnkBg1/7THirbC7I/szg9LpAMaNcfbpWoH8P85hEC78wOT5NqhoSZksFMuBGe2Y5cwPvsEAW+GQfg9R0lhObmTgPBrEkCL3IE5lf8wDBgGLK53J6DJlTBqksdzEFpJj1WxNvFoV4fmFhOg19hABSdpt4kkBABhcgubjvx+yLv9oxpJpY3e+Lyg6MlV9sm+SkAPnH8nzmTmQdXafuUg87RQ4Wqw0XSn1c9Y7PH45qCVQFNPeeBFK9FVFYTOLk9NO6DioxOVetNGvASy6xd/Xdrzl9kDS5D/XomMeg0R5AV1TwI0VS2Lruop/QLWowgQqKJrdIdE44w6nJw/owqjDYBLDUD+UvwdKQoj1eiAK9cq7zF4hkZHWjh2aq/Fub5tiMm+jBNdFZDX8+3OGKikhYSO2w8TL55fk15yJWmhEoH3iyak87fbfGsNqRsscTSF1A0xH+BPyNL3Jk0IFh2CpQrZmnMxQX38srlTPPVTXBj6Ykv8xbHq+PiLHrclNf3HSCRd72PMhbtN4I469f9fH7OsRnQj6FtDs7jJzMN402jBoiHSBogX0mocKGXQB1bSGHFx201cRdZu2oMqu2B3ct+7wL53A==")

	// insert rows - auto increment PKs will be set properly after the insert
	err = dbmap.Insert(&p1, &p2)
	checkErr(err, "Insert failed")

	// use convenience SelectInt
	//count, err := dbmap.SelectInt("select count(*) from drones")
	//checkErr(err, "select count(*) failed")
	//log.Println("Rows after inserting:", count)

	// update a row
	//p2.Token = "Go 1.2 is better than ever"
	//count, err = dbmap.Update(&p2)
	//checkErr(err, "Update failed")
	//log.Println("Rows updated:", count)

	// fetch one row - note use of "drone_id" instead of "Id" since column is aliased
	//
	// Dronegres users should use $1 instead of ? placeholders
	// See 'Known Issues' below
	//
	//err = dbmap.SelectOne(&p2, "select * from drones where drone_id=?", p2.Id)
	//checkErr(err, "SelectOne failed")
	//log.Println("p2 row:", p2)

	// fetch all rows
	//var drones []Drone
	//_, err = dbmap.Select(&drones, "select * from drones order by drone_id")
	//checkErr(err, "Select failed")
	//log.Println("All rows:")
	//for x, p := range drones {
	//	log.Printf("    %d: %v\n", x, p)
	//}

	// delete row by PK
	//count, err = dbmap.Delete(&p1)
	//checkErr(err, "Delete failed")
	//log.Println("Rows deleted:", count)

	// delete row manually via Exec
	//_, err = dbmap.Exec("delete from drones where drone_id=?", p1.Id)
	//checkErr(err, "Exec failed")

	// confirm count is zero
	//count, err = dbmap.SelectInt("select count(*) from drones")
	//checkErr(err, "select count(*) failed")
	//log.Println("Row count - should be zero:", count)

	//log.Println("Done!")
}

func newDrone(title, body string) Drone {
	return Drone{
		Created:   time.Now().UnixNano(),
		Token:     title,
		SecretKey: body,
	}
}

func initDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	db, err := sql.Open("sqlite3", "/tmp/dronevery_db.bin")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'drone' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Drone{}, "drones").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
