package opname

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// Ops friendry name generator.
//
// Safe to call in multi-threaded manner.
type Generator interface {
	// The length of generated names is less than or equal to 28.
	// The prefix of Gen() is always Prefix().
	Gen() string
	// Returns the prefix of the Generator instance.
	Prefix() string
}

type generator struct {
	prefix string
	dict   []string
}

var _ Generator = &generator{}

// Create new Generator whose Prefix() is prefix.
func New(prefix string, opts ...Option) (Generator, error) {
	if !validPrefix(prefix) {
		return nil, fmt.Errorf("prefix length must satisfy 1 <= length <= %d", MaxPrefixSize)
	}
	if err := validDict(dict); err != nil {
		return nil, err
	}
	g := generator{prefix, dict}
	for _, opt := range opts {
		if err := opt(&g); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}
	return &g, nil
}

// Option for Generator constructor.
type Option func(*generator) error

// Specify nick name dictionary of created Generator.
func NicknameDict(d []string) Option {
	return func(g *generator) error {
		if err := validDict(d); err != nil {
			return err
		}
		g.dict = d
		return nil
	}
}

func (g *generator) Gen() string {
	t := time.Now()
	source := rand.NewSource(t.UnixNano())
	r := rand.New(source)
	nickname := dict[r.Intn(len(dict))]

	return format(g.prefix, t, nickname)
}

func format(prefix string, t time.Time, nickname string) string {
	return fmt.Sprintf("%s%s%s%s",
		prefix,
		t.Format("20060102"),
		t.Format("150405"),
		nickname,
	)
}

func (g *generator) Prefix() string {
	return g.prefix
}

// Max size of generated name
const MaxSize = 28
const datetimeSize = 4 + 4

// Max prefix size of generator
const MaxPrefixSize = 4

// Max size of nick name
const MaxNicknameSize = MaxSize - datetimeSize - MaxPrefixSize

var (
	prefixRegexp   = regexp.MustCompile(fmt.Sprintf(`^[a-z][a-z0-9]{0,%d}$`, MaxPrefixSize-1))
	nicknameRegexp = regexp.MustCompile(fmt.Sprintf(`^[a-z0-9]{0,%d}[a-z]$`, MaxNicknameSize-1))
)

func validPrefix(s string) bool {
	return prefixRegexp.Match([]byte(s))
}

func validNichname(s string) bool {
	return nicknameRegexp.Match([]byte(s))
}

func validDict(ss []string) error {
	for _, s := range ss {
		if !validNichname(s) {
			return fmt.Errorf("invalid nichname: %s", s)
		}
	}
	return nil
}

