1
00:00:02,500 --> 00:00:08,990
today I'd like to talk about NGO which

2
00:00:07,370 --> 00:00:10,849
is interesting especially interesting

3
00:00:08,990 --> 00:00:12,709
for us in this course because course NGO

4
00:00:10,849 --> 00:00:14,450
is the language at the labs you're all

5
00:00:12,709 --> 00:00:17,090
going to do the labs in and so I want to

6
00:00:14,450 --> 00:00:19,280
focus today particularly on some of the

7
00:00:17,090 --> 00:00:22,940
machinery that sort of most useful in

8
00:00:19,280 --> 00:00:26,170
the labs and in most particular to

9
00:00:22,940 --> 00:00:29,600
distributed programming um first of all

10
00:00:26,170 --> 00:00:32,180
you know it's worth asking why we use go

11
00:00:29,600 --> 00:00:33,890
in this class in fact we could have used

12
00:00:32,180 --> 00:00:35,899
any one of a number of other system

13
00:00:33,890 --> 00:00:38,210
style languages plenty languages like

14
00:00:35,899 --> 00:00:40,730
Java or C sharp or even Python that

15
00:00:38,210 --> 00:00:42,859
provide the kind of facilities we need

16
00:00:40,730 --> 00:00:47,079
and indeed we used to use C++ in this

17
00:00:42,859 --> 00:00:49,010
class and it worked out fine it'll go

18
00:00:47,079 --> 00:00:50,719
indeed like many other languages

19
00:00:49,010 --> 00:00:51,649
provides a bunch of features which are

20
00:00:50,719 --> 00:00:53,409
particularly convenient

21
00:00:51,649 --> 00:00:56,659
that's good support for threads and

22
00:00:53,409 --> 00:00:59,239
locking and synchronization between

23
00:00:56,659 --> 00:01:01,789
threads which we use a lot it is a

24
00:00:59,239 --> 00:01:04,879
convenient remote procedure call package

25
00:01:01,789 --> 00:01:06,380
which doesn't sound like much but it

26
00:01:04,879 --> 00:01:09,950
actually turns out to be a significant

27
00:01:06,380 --> 00:01:11,210
constraint from in languages like C++

28
00:01:09,950 --> 00:01:13,579
for example it's actually a bit hard to

29
00:01:11,210 --> 00:01:14,930
find a convenient easy to use remote

30
00:01:13,579 --> 00:01:16,570
procedure call package and of course we

31
00:01:14,930 --> 00:01:18,409
use it all the time in this course or

32
00:01:16,570 --> 00:01:22,610
programs and different machines to talk

33
00:01:18,409 --> 00:01:25,100
to each other unlike C++ go is type safe

34
00:01:22,610 --> 00:01:27,619
and memory safe that is it's pretty hard

35
00:01:25,100 --> 00:01:29,180
to write a program that due to a bug

36
00:01:27,619 --> 00:01:31,340
scribbles over some random piece of

37
00:01:29,180 --> 00:01:34,789
memory and then causes the program to do

38
00:01:31,340 --> 00:01:36,320
mysterious things and that just

39
00:01:34,789 --> 00:01:39,680
eliminates a big class of bugs similarly

40
00:01:36,320 --> 00:01:41,869
it's garbage collected which means you

41
00:01:39,680 --> 00:01:44,509
never in danger of priam the same memory

42
00:01:41,869 --> 00:01:46,340
twice or free memory that's still in use

43
00:01:44,509 --> 00:01:48,859
or something the garbage vector just

44
00:01:46,340 --> 00:01:51,920
frees things when they stop being used

45
00:01:48,859 --> 00:01:54,829
and one thing it's maybe not obvious

46
00:01:51,920 --> 00:01:56,719
until you played around with just this

47
00:01:54,829 --> 00:01:58,640
kind of programming before but the

48
00:01:56,719 --> 00:02:01,880
combination of threads and garbage

49
00:01:58,640 --> 00:02:03,200
collection is particularly important one

50
00:02:01,880 --> 00:02:06,109
of the things that goes wrong in a non

51
00:02:03,200 --> 00:02:08,899
garbage collected language like C++ if

52
00:02:06,109 --> 00:02:10,690
you use threads is that it's always a

53
00:02:08,899 --> 00:02:13,430
bit of a puzzle and requires a bunch of

54
00:02:10,690 --> 00:02:14,190
bookkeeping to figure out when the last

55
00:02:13,430 --> 00:02:15,660
thread

56
00:02:14,190 --> 00:02:17,340
that's using a shared object has

57
00:02:15,660 --> 00:02:19,530
finished using that object because only

58
00:02:17,340 --> 00:02:20,910
then can you free the object as you end

59
00:02:19,530 --> 00:02:22,620
up writing quite a bit of coat it's like

60
00:02:20,910 --> 00:02:24,480
manually the programmer it's about a

61
00:02:22,620 --> 00:02:26,580
bunch of code to manually you know do

62
00:02:24,480 --> 00:02:28,470
reference counting or something in order

63
00:02:26,580 --> 00:02:30,030
to figure out you know when the last

64
00:02:28,470 --> 00:02:32,460
thread stopped using an object and

65
00:02:30,030 --> 00:02:34,710
that's just a pain and that problem

66
00:02:32,460 --> 00:02:36,560
completely goes away if you use garbage

67
00:02:34,710 --> 00:02:39,390
collection like we haven't go

68
00:02:36,560 --> 00:02:41,460
and finally the language is simple much

69
00:02:39,390 --> 00:02:44,640
simpler than C++ one of the problems

70
00:02:41,460 --> 00:02:47,250
with using C++ is that often if you made

71
00:02:44,640 --> 00:02:51,420
an error you know maybe even just a typo

72
00:02:47,250 --> 00:02:53,730
the the error message you get back from

73
00:02:51,420 --> 00:02:56,160
the compiler is so complicated that in

74
00:02:53,730 --> 00:02:57,510
C++ it's usually not worth trying to

75
00:02:56,160 --> 00:02:59,520
figure out what the error message meant

76
00:02:57,510 --> 00:03:01,470
and I find it's always just much quicker

77
00:02:59,520 --> 00:03:02,670
to go look at the line number and try to

78
00:03:01,470 --> 00:03:04,170
guess what the error must have been

79
00:03:02,670 --> 00:03:04,800
because the language is far too

80
00:03:04,170 --> 00:03:07,140
complicated

81
00:03:04,800 --> 00:03:09,710
whereas go is you know probably doesn't

82
00:03:07,140 --> 00:03:11,580
have a lot of people's favorite features

83
00:03:09,710 --> 00:03:14,940
but it's relatively straightforward

84
00:03:11,580 --> 00:03:17,250
language okay so at this point you're

85
00:03:14,940 --> 00:03:19,290
both on the tutorial if you're looking

86
00:03:17,250 --> 00:03:21,630
for sort of you know what to look at

87
00:03:19,290 --> 00:03:23,190
next to learn about the language a good

88
00:03:21,630 --> 00:03:25,590
place to look is the document titled

89
00:03:23,190 --> 00:03:30,390
effective go which you know you can find

90
00:03:25,590 --> 00:03:33,110
by searching the web all right the first

91
00:03:30,390 --> 00:03:35,940
thing I want to talk about is threads

92
00:03:33,110 --> 00:03:39,000
the reason why we care a lot about

93
00:03:35,940 --> 00:03:41,459
threads in this course is that threads

94
00:03:39,000 --> 00:03:44,370
are the sort of main tool we're going to

95
00:03:41,459 --> 00:03:47,340
be using to manage concurrency in

96
00:03:44,370 --> 00:03:49,920
programs and concurrency is a particular

97
00:03:47,340 --> 00:03:52,440
interest in distributed programming

98
00:03:49,920 --> 00:03:53,820
because it's often the case that one

99
00:03:52,440 --> 00:03:55,890
program actually needs to talk to a

100
00:03:53,820 --> 00:03:58,230
bunch of other computers you know client

101
00:03:55,890 --> 00:04:00,300
may talk to many servers or a server may

102
00:03:58,230 --> 00:04:02,430
be serving requests at the same time on

103
00:04:00,300 --> 00:04:04,830
behalf of many different clients and so

104
00:04:02,430 --> 00:04:06,030
we need a way to say oh you know I'm my

105
00:04:04,830 --> 00:04:07,410
program really has seven different

106
00:04:06,030 --> 00:04:10,110
things going on because it's talking to

107
00:04:07,410 --> 00:04:12,450
seven different clients and I want a

108
00:04:10,110 --> 00:04:14,459
simple way to allow it to do these seven

109
00:04:12,450 --> 00:04:16,859
different things you know without too

110
00:04:14,459 --> 00:04:19,590
much complex programming I mean sort of

111
00:04:16,859 --> 00:04:21,630
thrust threads are the answer so these

112
00:04:19,590 --> 00:04:24,419
are the things that the go documentation

113
00:04:21,630 --> 00:04:26,789
calls go routines which I call threads

114
00:04:24,419 --> 00:04:27,880
they're go routines are really this same

115
00:04:26,789 --> 00:04:32,560
as what everybody else calls

116
00:04:27,880 --> 00:04:36,600
Red's so the way to think of threads is

117
00:04:32,560 --> 00:04:43,120
that you have a program of one program

118
00:04:36,600 --> 00:04:46,060
and one address space I'm gonna draw a

119
00:04:43,120 --> 00:04:48,250
box to sort of denote an address space

120
00:04:46,060 --> 00:04:51,520
and within that address space in a

121
00:04:48,250 --> 00:04:54,550
serial program without threads you just

122
00:04:51,520 --> 00:04:57,040
have one thread of execution executing

123
00:04:54,550 --> 00:05:00,100
code in that address space one program

124
00:04:57,040 --> 00:05:02,080
counter one set of registers one stack

125
00:05:00,100 --> 00:05:04,180
that are sort of describing the current

126
00:05:02,080 --> 00:05:06,160
state of the execution in a threaded

127
00:05:04,180 --> 00:05:09,190
program like a go program you could have

128
00:05:06,160 --> 00:05:10,870
multiple threads and you know I got raw

129
00:05:09,190 --> 00:05:13,480
it as multiple squiggly lines and when

130
00:05:10,870 --> 00:05:16,120
each line represents really is a

131
00:05:13,480 --> 00:05:17,440
separate if the especially if the

132
00:05:16,120 --> 00:05:19,630
threads are executing at the same time

133
00:05:17,440 --> 00:05:21,550
but a separate program counter a

134
00:05:19,630 --> 00:05:24,220
separate set of registers and a separate

135
00:05:21,550 --> 00:05:26,230
stack for each of the threads so that

136
00:05:24,220 --> 00:05:28,330
they can have a sort of their own thread

137
00:05:26,230 --> 00:05:31,810
of control and be executing each thread

138
00:05:28,330 --> 00:05:33,310
in a different part of the program and

139
00:05:31,810 --> 00:05:35,530
so hidden here is that for every stack

140
00:05:33,310 --> 00:05:41,170
now there's a syrupy thread there's a

141
00:05:35,530 --> 00:05:44,530
stack that it's executing on the stacks

142
00:05:41,170 --> 00:05:46,720
are actually in in the one address space

143
00:05:44,530 --> 00:05:47,860
of the program so even though each stack

144
00:05:46,720 --> 00:05:51,220
each thread has its own stack

145
00:05:47,860 --> 00:05:52,240
technically the they're all in the same

146
00:05:51,220 --> 00:05:53,800
address space and different threads

147
00:05:52,240 --> 00:05:55,960
could refer to each other stacks if they

148
00:05:53,800 --> 00:05:59,050
knew the right addresses although you

149
00:05:55,960 --> 00:06:01,510
typically don't do that and then go when

150
00:05:59,050 --> 00:06:02,890
you even the main program you know when

151
00:06:01,510 --> 00:06:05,110
you first start up the program and it

152
00:06:02,890 --> 00:06:06,490
runs in main that's also it's just a go

153
00:06:05,110 --> 00:06:14,440
routine and can do all the things that

154
00:06:06,490 --> 00:06:17,850
go teens can do all right so as I

155
00:06:14,440 --> 00:06:21,730
mentioned one of the big reasons is to

156
00:06:17,850 --> 00:06:25,030
allow different parts of the program to

157
00:06:21,730 --> 00:06:27,550
sort of be in its own point in in a

158
00:06:25,030 --> 00:06:31,510
different activity so I usually refer to

159
00:06:27,550 --> 00:06:36,580
that as IO concurrency for historical

160
00:06:31,510 --> 00:06:38,050
reasons and the reason I call it IO

161
00:06:36,580 --> 00:06:39,580
concurrency is that in the old days

162
00:06:38,050 --> 00:06:41,260
where this first came up is that oh you

163
00:06:39,580 --> 00:06:41,470
might have one thread that's waiting to

164
00:06:41,260 --> 00:06:43,240
read

165
00:06:41,470 --> 00:06:44,410
from the disk and while it's waiting to

166
00:06:43,240 --> 00:06:46,330
reach from the disk you'd like to have a

167
00:06:44,410 --> 00:06:49,000
second thread that maybe can compute or

168
00:06:46,330 --> 00:06:50,560
read somewhere else in the disk or send

169
00:06:49,000 --> 00:06:54,490
a message in the network and wait for

170
00:06:50,560 --> 00:06:57,250
reply so and so I open currencies one of

171
00:06:54,490 --> 00:07:00,190
the things that threads by you for us it

172
00:06:57,250 --> 00:07:01,690
would usually mean I can I open currency

173
00:07:00,190 --> 00:07:04,090
we usually mean I can have one program

174
00:07:01,690 --> 00:07:06,010
that has launched or removed procedure

175
00:07:04,090 --> 00:07:08,140
calls requests to different servers on

176
00:07:06,010 --> 00:07:10,570
the network and is waiting for many

177
00:07:08,140 --> 00:07:13,180
replies at the same time that's how

178
00:07:10,570 --> 00:07:15,040
it'll come up for us and you know the

179
00:07:13,180 --> 00:07:17,110
way you would do that with threads is

180
00:07:15,040 --> 00:07:18,790
that you would create one thread for

181
00:07:17,110 --> 00:07:21,190
each of the remote procedure calls that

182
00:07:18,790 --> 00:07:23,530
you wanted to launch that thread would

183
00:07:21,190 --> 00:07:26,380
have code that you know sent the remote

184
00:07:23,530 --> 00:07:27,820
procedure call request message and sort

185
00:07:26,380 --> 00:07:29,320
of waited at this point in the thread

186
00:07:27,820 --> 00:07:31,330
and then finally when the reply came

187
00:07:29,320 --> 00:07:33,220
back the thread would continue executing

188
00:07:31,330 --> 00:07:34,990
and using threads allows us to have

189
00:07:33,220 --> 00:07:36,250
multiple threads that all launch

190
00:07:34,990 --> 00:07:38,440
requests into the network at the same

191
00:07:36,250 --> 00:07:40,690
time they all wait or they don't have to

192
00:07:38,440 --> 00:07:41,890
do it at the same time they can you know

193
00:07:40,690 --> 00:07:43,090
execute the different parts of this

194
00:07:41,890 --> 00:07:45,370
whenever they feel like it

195
00:07:43,090 --> 00:07:49,720
so that's i/o concurrency sort of

196
00:07:45,370 --> 00:07:53,950
overlapping of the progress of different

197
00:07:49,720 --> 00:07:57,060
activities and allowing one activity is

198
00:07:53,950 --> 00:08:00,330
waiting other activities can proceed

199
00:07:57,060 --> 00:08:02,740
another big reason to use threads is

200
00:08:00,330 --> 00:08:08,680
multi-core parallelism which I'll just

201
00:08:02,740 --> 00:08:10,270
call parallelism and here the thing

202
00:08:08,680 --> 00:08:11,890
where we'd be trying to achieve with

203
00:08:10,270 --> 00:08:13,540
threads is if you have a multi-core

204
00:08:11,890 --> 00:08:15,760
machine like I'm sure all of you do in

205
00:08:13,540 --> 00:08:17,470
your laptops if you have a sort of

206
00:08:15,760 --> 00:08:19,150
compute heavy job that needs a lot of

207
00:08:17,470 --> 00:08:21,060
CPU cycles wouldn't it be nice if you

208
00:08:19,150 --> 00:08:23,860
could have one program that could use

209
00:08:21,060 --> 00:08:25,630
CPU cycles on all of the cores of the

210
00:08:23,860 --> 00:08:27,580
machine and indeed if you write a

211
00:08:25,630 --> 00:08:30,040
multi-threaded go if you launch multiple

212
00:08:27,580 --> 00:08:31,570
go routines and go and they do something

213
00:08:30,040 --> 00:08:33,940
computer intensive like sit there in a

214
00:08:31,570 --> 00:08:36,669
loop and you know compute digits of pi

215
00:08:33,940 --> 00:08:38,520
or something then up to the limit of the

216
00:08:36,669 --> 00:08:41,260
number of cores in the physical machine

