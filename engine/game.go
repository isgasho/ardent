package engine

const (
	FlagResizable = 1 << iota
	FlagRunsInBackground
)

// Game is an engine instance.
type Game interface {
	// Run starts running the game.
	Run() error

	// AddRenderer adds a renderer
	// to the game's draw stack. Renderers will
	// be applied in the order they are added.
	AddRenderer(...Renderer)

	// IsFullscreen returns the fullscreen state of the game.
	IsFullscreen() bool
	// SetFullscreen sets the fullscreen state of the game.
	SetFullscreen(bool)
	// IsVsync returns the vsync state of the game.
	IsVsync() bool
	// SetVsync sets the vsync state of the game.
	SetVsync(bool)
	// IsFocused returns the focused state of the game.
	IsFocused() bool

	Component
	Input
}
