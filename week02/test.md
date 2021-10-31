```go

func QueryUsersById(id int) (User, error) {
	var user User
	row := Db.QueryRow("select id ,name from users where id = ?", id)
	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		return user, errors.Wrap(err, "QueryUsersById err")
	}
	return user, nil
}

func main() {
	user ,err := QueryUsersById("12")
	if err != nil{
		fmt.Printf("query user err : %+v",err)
		return
	}
	fmt.Println("query user : ",user)

}

```