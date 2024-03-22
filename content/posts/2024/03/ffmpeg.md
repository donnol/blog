# ffmpeg

## 图片加水印

`ffmpeg -i C:\Users\Pictures\1.png -vf "drawtext=fontfile=simhei.ttf:fontcolor=blue:fontsize=100:text='EGHI':x=W-tw-100:y=H-th-100:shadowy=2" C:\Users\Pictures\8x.jpg`

## 截图

windows: `ffmpeg -f gdigrab -s 500x500 -offset_x 100 -offset_y 100 -i desktop -frames:v 1 C:\Users\Pictures\screen.png`

## 录屏

windows: `ffmpeg -f gdigrab -framerate 20 -i desktop C:\Users\Pictures\out.avi`

## 分割

`ffmpeg -i C:\Users\Pictures\output.mp4 -f segment -segment_time 60 -segment_format mpegts -segment_list C:\Users\Pictures\video_name.m3u8 -c copy -bsf:v h264_mp4toannexb -map 0 C:\Users\Pictures\course-%04d.ts`

将视频每60秒分割为一个`ts`文件。其中`video_name.m3u8`文件为ts清单，播放器会根据它按序播放`ts`文件：

```m3u8
#EXTM3U
#EXT-X-VERSION:3
#EXT-X-MEDIA-SEQUENCE:0
#EXT-X-ALLOW-CACHE:YES
#EXT-X-TARGETDURATION:63
#EXTINF:62.600000,
course-0000.ts
#EXTINF:1.200000,
course-0001.ts
#EXT-X-ENDLIST
```

为什么要这样分割呢？主要是为了提高并发。

> 关于m3u8文件，其实就是多个ts文件的`播放清单`，浏览器或者播放器可以直接解析这种文件，并把多个ts文件组合起来播放。
>
> 只需对ts文件做主从，负载均衡。这样把一个大的视频文件，分成多个小的ts文件，就可以减少带宽的性能消耗，避免出现性能问题。