217
00:08:38,520 --> 00:08:43,690
your threads will run truly in parallel

218
00:08:41,260 --> 00:08:45,750
and if you launch you know two threads

219
00:08:43,690 --> 00:08:48,370
instead of one you'll get twice as many

220
00:08:45,750 --> 00:08:51,070
you'll be able to use twice as many CPU

221
00:08:48,370 --> 00:08:53,170
cycles per second so this is very

222
00:08:51,070 --> 00:08:54,740
important to some people it's not a big

223
00:08:53,170 --> 00:08:57,440
deal on this course

224
00:08:54,740 --> 00:08:59,480
be it's rare that we'll sort of think

225
00:08:57,440 --> 00:09:01,640
specifically about this kind of

226
00:08:59,480 --> 00:09:05,390
parallelism in the real world though of

227
00:09:01,640 --> 00:09:06,890
building things like servers to form

228
00:09:05,390 --> 00:09:09,770
parts of the distributed systems it can

229
00:09:06,890 --> 00:09:11,780
sometimes be extremely important to be

230
00:09:09,770 --> 00:09:13,550
able to have the server be able to run

231
00:09:11,780 --> 00:09:15,140
threads and harness the CPU power of a

232
00:09:13,550 --> 00:09:18,740
lot of cores just because the load from

233
00:09:15,140 --> 00:09:22,850
clients can often be pretty high okay so

234
00:09:18,740 --> 00:09:25,310
parallelism is a second reason why

235
00:09:22,850 --> 00:09:27,170
threads are quite a bit interested in

236
00:09:25,310 --> 00:09:29,540
distributed systems and a third reason

237
00:09:27,170 --> 00:09:32,780
which is maybe a little bit less

238
00:09:29,540 --> 00:09:35,510
important is there's some there's times

239
00:09:32,780 --> 00:09:38,630
when you really just want to be able to

240
00:09:35,510 --> 00:09:39,800
do something in the background or you

241
00:09:38,630 --> 00:09:42,770
know there's just something you need to

242
00:09:39,800 --> 00:09:45,260
do periodically and you don't want to

243
00:09:42,770 --> 00:09:47,420
have to sort of in the main part of your

244
00:09:45,260 --> 00:09:49,190
program sort of insert checks to say

245
00:09:47,420 --> 00:09:51,080
well should I be doing this things that

246
00:09:49,190 --> 00:09:52,370
should happen every second or so you

247
00:09:51,080 --> 00:09:54,380
just like to be able to fire something

248
00:09:52,370 --> 00:09:56,060
up that every second does whatever the

249
00:09:54,380 --> 00:10:00,500
periodic thing is so there's some

250
00:09:56,060 --> 00:10:03,380
convenience reasons and an example which

251
00:10:00,500 --> 00:10:05,210
will come up for you is it's often the

252
00:10:03,380 --> 00:10:07,250
case that some you know a master server

253
00:10:05,210 --> 00:10:09,350
may want to check periodically whether

254
00:10:07,250 --> 00:10:10,580
its workers are still alive because one

255
00:10:09,350 --> 00:10:12,170
of them is died you know you want to

256
00:10:10,580 --> 00:10:14,540
launch that work on another machine like

257
00:10:12,170 --> 00:10:17,420
MapReduce might do that and one way to

258
00:10:14,540 --> 00:10:19,100
arrange sort of oh do this check every

259
00:10:17,420 --> 00:10:21,680
second every minute you know send a

260
00:10:19,100 --> 00:10:24,170
message to the worker are you alive is

261
00:10:21,680 --> 00:10:25,700
to fire off a go routine that just sits

262
00:10:24,170 --> 00:10:26,990
in a loop that sleeps for a second and

263
00:10:25,700 --> 00:10:28,700
then does the periodic thing and then

264
00:10:26,990 --> 00:10:31,400
sleeps for a second again and so in the

265
00:10:28,700 --> 00:10:36,560
labs you'll end up firing off these kind

266
00:10:31,400 --> 00:10:42,470
of threads quite a bit yes is the

267
00:10:36,560 --> 00:10:44,840
overhead worth it yes the overhead is

268
00:10:42,470 --> 00:10:46,700
really pretty small for this stuff I

269
00:10:44,840 --> 00:10:50,150
mean you know it depends on how many you

270
00:10:46,700 --> 00:10:52,040
create a million threads that he sit in

271
00:10:50,150 --> 00:10:53,960
a loop waiting for a millisecond and

272
00:10:52,040 --> 00:10:56,330
then send a network message that's

273
00:10:53,960 --> 00:10:59,200
probably a huge load on your machine but

274
00:10:56,330 --> 00:11:01,400
if you create you know ten threads that

275
00:10:59,200 --> 00:11:04,850
sleep for a second and do a little bit

276
00:11:01,400 --> 00:11:06,690
of work it's probably not a big deal at

277
00:11:04,850 --> 00:11:10,019
all and it's

278
00:11:06,690 --> 00:11:13,980
I guarantee you the programmer time you

279
00:11:10,019 --> 00:11:16,199
say by not having to sort of mush

280
00:11:13,980 --> 00:11:19,319
together they're different different

281
00:11:16,199 --> 00:11:21,449
activities into one line of code it's

282
00:11:19,319 --> 00:11:26,160
it's worth the small amount of CPU cost

283
00:11:21,449 --> 00:11:27,600
almost always still you know you will if

284
00:11:26,160 --> 00:11:30,449
you're unlucky you'll discover in the

285
00:11:27,600 --> 00:11:32,519
labs that some loop of yours is not

286
00:11:30,449 --> 00:11:35,730
sleeping long enough or are you fired

287
00:11:32,519 --> 00:11:37,589
off a bunch of these and never made them

288
00:11:35,730 --> 00:11:41,329
exit for example and they just

289
00:11:37,589 --> 00:11:43,680
accumulate so you can push it too far

290
00:11:41,329 --> 00:11:46,410
okay so these are the reasons that the

291
00:11:43,680 --> 00:11:47,730
main reasons that people like threads a

292
00:11:46,410 --> 00:11:50,370
lot and that will use threads in this

293
00:11:47,730 --> 00:12:01,860
class any other questions about threads

294
00:11:50,370 --> 00:12:03,569
in general by asynchronous program you

295
00:12:01,860 --> 00:12:06,779
mean like a single thread of control

296
00:12:03,569 --> 00:12:09,509
that keeps state about many different

297
00:12:06,779 --> 00:12:12,209
activities yeah so this is a good

298
00:12:09,509 --> 00:12:13,439
question actually there is you know what

299
00:12:12,209 --> 00:12:15,089
would happen if we didn't have threads

300
00:12:13,439 --> 00:12:16,800
or we'd for some reason we didn't want

301
00:12:15,089 --> 00:12:18,899
to use threats like how would we be able

302
00:12:16,800 --> 00:12:21,209
to write a program that could you know a

303
00:12:18,899 --> 00:12:23,339
server that could talk to many different

304
00:12:21,209 --> 00:12:24,509
clients at the same time or a client

305
00:12:23,339 --> 00:12:26,100
that could talk to him any servers right

306
00:12:24,509 --> 00:12:29,069
what what tools could be used and it

307
00:12:26,100 --> 00:12:36,420
turns out there is sort of another line

308
00:12:29,069 --> 00:12:37,980
of another kind of another major style

309
00:12:36,420 --> 00:12:38,339
of how do you structure these programs

310
00:12:37,980 --> 00:12:40,019
called

311
00:12:38,339 --> 00:12:42,529
you call the asynchronous program I

312
00:12:40,019 --> 00:12:45,899
might call it a vent driven programming

313
00:12:42,529 --> 00:12:50,850
so sort of or you could use a vent

314
00:12:45,899 --> 00:12:52,709
prevent programming and the the general

315
00:12:50,850 --> 00:12:54,420
structure of an event-driven program is

316
00:12:52,709 --> 00:12:57,380
usually that it has a single thread and

317
00:12:54,420 --> 00:13:01,380
a single loop and what that loop does is

318
00:12:57,380 --> 00:13:03,779
sits there and waits for any input or

319
00:13:01,380 --> 00:13:05,699
sort of any event that might trigger

320
00:13:03,779 --> 00:13:07,800
processing so an event might be the

321
00:13:05,699 --> 00:13:10,889
arrival of a request from a client or a

322
00:13:07,800 --> 00:13:12,629
timer going off or if you're building a

323
00:13:10,889 --> 00:13:14,189
Window System protect many Windows

324
00:13:12,629 --> 00:13:16,019
systems on your laptops I've driven

325
00:13:14,189 --> 00:13:17,639
written an event-driven style where what

326
00:13:16,019 --> 00:13:18,339
they're waiting for is like key clicks

327
00:13:17,639 --> 00:13:20,259
or Mouse move

328
00:13:18,339 --> 00:13:21,550
or something so you might have a single

329
00:13:20,259 --> 00:13:23,350
in an event-driven program it of a

330
00:13:21,550 --> 00:13:25,269
single threat of control sits an aloof

331
00:13:23,350 --> 00:13:27,160
waits for input and whenever it gets an

332
00:13:25,269 --> 00:13:28,749
input like a packet it figures out oh

333
00:13:27,160 --> 00:13:31,540
you know which client did this packet

334
00:13:28,749 --> 00:13:34,839
come from and then it'll have a table of

335
00:13:31,540 --> 00:13:38,199
sort of what the state is of whatever

336
00:13:34,839 --> 00:13:40,509
activity its managing for that client

337
00:13:38,199 --> 00:13:42,220
and it'll say oh gosh I was in the

338
00:13:40,509 --> 00:13:44,110
middle of reading such-and-such a file

339
00:13:42,220 --> 00:13:45,850
you know now it's asked me to read the

340
00:13:44,110 --> 00:13:55,480
next block I'll go and be the next block

341
00:13:45,850 --> 00:13:57,249
and return it and threats are generally

342
00:13:55,480 --> 00:14:00,069
more convenient because they allow you

343
00:13:57,249 --> 00:14:01,689
to really you know it's much easier to

344
00:14:00,069 --> 00:14:03,610
write sequential just like straight

345
00:14:01,689 --> 00:14:05,649
lines of control code that does you know

346
00:14:03,610 --> 00:14:07,209
computes sends a message waits for

347
00:14:05,649 --> 00:14:09,009
response whatever it's much easier to

348
00:14:07,209 --> 00:14:13,509
write that kind of code in a thread than

349
00:14:09,009 --> 00:14:16,269
it is to chop up whatever the activity

350
00:14:13,509 --> 00:14:19,089
is into a bunch of little pieces that

351
00:14:16,269 --> 00:14:23,379
can sort of be activated one at a time

352
00:14:19,089 --> 00:14:29,110
by one of these event-driven loops that

353
00:14:23,379 --> 00:14:31,089
said the well and so one problem with

354
00:14:29,110 --> 00:14:32,889
the scheme is that it's it's a little

355
00:14:31,089 --> 00:14:34,660
bit of a pain to program another

356
00:14:32,889 --> 00:14:36,639
potential defect is that while you get

357
00:14:34,660 --> 00:14:38,769
io concurrency from this approach you

358
00:14:36,639 --> 00:14:39,910
don't get CPU parallelism so if you're

359
00:14:38,769 --> 00:14:42,730
writing a busy server that would really

360
00:14:39,910 --> 00:14:45,610
like to keep you know 32 cores busy on a

361
00:14:42,730 --> 00:14:49,149
big server machine you know a single

362
00:14:45,610 --> 00:14:50,620
loop is you know it's it's not a very

363
00:14:49,149 --> 00:14:55,029
natural way to harness more than one

364
00:14:50,620 --> 00:14:56,620
core on the other hand the overheads of

365
00:14:55,029 --> 00:14:59,259
adventure and programming are generally

366
00:14:56,620 --> 00:15:01,660
quite a bit less than threads you know

367
00:14:59,259 --> 00:15:05,379
Ed's are pretty cheap but each one of

368
00:15:01,660 --> 00:15:07,329
these threads is sitting on a stack you

369
00:15:05,379 --> 00:15:09,459
know stack is a kilobyte or a kilobytes

370
00:15:07,329 --> 00:15:10,959
or something you know if you have 20 of

371
00:15:09,459 --> 00:15:12,970
these threads who cares if you have a

372
00:15:10,959 --> 00:15:14,009
million of these threads then it's

373
00:15:12,970 --> 00:15:17,769
starting to be a huge amount of memory

374
00:15:14,009 --> 00:15:19,389
and you know maybe the scheduling

375
00:15:17,769 --> 00:15:20,980
bookkeeping for deciding what the thread

376
00:15:19,389 --> 00:15:23,259
to run next might also start you know

377
00:15:20,980 --> 00:15:25,899
you now have list scheduling lists with

378
00:15:23,259 --> 00:15:28,420
a thousand threads in them the threads

379
00:15:25,899 --> 00:15:30,519
can start to get quite expensive so if

380
00:15:28,420 --> 00:15:31,960
you are in a position where you need to

381
00:15:30,519 --> 00:15:33,520
have a single server that sir

382
00:15:31,960 --> 00:15:35,050
you know a million clients and has to

383
00:15:33,520 --> 00:15:37,920
sort of keep a little bit of state for

384
00:15:35,050 --> 00:15:39,370
each of a million clients this could be

385
00:15:37,920 --> 00:15:43,900
expensive

386
00:15:39,370 --> 00:15:46,060
and it's easier to write a very you know

387
00:15:43,900 --> 00:15:47,440
at some expense in programmer time it's

388
00:15:46,060 --> 00:15:50,560
easier to write a really stripped-down

389
00:15:47,440 --> 00:15:51,940
efficient low overhead service in a

390
00:15:50,560 --> 00:16:15,390
venture than programming just a lot more

391
00:15:51,940 --> 00:16:18,040
work are you asking my JavaScript I

392
00:16:15,390 --> 00:16:20,950
don't know the question is whether

393
00:16:18,040 --> 00:16:25,870
JavaScript has multiple cores executing

394
00:16:20,950 --> 00:16:27,250
your does anybody know depends on the

395
00:16:25,870 --> 00:16:29,230
implementation yeah so I don't know I

396
00:16:27,250 --> 00:16:31,570
mean it's a natural thought though even

397
00:16:29,230 --> 00:16:33,400
in you know even an NGO you might well

398
00:16:31,570 --> 00:16:35,170
want to have if you knew your machine

399
00:16:33,400 --> 00:16:37,000
had eight cores if you wanted to write

400
00:16:35,170 --> 00:16:39,490
the world's most efficient whatever

401
00:16:37,000 --> 00:16:42,850
server you could fire up eight threads

402
00:16:39,490 --> 00:16:47,350
and on each of the threads run sort of

403
00:16:42,850 --> 00:16:49,660
stripped-down event-driven loop just you

404
00:16:47,350 --> 00:16:51,160
know sort of one event loop Recor and

405
00:16:49,660 --> 00:16:54,700
that you know that would be a way to get

406
00:16:51,160 --> 00:16:59,160
both parallelism and to the bio

407
00:16:54,700 --> 00:16:59,160
concurrency yes

408
00:17:05,060 --> 00:17:07,950
okay so the question is what's the

409
00:17:06,870 --> 00:17:11,190
difference between threads and processes

410
00:17:07,950 --> 00:17:14,880
so usually on a like a UNIX machine a

411
00:17:11,190 --> 00:17:16,650
process is a single program that you're

412
00:17:14,880 --> 00:17:18,690
running and a sort of single address

413
00:17:16,650 --> 00:17:22,110
space a single bunch of memory for the

414
00:17:18,690 --> 00:17:23,640
process and inside a process you might

415
00:17:22,110 --> 00:17:25,290
have multiple threads and when you ready

416
00:17:23,640 --> 00:17:28,350
to go program and you run the go program

417
00:17:25,290 --> 00:17:31,890
running the go program creates one unix

418
00:17:28,350 --> 00:17:35,880
process and one sort of memory area and

419
00:17:31,890 --> 00:17:37,380
then when your go process creates go

420
00:17:35,880 --> 00:17:40,950
routines those are so sitting inside

421
00:17:37,380 --> 00:17:42,990
that one process so I'm not sure that's

422
00:17:40,950 --> 00:17:45,360
really an answer but just historically

423
00:17:42,990 --> 00:17:47,490
the operating systems have provided like

424
00:17:45,360 --> 00:17:49,580
this big box is the process that's

425
00:17:47,490 --> 00:17:52,050
implemented by the operating system and

426
00:17:49,580 --> 00:17:53,940
the individual and some of the operating

427
00:17:52,050 --> 00:17:56,550
system does not care what happens inside

428
00:17:53,940 --> 00:17:59,130
your process what language you use none

429
00:17:56,550 --> 00:18:00,810
of the operating systems business but

430
00:17:59,130 --> 00:18:03,510
inside that process you can run lots of

431
00:18:00,810 --> 00:18:05,220
threads now you know if you run more

432
00:18:03,510 --> 00:18:06,630
than one process in your machine you

