# colormap

![](images/palettes.png)

## Usage

A Colormap is created by passing it one of the available color palettes. A ToRGBA method then accepts a normalized x value (```float64 [0.0,1.0]```) as well as an alpha value (either ```float64 [0.0,1.0]``` or ```int/uint8 [0,255]```) and returns a new pixel value.


```Go
inputFile, _ := os.Open("images/gray8.png")
defer inputFile.Close()

img, _, _ := image.Decode(inputFile)

cm := colormap.Colormap{Palette: palette.Magma{}}

size := img.Bounds().Size()
rect := image.Rect(0, 0, size.X, size.Y)
newImg := image.NewRGBA(rect)

for y := 0; y < size.Y; y++ {
    for x := 0; x < size.X; x++ {
        grayPixel := img.At(x, y)
        p := color.GrayModel.Convert(grayPixel).(color.Gray).Y
        px := cm.ToRGBA(p, 255)
        newImgnewImg.Set(x, y, px)
    }
}

outputfile, _ := os.Create("magma.png")
png.Encode(outputfile, newImg)
```

## References

The color maps ```inferno```, ```masma```, ```plasma```, ```viridis``` were created by Stéfan van der Walt ([@stefanv](https://github.com/stefanv)) and Nathaniel Smith ([@njsmith](https://github.com/njsmith)). More information is available [here](https://bids.github.io/colormap/) and palette data can be found at [BIDS/colormap](https://github.com/BIDS/colormap).

The color maps ```crest```, ```flare```, ```icefire```, ```mako```, ```rocket```, and ```vlag``` were created for the Python statistical data visualization package [```Seaborn```](https://github.com/mwaskom/seaborn). More info on ```Seaborn``` color palettes can be found at their [website](https://seaborn.pydata.org/tutorial/color_palettes.html).

The color map ```cividis``` was developed by Jamie R. Nuñez, Christopher R. Anderton, and Ryan S. Renslow. More info can be found in their [paper](https://journals.plos.org/plosone/article?id=10.1371/journal.pone.0199239).

The color map ```turbo``` was developed by Anton Mikhailov ([@antonthefirst](https://github.com/antonthefirst)). More info can be found at the [Google AI Blog](https://ai.googleblog.com/2019/08/turbo-improved-rainbow-colormap-for.html).

The color mapping code found here is a port of a [Python gist](https://gist.github.com/mikhailov-work/ee72ba4191942acecc03fe6da94fc73f) provided by Anton Mikhailov ([@antonthefirst](https://github.com/antonthefirst)).

Original image above courtesy of NASA.
