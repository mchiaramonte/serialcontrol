# Remote RS232 control for Video Storm CMX devices (or at least the CMX88)
Very simple web-based interface to control my Video Storm CMX88. Currently runs on a Raspberry Pi W Zero 2

## Background
The Video Storm CMX88 is a very cool analog component video matrix switch and, as the name implies, it supports 8 inputs and 8 outputs. I got it off eBay for very cheap (sans power supply) and I didn't need it for all 8 outputs, but rather for taking one of the 8 inputs and sending the video to my CRT and Retrotink simultaneously. The switch also supports different outputs, but my main use-case was simultaneous output to the two aforementioned devices. The only thing I don't like about this solution is that the CMX88 only allows for forward advancement through the front menu system meaning that if you accidentally miss your output, you have to cycle through the whole menu to get back to it. Basically, just annoying more than anything else.

The other neat thing about this switch is that it supports [control via RS232](http://www.video-storm.com/manuals/CMX%20rs232.pdf), so I picked up a USB to Serial adapter and had intentions, at first, of using my MacBook whenever I wanted to change inputs/outputs. However, I quickly realized that this was less than ideal because it meant having the laptop handy at that moment and that I also have my USB-C adapter (since the serial adapter is USB A). Yeah, I can do it when I want to but somehow it felt less cumbersome to just press the buttons on the front of the switch most of the time. After messing around with the USB adapter on a Raspberry Pi W Zero 2 I realized there was a better way...

## Implementation
This is the first project I've ever worked on in Go. So, first and foremost, if there's anything in here that doesn't look Go-like or like it was done by an experienced Go programmer- guess what, it's because I'm not an experienced Go programmer.

Anyway, the implementation is dead simple. I'm using the `go.bug.st/serial.v1` module because that was the first thing that appeared in Google's amazing search trash heap. Using this, I query for the serial ports and simply select the first one found since it's not possible (in my setup anyway) for there to be multiple. The only reason I bother querying is on the off chance that something happens and the device path somehow changes. So, assuming there's at least one serial device, it picks the first one and attempts to open a port.

After that, it uses gin from `github.com/gin-gonic/gin` to host a simple index.html page for the web-based control and also an end-point for sending the request at `/video` using PUT. The request input is a simple JSON object:

```
{
    "Input": <input int>
    "Output": <output int>
}
```

And the index.html page is also dead simple in order to avoid needing any additional files, transpilers, frameworks, etc. It's just raw HTML and JavaScript right in the single file. It uses `fetch` to send the PUT request to the `/video` endpoint.

## Service Details/Assumptions
I have included a start script (start.sh) in case I ever need it to do something besides running the compiled service. There's a `control.service` file for running the app as a service in Raspbian. The expectation is that there is a user `control` created, it is in the `dialout` group in order to have access to the serial port, and that the service files are in `/opt/control`.