433
00:18:05,220 --> 00:18:09,570
know you run more than one program I can

434
00:18:06,630 --> 00:18:12,360
edit or compiler the operating system

435
00:18:09,570 --> 00:18:13,620
keep quite separate right you're your

436
00:18:12,360 --> 00:18:15,690
editor and your compiler each have

437
00:18:13,620 --> 00:18:16,830
memory but it's not the same memory that

438
00:18:15,690 --> 00:18:18,810
are not allowed to look at each other

439
00:18:16,830 --> 00:18:20,040
memory there's not much interaction

440
00:18:18,810 --> 00:18:22,470
between different processes so you

441
00:18:20,040 --> 00:18:24,030
redditor may have threads and your

442
00:18:22,470 --> 00:18:27,480
compiler may have threads but they're

443
00:18:24,030 --> 00:18:29,190
just in different worlds so within any

444
00:18:27,480 --> 00:18:31,890
one program the threads can share memory

445
00:18:29,190 --> 00:18:33,720
and can synchronize with channels and

446
00:18:31,890 --> 00:18:38,240
use mutexes and stuff but between

447
00:18:33,720 --> 00:18:41,310
processes there's just no no interaction

448
00:18:38,240 --> 00:18:45,140
that's just a traditional structure of

449
00:18:41,310 --> 00:18:45,140
these this kind of software

450
00:18:45,509 --> 00:18:48,509
yeah

451
00:18:53,370 --> 00:18:59,190
so the question is when a context switch

452
00:18:55,480 --> 00:18:59,190
happens does it happened for all threads

453
00:19:08,010 --> 00:19:12,580
okay so let's let's imagine you have a

454
00:19:10,420 --> 00:19:14,530
single core machine that's really only

455
00:19:12,580 --> 00:19:19,900
running and as just doing one thing at a

456
00:19:14,530 --> 00:19:21,580
time maybe the right way to think about

457
00:19:19,900 --> 00:19:22,960
it is that you're going to be you're

458
00:19:21,580 --> 00:19:27,520
running multiple processes on your

459
00:19:22,960 --> 00:19:31,180
machine the operating system will give

460
00:19:27,520 --> 00:19:33,430
the CPU sort of time slicing back and

461
00:19:31,180 --> 00:19:35,950
forth between these two programs so when

462
00:19:33,430 --> 00:19:37,300
the hardware timer ticks and the

463
00:19:35,950 --> 00:19:38,950
operating systems decides it's time to

464
00:19:37,300 --> 00:19:40,480
take away the CPU from the currently

465
00:19:38,950 --> 00:19:44,700
running process and give it to another

466
00:19:40,480 --> 00:19:44,700
process that's done at a process level

467
00:19:48,600 --> 00:19:55,240
it's complicated all right let me let me

468
00:19:52,330 --> 00:19:57,730
let me restart this these the threads

469
00:19:55,240 --> 00:20:00,070
that we use are based on threads that

470
00:19:57,730 --> 00:20:02,080
are provided by the operating system in

471
00:20:00,070 --> 00:20:03,910
the end and when the outer needs to some

472
00:20:02,080 --> 00:20:06,520
context switches its switching between

473
00:20:03,910 --> 00:20:08,620
the threads that it knows about so in a

474
00:20:06,520 --> 00:20:09,820
situation like this the operating system

475
00:20:08,620 --> 00:20:11,560
might know that there are two threads

476
00:20:09,820 --> 00:20:13,600
here in this process and three threads

477
00:20:11,560 --> 00:20:15,190
in this process and when the timer ticks

478
00:20:13,600 --> 00:20:16,870
the operating system will based on some

479
00:20:15,190 --> 00:20:18,730
scheduling algorithm pick a different

480
00:20:16,870 --> 00:20:21,130
thread to run it might be a different

481
00:20:18,730 --> 00:20:22,270
thread in this process or one of the

482
00:20:21,130 --> 00:20:25,690
threads in this process

483
00:20:22,270 --> 00:20:27,780
in addition go cleverly multiplex as

484
00:20:25,690 --> 00:20:29,860
many go routines on top of single

485
00:20:27,780 --> 00:20:32,590
operating system threads to reduce

486
00:20:29,860 --> 00:20:34,480
overhead so it's really probably a two

487
00:20:32,590 --> 00:20:37,330
stages of scheduling the operating

488
00:20:34,480 --> 00:20:40,030
system picks which big thread to run and

489
00:20:37,330 --> 00:20:43,290
then within that process go may have a

490
00:20:40,030 --> 00:20:43,290
choice of go routines to run

491
00:20:45,820 --> 00:20:55,539
all right okay so threads are convenient

492
00:20:53,409 --> 00:20:57,549
because a lot of times they allow you to

493
00:20:55,539 --> 00:20:59,919
write the code for each thread just as

494
00:20:57,549 --> 00:21:04,539
if it were a pretty ordinary sequential

495
00:20:59,919 --> 00:21:10,529
program however there are in fact some

496
00:21:04,539 --> 00:21:10,529
challenges with writing threaded code

497
00:21:15,149 --> 00:21:18,999
one is what to do about shared data one

498
00:21:17,679 --> 00:21:20,649
of the really cool things about the

499
00:21:18,999 --> 00:21:22,359
threading model is that these threads

500
00:21:20,649 --> 00:21:24,309
share the same address space they share

501
00:21:22,359 --> 00:21:26,919
memory if one thread creates an object

502
00:21:24,309 --> 00:21:29,139
in memory you can let other threads use

503
00:21:26,919 --> 00:21:30,489
it right you can have a array or

504
00:21:29,139 --> 00:21:31,419
something that all the different threads

505
00:21:30,489 --> 00:21:33,879
are reading and writing and that

506
00:21:31,419 --> 00:21:35,080
sometimes critical right if you you know

507
00:21:33,879 --> 00:21:36,309
if you're keeping some interesting state

508
00:21:35,080 --> 00:21:39,099
you know maybe you have a cache of

509
00:21:36,309 --> 00:21:40,840
things that your server your cache and

510
00:21:39,099 --> 00:21:42,309
memory when a thread is handling a

511
00:21:40,840 --> 00:21:43,809
client request it's gonna first look in

512
00:21:42,309 --> 00:21:45,399
that cache but the shared cache and each

513
00:21:43,809 --> 00:21:48,009
thread reads it and the threads may

514
00:21:45,399 --> 00:21:49,179
write the cache to update it when they

515
00:21:48,009 --> 00:21:51,220
have new information to stick in the

516
00:21:49,179 --> 00:21:55,210
cache so it's really cool you can share

517
00:21:51,220 --> 00:21:57,970
that memory but it turns out that it's

518
00:21:55,210 --> 00:21:59,739
very very easy to get bugs if you're not

519
00:21:57,970 --> 00:22:02,590
careful and you're sharing memory

520
00:21:59,739 --> 00:22:05,529
between threads so a totally classic

521
00:22:02,590 --> 00:22:07,629
example is you know supposing your

522
00:22:05,529 --> 00:22:09,669
thread so you have a global variable N

523
00:22:07,629 --> 00:22:11,440
and that's shared among the different

524
00:22:09,669 --> 00:22:17,559
threads and a thread just wants to

525
00:22:11,440 --> 00:22:20,710
increment n right but itself this is

526
00:22:17,559 --> 00:22:22,239
likely to be an invitation to bugs right

527
00:22:20,710 --> 00:22:25,119
if you don't do anything special around

528
00:22:22,239 --> 00:22:27,519
this code and the reason is that you

529
00:22:25,119 --> 00:22:29,919
know whenever you write code in a thread

530
00:22:27,519 --> 00:22:31,899
that you you know is accessing reading

531
00:22:29,919 --> 00:22:33,669
or writing data that's shared with other

532
00:22:31,899 --> 00:22:35,259
threads you know there's always the

533
00:22:33,669 --> 00:22:36,999
possibility and you got to keep in mind

534
00:22:35,259 --> 00:22:39,549
that some other thread may be looking at

535
00:22:36,999 --> 00:22:41,739
the data or modifying the data at the

536
00:22:39,549 --> 00:22:43,899
same time so the obvious problem with

537
00:22:41,739 --> 00:22:46,119
this is that maybe thread 1 is executing

538
00:22:43,899 --> 00:22:47,859
this code and thread 2 is actually in

539
00:22:46,119 --> 00:22:50,799
the same function in a different thread

540
00:22:47,859 --> 00:22:53,049
executing the very same code right and

541
00:22:50,799 --> 00:22:54,369
remember I'm imagining the N is a global

542
00:22:53,049 --> 00:22:57,039
variable so they're talking about the

543
00:22:54,369 --> 00:22:58,269
same n so what this boils down to you

544
00:22:57,039 --> 00:22:59,049
know you're not actually running this

545
00:22:58,269 --> 00:23:01,749
code you're running

546
00:22:59,049 --> 00:23:03,879
machine code the compiler produced and

547
00:23:01,749 --> 00:23:06,720
what that machine code does is it you

548
00:23:03,879 --> 00:23:13,210
know it loads X into a register

549
00:23:06,720 --> 00:23:18,129
you know adds one to the register and

550
00:23:13,210 --> 00:23:21,009
then stores that register into X with

551
00:23:18,129 --> 00:23:23,080
where X is a address of some location

552
00:23:21,009 --> 00:23:24,340
and ran so you know you can count on

553
00:23:23,080 --> 00:23:25,929
both of the threads

554
00:23:24,340 --> 00:23:28,720
they're both executing this line of code

555
00:23:25,929 --> 00:23:31,090
you know they both load the variable X

556
00:23:28,720 --> 00:23:32,470
into a register effect starts out at 0

557
00:23:31,090 --> 00:23:33,970
that means they both load at 0

558
00:23:32,470 --> 00:23:35,619
they both increment that register so

559
00:23:33,970 --> 00:23:37,809
they get one and they both store one

560
00:23:35,619 --> 00:23:39,999
back to memory and now two threads of

561
00:23:37,809 --> 00:23:44,019
incremented n and the resulting value is

562
00:23:39,999 --> 00:23:45,489
1 which well who knows what the

563
00:23:44,019 --> 00:23:47,289
programmer intended maybe that's what

564
00:23:45,489 --> 00:23:49,149
the programmer wanted but chances are

565
00:23:47,289 --> 00:24:02,710
not right chances are the programmer

566
00:23:49,149 --> 00:24:04,239
wanted to not 1 some some instructions

567
00:24:02,710 --> 00:24:09,639
are atomic so the question is a very

568
00:24:04,239 --> 00:24:11,259
good question which it's whether

569
00:24:09,639 --> 00:24:13,269
individual instructions are atomic so

570
00:24:11,259 --> 00:24:20,559
the answer is some are and some aren't

571
00:24:13,269 --> 00:24:23,440
so a store a 32-bit store is likely the

572
00:24:20,559 --> 00:24:25,749
extremely likely to be atomic in the

573
00:24:23,440 --> 00:24:27,960
sense that if 2 processors store at the

574
00:24:25,749 --> 00:24:30,730
same time to the same memory address

575
00:24:27,960 --> 00:24:33,159
32-bit values well you'll end up with is

576
00:24:30,730 --> 00:24:35,169
either the 32 bits from one processor or

577
00:24:33,159 --> 00:24:38,200
the 32 bits from the other processor but

578
00:24:35,169 --> 00:24:40,450
not a mixture other sizes it's not so

579
00:24:38,200 --> 00:24:41,980
clear like one byte stores it depends on

580
00:24:40,450 --> 00:24:44,440
the CPU you using because a one byte

581
00:24:41,980 --> 00:24:47,919
store is really almost certainly a 32

582
00:24:44,440 --> 00:24:50,350
byte load and then a modification of 8

583
00:24:47,919 --> 00:24:51,820
bits and a 32 byte store but it depends

584
00:24:50,350 --> 00:24:54,609
on the processor and more complicated

585
00:24:51,820 --> 00:24:55,899
instructions like increment your

586
00:24:54,609 --> 00:24:57,759
microprocessor may well have an

587
00:24:55,899 --> 00:25:00,220
increment instruction that can directly

588
00:24:57,759 --> 00:25:04,239
increment some memory location like

589
00:25:00,220 --> 00:25:05,619
pretty unlikely to be atomic although

590
00:25:04,239 --> 00:25:07,890
there's atomic versions of some of these

591
00:25:05,619 --> 00:25:12,740
instructions

592
00:25:07,890 --> 00:25:16,679
so there's no way all right so this is

593
00:25:12,740 --> 00:25:20,490
this is a just classic danger and it's

594
00:25:16,679 --> 00:25:22,679
usually called a race I'm gonna come up

595
00:25:20,490 --> 00:25:25,280
a lot is you're gonna do a lot of

596
00:25:22,679 --> 00:25:27,750
threaded programming with shared state

597
00:25:25,280 --> 00:25:30,570
race I think refers to as some ancient

598
00:25:27,750 --> 00:25:33,660
class of bugs involving electronic

599
00:25:30,570 --> 00:25:35,040
circuits but for us that you know the

600
00:25:33,660 --> 00:25:37,770
reason why it's called a race is because

601
00:25:35,040 --> 00:25:43,110
if one of the CPUs have started

602
00:25:37,770 --> 00:25:44,790
executing this code and the other one

603
00:25:43,110 --> 00:25:46,320
the others thread is sort of getting

604
00:25:44,790 --> 00:25:48,090
close to this code it's sort of a race

605
00:25:46,320 --> 00:25:50,309
as to whether the first processor can

606
00:25:48,090 --> 00:25:53,370
finish and get to the store before the

607
00:25:50,309 --> 00:25:54,809
second processor start status execute

608
00:25:53,370 --> 00:25:57,210
the load if the first processor actually

609
00:25:54,809 --> 00:25:59,070
manages it to do the store before the

610
00:25:57,210 --> 00:26:00,750
second processor gets to the load then

611
00:25:59,070 --> 00:26:03,450
the second processor will see the stored

612
00:26:00,750 --> 00:26:07,190
value and the second processor will load

613
00:26:03,450 --> 00:26:07,190
one and add one to it in store two

614
00:26:07,370 --> 00:26:13,679
that's how you can justify this

615
00:26:10,200 --> 00:26:15,270
terminology okay and so the way you

616
00:26:13,679 --> 00:26:16,890
solve this certainly something this

617
00:26:15,270 --> 00:26:19,470
simple is you insert locks

618
00:26:16,890 --> 00:26:23,910
you know you as a programmer you have

619
00:26:19,470 --> 00:26:26,270
some strategy in mind for locking the

620
00:26:23,910 --> 00:26:28,830
data you can say well you know this

621
00:26:26,270 --> 00:26:31,500
piece of shared data can only be used

622
00:26:28,830 --> 00:26:33,090
when such-and-such a lock is held and

623
00:26:31,500 --> 00:26:36,660
you'll see this and you may have used

624
00:26:33,090 --> 00:26:39,799
this in the tutorial the go calls locks

625
00:26:36,660 --> 00:26:44,130
mutexes so what you'll see is a mule Ock

626
00:26:39,799 --> 00:26:48,320
before a sequence of code that uses

627
00:26:44,130 --> 00:26:52,020
shared data and you unlock afterwards

628
00:26:48,320 --> 00:26:53,340
and then whichever two threads execute

629
00:26:52,020 --> 00:26:56,220
this when it to everyone is lucky enough

630
00:26:53,340 --> 00:26:57,600
to get the lock first gets to do all

631
00:26:56,220 --> 00:26:59,850
this stuff and finish before the other

632
00:26:57,600 --> 00:27:02,340
one is allowed to proceed and so you can

633
00:26:59,850 --> 00:27:05,460
think of wrapping a some code in a lock

634
00:27:02,340 --> 00:27:07,020
as making a bunch of you know remember

635
00:27:05,460 --> 00:27:10,380
this even though it's one line it's

636
00:27:07,020 --> 00:27:13,669
really three distinct operations you can

637
00:27:10,380 --> 00:27:16,040
think of a lock as causing this sort of

638
00:27:13,669 --> 00:27:18,240
multi-step code sequence to be atomic

639
00:27:16,040 --> 00:27:21,320
with respect to other people who have to

640
00:27:18,240 --> 00:27:21,320
lock yes

641
00:27:26,370 --> 00:27:30,970
should you can you repeat the

642
00:27:30,580 --> 00:27:37,090
question

643
00:27:30,970 --> 00:27:39,040
oh that's a great question the question

644
00:27:37,090 --> 00:27:41,350
was how does go know which variable

645
00:27:39,040 --> 00:27:43,090
we're walking right here of course is

646
00:27:41,350 --> 00:27:45,910
only one variable but maybe we're saying

647
00:27:43,090 --> 00:27:47,440
an equals x plus y really threes few

648
00:27:45,910 --> 00:27:52,890
different variables and the answer is

649
00:27:47,440 --> 00:27:55,030
that go has no idea it's not there's no

650
00:27:52,890 --> 00:27:58,030
Association at all

651
00:27:55,030 --> 00:28:00,640
anywhere between this lock so this new

