package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"taskTracker/tskMng"
)

func main() {

	var strname string
	var tb tskMng.TaskBoard

	f, err := os.Open("taskdata.json")
	if err != nil {
		fmt.Println(err)
	}
	f.Close()

	data, _ := os.ReadFile("taskdata.json")
	json.Unmarshal(data, &tb)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strname = scanner.Text()

	user := tskMng.NewUser(0, strname)

	tb.TaskManageCLI(user)
	f, _ = os.Create("taskdata.json")
	dataToJson, _ := json.MarshalIndent(tb, "", "    ")
	f.WriteString(string(dataToJson))

	fmt.Print(string(dataToJson))

}
