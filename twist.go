/* File twist.go
 * created by Vincenzo De Florio on Die Apr 12 12:18:52 CEST 2016
 * translated from twist.c, by the same author
 */

package	main
import	(
	"fmt"
	"os"
	"log"
	"strconv"
	"io"
)

func Strchr(s []byte, delim byte) *byte {
	var i int
	for i=0; i<len(s); i++ {
		if s[i] == byte {
			return &(s[i])
		}
	}
	return nil
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

const ERR1 = -1
const MAXTWEEPS = 500
const SIZEOFPICTURE = 24


func assert (condition bool) () {
	if condition == false {
		log.Fatalln("assertion failed") // Fatal(err)
		// Fatalln calls os.Exit(1)
	}
}


func main () {
	// #CLANG argc argv
	argv := os.Args
	argc := len(argv)

	Verbose := false
	var err error
	var i, l, sum, tweet_length int
	var tweet, tweeps_file string
	var tf *File
	var out_file string = nil
	// char *out_file = NULL
	var dontTweet bool
	// char dontTweet = 0
	//char first_time=1
	var first_time bool = false
	//int ntweeps = MAXTWEEPS
	ntweeps := MAXTWEEPS
	var tweeps, tweep *string
	// char **tweeps, **tweep
	//char * line = NULL
	var line string = nil
	len := 0
	// size_t len = 0
	// ssize_t read
	var read int
	// int n
	var n int
	// char *p
	var p *byte
	var buffer string // char buffer[141]
	// char *picture_file = NULL
	var picture_file string = nil
	//char format_string[64]
	var format_string string

	if argc == 1 {
		usage()
		return
	}
	for i=1; i<argc; i++ {
		if argv[i][0] == '-' {
			switch argv[i][1] {
			case 'm':  // message
				// instead of tweet = strdup(argv[++i]), I just do...
				i++
				tweet = argv[i]	// note the ":=" operator

			case 'l':  // tweep list
				i++
				tweeps_file = argv[i]
			case 's':  // simulate (do not tweet)
				dontTweet = 1
			case 'v':  // verbose
				Verbose = 1
			case '#':  // max no. of tweeps
				i++
				// #CLANG sscanf(argv[i], "%d", &ntweeps)
				// ntweeps = strconv.ParseInt(argv[i],10,32)
// i, err := strconv.Atoi("-42")
// s := strconv.Itoa(-42)
				ntweeps, err = strconv.Atoi(argv[i])
				if err != nil {
					usage()
					os.Exit(ERR1)
				}
				assert (ntweeps > 0)
			case 'o': // output file
				i++
				out_file = argv[i]
			case 'p': // picture file
				i++
				picture_file = argv[i]
			default:
				usage()
				os.Exit(ERR1)
			}
		}
	}

	// file system operations
	// tf = fopen(tweeps_file, "r")
	tf, err = os.Open(tweeps_file)
	if (err != nil) {
		log.Fatal(err)
	}
	// assert (err != NULL)

	tweet_length = len(tweet)
	assert (tweet_length < 140)

	// tweeps = (char **) malloc(ntweeps * sizeof(char *))
	tweeps = make([]string, ntweeps)
	
	for n=0; ; n++ {
		line, err = ReadString('\n')
		if err == io.EOF {
			break
		}
		tweeps[n] = line
		p = &tweeps[n]
		p = Strchr(p, '\n')
		assert(p != nil)
		*p = ' '; // THIS IS A SPACE, USED TO SEPARATE TWEEPS ON THE OUTPUT TWEET
		if Verbose {
			fmt.Printf("Tweep[%d] = %s\n", n, tweeps[n])
		}
	}

	tf.Close()
	assert (err != nil)
	// if line free(line)  NO NEED FOR THIS

	tf, err = os.Create(out_file)
	check(err)

	// constructing the tweets
	tweep = tweeps
	first_time=true

	for i=0; i<n;  {
		buffer = nil
		// sum = (picture_file == NULL)? 0:SIZEOFPICTURE
		// strcpy(buffer, tweet)
		// strcat(buffer, " ")
		for sum=tweet_length+1;    i<n;    i=i+1 {
			l = len(tweep[i])
			// NOTA: COSI' FUNZIONA if sum+l>137 break
			if sum+l>138 { break }
			strcat(buffer, tweep[i])
			//sum = strlen(buffer)
			if verbose { fmt.Printf("buffer=%s\n", buffer) }
			sum += l
			//++tweep
		}

		// only single space chars please
		for ;; {
			//s = buffer
			//p=strstr(s, "  ")
			// func Index(s, sep string) int
			// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
			idx := Index(s, "  ")
			if idx == -1 { break }
			buffer = buffer[0:idx] + buffer[idx+1:]
		}

		len = len(buffer)+tweet_length+1
		assert(len<=139)

		if first_time==true {
			fmt.Printf("TWEET: %s %s (nchar=%zd)\n", tweet, buffer, len)
			first_time=false
			//sprintf(format_string, "sudo t update '%%s %%s' -f %s '%%s %%s'", picture_file
			fmt.Fprintf(tf, "sudo t update '%s %s'\n", tweet, buffer)
		} else {
			fmt.Printf("TWEET: %s%s (nchar=%zd)\n", buffer, tweet, len)
			fmt.Fprintf(tf, "sudo t update '%s%s'\n", buffer, tweet)
		}
	}

	// closings
	tf.Close()
	//for i=0; i<n; i++ free(tweeps[i]); 
	//free(tweeps)
}

func usage() {
	fmt.Fprintln(os.Stderr,  "twist: Error in the calling arguments\n")
	fmt.Fprintln(os.Stderr,  "usage: twist -m messageToTweet -l tweepList -s -v\n")
	fmt.Fprintln(os.Stderr,  "       -# maxNumberOfTweeps -o outputScript\n")
	//fmt.Fprintln(os.Stderr,  "       -# maxNumberOfTweeps -o outputScript -p pictureFilename\n")
	fmt.Fprintln(os.Stderr,  "       -s: simulate (do not tweet), -v: verbose\n")
}
/* End of file twist.go */