652
00:27:58,030 --> 00:28:04,600
thing is a variable which is a tight

653
00:28:00,640 --> 00:28:07,630
mutex there's just there's no

654
00:28:04,600 --> 00:28:10,720
association in the language between the

655
00:28:07,630 --> 00:28:12,340
lock and any variables the associations

656
00:28:10,720 --> 00:28:14,470
in the programmers head so as a

657
00:28:12,340 --> 00:28:18,670
programmer you need to say oh here's a

658
00:28:14,470 --> 00:28:20,770
bunch of shared data and any time you

659
00:28:18,670 --> 00:28:22,810
modify any of it you know here's a

660
00:28:20,770 --> 00:28:24,280
complex data structure say a tree or an

661
00:28:22,810 --> 00:28:26,560
expandable hash table or something

662
00:28:24,280 --> 00:28:27,820
anytime you're going to modify it and of

663
00:28:26,560 --> 00:28:29,710
course a tree is composed many many

664
00:28:27,820 --> 00:28:30,880
objects anytime you got to modify

665
00:28:29,710 --> 00:28:32,290
anything that's associated with this

666
00:28:30,880 --> 00:28:34,540
data structure you have to hold such and

667
00:28:32,290 --> 00:28:36,580
such a lock right and of course is many

668
00:28:34,540 --> 00:28:37,630
objects and instead of objects changes

669
00:28:36,580 --> 00:28:40,120
because you might allocate new tree

670
00:28:37,630 --> 00:28:41,980
nodes but it's really the programmer who

671
00:28:40,120 --> 00:28:44,710
sort of works out a strategy for

672
00:28:41,980 --> 00:28:47,470
ensuring that the data structure is used

673
00:28:44,710 --> 00:28:50,260
by only one core at a time and so it

674
00:28:47,470 --> 00:28:51,670
creates the one or maybe more locks and

675
00:28:50,260 --> 00:28:53,020
there's many many locking strategies you

676
00:28:51,670 --> 00:28:57,580
could apply to a tree you can imagine a

677
00:28:53,020 --> 00:28:58,780
tree with a lock for every tree node the

678
00:28:57,580 --> 00:29:00,670
programmer works out the strategy

679
00:28:58,780 --> 00:29:02,230
allocates the locks and keeps in the

680
00:29:00,670 --> 00:29:04,960
programmers head the relationship to the

681
00:29:02,230 --> 00:29:07,690
data but go for go it's this is this

682
00:29:04,960 --> 00:29:10,930
lock it's just like a very simple thing

683
00:29:07,690 --> 00:29:12,970
there's a lock object the first thread

684
00:29:10,930 --> 00:29:14,770
that calls lock gets the lock other

685
00:29:12,970 --> 00:29:18,540
threads have to wait until none locks

686
00:29:14,770 --> 00:29:18,540
and that's all go knows

687
00:29:18,800 --> 00:29:21,160
yeah

688
00:29:23,990 --> 00:29:29,040
does it not lock all variables that are

689
00:29:26,730 --> 00:29:30,690
part of the object go doesn't know

690
00:29:29,040 --> 00:29:33,720
anything about the relationship between

691
00:29:30,690 --> 00:29:37,710
variables and locks so when you acquire

692
00:29:33,720 --> 00:29:41,280
that lock when you have code that calls

693
00:29:37,710 --> 00:29:44,220
lock exactly what it is doing it is

694
00:29:41,280 --> 00:29:47,280
acquiring this lock and that's all this

695
00:29:44,220 --> 00:29:49,260
does and anybody else who tries to lock

696
00:29:47,280 --> 00:29:55,170
objects so somewhere else who would have

697
00:29:49,260 --> 00:29:56,580
declared you know mutex knew all right

698
00:29:55,170 --> 00:29:58,410
and this mu refers to some particular

699
00:29:56,580 --> 00:30:01,350
lock object no and there me many many

700
00:29:58,410 --> 00:30:04,620
locks right all this does is acquires

701
00:30:01,350 --> 00:30:06,000
this lock and anybody else who wants to

702
00:30:04,620 --> 00:30:09,120
acquire it has to wait until we unlock

703
00:30:06,000 --> 00:30:11,970
this lock that's totally up to us as

704
00:30:09,120 --> 00:30:33,570
programmers what we were protecting with

705
00:30:11,970 --> 00:30:38,100
that lock so the question is is it

706
00:30:33,570 --> 00:30:39,660
better to have the lock be a private the

707
00:30:38,100 --> 00:30:42,890
private business of the data structure

708
00:30:39,660 --> 00:30:44,820
like supposing it a zoning map yeah and

709
00:30:42,890 --> 00:30:46,860
you know you would hope although it's

710
00:30:44,820 --> 00:30:49,860
not true that map internally would have

711
00:30:46,860 --> 00:30:54,240
a lock protecting it and that's a

712
00:30:49,860 --> 00:30:56,340
reasonable strategy would be to have I

713
00:30:54,240 --> 00:30:58,350
mean what would be to have it if you

714
00:30:56,340 --> 00:30:59,820
define a data structure that needs to be

715
00:30:58,350 --> 00:31:01,650
locked to have the lock be sort of

716
00:30:59,820 --> 00:31:03,210
interior that have each of the data

717
00:31:01,650 --> 00:31:04,860
structures methods be responsible for

718
00:31:03,210 --> 00:31:06,840
acquiring that lock and the user the

719
00:31:04,860 --> 00:31:09,210
data structure may never know that

720
00:31:06,840 --> 00:31:10,830
that's pretty reasonable and the only

721
00:31:09,210 --> 00:31:15,930
point at which that breaks down is that

722
00:31:10,830 --> 00:31:18,090
um well it's a couple things one is if

723
00:31:15,930 --> 00:31:20,910
the programmer knew that the data was

724
00:31:18,090 --> 00:31:22,200
never shared they might be bummed that

725
00:31:20,910 --> 00:31:23,400
they were paying the lock overhead for

726
00:31:22,200 --> 00:31:25,820
something they knew didn't need to be

727
00:31:23,400 --> 00:31:31,140
locked so that's one potential problem

728
00:31:25,820 --> 00:31:33,060
the other is that if you if there's any

729
00:31:31,140 --> 00:31:35,040
inter data structure of dependencies so

730
00:31:33,060 --> 00:31:36,020
we have two data structures each with

731
00:31:35,040 --> 00:31:38,710
locks and

732
00:31:36,020 --> 00:31:41,080
and they maybe use each other then

733
00:31:38,710 --> 00:31:45,160
there's a risk of cycles and deadlocks

734
00:31:41,080 --> 00:31:47,600
right and the deadlocks can be solved

735
00:31:45,160 --> 00:31:52,370
but the usual solutions to deadlocks

736
00:31:47,600 --> 00:31:54,290
requires lifting the locks out of out of

737
00:31:52,370 --> 00:31:56,000
the implementations up into the calling

738
00:31:54,290 --> 00:31:59,180
code I will talk about that some point

739
00:31:56,000 --> 00:32:00,710
but it's not a it's a good idea to hide

740
00:31:59,180 --> 00:32:11,600
the locks but it's not always a good

741
00:32:00,710 --> 00:32:14,060
idea all right okay so one problem you

742
00:32:11,600 --> 00:32:16,370
run into with threads is these races and

743
00:32:14,060 --> 00:32:18,020
generally you solve them with locks okay

744
00:32:16,370 --> 00:32:19,430
or actually there's two big strategies

745
00:32:18,020 --> 00:32:22,330
one is you figure out some locking

746
00:32:19,430 --> 00:32:25,180
strategy for making access to the data

747
00:32:22,330 --> 00:32:29,170
single thread one thread at a time or

748
00:32:25,180 --> 00:32:32,420
yury you fix your code to not share data

749
00:32:29,170 --> 00:32:35,540
if you can do that it's that's probably

750
00:32:32,420 --> 00:32:38,330
better because it's less complex all

751
00:32:35,540 --> 00:32:40,100
right so another issue that shows up

752
00:32:38,330 --> 00:32:44,780
with leads threads is called

753
00:32:40,100 --> 00:32:46,670
coordination when we're doing locking

754
00:32:44,780 --> 00:32:48,530
the different threads involved probably

755
00:32:46,670 --> 00:32:49,820
have no idea that the other ones exist

756
00:32:48,530 --> 00:32:51,650
they just want to like be able to get

757
00:32:49,820 --> 00:32:53,960
out the data without anybody else

758
00:32:51,650 --> 00:32:55,220
interfering but there are also cases

759
00:32:53,960 --> 00:32:56,870
where you need where you do

760
00:32:55,220 --> 00:32:58,370
intentionally want different threads to

761
00:32:56,870 --> 00:32:59,930
interact I want to wait for you

762
00:32:58,370 --> 00:33:01,550
maybe you're producing some data you

763
00:32:59,930 --> 00:33:03,140
know you're a different thread than me

764
00:33:01,550 --> 00:33:05,900
you're you're producing data I'm gonna

765
00:33:03,140 --> 00:33:09,980
wait until you've generated the data

766
00:33:05,900 --> 00:33:11,180
before I read it right or you launch a

767
00:33:09,980 --> 00:33:12,440
bunch of threads to say you crawl the

768
00:33:11,180 --> 00:33:14,180
web and you want to wait for all those

769
00:33:12,440 --> 00:33:16,970
fits to finish so there's times when we

770
00:33:14,180 --> 00:33:18,260
intentionally want different to us to

771
00:33:16,970 --> 00:33:18,770
interact with each other to wait for

772
00:33:18,260 --> 00:33:23,320
each other

773
00:33:18,770 --> 00:33:23,320
and that's usually called coordination

774
00:33:23,440 --> 00:33:28,250
and there's a bunch of as you probably

775
00:33:26,630 --> 00:33:31,220
know from having done the tutorial

776
00:33:28,250 --> 00:33:34,299
there's a bunch of techniques in go for

777
00:33:31,220 --> 00:33:37,029
doing this like channels

778
00:33:34,299 --> 00:33:38,950
which are really about sending data from

779
00:33:37,029 --> 00:33:42,039
one threat to another and breeding but

780
00:33:38,950 --> 00:33:45,239
they did to be sent there's also other

781
00:33:42,039 --> 00:33:48,940
stuff that more special purpose things

782
00:33:45,239 --> 00:33:51,549
like there's a idea called condition

783
00:33:48,940 --> 00:33:53,229
variables which is great if there's some

784
00:33:51,549 --> 00:33:54,489
thread out there and you want to kick it

785
00:33:53,229 --> 00:33:55,839
period you're not sure if the other

786
00:33:54,489 --> 00:33:57,459
threads even waiting for you but if it

787
00:33:55,839 --> 00:33:59,499
is waiting for you you just like to give

788
00:33:57,459 --> 00:34:01,509
it a kick so it can well know that it

789
00:33:59,499 --> 00:34:06,279
should continue whatever it's doing and

790
00:34:01,509 --> 00:34:08,260
then there's wait group which is

791
00:34:06,279 --> 00:34:10,450
particularly good for launching a a

792
00:34:08,260 --> 00:34:14,470
known number of go routines and then

793
00:34:10,450 --> 00:34:16,059
waiting for them Dolph to finish and a

794
00:34:14,470 --> 00:34:23,619
final piece of damage that comes up with

795
00:34:16,059 --> 00:34:25,929
threads deadlock the deadlock refers to

796
00:34:23,619 --> 00:34:29,220
the general problem that you sometimes

797
00:34:25,929 --> 00:34:32,829
run into where one thread

798
00:34:29,220 --> 00:34:35,859
you know thread this thread is waiting

799
00:34:32,829 --> 00:34:37,750
for thread two to produce something so

800
00:34:35,859 --> 00:34:41,169
you know it's draw an arrow to say

801
00:34:37,750 --> 00:34:42,279
thread one is waiting for thread two you

802
00:34:41,169 --> 00:34:43,990
know for example thread one may be

803
00:34:42,279 --> 00:34:46,690
waiting for thread two to release a lock

804
00:34:43,990 --> 00:34:48,279
or to send something on the channel or

805
00:34:46,690 --> 00:34:51,579
to you know decrement something in a

806
00:34:48,279 --> 00:34:55,210
wait group however unfortunately maybe T

807
00:34:51,579 --> 00:34:57,900
two is waiting for thread thread one to

808
00:34:55,210 --> 00:35:00,069
do something and this is particularly

809
00:34:57,900 --> 00:35:01,839
common in the case of locks its thread

810
00:35:00,069 --> 00:35:05,740
one acquires lock a and thread to

811
00:35:01,839 --> 00:35:07,690
acquire lock be so thread one is

812
00:35:05,740 --> 00:35:11,230
acquired lock a throw two is required

813
00:35:07,690 --> 00:35:14,589
lot B and then next thread one needs to

814
00:35:11,230 --> 00:35:15,910
lock B also that is hold two locks which

815
00:35:14,589 --> 00:35:17,289
sometimes shows up and it just so

816
00:35:15,910 --> 00:35:19,990
happens that thread two needs to hold

817
00:35:17,289 --> 00:35:21,490
block hey that's a deadlock all right at

818
00:35:19,990 --> 00:35:23,410
least grab their first lock and then

819
00:35:21,490 --> 00:35:24,520
proceed down to where they need their

820
00:35:23,410 --> 00:35:26,740
second lock and now they're waiting for

821
00:35:24,520 --> 00:35:28,779
each other forever right neither can

822
00:35:26,740 --> 00:35:33,569
proceed neither then can release the

823
00:35:28,779 --> 00:35:36,190
lock and usually just nothing happens so

824
00:35:33,569 --> 00:35:37,420
if your program just kind of grinds to a

825
00:35:36,190 --> 00:35:40,000
halt and doesn't seem to be doing

826
00:35:37,420 --> 00:35:42,960
anything but didn't crash deadlock is

827
00:35:40,000 --> 00:35:42,960
it's one thing to check

828
00:35:43,799 --> 00:35:53,859
okay all right let's look at the web

829
00:35:48,339 --> 00:36:00,430
crawler from the tutorial as an example

830
00:35:53,859 --> 00:36:03,730
of some of this threading stuff I have a

831
00:36:00,430 --> 00:36:07,630
couple of two solutions and different

832
00:36:03,730 --> 00:36:09,430
styles are really three solutions in

833
00:36:07,630 --> 00:36:10,779
different styles to allow us to talk a

834
00:36:09,430 --> 00:36:13,539
bit about the details of some of this

835
00:36:10,779 --> 00:36:16,119
thread programming so first of all you

836
00:36:13,539 --> 00:36:18,490
all probably know web crawler its job is

837
00:36:16,119 --> 00:36:20,470
you give it the URL of a page that it

838
00:36:18,490 --> 00:36:23,109
starts at and you know many web pages

839
00:36:20,470 --> 00:36:24,849
have links to other pages so what a web

840
00:36:23,109 --> 00:36:27,609
crawler is trying to do is if that's the

841
00:36:24,849 --> 00:36:29,529
first page extract all the URLs that

842
00:36:27,609 --> 00:36:32,049
were mentioned that pages links you know

843
00:36:29,529 --> 00:36:33,730
fetch the pages they point to look at

844
00:36:32,049 --> 00:36:35,740
all those pages for the ules are all

845
00:36:33,730 --> 00:36:38,380
those but all urls that they refer to

846
00:36:35,740 --> 00:36:41,230
and keep on going until it's fetched all

847
00:36:38,380 --> 00:36:45,490
the pages in the web let's just say and

848
00:36:41,230 --> 00:36:52,150
then it should stop in addition the the

849
00:36:45,490 --> 00:36:53,380
graph of pages and URLs is cyclic that

850
00:36:52,150 --> 00:36:55,420
is if you're not careful

851
00:36:53,380 --> 00:36:57,099
um you may end up following if you don't

852
00:36:55,420 --> 00:36:59,019
remember oh I've already fetched this

853
00:36:57,099 --> 00:37:01,450
web page already you may end up

854
00:36:59,019 --> 00:37:03,579
following cycles forever and you know

855
00:37:01,450 --> 00:37:05,559
your crawler will never finish so one of

856
00:37:03,579 --> 00:37:08,140
the jobs of the crawler is to remember

857
00:37:05,559 --> 00:37:10,779
the set of pages that is already crawled

858
00:37:08,140 --> 00:37:15,819
or already even started a fetch for and

859
00:37:10,779 --> 00:37:17,200
to not start a second fetch for any page

860
00:37:15,819 --> 00:37:18,849
that it's already started fetching on

861
00:37:17,200 --> 00:37:21,759
and you can think of that as sort of

862
00:37:18,849 --> 00:37:25,269
imposing a tree structure finding a sort

863
00:37:21,759 --> 00:37:31,809
of tree shaped subset of the cyclic

864
00:37:25,269 --> 00:37:33,430
graph of actual web pages okay so we

865
00:37:31,809 --> 00:37:37,059
want to avoid cycles we want to be able

