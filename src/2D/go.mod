module CoreCascade2D

require filebuffer v0.0.0 // indirect

require color v0.0.0

require vector v0.0.0

require (
	github.com/chewxy/math32 v1.11.1 // indirect
	linear_image v0.0.0
)

replace filebuffer => ../libs/filebuffer

replace color => ../libs/color

replace vector => ../libs/vector

replace linear_image => ./../libs/linear_image

go 1.23.9
