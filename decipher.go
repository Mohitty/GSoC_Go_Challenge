package main

import "fmt"
import "io/ioutil"
import "encoding/hex"

func decodeString(value string, m map[uint8]int) (k byte, plaintext string) {
	//convert the given hexadecimal encoded string to byte array.
    	decodedvalue, err := hex.DecodeString(value)
    	if err != nil {
    		panic(err)
    	}
        //resultBytes contains the result of XOR of decodedvalue and key. 
    	resultBytes := make([]byte, len(decodedvalue))
	
	//score contains the max score.
    	score := 0
	
        //trying all possible ascii keys.
    	key := byte(0)
    	i := byte(0)

	for ; ;  {
	    key = byte(i)
	    for j := 0; j<len(decodedvalue); j++ {
	    	resultBytes[j] = decodedvalue[j] ^ key
	    }
	    	
    	    temp_score := 0
    	    for x :=0 ; x<len(resultBytes); x++ {
    		temp_score += m[resultBytes[x]]
    	    }
    	    if temp_score>score {
    		k = i
    		score = temp_score
    		plaintext = string(resultBytes)
    	    }
		
	    if i==255 {
	    	break
	    } else {
	    	i++
	    }
	}
	return 
}

func main() {
	//read file and calculate the frequency of characters in the text.
	b, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Print(err)
	}
	m := make(map[uint8]int)
	for i := 0; i<len(b); i++ {
		_, ok := m[b[i]]
		if ok==true {
			m[b[i]] += 1
		} else {
			m[b[i]] = 1
		}
	}
	
	value := "20342824471508040c1447060903471e081247061502471102151e47040b02110215"
    
        //decipher the encoded text "value".
	k, plaintext := decodeString(value, m)

	fmt.Println("key: ", string(k))
	fmt.Println("plaintext: ", plaintext)
}

