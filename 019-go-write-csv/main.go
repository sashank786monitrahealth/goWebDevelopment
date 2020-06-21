package main;

import(
  "encoding/csv"
  "fmt"
  "os"
  "log"
)

func main(){

var outerArr [][]string;
var innerArr []string;

innerArr = append(innerArr, "Ramesh");
innerArr = append(innerArr, "Kerala");
innerArr = append(innerArr, "25");
innerArr = append(innerArr, "Vanchikulam Road, Thrissur");

outerArr = append(outerArr,innerArr);
innerArr = nil;

innerArr = append(innerArr, "Ganesh");
innerArr = append(innerArr, "Karnataka");
innerArr = append(innerArr, "35");
innerArr = append(innerArr, "Philomena Church, Mysore");

outerArr = append(outerArr,innerArr);

fmt.Printf("outerArr = %v\n",outerArr);


// begin:

    file, err := os.Create("result.csv")
    checkError("Cannot create file", err)
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range outerArr {
        err := writer.Write(value)
        checkError("Cannot write to file", err)
    }

// end:

}


func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}



