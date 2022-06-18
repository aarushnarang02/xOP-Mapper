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

func csvfileToMap() (returnMap []map[int]string){

    filePath := "justNames.csv"
	
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


	s2 := rawData

	m := make(map[int]string)

	for i, _ := range s2{          // i -> column
		for j,_ := range s2[0]{    // j -> row
			if j==0 && i > 0{
			  //fmt.Println(i,s2[i][j])
			  m[i] = s2[i][j]
			  }
		}
		
		} 
		//fmt.Println(m[13])
		returnMap = append(returnMap,m)
	   return returnMap
	
}