866
00:37:33,430 --> 00:37:38,289
to not fetch a page twice it also it

867
00:37:37,059 --> 00:37:40,089
turns out that it just takes a long time

868
00:37:38,289 --> 00:37:42,160
to fetch a web page but it's good

869
00:37:40,089 --> 00:37:46,960
servers are slow and because the network

870
00:37:42,160 --> 00:37:48,670
has a long speed of light latency and so

871
00:37:46,960 --> 00:37:50,289
you definitely don't want to fetch pages

872
00:37:48,670 --> 00:37:54,190
one at a time unless you want to crawl

873
00:37:50,289 --> 00:37:56,030
to take many years so it pays enormous

874
00:37:54,190 --> 00:37:57,980
lead to fetch many pages that same

875
00:37:56,030 --> 00:37:59,480
I'm up to some limit right you want to

876
00:37:57,980 --> 00:38:01,490
keep on increasing the number of pages

877
00:37:59,480 --> 00:38:03,140
you fetch in parallel until the

878
00:38:01,490 --> 00:38:05,690
throughput you're getting in pages per

879
00:38:03,140 --> 00:38:07,910
second stops increasing that is running

880
00:38:05,690 --> 00:38:11,300
increase the concurrency until you run

881
00:38:07,910 --> 00:38:12,500
out of network capacity so we want to be

882
00:38:11,300 --> 00:38:15,980
able to launch multiple fetches in

883
00:38:12,500 --> 00:38:17,900
parallel and a final challenge which is

884
00:38:15,980 --> 00:38:19,400
sometimes the hardest thing to solve is

885
00:38:17,900 --> 00:38:21,620
to know when the crawl is finished

886
00:38:19,400 --> 00:38:24,110
and once we've crawled all the pages we

887
00:38:21,620 --> 00:38:25,370
want to stop and say we're done but we

888
00:38:24,110 --> 00:38:27,110
actually need to write the code to

889
00:38:25,370 --> 00:38:29,840
realize aha

890
00:38:27,110 --> 00:38:32,480
we've crawled every single page and for

891
00:38:29,840 --> 00:38:33,890
some solutions I've tried figuring out

892
00:38:32,480 --> 00:38:38,420
when you're done has turned out to be

893
00:38:33,890 --> 00:38:40,730
the hardest part all right so my first

894
00:38:38,420 --> 00:38:43,400
crawler is this serial crawler here and

895
00:38:40,730 --> 00:38:45,980
by the way this code is available on the

896
00:38:43,400 --> 00:38:48,500
website under crawler go on the schedule

897
00:38:45,980 --> 00:38:53,240
you won't look at it this wrist calls a

898
00:38:48,500 --> 00:38:58,120
serial crawler it effectively performs a

899
00:38:53,240 --> 00:39:02,870
depth-first search into the web graph

900
00:38:58,120 --> 00:39:04,490
and there is sort of one moderately

901
00:39:02,870 --> 00:39:06,650
interesting thing about it it keeps this

902
00:39:04,490 --> 00:39:08,630
map called fetched which is basically

903
00:39:06,650 --> 00:39:11,210
using as a set in order to remember

904
00:39:08,630 --> 00:39:12,830
which pages it's crawled and that's like

905
00:39:11,210 --> 00:39:16,250
the only interesting part of this you

906
00:39:12,830 --> 00:39:17,920
give it a URL that at line 18 if it's

907
00:39:16,250 --> 00:39:20,210
already fetched the URL it just returns

908
00:39:17,920 --> 00:39:22,420
if it doesn't fetch the URL it first

909
00:39:20,210 --> 00:39:26,060
remembers that it is now fetched it

910
00:39:22,420 --> 00:39:27,650
actually gets fetches that page and

911
00:39:26,060 --> 00:39:29,540
extracts the URLs that are in the page

912
00:39:27,650 --> 00:39:33,430
with the fetcher and then iterates over

913
00:39:29,540 --> 00:39:35,660
the URLs in that page and calls itself

914
00:39:33,430 --> 00:39:38,060
for every one of those pages and it

915
00:39:35,660 --> 00:39:40,010
passes to itself the way it it really

916
00:39:38,060 --> 00:39:43,070
has just a one table there's only one

917
00:39:40,010 --> 00:39:45,770
fetched map of course because you know

918
00:39:43,070 --> 00:39:47,330
when I call recursive crawl and it

919
00:39:45,770 --> 00:39:49,970
fetches a bunch of pages after it

920
00:39:47,330 --> 00:39:52,340
returns I want to be where you know the

921
00:39:49,970 --> 00:39:53,960
outer crawl instance needs to be aware

922
00:39:52,340 --> 00:39:56,330
that certain pages are already fetched

923
00:39:53,960 --> 00:39:58,250
so we depend very much on the fetched

924
00:39:56,330 --> 00:40:01,970
map being passed between the functions

925
00:39:58,250 --> 00:40:03,770
by reference instead of by copying so it

926
00:40:01,970 --> 00:40:05,330
so under the hood what must really be

927
00:40:03,770 --> 00:40:08,480
going on here is that go is passing a

928
00:40:05,330 --> 00:40:10,070
pointer to the map object

929
00:40:08,480 --> 00:40:12,830
to each of the calls of crawl so they

930
00:40:10,070 --> 00:40:15,740
all share the pointer to the same object

931
00:40:12,830 --> 00:40:22,760
and memory rather than copying rather

932
00:40:15,740 --> 00:40:24,140
than copying than that any questions so

933
00:40:22,760 --> 00:40:25,760
this code definitely does not solve the

934
00:40:24,140 --> 00:40:30,650
problem that was posed right because it

935
00:40:25,760 --> 00:40:33,200
doesn't launch parallel parallel fetches

936
00:40:30,650 --> 00:40:35,150
now so clue we need to insert goroutines

937
00:40:33,200 --> 00:40:37,880
somewhere in this code right to get

938
00:40:35,150 --> 00:40:41,480
parallel fetches so let's suppose just

939
00:40:37,880 --> 00:40:51,440
for chuckles dad we just start with the

940
00:40:41,480 --> 00:40:54,950
most lazy thing because why so I'm gonna

941
00:40:51,440 --> 00:40:57,470
just modify the code to run the

942
00:40:54,950 --> 00:41:00,470
subsidiary crawls each in its own go

943
00:40:57,470 --> 00:41:01,580
routine actually before I do that why

944
00:41:00,470 --> 00:41:04,070
don't I run the code just to show you

945
00:41:01,580 --> 00:41:07,670
what correct output looks like so hoping

946
00:41:04,070 --> 00:41:09,380
this other window Emad run the crawler

947
00:41:07,670 --> 00:41:10,970
it actually runs all three copies of the

948
00:41:09,380 --> 00:41:14,330
crawler and they all find exactly the

949
00:41:10,970 --> 00:41:16,100
same set of webpages so this is the

950
00:41:14,330 --> 00:41:19,220
output that we're hoping to see five

951
00:41:16,100 --> 00:41:20,750
lines five different web pages are are

952
00:41:19,220 --> 00:41:26,120
fetched prints a line for each one so

953
00:41:20,750 --> 00:41:28,130
let me now run the subsidiary crawls in

954
00:41:26,120 --> 00:41:35,540
their own go routines and run that code

955
00:41:28,130 --> 00:41:37,880
so what am I going to see the hope is to

956
00:41:35,540 --> 00:41:42,800
fetch these webpages in parallel for

957
00:41:37,880 --> 00:41:45,110
higher performance so okay so you're

958
00:41:42,800 --> 00:41:47,680
voting for only seeing one URL and why

959
00:41:45,110 --> 00:41:47,680
so why is that

960
00:41:50,980 --> 00:41:59,000
yeah yes that's exactly right you know

961
00:41:55,220 --> 00:42:00,559
after the after it's not gonna wait in

962
00:41:59,000 --> 00:42:02,450
this loop at line 26 it's gonna zip

963
00:42:00,559 --> 00:42:04,849
right through that loop I was gonna

964
00:42:02,450 --> 00:42:07,039
fetch 1p when the ferry first webpage at

965
00:42:04,849 --> 00:42:08,390
line 22 and then a loop it's gonna fly

966
00:42:07,039 --> 00:42:10,220
off the girl routines and immediately

967
00:42:08,390 --> 00:42:11,990
the scroll function is gonna return and

968
00:42:10,220 --> 00:42:13,789
if it was called from main main what was

969
00:42:11,990 --> 00:42:15,200
exit almost certainly before any of the

970
00:42:13,789 --> 00:42:16,880
routines was able to do any work at all

971
00:42:15,200 --> 00:42:19,690
so we'll probably just see the first web

972
00:42:16,880 --> 00:42:23,920
page and I'm gonna do when I run it

973
00:42:19,690 --> 00:42:26,660
you'll see here under serial that only

974
00:42:23,920 --> 00:42:28,730
the one web page was found now in fact

975
00:42:26,660 --> 00:42:30,799
since this program doesn't exit after

976
00:42:28,730 --> 00:42:32,269
the serial crawler those Guru T's are

977
00:42:30,799 --> 00:42:35,390
still running and they actually print

978
00:42:32,269 --> 00:42:37,819
their output down here interleaved with

979
00:42:35,390 --> 00:42:42,829
the next crawler example but

980
00:42:37,819 --> 00:42:45,829
nevertheless the codes just adding a go

981
00:42:42,829 --> 00:42:49,579
here absolutely doesn't work so let's

982
00:42:45,829 --> 00:42:52,190
get rid of that okay so now I want to

983
00:42:49,579 --> 00:42:55,789
show you a one style of concurrent

984
00:42:52,190 --> 00:42:59,750
crawler and I'm presenting to one of

985
00:42:55,789 --> 00:43:02,809
them written with shared data shared

986
00:42:59,750 --> 00:43:05,059
objects and locks it's the first one and

987
00:43:02,809 --> 00:43:08,990
another one written without shared data

988
00:43:05,059 --> 00:43:11,359
but with passing information along

989
00:43:08,990 --> 00:43:12,920
channels in order to coordinate the

990
00:43:11,359 --> 00:43:17,000
different threads so this is the shared

991
00:43:12,920 --> 00:43:18,980
data one or this is just one of many

992
00:43:17,000 --> 00:43:22,460
ways of building a web crawler using

993
00:43:18,980 --> 00:43:26,079
shared data so this code significantly

994
00:43:22,460 --> 00:43:31,130
more complicated than a serial crawler

995
00:43:26,079 --> 00:43:33,740
it creates a thread for each fetch it

996
00:43:31,130 --> 00:43:38,390
does alright but the huge difference is

997
00:43:33,740 --> 00:43:40,700
that it does with two things one it does

998
00:43:38,390 --> 00:43:44,690
the bookkeeping required to notice when

999
00:43:40,700 --> 00:43:47,900
all of the crawls have finished and it

1000
00:43:44,690 --> 00:43:49,849
handles the shared table of which URLs

1001
00:43:47,900 --> 00:43:53,809
have been crawled correctly so this code

1002
00:43:49,849 --> 00:43:59,349
still has this table of URLs and that's

1003
00:43:53,809 --> 00:44:06,130
this F dot fetched this F dot fetch

1004
00:43:59,349 --> 00:44:10,660
map at line 43 but this this table is

1005
00:44:06,130 --> 00:44:12,190
actually shared by all of the all of the

1006
00:44:10,660 --> 00:44:14,890
crawler threads and all the collar

1007
00:44:12,190 --> 00:44:16,900
threads are making or executing inside

1008
00:44:14,890 --> 00:44:18,609
concurrent mutex and so we still have

1009
00:44:16,900 --> 00:44:20,619
this sort of tree up in current mutexes

1010
00:44:18,609 --> 00:44:22,599
that's exploring different parts of the

1011
00:44:20,619 --> 00:44:25,660
web graph but each one of them was

1012
00:44:22,599 --> 00:44:28,720
launched as a as his own go routine

1013
00:44:25,660 --> 00:44:30,400
instead of as a function call but

1014
00:44:28,720 --> 00:44:32,859
they're all sharing this table of state

1015
00:44:30,400 --> 00:44:34,990
this table of test URLs because if one

1016
00:44:32,859 --> 00:44:36,970
go routine fetches a URL we don't want

1017
00:44:34,990 --> 00:44:40,420
another girl routine to accidentally

1018
00:44:36,970 --> 00:44:43,150
fetch the same URL and as you can see

1019
00:44:40,420 --> 00:44:48,250
here line 42 and 45 I've surrounded them

1020
00:44:43,150 --> 00:44:51,099
by the new taxes that are required to to

1021
00:44:48,250 --> 00:44:52,900
prevent a race that would occur if I

1022
00:44:51,099 --> 00:44:57,730
didn't add them new Texas so the danger

1023
00:44:52,900 --> 00:44:59,980
here is that at line 43 a thread is

1024
00:44:57,730 --> 00:45:02,680
checking of URLs already been fetched so

1025
00:44:59,980 --> 00:45:06,819
two threads happen to be following the

1026
00:45:02,680 --> 00:45:09,490
same URL now two calls to concurrent

1027
00:45:06,819 --> 00:45:11,140
mutex end up looking at the same URL

1028
00:45:09,490 --> 00:45:13,930
maybe because that URL was mentioned in

1029
00:45:11,140 --> 00:45:17,559
two different web pages if we didn't

1030
00:45:13,930 --> 00:45:18,819
have the lock they'd both access the

1031
00:45:17,559 --> 00:45:20,829
math table to see if the threaded and

1032
00:45:18,819 --> 00:45:23,650
then already if the URL had been already

1033
00:45:20,829 --> 00:45:27,069
fetched and they both get false at line

1034
00:45:23,650 --> 00:45:30,880
43 they both set the URLs entering the

1035
00:45:27,069 --> 00:45:32,380
table to true at line 44 and at 47 they

1036
00:45:30,880 --> 00:45:33,880
will both see that I already was false

1037
00:45:32,380 --> 00:45:37,030
and then they both go on to patch the

1038
00:45:33,880 --> 00:45:38,740
web page so we need the lock there and

1039
00:45:37,030 --> 00:45:41,410
the way to think about it I think is

1040
00:45:38,740 --> 00:45:44,020
that we want lines 43 and 44 to be

1041
00:45:41,410 --> 00:45:45,910
atomic that is we don't want some other

1042
00:45:44,020 --> 00:45:48,460
thread to to get in and be using the

1043
00:45:45,910 --> 00:45:50,349
table between 43 and 44 we we want to

1044
00:45:48,460 --> 00:45:52,690
read the current content each thread

1045
00:45:50,349 --> 00:45:55,780
wants to read the current table contents

1046
00:45:52,690 --> 00:45:57,309
and update it without any other thread

1047
00:45:55,780 --> 00:46:01,150
interfering and so that's what the locks

1048
00:45:57,309 --> 00:46:03,280
are doing for us okay so so actually any

1049
00:46:01,150 --> 00:46:05,940
questions about the about the locking

1050
00:46:03,280 --> 00:46:05,940
strategy here

1051
00:46:07,750 --> 00:46:13,670
all right once we check the URLs entry

1052
00:46:10,760 --> 00:46:15,320
in the table alliant 51 it just crawls

1053
00:46:13,670 --> 00:46:18,950
it just fetches that page in the usual

1054
00:46:15,320 --> 00:46:20,600
way and then the other thing interesting

1055
00:46:18,950 --> 00:46:35,450
thing that's going on is the launching

1056
00:46:20,600 --> 00:46:43,970
of the threads yes so the question is

1057
00:46:35,450 --> 00:46:47,120
what's with the F dot no no the MU it is

1058
00:46:43,970 --> 00:46:50,330
okay so there's a structure to find out

1059
00:46:47,120 --> 00:46:53,930
line 36 that sort of collects together

1060
00:46:50,330 --> 00:46:55,280
all the different stuff that all the

1061
00:46:53,930 --> 00:46:57,380
different state that we need to run this

1062
00:46:55,280 --> 00:46:58,820
crawl and here it's only two objects but

1063
00:46:57,380 --> 00:47:00,620
you know it could be a lot more and

1064
00:46:58,820 --> 00:47:02,030
they're only grouped together for

1065
00:47:00,620 --> 00:47:05,390
convenience there's no other

1066
00:47:02,030 --> 00:47:07,490
significance to the fact there's no deep

1067
00:47:05,390 --> 00:47:11,750
significance the fact that mu and fetch

1068
00:47:07,490 --> 00:47:14,690
store it inside the same structure and

1069
00:47:11,750 --> 00:47:15,890
that F dot is just sort of the syntax

1070
00:47:14,690 --> 00:47:17,180
are getting out one of the elements in

1071
00:47:15,890 --> 00:47:19,070
the structure so I just happened to put

1072
00:47:17,180 --> 00:47:21,080
them you in the structure because it

1073
00:47:19,070 --> 00:47:22,790
allows me to group together all the

1074
00:47:21,080 --> 00:47:25,600
stuff related to a crawl but that

1075
00:47:22,790 --> 00:47:28,880
absolutely does not mean that go

