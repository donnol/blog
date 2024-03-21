# ffmpeg

## 图片加水印

`ffmpeg -i C:\Users\Pictures\1.png -vf "drawtext=fontfile=simhei.ttf:fontcolor=blue:fontsize=100:text='EGHI':x=W-tw-100:y=H-th-100:shadowy=2" C:\Users\Pictures\8x.jpg`

## 截图

windows: `ffmpeg -f gdigrab -s 500x500 -offset_x 100 -offset_y 100 -i desktop -frames:v 1 C:\Users\Pictures\screen.png`

## 录屏

windows: `ffmpeg -f gdigrab -framerate 20 -i desktop C:\Users\Pictures\out.avi`
