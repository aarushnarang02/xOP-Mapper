package main

import ( 
     "fmt"
	 "encoding/csv"
	 "os"
)

func main(){
    
    m := csvfileToMap()
	fmt.Println(m)
}    

func csvfileToMap() (returnMap []map[string]map[string]string){

    filePath := "usingNames.csv"
	
	csvfile, err := os.Open(filePath)
	if err != nil{
		return nil
	}
    
	defer csvfile.Close()

	csvfile.Seek(0,0)
	reader := csv.NewReader(csvfile)

	rawData, err := reader.ReadAll()

	if err != nil{
		return nil}

   // header := []string{} //holds the first row header

	s2 := rawData
	//n := len(s2)
	//m := len(s2[0])
	//names := make(map[string]map[string]string)

	for i, _ := range s2{
		for j,_ := range s2[0]{
			if j==0 && i > 0{
			  //names[header[i]] = s2[i][j]
			  fmt.Println(s2[i][j])
			}
			//returnMap = append(returnMap, names)
		}
	} 
        	
	   return returnMap
	
}