1076
00:47:25,600 --> 00:47:30,950
associates the MU with that structure or

1077
00:47:28,880 --> 00:47:33,710
with the fetch map or anything it's just

1078
00:47:30,950 --> 00:47:35,090
a lock objects and just has a lock

1079
00:47:33,710 --> 00:47:37,930
function you can call and that's all

1080
00:47:35,090 --> 00:47:37,930
that's going on

1081
00:47:53,790 --> 00:48:00,520
so the question is how come in order to

1082
00:47:58,750 --> 00:48:02,440
pass something by reference I had to use

1083
00:48:00,520 --> 00:48:03,940
star here where it is when a in the

1084
00:48:02,440 --> 00:48:06,040
previous example when we were passing a

1085
00:48:03,940 --> 00:48:07,569
map we didn't have to use star that is

1086
00:48:06,040 --> 00:48:09,069
didn't have to pass a pointer I mean

1087
00:48:07,569 --> 00:48:15,339
that star notation you're seeing there

1088
00:48:09,069 --> 00:48:16,809
in mine 41 basically and he's saying

1089
00:48:15,339 --> 00:48:19,210
that we're passing a pointer to this

1090
00:48:16,809 --> 00:48:20,559
fetch state object and we want it to be

1091
00:48:19,210 --> 00:48:22,000
a pointer because we want there to be

1092
00:48:20,559 --> 00:48:23,710
one object in memory and all the

1093
00:48:22,000 --> 00:48:25,240
different go routines I want to use that

1094
00:48:23,710 --> 00:48:28,000
same object so they all need a pointer

1095
00:48:25,240 --> 00:48:29,410
to that same object so so we need to

1096
00:48:28,000 --> 00:48:30,940
find your own structure that's sort of

1097
00:48:29,410 --> 00:48:32,530
the syntax you use for passing a pointer

1098
00:48:30,940 --> 00:48:35,920
the reason why we didn't have to do it

1099
00:48:32,530 --> 00:48:39,000
with map is because although it's not

1100
00:48:35,920 --> 00:48:42,579
clear from the syntax a map is a pointer

1101
00:48:39,000 --> 00:48:45,069
it's just because it's built into the

1102
00:48:42,579 --> 00:48:50,530
language they don't make you put a star

1103
00:48:45,069 --> 00:48:52,420
there but what a map is is if you

1104
00:48:50,530 --> 00:48:55,319
declare a variable type map what that is

1105
00:48:52,420 --> 00:48:57,579
is a pointer to some data in the heap so

1106
00:48:55,319 --> 00:48:59,260
it was a pointer anyway and it's always

1107
00:48:57,579 --> 00:49:00,609
passed by reference do they you just

1108
00:48:59,260 --> 00:49:01,210
don't have to put the star and it does

1109
00:49:00,609 --> 00:49:03,609
it for you

1110
00:49:01,210 --> 00:49:06,130
so there's they're definitely map is

1111
00:49:03,609 --> 00:49:07,900
special you cannot define map in the

1112
00:49:06,130 --> 00:49:09,430
language it's it has to be built in

1113
00:49:07,900 --> 00:49:15,819
because there's some curious things

1114
00:49:09,430 --> 00:49:18,849
about it okay good okay so we fetch the

1115
00:49:15,819 --> 00:49:20,799
page now we want to fire off a crawl go

1116
00:49:18,849 --> 00:49:23,170
routine for each URL mentioned in the

1117
00:49:20,799 --> 00:49:26,440
page we just fetch so that's done in

1118
00:49:23,170 --> 00:49:29,890
line 56 on line 50 sisters loops over

1119
00:49:26,440 --> 00:49:32,950
the URLs that the fetch function

1120
00:49:29,890 --> 00:49:35,740
returned and for each one fires off a go

1121
00:49:32,950 --> 00:49:41,530
routine at line 58 and that lines that

1122
00:49:35,740 --> 00:49:43,990
func syntax in line 58 is a closure or a

1123
00:49:41,530 --> 00:49:46,599
sort of immediate function but that func

1124
00:49:43,990 --> 00:49:49,140
thing keyword is doing is to clearing a

1125
00:49:46,599 --> 00:49:53,280
function right there that we then call

1126
00:49:49,140 --> 00:49:53,280
so the way to read it maybe is

1127
00:49:53,740 --> 00:50:00,230
that if you can declare a function as a

1128
00:49:56,780 --> 00:50:03,349
piece of data as just func you know and

1129
00:50:00,230 --> 00:50:08,930
then you give the arguments and then you

1130
00:50:03,349 --> 00:50:12,500
give the body and that's a clears and so

1131
00:50:08,930 --> 00:50:14,000
this is an object now this is like it's

1132
00:50:12,500 --> 00:50:18,349
like when you type one when you have a

1133
00:50:14,000 --> 00:50:19,819
one or 23 or something you're declaring

1134
00:50:18,349 --> 00:50:21,079
a sort of constant object and this is

1135
00:50:19,819 --> 00:50:24,079
the way to define a constant function

1136
00:50:21,079 --> 00:50:25,730
and we do it here because we want to

1137
00:50:24,079 --> 00:50:27,290
launch a go routine that's gonna run

1138
00:50:25,730 --> 00:50:29,119
this function that we declared right

1139
00:50:27,290 --> 00:50:31,069
here and so we in order to make the go

1140
00:50:29,119 --> 00:50:33,140
routine we have to add a go in front to

1141
00:50:31,069 --> 00:50:35,059
say we want to go routine and then we

1142
00:50:33,140 --> 00:50:37,520
have to call the function because the go

1143
00:50:35,059 --> 00:50:39,140
syntax says the syntax of the go

1144
00:50:37,520 --> 00:50:40,910
keywords as you follow it by a function

1145
00:50:39,140 --> 00:50:43,460
name and arguments you want to pass that

1146
00:50:40,910 --> 00:50:50,900
function and so we're gonna pass some

1147
00:50:43,460 --> 00:50:52,670
arguments here and there's two reasons

1148
00:50:50,900 --> 00:50:55,069
we're doing this well really this one

1149
00:50:52,670 --> 00:50:57,790
reason we you know in some other

1150
00:50:55,069 --> 00:51:00,230
circumstance we could have just said go

1151
00:50:57,790 --> 00:51:01,400
concurrent mutex oh I concur mutex is

1152
00:51:00,230 --> 00:51:06,619
the name of the function we actually

1153
00:51:01,400 --> 00:51:08,119
want to call with this URL but we want

1154
00:51:06,619 --> 00:51:10,160
to do a few other things as well so we

1155
00:51:08,119 --> 00:51:12,170
define this little helper function that

1156
00:51:10,160 --> 00:51:15,740
first calls concurrent mutex for us with

1157
00:51:12,170 --> 00:51:17,119
the URL and then after them current

1158
00:51:15,740 --> 00:51:19,520
mutex is finished we do something

1159
00:51:17,119 --> 00:51:22,069
special in order to help us wait for all

1160
00:51:19,520 --> 00:51:24,920
the crawls to be done before the outer

1161
00:51:22,069 --> 00:51:27,380
function returns so that brings us to

1162
00:51:24,920 --> 00:51:29,569
the the weight group the weight group at

1163
00:51:27,380 --> 00:51:33,619
line 55 it's a just a data structure to

1164
00:51:29,569 --> 00:51:35,030
find by go to help with coordination and

1165
00:51:33,619 --> 00:51:39,290
the game with weight group is that

1166
00:51:35,030 --> 00:51:43,640
internally it has a counter and you call

1167
00:51:39,290 --> 00:51:46,549
weight group dot add like a line 57 to

1168
00:51:43,640 --> 00:51:48,619
increment the counter and we group done

1169
00:51:46,549 --> 00:51:50,900
to decrement it and then this weight

1170
00:51:48,619 --> 00:51:53,119
what this weight method called line 63

1171
00:51:50,900 --> 00:51:56,510
waits for the counter to get down to

1172
00:51:53,119 --> 00:51:59,329
zero so a weight group is a way to wait

1173
00:51:56,510 --> 00:52:02,540
for a specific number of things to

1174
00:51:59,329 --> 00:52:04,010
finish and it's useful in a bunch of

1175
00:52:02,540 --> 00:52:05,359
different situations here we're using it

1176
00:52:04,010 --> 00:52:05,920
to wait for the last go routine to

1177
00:52:05,359 --> 00:52:07,839
finish

1178
00:52:05,920 --> 00:52:11,200
because we add one to the weight group

1179
00:52:07,839 --> 00:52:13,119
for every go routine we create line 60

1180
00:52:11,200 --> 00:52:15,310
at the end of this function we've

1181
00:52:13,119 --> 00:52:18,130
declared decrement the counter in the

1182
00:52:15,310 --> 00:52:20,250
weight group and then line three weights

1183
00:52:18,130 --> 00:52:22,300
until all the decrements have finished

1184
00:52:20,250 --> 00:52:23,920
and so the reason why we declared this

1185
00:52:22,300 --> 00:52:26,530
little function was basically to be able

1186
00:52:23,920 --> 00:52:28,630
to both call concurrently text and call

1187
00:52:26,530 --> 00:52:39,760
dot that's really why we needed that

1188
00:52:28,630 --> 00:52:43,240
function so the question is what if one

1189
00:52:39,760 --> 00:52:45,820
of the subroutines fails and doesn't

1190
00:52:43,240 --> 00:52:49,210
reach the done line that's a darn good

1191
00:52:45,820 --> 00:52:51,070
question there is you know if I forget

1192
00:52:49,210 --> 00:52:53,440
the exact range of errors that will

1193
00:52:51,070 --> 00:52:55,150
cause the go routine to fail without

1194
00:52:53,440 --> 00:52:56,589
causing the program to feel maybe

1195
00:52:55,150 --> 00:52:57,790
divides by zero I don't know where

1196
00:52:56,589 --> 00:52:59,140
dereference is a nil pointer

1197
00:52:57,790 --> 00:53:04,570
not sure but there are certainly ways

1198
00:52:59,140 --> 00:53:06,910
for a function to fail and I have the go

1199
00:53:04,570 --> 00:53:08,890
routine die without having the program

1200
00:53:06,910 --> 00:53:12,130
die and that would be a problem for us

1201
00:53:08,890 --> 00:53:13,660
and so really the white right way to I'm

1202
00:53:12,130 --> 00:53:15,849
sure you had this in mind and asking the

1203
00:53:13,660 --> 00:53:18,520
question the right way to write this to

1204
00:53:15,849 --> 00:53:20,740
be sure that the done call is made no

1205
00:53:18,520 --> 00:53:27,180
matter why this guru team is finishing

1206
00:53:20,740 --> 00:53:31,540
would be to put a defer here which means

1207
00:53:27,180 --> 00:53:34,330
call done before the surrounding

1208
00:53:31,540 --> 00:53:36,130
function finishes and always call it no

1209
00:53:34,330 --> 00:53:42,119
matter why the surrounding function is

1210
00:53:36,130 --> 00:53:42,119
finished yes

1211
00:53:53,559 --> 00:54:00,650
and yes yeah so the question is how come

1212
00:53:58,789 --> 00:54:08,210
two users have done in different threads

1213
00:54:00,650 --> 00:54:10,640
aren't a race yeah so the answer must be

1214
00:54:08,210 --> 00:54:14,170
that internally dot a weight group has a

1215
00:54:10,640 --> 00:54:18,200
mutex or something like it that each of

1216
00:54:14,170 --> 00:54:19,970
Dunn's methods acquires before doing

1217
00:54:18,200 --> 00:54:22,789
anything else so that simultaneously

1218
00:54:19,970 --> 00:54:32,170
calls to a done to await groups methods

1219
00:54:22,789 --> 00:54:32,170
are trees we could to did a low class

1220
00:54:39,519 --> 00:54:45,440
yeah for certain leaf C++ and in C you

1221
00:54:43,880 --> 00:54:47,390
want to look at something called P

1222
00:54:45,440 --> 00:54:48,650
threads for C threads come in a library

1223
00:54:47,390 --> 00:54:51,710
they're not really part of the language

1224
00:54:48,650 --> 00:54:55,420
called P threads which they have these

1225
00:54:51,710 --> 00:55:04,450
are extremely traditional and ancient

1226
00:54:55,420 --> 00:55:04,450
primitives that all languages yeah

1227
00:55:06,630 --> 00:55:14,140
say it again you know not in this code

1228
00:55:12,220 --> 00:55:15,250
but you know you could imagine a use of

1229
00:55:14,140 --> 00:55:21,250
weight groups I mean weight groups just

1230
00:55:15,250 --> 00:55:22,990
count stuff and yeah yeah yeah weight

1231
00:55:21,250 --> 00:55:27,370
group doesn't really care what you're

1232
00:55:22,990 --> 00:55:45,850
pounding or why I mean you know this is

1233
00:55:27,370 --> 00:55:48,550
the most common way to see it use you're

1234
00:55:45,850 --> 00:55:54,780
wondering why you is passed as a

1235
00:55:48,550 --> 00:55:59,070
parameter to the function at 58 okay

1236
00:55:54,780 --> 00:56:01,450
yeah this is alright so the question is

1237
00:55:59,070 --> 00:56:05,890
okay so actually backing up a little bit

1238
00:56:01,450 --> 00:56:09,010
the rules for these for a function like

1239
00:56:05,890 --> 00:56:10,960
the one I'm defining on 58 is that if

1240
00:56:09,010 --> 00:56:14,050
the function body mentions a variable

1241
00:56:10,960 --> 00:56:17,470
that's declared in the outer function

1242
00:56:14,050 --> 00:56:19,000
but not shadowed then the the inner

1243
00:56:17,470 --> 00:56:20,650
functions use of that is the same

1244
00:56:19,000 --> 00:56:23,080
variable in the inner function as in the

1245
00:56:20,650 --> 00:56:26,380
outer function and so that's what's

1246
00:56:23,080 --> 00:56:28,780
happening with Fechter for example like

1247
00:56:26,380 --> 00:56:30,250
what is this variable here refer to what

1248
00:56:28,780 --> 00:56:32,980
does the Fechter variable refer to in

1249
00:56:30,250 --> 00:56:35,290
the inner function well it refers it's

1250
00:56:32,980 --> 00:56:37,480
the same variable as as the fetcher in

1251
00:56:35,290 --> 00:56:38,920
the outer function says just is that

1252
00:56:37,480 --> 00:56:40,510
variable and so when the inner function

1253
00:56:38,920 --> 00:56:42,310
refers to fetcher it just means it's

1254
00:56:40,510 --> 00:56:45,670
just referring the same variable as this

1255
00:56:42,310 --> 00:56:48,160
one here and the same with F f is it's

1256
00:56:45,670 --> 00:56:50,320
used here it's just is this variable so

1257
00:56:48,160 --> 00:56:55,990
you might think that we could get rid of

1258
00:56:50,320 --> 00:56:57,880
the this u argument that we're passing

1259
00:56:55,990 --> 00:56:59,860
and just have the inner function take no

1260
00:56:57,880 --> 00:57:04,530
arguments at all but just use the U that

1261
00:56:59,860 --> 00:57:04,530
was defined up on line 56 in the loop

1262
00:57:05,070 --> 00:57:09,910
and it'll be nice if we could do that

1263
00:57:07,390 --> 00:57:12,550
because save us some typing it turns out

1264
00:57:09,910 --> 00:57:16,060
not to work and the reason is that the

1265
00:57:12,550 --> 00:57:17,410
semantics of go of the for loop at line

1266
00:57:16,060 --> 00:57:21,850
56 is that the

1267
00:57:17,410 --> 00:57:23,620
for the updates the variable you so in

1268
00:57:21,850 --> 00:57:29,380
the first iteration of the for loop that

1269
00:57:23,620 --> 00:57:31,510
variable u contains some URL and when

1270
00:57:29,380 --> 00:57:34,150
you enter the second iteration before

1271
00:57:31,510 --> 00:57:37,750
the that variable this contents are

1272
00:57:34,150 --> 00:57:39,370
changed to be the second URL and that

1273
00:57:37,750 --> 00:57:41,530
means that the first go routine that we

1274
00:57:39,370 --> 00:57:43,000
launched that's just looking at the

1275
00:57:41,530 --> 00:57:46,930
outer if it we're looking at the outer

1276
00:57:43,000 --> 00:57:48,910
functions u variable the that first go

1277
00:57:46,930 --> 00:57:51,400
team we launched would see a different

1278
00:57:48,910 --> 00:57:53,800
value in the u variable after the outer

1279
00:57:51,400 --> 00:57:55,150
function it updated it and sometimes

1280
00:57:53,800 --> 00:57:58,000
that's actually what you want so for

1281
00:57:55,150 --> 00:58:01,600
example for for F and then particular F

1282
00:57:58,000 --> 00:58:04,960
dot fetched we interaction absolutely

1283
00:58:01,600 --> 00:58:06,490
wants to see changes to that map but for

1284
00:58:04,960 --> 00:58:09,070
you we don't want to see changes the

