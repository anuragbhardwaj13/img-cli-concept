
#### **About This Project**
Hey there! I was just experimenting with Go's `image` package and decided to try out some basic image manipulation tasks. This little program can:
- Rotate an image 90 degrees.
- Resize an image to any dimensions.
- Convert an image to PNG format.

It's a simple project for learning and fun, nothing fancy!

#### **How It Works**
1. The program reads an image file (like `test.jpg`) and decodes it.
2. You can uncomment the functions in the `main()` function to:
   - Rotate the image (`rotate90(img)`).
   - Resize the image (`resizeImage(img, newHeight, newWidth)`).
   - Save the image as a PNG (`saveAsPng(img)`).
3. The output files are saved in the same directory:
   - Rotated: `rotated.jpg`
   - Resized: `resize.png`
   - PNG version: `test.png`

#### **How to Run**
1. Drop your image file (e.g., `test.jpg`) in the same folder as the code.
2. Run the program with `go run main.go`.
3. Check out the processed images!

#### **Why I Did This**
I just wanted to mess around with Go's `image` package and see how it works. This is purely for practice, so don't expect anything too polished. 😊

Enjoy!