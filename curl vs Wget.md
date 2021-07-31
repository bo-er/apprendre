## [curl vs Wget](https://daniel.haxx.se/docs/curl-vs-wget.html)

# curl vs Wget

The main differences as I (Daniel Stenberg) see them. Please consider my bias towards [curl](https://curl.se/) since after all, curl is my baby - but I contribute to [Wget](http://www.gnu.org/software/wget/) as well.

Please let me know if you have other thoughts or comments on this document.

[File issues or pull-requests](https://github.com/bagder/docs) if you find problems or have improvements.

## What both commands do

- both are command line tools that can download contents from FTP, HTTP(S)
- both can send HTTP POST requests
- both support HTTP cookies
- both support metalink, HSTS and HTTP proxy
- both are designed to work without user interaction
- both are fully open source and free software
- both projects started in 1996 (under other names)
- both are portable and run on many operating systems

# How they differ

## curl

- _library_: curl is powered by _libcurl_ - a cross-platform library with a stable API that can be used by each and everyone. This difference is major since it creates a completely different attitude on how to do things internally. It is also slightly harder to make a library than a "mere" command line tool.
- _pipes_. curl works more like the traditional Unix cat command, it sends more stuff to stdout, and reads more from stdin in a "everything is a pipe" manner. Wget is more like cp, using the same analogue.
- _Single shot_: curl is basically made to do single-shot transfers of data. It transfers just the URLs that the user specifies, and does not contain any recursive downloading logic nor any sort of HTML parser.
- _More protocols_: curl supports FTP(S), GOPHER(S), HTTP(S), SCP, SFTP, TFTP, TELNET, DICT, LDAP(S), MQTT, FILE, POP3(S), IMAP(S), SMB(S), SMTP(S), RTMP and RTSP. Wget supports HTTP(S) and FTP.
- _More portable_: curl builds and runs on lots of more platforms than wget. For example: OS/400, TPF and other more "exotic" platforms that aren't straight-forward Unix clones. curl requires but a C89 compiler.
- _More SSL libraries_ and SSL support: curl can be built with one out of thirteen (13!) different SSL/TLS libraries, and it offers more control and wider support for protocol details.
- _HTTP auth_: curl supports more HTTP authentication methods, especially over HTTP proxies: Basic, Digest, NTLM and Negotiate
- _SOCKS_: curl supports SOCKS4 and SOCKS5 for proxy access. With local or proxy based name resolving.
- _Bidirectional_: curl offers upload and sending capabilities. Wget only offers plain HTTP POST support.
- _HTTP multipart/form-data_ sending, which allows users to do HTTP "upload" and in general emulate browsers and do HTTP automation to a wider extent.
- curl supports gzip, brotli, zstd and deflate Content-Encoding and does _automatic decompression_.
- curl offers and performs decompression of _Transfer-Encoded HTTP_, wget doesn't.
- curl supports _HTTP/2_, _HTTP/3_, _alt-svc_ and it does dual-stack connects using _Happy Eyeballs_.
- curl can do many transfers in parallel (`-Z`).
- _Much more developer activity_. While this can be debated, I consider three metrics here: mailing list activity, source code commit frequency and release frequency. Anyone following these two projects can see that the curl project has a lot higher pace in all these areas, and it has been so for 15+ years. [Compare on openhub](https://www.openhub.net/p/_compare?project_0=cURL&project_1=Wget).
- curl comes pre-installed on macOS and Windows 10. Wget does not.

## Wget

- Wget is _command line only_. There's no library.
- _Recursive!_: Wget's major strong side compared to curl is its ability to download recursively, or even just download everything that is referred to from a remote resource, be it a HTML page or a FTP directory listing.
- _Older_: Wget has traces back to its predecessor from [January 9, 1996](https://ftp.sunet.se/mirror/archive/ftp.sunet.se/pub/www/utilities/wget/old-versions/), while curl can be tracked back no earlier than to [November 11, 1996](https://curl.se/docs/history.html).
- _GPL_: Wget is _GPL v3_. curl is _MIT licensed_.
- _GNU_: Wget is part of the _GNU_ project and all copyrights are assigned to _FSF_. The curl project is entirely stand-alone and independent with no organization parenting at all with almost all copyrights owned by _Daniel_.
- Wget requires _no extra options_ to simply download a remote URL to a local file, while curl requires -o or -O.
- Wget supports only _GnuTLS or OpenSSL_ for SSL/TLS support.
- Wget supports only _Basic_ auth as the only auth type over HTTP proxy.
- Wget has no SOCKS support.
- Its ability to recover from a prematurely broken transfer and _continue downloading_ has no counterpart in curl.
- Wget enables more features by default: cookies, redirect-following, time stamping from the remote resource etc. With curl most of those features need to be explicitly enabled.
- There's a 'wget' in BusyBox, there's no curl there (it is not the actual wget, just a stripped down clone with the same name).
- Wget can be typed in using only the left hand on a qwerty keyboard!
- Wget requires a C99 compiler and also relies on gnulib.

# When to use which

Primarily: use the one that gets the job done for you.

Wget has (recursive) downloading powers that curl does not feature and it also handle download retries over unreliable connections possibly slightly more effective.

For just about everything else, curl is probably the more suitable tool.

# Additional Stuff

In recent years, **wget2** is worked on to become the replacement for wget. This comparison will eventually get wget2 details as well.

Two other capable tools with similar feature set include [aria2](https://aria2.github.io/) and [axel](https://github.com/axel-download-accelerator/axel) - try them out!

For a stricter feature by feature comparison (that also compares other similar tools), see the [curl comparison table](https://curl.se/docs/comparison-table.html)

# Thanks

Feedback and improvements by: Micah Cowan, Olemis Lang
