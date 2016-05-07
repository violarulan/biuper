package biuper

import (                                                                            
        "crypto/md5"                                                                
        "fmt"                                                                       
        "io"                                                                        
        "log"                                                                       
        "os"    
	"encoding/hex"                                                                    
        "os/exec"                                                                   
        "path/filepath"                                                             
        "strings"                                                                   
)                                                                                   
                                                                                    
func GetCurrentPath() string {                                                      
        file, _ := exec.LookPath(os.Args[0])                                        
        path, _ := filepath.Abs(file)                                               
        path = string(path[0:strings.LastIndex(path, "/")])                         
        return path                                                                 
}                                                                                   
                                                                                    
func GetFileMd5(filename string) string {                                                                       
        targetfile := GetCurrentPath() + "/" + filename                                 
                                                            
        file, inerr := os.Open(targetfile)                                            
        if inerr != nil {
		log.Fatal(inerr)
}                                                          
                md5h := md5.New()                                                   
                io.Copy(md5h, file)                                                 
                                     
            
	return fmt.Sprintf("%x", md5h.Sum([]byte("")))                                                                 
}     

func Md5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}
