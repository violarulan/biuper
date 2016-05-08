# biuper
Biu.moe Uploader

基本功能有了，过几天高考完了再弄别的qwq

UID 和 KEY 在 config.yaml 中定义


    root@whitemagi:~/BiuUploader/biuper# go run biuper.go -h
    Usage of /tmp/go-build133564160/command-line-arguments/_obj/exe/biuper:
        -F int
          Use -F 1 to Force Upload(default 0)
        -a string
          Use -a <album> (Album) (default "musicalbum_undefined")
        -f string
          Use -f <musicfile> (required) (default "musicfile_undefined")
        -r string
          Use -r <remark> (Remarks) (default "musicremark_undefined")
        -s string
          Use -s <singer> (Singer) (default "musicsinger_undefined")
        -t string
          Use -t <title> (required) (default "musictitle_undefined")
        -t1 string
          Use -t1 <type1> (1:anime(default), 2:galgame, 3: idol, 4: touhou project, 5:vocaloid, 6:doujin) (default "1")
        -t2 string
          Use -t2 <type2> (1:original(default), 2:instrumental, 3:absolute music, 4:cover version) (default "1")

    
    example: ./biuper -a "FINAL FANTASY XIV: BEFORE METEOR" -f "104. Answers.flac" -r "《最终幻想14：世界毁灭之前》最热插曲" -s "Susan Calloway" -t "Answers" -t1 "1" -t2 "1" -F 1
