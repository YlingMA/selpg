package main
import (
  "fmt"
  "os"
  "strings"
  "strconv"
  //"flag"
  "os/exec"
)
type selpg_args struct {
	start_page int
  end_page int
  in_filename string
  page_len int
  page_type string
  print_dest string
	//char in_filename[BUFSIZ];
	//char print_dest[BUFSIZ];
}
//var sp_args = selpg_args{}
var progname string
func usage() {
  fmt.Printf("\nUSAGE: ", progname, " -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]\n")
}
func main() {

  var sa = selpg_args{-1,-1,"",72,"l",""}
  var av = os.Args
  var ac = len(av)
  var _sa = &sa
  progname = av[0]
  /*flag.IntVar(&sa.start_page,"s", -1, "specify start page(>=1)")
  flag.IntVar(&sa.end_page,"e", -1, "specify end page(>=s)")
  flag.IntVar(&sa.page_len, "l", 72, "specify length of one page")
  page_type := flag.Bool("f", false, "you can't set -f and page length at the same time")
  print_dest := flag.String("d", "", "specify print dest.")
  flag.Usage = usage
  flag.Parse()*/
	process_args(ac, av, _sa)
	process_input(sa)
}

func process_args(ac int,av []string, psa *selpg_args) {
  var s1,s2 string
  var argno,i int

  if ac < 3 {
    fmt.Println(progname,": not enough arguments\n")
    go usage()
    os.Exit(1)
  }
  //start_page
  s1 = av[1]
  if strings.HasPrefix(s1, "-s")==false {
    fmt.Println(s1,"\n")
    fmt.Fprintf(os.Stderr, "1st arg should be -sstart_page\n")
    //fmt.Println(progname,": 1st arg should be -sstart_page\n")
    go usage()
    os.Exit(2)
  }
  i,error := strconv.Atoi(s1[2:])
  if error != nil{
    fmt.Println("字符串转换成整数失败1")
  }
  if i<1 || i>2147283646 {
    fmt.Fprintf(os.Stderr, "invalid start page\n")
    //fmt.Println(progname,": invalid start page ",i,"\n")
    go usage()
    os.Exit(3)
  }
  psa.start_page = i

  //end_page
  s1 = av[2]
  if strings.HasPrefix(s1, "-e")==false {
    fmt.Fprintf(os.Stderr, " 2nd arg should be -eend_page\n")
    //fmt.Println(progname,": 2nd arg should be -eend_page\n")
    go usage()
    os.Exit(4)
  }
  i,error = strconv.Atoi(s1[2:])
  if error != nil{
    fmt.Println("字符串转换成整数失败2")
  }
  if i<1 || i > 2147283646 || i<psa.start_page {
     fmt.Fprintf(os.Stderr, " invalid end page \n")
    //fmt.Println(progname,": invalid end page ",i ,"\n")
    go usage()
    os.Exit(5)
  }
  psa.end_page = i
  argno = 3
  for ;(argno <= (ac - 1) && av[argno][0] == '-'); {
    s1 = av[argno]
    switch s1[1] {
    case 'l':
      s2 = s1[2:]
      i,error = strconv.Atoi(s2)
      if error != nil{
        fmt.Println("字符串转换成整数失败3")
      }
      if i<1 || i > 2147283646 {
        fmt.Fprintf(os.Stderr, " invalid page length \n")
        //fmt.Println(progname, ": invalid page length ",s2,"\n")
        go usage()
        os.Exit(6)
      }
      psa.page_len = i
      argno = argno+1
      continue
      break

    case 'f':
      if s1 != "-f" {
        fmt.Fprintf(os.Stderr, "option should be \"-f\"\n")
        //fmt.Println(progname,": option should be \"-f\"\n")
        go usage()
        os.Exit(7)
      }
      psa.page_type="f"
      argno = argno+1
      continue
      break

    case 'd':
      s2 = s1[2:]
      if strings.Count(s2,"")-1 < 1 {
        fmt.Fprintf(os.Stderr, "-d option requires a printer destination\n")
        //fmt.Println(progname,": -d option requires a printer destination\n")
        go usage()
        os.Exit(8)
      }
      psa.print_dest = s2
      argno = argno+1
      continue
      break
    default:
      fmt.Fprintf(os.Stderr, "unknown option \n")
      //fmt.Println(progname,": unknown option ",s1,"\n")
      go usage()
      os.Exit(9)
      break

    }
  }
  if argno <= ac-1  {
    psa.in_filename = av[argno]
    //打开文件是否成功

    //assert

  }
}