var dict = []string{"dry", "dew", "bay", "hot", "icy", "fen", "wet", "dew", "icy", "dry", "fog", "bog", "wet", "wind", "calm", "pier", "dock", "port", "peal", "dune", "boom", "snow", "mild", "cool", "hail", "warm", "cold", "lake", "halo", "gust", "gale", "pond", "hazy", "mire", "rain", "heat", "flow", "airy", "mist", "haze", "smog", "tide", "thaw", "rime", "warm", "dewy", "arid", "flow", "cool", "roar", "boom", "bolt", "drip", "wave", "soak", "surf", "heat", "frost", "moist", "point", "tidal", "front", "jetty", "flash", "drift", "frost", "rainy", "whirl", "cloud", "clear", "clear", "chill", "flood", "humid", "clime", "flood", "radar", "solar", "draft", "glaze", "lunar", "swamp", "gusty", "spark", "vapor", "cloud", "snowy", "brisk", "balmy", "beach", "nippy", "shore", "muggy", "fresh", "marsh", "crisp", "blowy", "dusty", "dense", "soggy", "heavy", "foggy", "light", "smoky", "thick", "sunny", "crack", "shiny", "windy", "rainy", "storm", "chill", "drift", "coast", "sleet", "squall", "mirage", "freezy", "powder", "shovel", "arctic", "stormy", "static", "bright", "rumble", "breeze", "lagoon", "freeze", "wintry", "sparse", "breezy", "sultry", "frigid", "chilly", "frozen", "bitter", "steamy", "harbor", "drippy", "biting", "stormy", "drying", "cloudy", "frosty", "fogbow", "sundog", "shower", "pillar", "corona", "season", "icicle", "icecap", "aurora", "nimbus", "shiver", "puddle", "system", "funnel", "haboob", "levant", "cirrus", "nimbus", "meteor", "runoff", "mirage", "aurora", "kelvin", "bright", "zephyr", "puddle", "squall", "boreal", "tundra", "vortex", "icecap", "albedo", "stream", "icicle", "drafty", "static", "strike", "warmth", "splash", "slushy", "eclipse", "updraft", "current", "icycold", "frosted", "riptide", "climate", "drizzle", "monsoon", "estuary", "cyclone", "beaming", "tornado", "pouring", "flooded", "soaking", "splashy", "typhoon", "stratus", "cumulus", "rivulet", "cascade", "snowman", "sunbeam", "radiant", "trickle", "spatter", "celsius", "iceberg", "thunder", "rainbow", "celsius", "degrees", "eclipse", "drought", "glacier", "cyclone", "showery", "oldsnow", "chinook", "drizzle", "graupel", "melting", "glacial", "boiling", "thunder", "searing", "wetness", "snowcap", "whisper", "rainbow", "weather", "sunbeam", "tornado", "typhoon", "cyclone", "monsoon", "mistral", "cumulus", "thermal", "stratus", "windfarm", "sunlight", "undertow", "easterly", "forecast", "pressure", "climatic", "westerly", "downpour", "seashore", "blizzard", "clearsky", "snowfall", "overcast", "mudslide", "anabatic", "snowshoe", "chubasco", "tropical", "snowplow", "snowball", "sunshine", "overcast", "snowmelt", "meteoric", "snowfall", "heatwave", "icelayer", "icesheet", "humidity", "humidity", "frosting", "freezing", "icebound", "forecast", "blackice", "icesleet", "freezing", "whiteout", "volcanic", "eruption", "ashcloud", "blizzard", "heatwave", "raincoat", "moonbeam", "tropical", "sunlight", "electric", "snowline", "raindrum", "thundery", "raindrop", "borealis", "raingear", "raindrop", "spectrum", "moisture", "snowpack", "rainwear", "rainfall", "frostbite", "northerly", "rainslick", "driftwood", "iceflower", "cloudbank", "scorching", "lightning", "rainproof", "snowdrift", "australis", "celestial", "blizzardy", "temperate", "windblown", "starlight", "supercell", "meteorite", "dustcloud", "condition", "downdraft", "raincheck", "hailstorm", "hailstone", "satellite", "barometer", "meltwater", "freshsnow", "windspeed", "rainstorm", "jetstream", "windstorm", "rainmeter", "snowflake", "warmfront", "reservoir", "evalanche", "frostwork", "sparkling", "frostbite", "radiating", "lightning", "hurricane", "raincloud", "whirlwind", "frostline", "katabatic", "clearance", "coldfront", "sandstorm", "duststorm", "snowfield", "galeforce", "snowflake", "hurricane", "southerly", "icecrystal", "cloudburst", "sweltering", "fahrenheit", "refraction", "floodplain", "cloudscape", "anemometer", "fahrenheit", "hygrometer", "atmosphere", "glaciation", "waterspout", "breakwater", "overheated", "prediction", "glistening", "frostflower", "temperature", "thermometer", "weathercock", "altocumulus", "meteorology", "temperature", "seismometer", "anticyclone", "nimbostratus", "cyclogenesis", "thunderstorm", "cirrostratus", "cirrocumulus", "precipitation"}

func init() {
	if err := validDict(dict); err != nil {
		panic(err)
	}
}
