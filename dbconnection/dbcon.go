package dbconnection

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Member struct {
	ID        int    `json:"ID,omitempty"`
	UserName  string `json:"username,omitempty"`
	Gender    string `json:"gender,omitempty"`
	SkinType  string `json:"skintype,omitempty"`
	SkinColor string `json:"skincolor,omitempty"`
}

type DisplayAll struct {
	UserName   string `json:"USERNAME,omitempty"`
	Gender     string `json:"GENDER,omitempty"`
	SkinType   string `json:"SKINTYPE,omitempty"`
	SkinColor  string `json:"SKINCOLOR,omitempty"`
	DescReview string `json:"DESC_REVIEW,omitempty"`
	NoOfLikes  int    `json:"noOfLikes,omitempty"`
}

type Product struct {
	ProductID int `json:"productid"`
}

type LikeDislike struct {
	Option   string `json:"option"`
	ReviewID int    `json:"reviewID"`
	MemberID int    `json:"memberID"`
}

func InsertToMemberTable(table Member) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/k_tech")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Connection to DB successful")

	template := "INSERT INTO `k_tech`.`Member` (`ID_MEMBER`, `USERNAME`, `GENDER`, `SKINTYPE`, `SKINCOLOR`) VALUES ("

	id := strconv.Itoa(table.ID)

	queryToDB := template + "'" + id + "', '" + table.UserName + "', '" + table.Gender + "', '" + table.SkinType + "', '" + table.SkinColor + "')"
	insert, err := db.Query(queryToDB)

	if err != nil {
		panic(err)
	}

	defer insert.Close()
	fmt.Println("Insert to DB successful")
}

func UpdateToMemberTable(table Member) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/k_tech")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Connection to DB successful")

	id := strconv.Itoa(table.ID)
	templateGet2 := "\"" + id + "\""

	template := "UPDATE `k_tech`.`Member` SET "

	template2 := "WHERE ID_MEMBER = "

	temp := buildForUpdate(table)

	var querytoUpdate string

	for i := range temp {
		if temp[i] == "SKINTYPE" {
			querytoUpdate = template + temp[i] + " = '" + table.SkinType + "' " + template2 + templateGet2
		} else if temp[i] == "SKINCOLOR" {
			querytoUpdate = template + temp[i] + " = '" + table.SkinColor + "' " + template2 + templateGet2
		} else if temp[i] == "GENDER" {
			querytoUpdate = template + temp[i] + " = '" + table.Gender + "' " + template2 + templateGet2
		} else if temp[i] == "USERNAME" {
			querytoUpdate = template + temp[i] + " = '" + table.UserName + "' " + template2 + templateGet2
		} else {
			continue
		}
		update, err := db.Query(querytoUpdate)
		if err != nil {
			panic(err)
		}
		defer update.Close()
	}
	fmt.Println("Insert to DB successful")
}

func DeleteFromMemberTable(table Member) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/k_tech")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Connection to DB successful")

	template := "DELETE FROM `k_tech`.`Member` WHERE ID_MEMBER = "

	id := strconv.Itoa(table.ID)

	queryToDB := template + "\"" + id + "\""

	delete, err := db.Query(queryToDB)

	if err != nil {
		panic(err)
	}

	defer delete.Close()
	fmt.Println("Deleting successful")
}

func ViewAllMemberTable(table Member) []Member {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/k_tech")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Connection to DB successful")

	get, err := db.Query("SELECT * FROM k_tech.Member")

	if err != nil {
		panic(err)
	}

	var test []Member
	for get.Next() {
		err = get.Scan(&table.ID,
			&table.UserName,
			&table.Gender,
			&table.SkinType,
			&table.SkinColor)
		test = append(test, table)
	}
	fmt.Printf("%+v\n\n", test)

	defer get.Close()
	fmt.Println("Get Data successful")
	return test
}

func ViewProduct(id Product) []DisplayAll {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/k_tech")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Connection to DB successful")

	idStr := strconv.Itoa(id.ProductID)

	query := "SELECT m.USERNAME, m.GENDER, m.SKINTYPE, m.SKINCOLOR, r.DESC_REVIEW, count(l.ID_REVIEW) as numberOfLikes FROM Review_Product r join Member m on r.ID_MEMBER = m.ID_MEMBER join like_review l on r.ID_REVIEW = l.ID_REVIEW where r.ID_PRODUCT = "

	queryString2 := " group by l.ID_REVIEW"

	queryString := query + idStr + queryString2

	get, err := db.Query(queryString)

	if err != nil {
		panic(err)
	}

	var test []DisplayAll
	var temp DisplayAll
	for get.Next() {
		err = get.Scan(&temp.UserName,
			&temp.Gender,
			&temp.SkinType,
			&temp.SkinColor,
			&temp.DescReview,
			&temp.NoOfLikes)
		test = append(test, temp)
	}
	fmt.Printf("%+v\n\n", test)

	defer get.Close()
	fmt.Println("Get Data successful")
	return test
}

func InsertLikeorDislike(option LikeDislike) string {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/k_tech")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Connection to DB successful")
	fmt.Printf("%+v\n", option)
	reviewid := strconv.Itoa(option.ReviewID)

	if option.Option == "Like" {
		memberid := strconv.Itoa(option.MemberID)
		temp := "INSERT INTO `k_tech`.`Like_Review` (`ID_REVIEW`, `ID_MEMBER`) VALUES ('"
		queryLike := temp + reviewid + "', '" + memberid + "')"
		fmt.Printf("%s\n", queryLike)
		insert, err := db.Query(queryLike)
		if err != nil {
			panic(err)
		}
		defer insert.Close()
	} else if option.Option == "Dislike" {
		temp := "DELETE FROM `k_tech`.`Like_Review` WHERE (`ID_REVIEW` = '"
		queryLike := temp + reviewid + "')"
		fmt.Printf("%s\n", queryLike)
		del, err := db.Query(queryLike)
		if err != nil {
			panic(err)
		}
		defer del.Close()
	} else {
		return "Please Choose Like or Dislike option"
	}
	fmt.Println("Update Data successful")
	return "Update Successful"
}