func process_input(sa selpg_args){
  //var s1 string
  //var crc string
  //var c int
  //var line string
  var line_ctr int
  var page_ctr int
  //var inbuf string


  // FILE *fin; /* input stream */
	// FILE *fout; /* output stream */
	// char s1[BUFSIZ]; /* temp string var */
	// char *crc; /* for char ptr return code */
	// int c; /* to read 1 char */
	// char line[BUFSIZ];
	// int line_ctr; /* line counter */
	// int page_ctr; /* page counter */
	// char inbuf[INBUFSIZ]; /* for better performance on input stream */
  buf := make([]byte, 10)
  if sa.in_filename == "" {
    fin := os.Stdin
    if fin == nil {
      fmt.Fprintf(os.Stderr, "could not get filename at line 185 \n")
      os.Exit(11)
    }
  } else {
    fin,err := os.Open(sa.in_filename)
    if err != nil {
            fmt.Println(sa.in_filename,err)
            return
    }
    if fin == nil {
      fmt.Fprintf(os.Stderr, "could not open input file at line 185 \n")
      //fmt.Println(progname,": could not open input file \"",sa.in_filename,"\"\n")
      os.Exit(12)
    }
    fin.Close()
  }
  if sa.print_dest=="" {
    //fout := os.Stdout
  } else {
    //287

    var s1 string= "-d"+sa.print_dest


    cmd := exec.Command("lp", s1)
    cmd.Run()
    /*              if fout == nil {
      fmt.Println(progname,": could not open pipe to \"",s1,"%s\"\n")
      os.Exit(13)
    }*/

  }
  if sa.page_type == "l" {
    line_ctr = 0
    page_ctr = 1
    //fmt.Println(sa.in_filename,"\n")
    fin,err := os.Open(sa.in_filename)
    if err != nil {
      //fmt.Fprintf(os.Stderr, err)
            fmt.Println(sa.in_filename,err)
            return
    }
    for{
      n, _ := fin.Read(buf)
      if 0 == n {
        break
      }
      for i:=0;i<n;i++{
        if buf[i] == '\n' {
          line_ctr++
          if line_ctr > sa.page_len {
    				page_ctr++
    				line_ctr = 1
    			}
        }
      }
      if (page_ctr >= sa.start_page) && (page_ctr <= sa.end_page) {
        os.Stdout.Write(buf[:n])
      }
    }
  } else {
    page_ctr = 1
    fin,err := os.Open(sa.in_filename)
    if err != nil {
      //fmt.Fprintf(os.Stderr, err+" at line 238\n")
            fmt.Println(sa.in_filename,err)
            return
    }
    for {
      n, _ := fin.Read(buf)
      if 0 == n {
        break
      }
      for i:=0;i<n;i++{
        if buf[i] == '\f' {
          page_ctr++
        }
      }
      os.Stdout.Write(buf[:n])
    }
    fin.Close()
  }
  if page_ctr < sa.start_page {
    fmt.Fprintf(os.Stderr, "start_page greater than total pages,no output written at line 257\n")
    //fmt.Println(progname,": start_page (", sa.start_page, ") greater than total pages (", page_ctr,"),no output written\n")
  }
  if page_ctr < sa.end_page {
    fmt.Fprintf(os.Stderr, "end_page greater than total pages,,less output than expected at line 261\n")
    //fmt.Println(progname,": end_page (",sa.end_page,") greater than total pages (", page_ctr,"),less output than expected\n")
  }
  //355-362
  fmt.Fprintf(os.Stderr, "Program is done! at line 265\n")
  //fmt.Println(progname,": done\n")
}