1285
00:58:06,490 --> 00:58:12,130
first go routine we spawn should read

1286
00:58:09,070 --> 00:58:13,810
the first URL not the second URL so we

1287
00:58:12,130 --> 00:58:16,080
want that go routine to have a copy you

1288
00:58:13,810 --> 00:58:18,370
have its own private copy of the URL and

1289
00:58:16,080 --> 00:58:20,560
you know is we could have done it in

1290
00:58:18,370 --> 00:58:22,150
other ways we could have but the way

1291
00:58:20,560 --> 00:58:25,630
this code happens to do it to produce

1292
00:58:22,150 --> 00:58:31,860
the copy private to that inner function

1293
00:58:25,630 --> 00:58:31,860
is by passing the URLs in argument yes

1294
00:58:34,450 --> 00:58:51,200
yeah if we have passed the address of

1295
00:58:36,890 --> 00:58:52,190
you yeah then it uh it's actually I

1296
00:58:51,200 --> 00:58:54,049
don't know how strings work but it is

1297
00:58:52,190 --> 00:59:00,140
absolutely giving you your own private

1298
00:58:54,049 --> 00:59:08,950
copy of the variable you get your own

1299
00:59:00,140 --> 00:59:08,950
copy of the variable and it yeah

1300
00:59:26,500 --> 00:59:33,860
are you saying we don't need to play

1301
00:59:28,850 --> 00:59:35,270
this trick in the code we definitely

1302
00:59:33,860 --> 00:59:37,700
need to play this trick in the code and

1303
00:59:35,270 --> 00:59:39,170
what's going on is this it's so the

1304
00:59:37,700 --> 00:59:41,960
question is Oh strings are immutable

1305
00:59:39,170 --> 00:59:43,850
strings are immutable right yeah so how

1306
00:59:41,960 --> 00:59:45,110
kind of strings are immutable how can

1307
00:59:43,850 --> 00:59:47,720
the outer function change the string

1308
00:59:45,110 --> 00:59:49,700
there should be no problem the problem

1309
00:59:47,720 --> 00:59:51,500
is not that the string is changed the

1310
00:59:49,700 --> 00:59:56,150
problem is that the variable U is

1311
00:59:51,500 --> 00:59:57,860
changed so the when the inner function

1312
00:59:56,150 --> 00:59:59,000
mentions a variable that's defined in

1313
00:59:57,860 --> 01:00:01,070
the outer function it's referring to

1314
00:59:59,000 --> 01:00:03,320
that variable and the variables current

1315
01:00:01,070 --> 01:00:06,590
value so when you if you have a string

1316
01:00:03,320 --> 01:00:09,110
variable that has has a in it and then

1317
01:00:06,590 --> 01:00:10,520
you assign B to that string variable

1318
01:00:09,110 --> 01:00:12,530
you're not over writing the string

1319
01:00:10,520 --> 01:00:15,950
you're changing the variable to point to

1320
01:00:12,530 --> 01:00:18,680
a different string and and because the

1321
01:00:15,950 --> 01:00:21,290
for loop changes the U variable to point

1322
01:00:18,680 --> 01:00:22,970
to a different string you know that

1323
01:00:21,290 --> 01:00:24,680
change to you would be visible inside

1324
01:00:22,970 --> 01:00:26,390
the inner function and therefore the

1325
01:00:24,680 --> 01:00:29,260
inner function needs its own copy of the

1326
01:00:26,390 --> 01:00:29,260
variable

1327
01:00:36,150 --> 01:00:42,000
essentially make a copy of that so that

1328
01:00:50,250 --> 01:00:54,670
okay but that is what we're doing in

1329
01:00:53,110 --> 01:00:56,440
this code and that's that is why this

1330
01:00:54,670 --> 01:00:59,080
code works okay

1331
01:00:56,440 --> 01:01:00,400
the proposal or the broken code that

1332
01:00:59,080 --> 01:01:02,850
we're not using here I will show you the

1333
01:01:00,400 --> 01:01:02,850
broken code

1334
01:01:44,060 --> 01:01:47,700
this is just like a horrible detail but

1335
01:01:46,440 --> 01:01:50,850
it is unfortunately one that you'll run

1336
01:01:47,700 --> 01:01:52,200
into while doing the labs so you should

1337
01:01:50,850 --> 01:01:54,690
be at least where that there's a problem

1338
01:01:52,200 --> 01:02:12,170
and when you run into it maybe you can

1339
01:01:54,690 --> 01:02:15,870
try to figure out the details okay

1340
01:02:12,170 --> 01:02:18,090
that's a great question so so the

1341
01:02:15,870 --> 01:02:19,770
question is you know if you have an

1342
01:02:18,090 --> 01:02:21,180
inner function just a repeated if you

1343
01:02:19,770 --> 01:02:23,430
have an inner function that refers to a

1344
01:02:21,180 --> 01:02:25,980
variable in the surrounding function but

1345
01:02:23,430 --> 01:02:28,590
the surrounding function returns what is

1346
01:02:25,980 --> 01:02:30,030
the inner functions variable referring

1347
01:02:28,590 --> 01:02:32,460
to anymore since the outer function is

1348
01:02:30,030 --> 01:02:35,190
as returned and the answer is that go

1349
01:02:32,460 --> 01:02:38,160
notices go analyzes your inner functions

1350
01:02:35,190 --> 01:02:39,870
or these are called closures go analyzes

1351
01:02:38,160 --> 01:02:41,580
them the compiler analyze them says aha

1352
01:02:39,870 --> 01:02:42,570
oh this disclosure this inner function

1353
01:02:41,580 --> 01:02:44,310
is using a variable in the outer

1354
01:02:42,570 --> 01:02:47,580
function we're actually gonna and the

1355
01:02:44,310 --> 01:02:50,670
compiler will allocate heat memory to

1356
01:02:47,580 --> 01:02:52,590
hold the variable the you know the

1357
01:02:50,670 --> 01:02:55,230
current value of the variable and both

1358
01:02:52,590 --> 01:02:58,350
functions will refer to that that little

1359
01:02:55,230 --> 01:02:59,760
area heap that has the barrel so it

1360
01:02:58,350 --> 01:03:01,590
won't be allocated the variable won't be

1361
01:02:59,760 --> 01:03:03,180
on the stack as you might expect it's

1362
01:03:01,590 --> 01:03:04,980
moved to the heap if if the compiler

1363
01:03:03,180 --> 01:03:06,060
sees that it's using a closure and then

1364
01:03:04,980 --> 01:03:07,950
when the outer function returns the

1365
01:03:06,060 --> 01:03:09,840
object is still there in the heap the

1366
01:03:07,950 --> 01:03:11,820
inner function can still get at it and

1367
01:03:09,840 --> 01:03:13,440
then the garbage collector is

1368
01:03:11,820 --> 01:03:15,540
responsible for noticing that the last

1369
01:03:13,440 --> 01:03:18,540
function to refer to this little piece

1370
01:03:15,540 --> 01:03:24,769
of heat that's exited returned and to

1371
01:03:18,540 --> 01:03:29,309
free it only then okay okay

1372
01:03:24,769 --> 01:03:30,629
okay so wait group wait group is maybe

1373
01:03:29,309 --> 01:03:32,549
the more important thing here that the

1374
01:03:30,629 --> 01:03:35,719
technique that this code uses to wait

1375
01:03:32,549 --> 01:03:37,739
for all the all this level of crawls to

1376
01:03:35,719 --> 01:03:39,599
finished all its direct chill and the

1377
01:03:37,739 --> 01:03:41,279
finish is the wait group of course

1378
01:03:39,599 --> 01:03:44,909
there's many of these wait groups one

1379
01:03:41,279 --> 01:03:46,289
per call two concurrent mutex each call

1380
01:03:44,909 --> 01:03:49,519
that concurrent mutex just waits for its

1381
01:03:46,289 --> 01:03:53,609
own children to finish and then returns

1382
01:03:49,519 --> 01:03:54,479
okay so back to the lock actually

1383
01:03:53,609 --> 01:03:56,279
there's one more thing I want to talk

1384
01:03:54,479 --> 01:03:57,949
about with a lock and that is to explore

1385
01:03:56,279 --> 01:04:00,689
what would happen if we hadn't locked

1386
01:03:57,949 --> 01:04:02,369
right I'm claiming oh you know you don't

1387
01:04:00,689 --> 01:04:05,209
lock you're gonna get these races you're

1388
01:04:02,369 --> 01:04:11,039
gonna get incorrect execution whatever

1389
01:04:05,209 --> 01:04:14,519
let's give it a shot I'm gonna I'm gonna

1390
01:04:11,039 --> 01:04:17,159
comment out the locks and the question

1391
01:04:14,519 --> 01:04:24,179
is what happens if I run the code with

1392
01:04:17,159 --> 01:04:26,459
no locks what am I gonna see so we may

1393
01:04:24,179 --> 01:04:28,589
see a ru or I'll call twice or I fetch

1394
01:04:26,459 --> 01:04:31,649
twice yeah that's yeah that would be the

1395
01:04:28,589 --> 01:04:34,799
error you might expect alright so I'll

1396
01:04:31,649 --> 01:04:36,269
run it without locks and we're looking

1397
01:04:34,799 --> 01:04:38,459
at the concurrent map the one in the

1398
01:04:36,269 --> 01:04:40,369
middle this time it doesn't seem to have

1399
01:04:38,459 --> 01:04:49,139
fetched anything twice it's only five

1400
01:04:40,369 --> 01:04:50,719
run again gosh so far genius so maybe

1401
01:04:49,139 --> 01:04:52,499
we're wasting our time with those locks

1402
01:04:50,719 --> 01:04:57,419
yeah never seems to go wrong I've

1403
01:04:52,499 --> 01:05:00,269
actually never seem to go wrong so the

1404
01:04:57,419 --> 01:05:03,329
code is nevertheless wrong and someday

1405
01:05:00,269 --> 01:05:04,559
it will fail okay the problem is that

1406
01:05:03,329 --> 01:05:06,179
you know this is only a couple of

1407
01:05:04,559 --> 01:05:07,979
instructions here and so the chances of

1408
01:05:06,179 --> 01:05:09,539
these two threads which are maybe

1409
01:05:07,979 --> 01:05:12,269
hundreds of instructions happening to

1410
01:05:09,539 --> 01:05:14,489
stumble on this you know the same couple

1411
01:05:12,269 --> 01:05:17,729
of instructions at the same time is

1412
01:05:14,489 --> 01:05:20,459
quite low and indeed and and this is a

1413
01:05:17,729 --> 01:05:23,519
real bummer about buggy code with races

1414
01:05:20,459 --> 01:05:25,289
is that it usually works just fine but

1415
01:05:23,519 --> 01:05:28,020
it probably won't work when the customer

1416
01:05:25,289 --> 01:05:30,510
runs it on their computer

1417
01:05:28,020 --> 01:05:32,940
so it's actually bad news for us right

1418
01:05:30,510 --> 01:05:34,920
what do we you know it it can be in

1419
01:05:32,940 --> 01:05:37,170
complex programs quite difficult to

1420
01:05:34,920 --> 01:05:39,390
figure out if you have a race right and

1421
01:05:37,170 --> 01:05:41,910
you might you may have code that just

1422
01:05:39,390 --> 01:05:44,610
looks completely reasonable that is in

1423
01:05:41,910 --> 01:05:47,970
fact sort of unknown to you using shared

1424
01:05:44,610 --> 01:05:50,040
variables and the answer is you really

1425
01:05:47,970 --> 01:05:53,580
the only way to find races in practice

1426
01:05:50,040 --> 01:05:55,890
to be is you automated tools and luckily

1427
01:05:53,580 --> 01:06:00,119
go actually gives us this pretty good

1428
01:05:55,890 --> 01:06:04,619
race detector built-in to go and you

1429
01:06:00,119 --> 01:06:06,540
should use it so if you pass the - race

1430
01:06:04,619 --> 01:06:09,710
flag when you have to get your go

1431
01:06:06,540 --> 01:06:11,760
program and run this race detector which

1432
01:06:09,710 --> 01:06:16,650
well I'll run the race detector and

1433
01:06:11,760 --> 01:06:19,680
we'll see so it emits an error message

1434
01:06:16,650 --> 01:06:21,420
from us it's found a race and it

1435
01:06:19,680 --> 01:06:23,550
actually tells us exactly where the race

1436
01:06:21,420 --> 01:06:25,260
happened so there's a lot of junk in

1437
01:06:23,550 --> 01:06:28,320
this output but the really critical

1438
01:06:25,260 --> 01:06:30,180
thing is that the race detector realize

1439
01:06:28,320 --> 01:06:32,790
that we had read a variable that's what

1440
01:06:30,180 --> 01:06:35,670
this read is that was previously written

1441
01:06:32,790 --> 01:06:37,770
and there was no intervening release and

1442
01:06:35,670 --> 01:06:40,200
acquire of a lock that's what that's

1443
01:06:37,770 --> 01:06:43,710
what this means furthermore it tells us

1444
01:06:40,200 --> 01:06:49,290
the line number so it's told us that the

1445
01:06:43,710 --> 01:06:51,660
read was a line 43 and the write the

1446
01:06:49,290 --> 01:06:53,190
previous write was at line 44 and indeed

1447
01:06:51,660 --> 01:06:56,220
we look at the code and the read isn't

1448
01:06:53,190 --> 01:06:58,170
line 43 and the right is at lying 44 so

1449
01:06:56,220 --> 01:07:00,869
that means that one thread did a write

1450
01:06:58,170 --> 01:07:02,520
at line 44 and then without any

1451
01:07:00,869 --> 01:07:05,340
intervening lock and another thread came

1452
01:07:02,520 --> 01:07:07,560
along and read that written data at line

1453
01:07:05,340 --> 01:07:10,020
43 that's basically what the race

1454
01:07:07,560 --> 01:07:11,820
detector is looking for the way it works

1455
01:07:10,020 --> 01:07:15,000
internally is it allocates sort of

1456
01:07:11,820 --> 01:07:16,260
shadow memory now lucky some you know it

1457
01:07:15,000 --> 01:07:17,460
uses a huge amount of memory and

1458
01:07:16,260 --> 01:07:19,710
basically for every one of your memory

1459
01:07:17,460 --> 01:07:21,600
locations the race detector is allocated

1460
01:07:19,710 --> 01:07:24,330
a little bit of memory itself in which

1461
01:07:21,600 --> 01:07:26,400
it keeps track of which threads recently

1462
01:07:24,330 --> 01:07:28,590
read or wrote every single memory

1463
01:07:26,400 --> 01:07:30,810
location and then when and it also to

1464
01:07:28,590 --> 01:07:32,609
keep tracking keeping track of when

1465
01:07:30,810 --> 01:07:35,430
threads acquiring release locks and do

1466
01:07:32,609 --> 01:07:37,980
other synchronization activities that it

1467
01:07:35,430 --> 01:07:39,180
knows forces but force threads to not

1468
01:07:37,980 --> 01:07:40,980
run

1469
01:07:39,180 --> 01:07:42,450
and if the race detector driver sees a

1470
01:07:40,980 --> 01:07:45,210
ha there was a memory location that was

1471
01:07:42,450 --> 01:07:49,160
written and then read with no

1472
01:07:45,210 --> 01:08:06,600
intervening market it'll raise an error

1473
01:07:49,160 --> 01:08:12,170
yes I believe it is not perfect yeah I

1474
01:08:06,600 --> 01:08:15,180
have to think about it what one

1475
01:08:12,170 --> 01:08:18,900
certainly one way it is not perfect is

1476
01:08:15,180 --> 01:08:21,270
that if you if you don't execute some

1477
01:08:18,900 --> 01:08:25,110
code the race detector doesn't know

1478
01:08:21,270 --> 01:08:27,990
anything about it so it's not analyzing

1479
01:08:25,110 --> 01:08:29,220
it's not doing static analysis the

1480
01:08:27,990 --> 01:08:31,770
racing sectors not looking at your

1481
01:08:29,220 --> 01:08:33,390
source and making decisions based on the

1482
01:08:31,770 --> 01:08:35,700
source it's sort of watching what

1483
01:08:33,390 --> 01:08:37,650
happened at on this particular run of

1484
01:08:35,700 --> 01:08:39,330
the program and so if this particular

1485
01:08:37,650 --> 01:08:42,450
run of the program didn't execute some

1486
01:08:39,330 --> 01:08:44,370
code that happens to read or write

1487
01:08:42,450 --> 01:08:46,500
shared data then the race detector will

1488
01:08:44,370 --> 01:08:48,270
never know and there could be erased

1489
01:08:46,500 --> 01:08:49,320
there so that's certainly something to

1490
01:08:48,270 --> 01:08:50,580
watch out for so you know if you're

1491
01:08:49,320 --> 01:08:53,100
serious about the race detector you need

1492
01:08:50,580 --> 01:08:55,620
to set up sort of testing apparatus that

1493
01:08:53,100 --> 01:08:59,340
tries to make sure all all the code is

