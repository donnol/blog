package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func medianFilter(img image.Image, kernelSize int) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// 创建一个新的图像作为结果
	result := image.NewRGBA(bounds)

	// 填充图像边缘
	paddedImg := padImage(img, kernelSize/2)

	// 遍历图像每个像素
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// 获取窗口内的像素值
			window := getWindow(paddedImg, x, y, kernelSize)

			// 计算窗口内像素值的排序
			sortedWindow := sortPixels(window)

			// 获取中值
			median := sortedWindow[len(sortedWindow)/2]

			// 将中值赋值给输出图像
			result.Set(x, y, median)
		}
	}

	return result
}

// 填充图像边缘
func padImage(img image.Image, padding int) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// 创建一个新的图像，包含填充区域
	paddedImg := image.NewRGBA(image.Rect(0, 0, width+padding*2, height+padding*2))

	// 将原始图像复制到新图像的中心
	draw.Draw(paddedImg, paddedImg.Bounds(), img, image.Point{padding, padding}, draw.Src)

	return paddedImg
}

// 获取窗口内的像素值
func getWindow(img image.Image, x, y, kernelSize int) []color.Color {
	window := make([]color.Color, 0)
	for i := y - kernelSize/2; i <= y+kernelSize/2; i++ {
		for j := x - kernelSize/2; j <= x+kernelSize/2; j++ {
			if i >= 0 && i < img.Bounds().Dy() && j >= 0 && j < img.Bounds().Dx() {
				window = append(window, img.At(j, i))
			}
		}
	}
	return window
}

// 对像素颜色值进行排序
func sortPixels(pixels []color.Color) []color.Color {
	// 比较函数，根据像素的亮度进行排序
	compare := func(i, j int) bool {
		r1, g1, b1, _ := pixels[i].RGBA()
		r2, g2, b2, _ := pixels[j].RGBA()
		return (r1+g1+b1)/3 < (r2+g2+b2)/3
	}

	// 使用快速排序对像素颜色值进行排序
	quickSort(pixels, 0, len(pixels)-1, compare)

	return pixels
}

// 快速排序算法
func quickSort(arr []color.Color, low, high int, compare func(i, j int) bool) {
	if low < high {
		pivot := partition(arr, low, high, compare)
		quickSort(arr, low, pivot-1, compare)
		quickSort(arr, pivot+1, high, compare)
	}
}

// 快速排序的分区操作
func partition(arr []color.Color, low, high int, compare func(i, j int) bool) int {
	i := low - 1
	for j := low; j < high; j++ {
		if compare(j, high) {
			i++
			swap(arr, i, j)
		}
	}
	swap(arr, i+1, high)
	return i + 1
}

// 交换两个元素
func swap(arr []color.Color, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	// 读取输入图像
	file, err := os.Open("./testdata/AOP.png")
	if err != nil {
		fmt.Println("打开输入图像失败:", err)
		return
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("解码图像失败:", err)
		return
	}

	// 应用中值滤波器
	filteredImg := medianFilter(img, 3) // 使用 3x3 的窗口

	// 保存处理后的图像
	output, err := os.Create("./testdata/output.png")
	if err != nil {
		fmt.Println("创建输出文件失败:", err)
		return
	}
	defer output.Close()

	err = png.Encode(output, filteredImg)
	if err != nil {
		fmt.Println("编码输出图像失败:", err)
		return
	}

	fmt.Println("中值滤波完成，已保存至 output.png")
}
