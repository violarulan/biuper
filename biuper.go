package main

import (
    "fmt"
    "flag"
    "log"
    "os"
    "encoding/json"
    "github.com/masahoshiro/biuper/s"
)

var (
    logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
    music_file *string = flag.String("f", "musicfile_undefined", "Use -f <musicfile> (required)")
    music_title *string = flag.String("t", "musictitle_undefined", "Use -t <title> (required)")
    music_singer *string = flag.String("s", "musicsinger_undefined", "Use -s <singer> (Singer)")
    music_album *string = flag.String("a", "musicalbum_undefined", "Use -a <album> (Album)")
    music_remark *string = flag.String("r", "musicremark_undefined", "Use -r <remark> (Remarks)")
    music_type1 *string = flag.String("t1", "1", "Use -t1 <type1> (1:anime(default), 2:galgame, 3: idol, 4: touhou project, 5:vocaloid, 6:doujin)")
    music_type2 *string = flag.String("t2", "1", "Use -t2 <type2> (1:original(default), 2:instrumental, 3:absolute music, 4:cover version)")
    force_upload *int = flag.Int("F", 0, "Use -F 1 to Force Upload(default 0)")
)

func main() {
    flag.Parse()

    if (*music_file == "musicfile_undefined") || *music_title == "musictitle_undefined" {
        logger.Fatal("[ERROR] Some necessary params missing. Did you forget title or filename?")
        os.Exit(1)
    }

    key := biuper.ReadConf("common.key")
    uid := biuper.ReadConf("common.uid")

    filename := *music_file
    songinfo := make(map[string]string)

    songinfo["uid"] = uid
    songinfo["filemd5"] = biuper.GetFileMd5(filename)
    songinfo["title"] = *music_title
    songinfo["singer"] = *music_singer
    songinfo["album"] = *music_album
    songinfo["remark"] = *music_remark
    //songinfo["sign"] = "fd14711abb548ae9acfe50d0f0f12f00"
    songinfo["type1"] = *music_type1
    songinfo["type2"] = *music_type2
    songinfo["force"] = fmt.Sprintf("%d", *force_upload)
    songinfo["sign"] = biuper.Md5(songinfo["uid"]+songinfo["filemd5"]+songinfo["title"]+songinfo["singer"]+songinfo["album"]+songinfo["remark"]+key)

    var dat map[string]interface{}
    
    logger.Println(fmt.Sprintf("[INFO] Requesting UpToken ..."))
    resp := biuper.FormPost(songinfo, "https://api.biu.moe/Api/createSong")
   
    json.Unmarshal([]byte(resp), &dat)

    if dat["success"] == true {
        logger.Println(fmt.Sprintf("[INFO] Get UpToken: %s", dat["token"]))
        params := make(map[string]string)
        params["key"] = biuper.GetFileMd5(filename)
        params["x:md5"] = biuper.GetFileMd5(filename)
        params["token"] = fmt.Sprintf("%s", dat["token"]) 
        biuper.Upload(params, filename)
    } else {
        logger.Fatal(fmt.Sprintf("[ERROR] Get UpToken failed: %s", resp))
        os.Exit(1)
    }

}