1494
01:08:55,620 --> 01:09:01,620
executed but it's it's it's very good

1495
01:08:59,340 --> 01:09:07,830
and you just have to use it for your 8

1496
01:09:01,620 --> 01:09:09,300
to 4 lives okay so this is race here and

1497
01:09:07,830 --> 01:09:12,330
of course the race didn't actually occur

1498
01:09:09,300 --> 01:09:14,370
what the race editor did not see was the

1499
01:09:12,330 --> 01:09:17,370
actual interleaving simultaneous

1500
01:09:14,370 --> 01:09:18,900
execution of some sensitive code right

1501
01:09:17,370 --> 01:09:21,860
it didn't see two threads literally

1502
01:09:18,900 --> 01:09:23,970
execute lines 43 and 44 at the same time

1503
01:09:21,860 --> 01:09:25,140
and as we know from having run the

1504
01:09:23,970 --> 01:09:28,320
things by hand that apparently doesn't

1505
01:09:25,140 --> 01:09:29,880
happen only with low probability all it

1506
01:09:28,320 --> 01:09:31,530
saw was at one point that was a right

1507
01:09:29,880 --> 01:09:37,620
and they made me much later there was a

1508
01:09:31,530 --> 01:09:39,180
read with no intervening walk and so

1509
01:09:37,620 --> 01:09:41,540
enact in that sense it can sort of

1510
01:09:39,180 --> 01:09:47,630
detect races that didn't actually happen

1511
01:09:41,540 --> 01:09:47,630
or didn't really cause bugs okay

1512
01:09:49,540 --> 01:09:57,550
okay one final question about this this

1513
01:09:52,550 --> 01:09:57,550
crawler how many threads does it create

1514
01:10:03,639 --> 01:10:24,969
yeah and how many concurrent threads

1515
01:10:10,119 --> 01:10:27,159
could there be yeah so a defect in this

1516
01:10:24,969 --> 01:10:28,719
crawler is that there's no obvious bound

1517
01:10:27,159 --> 01:10:30,580
on the number of simultaneous threads

1518
01:10:28,719 --> 01:10:32,800
that might create you know with the test

1519
01:10:30,580 --> 01:10:34,659
case which only has five URLs big

1520
01:10:32,800 --> 01:10:36,610
whoopee but if you're crawling a real

1521
01:10:34,659 --> 01:10:38,380
wheel web with you know I don't know are

1522
01:10:36,610 --> 01:10:40,389
there billions of URLs out there maybe

1523
01:10:38,380 --> 01:10:41,380
not we certainly don't want to be in a

1524
01:10:40,389 --> 01:10:43,380
position where the crawler might

1525
01:10:41,380 --> 01:10:46,119
accidentally create billions of threads

1526
01:10:43,380 --> 01:10:47,889
because you know thousands of threads

1527
01:10:46,119 --> 01:10:51,070
it's just fine billions of threads it's

1528
01:10:47,889 --> 01:10:54,280
not okay because each one sits on some

1529
01:10:51,070 --> 01:10:56,080
amount of memory so a you know there's

1530
01:10:54,280 --> 01:10:58,270
probably many defects in real life for

1531
01:10:56,080 --> 01:11:00,219
this crawler but one at the level we're

1532
01:10:58,270 --> 01:11:01,600
talking about is that it does create too

1533
01:11:00,219 --> 01:11:03,280
many threads and really ought to have a

1534
01:11:01,600 --> 01:11:04,630
way of saying well you can create 20

1535
01:11:03,280 --> 01:11:06,520
threads or 100 threads or a thousand

1536
01:11:04,630 --> 01:11:08,199
threads but no more so one way to do

1537
01:11:06,520 --> 01:11:11,050
that would be to pre create a pool a

1538
01:11:08,199 --> 01:11:13,179
fixed size pool of workers and have the

1539
01:11:11,050 --> 01:11:14,830
workers just iteratively look for

1540
01:11:13,179 --> 01:11:18,159
another URL to crawl crawl that URL

1541
01:11:14,830 --> 01:11:21,070
rather than creating a new thread for

1542
01:11:18,159 --> 01:11:23,230
each URL okay so next up I want to talk

1543
01:11:21,070 --> 01:11:25,600
about a another crawler that's

1544
01:11:23,230 --> 01:11:28,780
implemented and a significantly

1545
01:11:25,600 --> 01:11:31,869
different way using channels instead of

1546
01:11:28,780 --> 01:11:33,489
shared memory it's a member on the mutex

1547
01:11:31,869 --> 01:11:34,840
call or I just said there is this table

1548
01:11:33,489 --> 01:11:36,550
of URLs that are called that's shared

1549
01:11:34,840 --> 01:11:40,330
between all the threads and asked me

1550
01:11:36,550 --> 01:11:44,440
locked this version does not have such a

1551
01:11:40,330 --> 01:11:52,510
table does not share memory and does not

1552
01:11:44,440 --> 01:11:55,060
need to use locks okay so this one the

1553
01:11:52,510 --> 01:11:57,790
instead there's basically a master

1554
01:11:55,060 --> 01:12:00,699
thread that's his master function on a

1555
01:11:57,790 --> 01:12:02,800
decent 986 and it has a table but the

1556
01:12:00,699 --> 01:12:06,690
table is private to the master function

1557
01:12:02,800 --> 01:12:09,219
and what the master function is doing is

1558
01:12:06,690 --> 01:12:11,409
instead of sort of basically creating a

1559
01:12:09,219 --> 01:12:13,330
tree of functions that corresponds to

1560
01:12:11,409 --> 01:12:17,940
the exploration of the graph which the

1561
01:12:13,330 --> 01:12:21,880
previous crawler did this one fires off

1562
01:12:17,940 --> 01:12:23,770
one ute one guru team per URL that it's

1563
01:12:21,880 --> 01:12:26,560
fetches and that but it's only the

1564
01:12:23,770 --> 01:12:28,300
master only the one master that's

1565
01:12:26,560 --> 01:12:30,130
creating these threads so we don't have

1566
01:12:28,300 --> 01:12:35,199
a tree of functions creating threads we

1567
01:12:30,130 --> 01:12:37,540
just have the one master okay so it

1568
01:12:35,199 --> 01:12:41,550
creates its own private map a line 88

1569
01:12:37,540 --> 01:12:44,650
this record what it's fetched and then

1570
01:12:41,550 --> 01:12:46,900
it also creates a channel just a single

1571
01:12:44,650 --> 01:12:49,120
channel that all of its worker threads

1572
01:12:46,900 --> 01:12:50,699
are going to talk to and the idea is

1573
01:12:49,120 --> 01:12:53,040
that it's gonna fire up a worker thread

1574
01:12:50,699 --> 01:12:55,630
and each worker thread that it fires up

1575
01:12:53,040 --> 01:12:58,150
when it finished such as fetching the

1576
01:12:55,630 --> 01:13:00,250
page will send exactly one item back to

1577
01:12:58,150 --> 01:13:03,219
the master on the channel and that item

1578
01:13:00,250 --> 01:13:07,960
will be a list of the URLs in the page

1579
01:13:03,219 --> 01:13:10,420
that that worker thread fetched so the

1580
01:13:07,960 --> 01:13:13,989
master sits in a loop we're in line

1581
01:13:10,420 --> 01:13:16,780
eighty nine is reading entries from the

1582
01:13:13,989 --> 01:13:20,469
channel and so we have to imagine that

1583
01:13:16,780 --> 01:13:22,840
it's started up some workers in advance

1584
01:13:20,469 --> 01:13:24,489
and now it's reading the information the

1585
01:13:22,840 --> 01:13:26,830
URL lists that those workers send back

1586
01:13:24,489 --> 01:13:28,810
and each time he gets a URL is sitting

1587
01:13:26,830 --> 01:13:32,620
on land eighty nine it then loops over

1588
01:13:28,810 --> 01:13:36,100
the URLs in that URL list from a single

1589
01:13:32,620 --> 01:13:39,969
page fetch align ninety and if the URL

1590
01:13:36,100 --> 01:13:42,190
hasn't already been fetched it fires off

1591
01:13:39,969 --> 01:13:44,800
a new worker at line 94 to fetch that

1592
01:13:42,190 --> 01:13:47,320
URL and if we look at the worker code

1593
01:13:44,800 --> 01:13:51,130
online starting line 77 basically calls

1594
01:13:47,320 --> 01:13:53,710
his fetcher and then sends a message on

1595
01:13:51,130 --> 01:13:57,929
the channel a line 80 or 82 saying

1596
01:13:53,710 --> 01:14:01,170
here's the URLs in the page they fetched

1597
01:13:57,929 --> 01:14:03,639
and notice that now that the maybe

1598
01:14:01,170 --> 01:14:07,989
interesting thing about this is that the

1599
01:14:03,639 --> 01:14:10,210
worker threads don't share any objects

1600
01:14:07,989 --> 01:14:11,530
there's no shared object between the

1601
01:14:10,210 --> 01:14:12,850
workers and the master so we don't have

1602
01:14:11,530 --> 01:14:16,360
to worry about locking we don't have to

1603
01:14:12,850 --> 01:14:18,940
worry about rhesus instead this is a

1604
01:14:16,360 --> 01:14:21,100
example of sort of communicating

1605
01:14:18,940 --> 01:14:25,620
information instead of getting at it

1606
01:14:21,100 --> 01:14:25,620
through shared memory yes

1607
01:14:33,930 --> 01:14:40,810
yeah yeah so the observation is that the

1608
01:14:38,140 --> 01:14:42,250
code appears but the workers are the

1609
01:14:40,810 --> 01:14:47,130
observation is the workers are modifying

1610
01:14:42,250 --> 01:14:47,130
ch while the Masters reading it and

1611
01:14:49,170 --> 01:14:54,160
that's not the way the go authors would

1612
01:14:51,520 --> 01:14:55,360
like you to think about this the way

1613
01:14:54,160 --> 01:14:58,030
they want you to think about this is

1614
01:14:55,360 --> 01:15:00,880
that CH is a channel and the channel has

1615
01:14:58,030 --> 01:15:03,070
send and receive operations and the

1616
01:15:00,880 --> 01:15:05,260
workers are sending on the channel while

1617
01:15:03,070 --> 01:15:09,250
the master receives on the channel and

1618
01:15:05,260 --> 01:15:11,050
that's perfectly legal the channel is

1619
01:15:09,250 --> 01:15:12,790
happy I mean what that really means is

1620
01:15:11,050 --> 01:15:15,610
that the internal implementation of

1621
01:15:12,790 --> 01:15:19,000
channel has a mutex in it and the

1622
01:15:15,610 --> 01:15:20,860
channel operations are careful to take

1623
01:15:19,000 --> 01:15:22,449
out the mutex when they're messing with

1624
01:15:20,860 --> 01:15:24,190
the channels internal data to ensure

1625
01:15:22,449 --> 01:15:27,580
that it doesn't actually have any

1626
01:15:24,190 --> 01:15:29,290
reasons in it but yeah channels are sort

1627
01:15:27,580 --> 01:15:30,400
of protected against concurrency and

1628
01:15:29,290 --> 01:15:34,680
you're allowed to use them concurrently

1629
01:15:30,400 --> 01:15:34,680
from different threads yes

1630
01:15:36,389 --> 01:15:43,190
over the channel receive yes

1631
01:15:53,810 --> 01:15:58,850
we don't need to close the channel I

1632
01:15:56,260 --> 01:16:00,590
mean okay the the break statement is

1633
01:15:58,850 --> 01:16:03,160
about when the crawl has completely

1634
01:16:00,590 --> 01:16:06,410
finished and we fetched every single URL

1635
01:16:03,160 --> 01:16:09,230
right because hey what's going on is the

1636
01:16:06,410 --> 01:16:13,190
master is keeping I mean this n value is

1637
01:16:09,230 --> 01:16:14,860
private value and a master every time it

1638
01:16:13,190 --> 01:16:17,360
fires off a worker at increments the end

1639
01:16:14,860 --> 01:16:20,480
though every worker it starts since

1640
01:16:17,360 --> 01:16:21,920
exactly one item on the channel and so

1641
01:16:20,480 --> 01:16:23,120
every time the master reads an item off

1642
01:16:21,920 --> 01:16:24,920
the channel it knows that one of his

1643
01:16:23,120 --> 01:16:29,060
workers is finished and when the number

1644
01:16:24,920 --> 01:16:32,870
of outstanding workers goes to zero then

1645
01:16:29,060 --> 01:16:34,520
we're done and we don't once the number

1646
01:16:32,870 --> 01:16:36,500
of outstanding workers goes to zero then

1647
01:16:34,520 --> 01:16:40,370
the only reference to the channel is

1648
01:16:36,500 --> 01:16:41,780
from the master or from oh really from

1649
01:16:40,370 --> 01:16:43,460
the code that calls the master and so

1650
01:16:41,780 --> 01:16:45,260
the garbage collector will very soon see

1651
01:16:43,460 --> 01:16:48,680
that the channel has no references to it

1652
01:16:45,260 --> 01:16:50,060
and will free the channel so in this

1653
01:16:48,680 --> 01:16:53,630
case sometimes you need to close

1654
01:16:50,060 --> 01:16:56,170
channels but actually I rarely have to

1655
01:16:53,630 --> 01:16:56,170
close channels

1656
01:17:03,150 --> 01:17:06,050
he said again

1657
01:17:09,749 --> 01:17:16,389
so the question is alright so you can

1658
01:17:12,219 --> 01:17:19,949
see at line 106 before calling master

1659
01:17:16,389 --> 01:17:25,659
concurrent channel sort of fires up one

1660
01:17:19,949 --> 01:17:26,710
shoves one URL into the channel and it's

1661
01:17:25,659 --> 01:17:28,059
to sort of get the whole thing started

1662
01:17:26,710 --> 01:17:29,469
because the code for master was written

1663
01:17:28,059 --> 01:17:31,749
you know the master goes right into

1664
01:17:29,469 --> 01:17:33,489
reading from the channel line 89 so

1665
01:17:31,749 --> 01:17:36,550
there better be something in the channel

1666
01:17:33,489 --> 01:17:38,260
otherwise line 89 would block forever so

1667
01:17:36,550 --> 01:17:42,550
if it weren't for that little code at

1668
01:17:38,260 --> 01:17:44,050
line 107 the for loop at 89 would block

1669
01:17:42,550 --> 01:17:54,460
reading from the channel forever and

1670
01:17:44,050 --> 01:17:56,289
this code wouldn't work well yeah so the

1671
01:17:54,460 --> 01:17:57,510
observation is gosh you know wouldn't it

1672
01:17:56,289 --> 01:17:59,530
be nice to be able to write code that

1673
01:17:57,510 --> 01:18:01,570
would be able to notice if there's

1674
01:17:59,530 --> 01:18:03,219
nothing waiting on the channel and you

1675
01:18:01,570 --> 01:18:05,019
can if you look up the Select statement

1676
01:18:03,219 --> 01:18:06,579
it's much more complicated than this but

1677
01:18:05,019 --> 01:18:09,460
there is the Select statement which

1678
01:18:06,579 --> 01:18:11,139
allows you to proceed to not block if

1679
01:18:09,460 --> 01:18:13,590
something if there's nothing waiting on

1680
01:18:11,139 --> 01:18:13,590
the channel

1681
01:18:44,590 --> 01:19:02,600
because the work resin finish okay sorry

1682
01:18:59,600 --> 01:19:03,830
to the first question is there I think

1683
01:19:02,600 --> 01:19:05,630
what you're really worried about is

1684
01:19:03,830 --> 01:19:09,110
whether we're actually able to launch

1685
01:19:05,630 --> 01:19:37,220
parallel so the very first step won't be

1686
01:19:09,110 --> 01:19:40,270
in parallel because there's an exit

1687
01:19:37,220 --> 01:19:44,450
owner the for-loop weights in at line 89

1688
01:19:40,270 --> 01:19:47,660
that's not okay that for loop at line 89

1689
01:19:44,450 --> 01:19:49,190
is does not just loop over the current

1690
01:19:47,660 --> 01:19:54,230
contents of the channel and then quit

1691
01:19:49,190 --> 01:19:58,100
that is the for loop at 89 is going to

1692
01:19:54,230 --> 01:19:59,780
read it may never exit but it's gonna

1693
01:19:58,100 --> 01:20:01,130
read it's just going to keep waiting

1694
01:19:59,780 --> 01:20:04,280
until something shows up in the channel

1695
01:20:01,130 --> 01:20:10,250
so if you don't hit the break at line 99

1696
01:20:04,280 --> 01:20:12,440
the for loop own exit yeah alright I'm

1697
01:20:10,250 --> 01:20:15,800
afraid we're out of time we'll continue

1698
01:20:12,440 --> 01:20:18,260
this actually we have a presentation

1699
01:20:15,800 --> 01:20:20,950
scheduled by the TAS which I'll talk

1700
01:20:18,260 --> 01:20:20,950
more about go

