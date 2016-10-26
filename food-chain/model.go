package foodchain

const (
	die         = "I don't know why she swallowed the fly. Perhaps she'll die."
	dead        = "She's dead, of course!"
	testVersion = 2
)

type animal struct {
	name, action, resolve string
}

var animals = []animal{
	{
		name: "fly",
	},
	{
		name:    "spider",
		action:  "It wriggled and jiggled and tickled inside her.",
		resolve: "She swallowed the spider to catch the fly.",
	},
	{
		name:    "bird",
		action:  "How absurd to swallow a bird!",
		resolve: "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.",
	},
	{
		name:    "cat",
		action:  "Imagine that, to swallow a cat!",
		resolve: "She swallowed the cat to catch the bird.",
	},
	{
		name:    "dog",
		action:  "What a hog, to swallow a dog!",
		resolve: "She swallowed the dog to catch the cat.",
	},
	{
		name:    "goat",
		action:  "Just opened her throat and swallowed a goat!",
		resolve: "She swallowed the goat to catch the dog.",
	},
	{
		name:    "cow",
		action:  "I don't know how she swallowed a cow!",
		resolve: "She swallowed the cow to catch the goat.",
	},
	{
		name: "horse",
	},
}
