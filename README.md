# biuper
Biu.moe Uploader

基本功能有了，过几天高考完了再弄别的qwq

UID 和 KEY 在 config.yaml 中定义


    root@whitemagi:~/BiuUploader# ./biuper -h
    
    Usage of ./biuper:
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
    	  Use -t1 <type1> (1:anime, 2:galgame, 3: idol, 4: touhou project, 5:vocaloid, 6:doujin) (default "musictype1_undefined")
    -t2 string
    	  Use -t2 <type2> (1:original, 2:instrumental, 3:absolute music, 4:cover version) (default "musictype2_undefined")
    
    example: ./biuper -a "FINAL FANTASY XIV: BEFORE METEOR" -f "104. Answers.flac" -r "《最终幻想14：世界毁灭之前》最热插曲" -s "Susan Calloway" -t "Answers" -t1 "1" -t2 "1"
