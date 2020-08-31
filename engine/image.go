package engine

// An Image represents a single unchanging
// image that can be applied to a renderer.
type Image interface {
	// Translate sets the x y
	// translation of the image
	// relative to the origin.
	Translate(float64, float64)

	// Scale sets the x y scale
	// of the image relative to the origin.
	Scale(float64, float64)

	// Rotate sets the rotation
	// in radians relative to the origin.
	Rotate(float64)

	// Size returns the size of the image.
	Size() (int, int)
}