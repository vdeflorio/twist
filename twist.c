/* File twist.c
 * created/modified on Son Nov  8 17:59:09 CET 2015

 * by creac for Son Nov  8 17:59:09 CET 2015

 */

#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>
#include <string.h>
#include <assert.h>
#include <stdlib.h>

#  define STRERR1  "Error in the arguments"
#  define STRERR2  "Can't open input file(s)"
#  define STRERR3  "Can't stat input file(s)"
#  define STRERR4  "Can't open output file"
#  define STRERR5  "Synch error -- empty file?"
#  define FORMAT   "USAGE: "
#  define STRWARN1 "Warning: files have not the same size"
#  define STRMSG1  "Processing starts..."
#  define STRMSG2  "Processing ends."
#  define STRMSG3  "Remember"

#define ERR1	-1
#define MAXTWEEPS 500
#define SIZEOFPICTURE    24	// amount of characters corresponding to a picture

char verbose;

void usage(void);

int main (int argc, char *argv[])
{
	int i, l, sum, tweet_length;
	char *tweet;
	char *tweeps_file; FILE *tf;
	char *out_file = NULL;
	char dontTweet = 0;
	char first_time=1;
	int ntweeps = MAXTWEEPS;
	char **tweeps, **tweep;
	char * line = NULL;
	size_t len = 0;
	ssize_t read;
	int n;
	char *p;
	char buffer[141];
	char *picture_file = NULL;
	char format_string[64];

	if (argc == 1) {
		usage();
		return 0;
	}
	for (i=1; i<argc; i++) {
		if (argv[i][0] == '-')
		switch(argv[i][1]) {
		case 'm':  // message
			tweet = strdup(argv[++i]);
			break;
		case 'l':  // tweep list
			tweeps_file = argv[++i];
			break;
		case 's':  // simulate (do not tweet)
			dontTweet = 1;
			break;
		case 'v':  // verbose
			verbose = 1;
			break;
		case '#':  // max no. of tweeps
			sscanf(argv[++i], "%d", &ntweeps);
			assert(ntweeps > 0);
			break;
		case 'o': // output file
			out_file = argv[++i];
			break;
		case 'p': // picture file
			picture_file = argv[++i];
			break;
		default: usage(); return ERR1;
		}
	}

	// file system operations
	tf = fopen(tweeps_file, "r");
	assert (tf != NULL);

	tweet_length = strlen(tweet);
	assert (tweet_length < 140);

	tweeps = (char **) malloc(ntweeps * sizeof(char *));
	
	for (n=0; ((read = getline(&line, &len, tf)) != -1); n++) {
		p = tweeps[n] = strdup(line);
		p = strchr(p, '\n');
		assert(p != NULL);
		*p = ' '; // THIS IS A SPACE, USED TO SEPARATE TWEEPS ON THE OUTPUT TWEET
		if (verbose) printf("Tweep[%d] = %s\n", n, tweeps[n]);
	}

	fclose(tf);
	if (line) free(line);

	tf = fopen(out_file, "w");
	assert (tf != NULL);

	// constructing the tweets
	tweep = tweeps;
	i=0;
	first_time=1;

	while (i<n) {
		*buffer = '\0';
		// sum = (picture_file == NULL)? 0:SIZEOFPICTURE;
		// strcpy(buffer, tweet);
		// strcat(buffer, " ");
		for (sum=tweet_length+1;    i<n;    i++, tweep++) {

			l = strlen(*tweep);
			// NOTA: COSI' FUNZIONA if (sum+l>137) break;
			if (sum+l>138) break;
			strcat(buffer, *tweep);
			//sum = strlen(buffer);
			if (verbose) printf("buffer=%s\n", buffer);
			sum += l;
		}

		// only single space chars please
		do { char *s;
			s = strdup(buffer);
			p=strstr(s, "  ");
			if (p==NULL) break;
			*p = '\0';
			sprintf(buffer, "%s%s", s, p+1);
			free(s);
		} while (1);

		len = strlen(buffer)+tweet_length+1;
		assert(len<=139);

		if (first_time)
		{
		printf("TWEET: %s %s (nchar=%zd)\n", tweet, buffer, len), first_time=0;
		//sprintf(format_string, "sudo t update '%%s %%s' -f %s '%%s %%s'", picture_file
		fprintf(tf, "sudo t update '%s %s'\n", tweet, buffer);
		}
		else
		{
		printf("TWEET: %s%s (nchar=%zd)\n", buffer, tweet, len);
		fprintf(tf, "sudo t update '%s%s'\n", buffer, tweet);
		}
	}

	// closings
	fclose(tf);
	for (i=0; i<n; i++) free(tweeps[i]); 
	free(tweeps);
}

void usage(void)
{
fprintf(stderr, "twist: %s\n", STRERR1);
fprintf(stderr, "usage: twist -m messageToTweet -l tweepList -s -v\n");
fprintf(stderr, "       -# maxNumberOfTweeps -o outputScript\n");
//fprintf(stderr, "       -# maxNumberOfTweeps -o outputScript -p pictureFilename\n");
fprintf(stderr, "       -s: simulate (do not tweet), -v: verbose\n");
}
/* End of file twist.c */
