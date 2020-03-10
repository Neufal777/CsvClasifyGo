# Read Csv file [ top 50 Spotify ]

## Usage

```golang
s := &Song{}

files := []string{
	"top50.csv",
}

s.readList(files)
```

## Output

A "Song" struct type with all the information
```

&{3 boyfriend (with Social House) Ariana Grande dance pop 190 80 40 -4 16 70 186 12 46 85}
&{4 Beautiful People (feat. Khalid) Ed Sheeran pop 93 65 64 -8 8 55 198 12 19 86}
&{5 Goodbyes (Feat. Young Thug) Post Malone dfw rap 150 65 58 -4 11 18 175 45 7 94}
&{6 I Don't Care (with Justin Bieber) Ed Sheeran pop 102 68 80 -5 9 84 220 9 4 84}
&{7 Ransom Lil Tecca trap music 180 64 75 -6 7 23 131 2 29 92}
&{8 How Do You Sleep? Sam Smith pop 111 68 48 -5 8 35 202 15 9 90}
&{9 Old Town Road - Remix Lil Nas X country rap 136 62 88 -6 11 64 157 5 10 87}
&{10 bad guy Billie Eilish electropop 135 43 70 -11 10 56 194 33 38 95}
&{11 Callaita Bad Bunny reggaeton 176 62 61 -5 24 24 251 60 31 93}
&{12 Loco Contigo (feat. J. Balvin & Tyga) DJ Snake dance pop 96 71 82 -4 15 38 185 28 7 86}
&{13 Someone You Loved Lewis Capaldi pop 110 41 50 -6 11 45 182 75 3 88}
&{14 Otro Trago - Remix Sech panamanian pop 176 79 73 -2 6 76 288 7 20 87}
&{15 Money In The Grave (Drake ft. Rick Ross) Drake canadian hip hop 101 50 83 -4 12 10 205 10 5 92}
```
