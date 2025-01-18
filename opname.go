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
	if err := validDict(defaultDict); err != nil {
		return nil, err
	}
	g := generator{prefix, defaultDict}
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
	nickname := defaultDict[r.Intn(len(defaultDict))]

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

var defaultDict = []string{"bay", "bog", "dew", "dry", "fen", "fog", "hot", "icy", "wet", "airy", "arid", "bolt", "boom", "calm", "cold", "cool", "dewy", "dock", "drip", "dune", "flow", "gale", "gust", "hail", "halo", "haze", "hazy", "heat", "lake", "mild", "mire", "mist", "peal", "pier", "pond", "port", "rain", "rime", "roar", "smog", "snow", "soak", "surf", "thaw", "tide", "warm", "wave", "wind", "balmy", "beach", "blowy", "brisk", "chill", "clear", "clime", "cloud", "coast", "crack", "crisp", "dense", "draft", "drift", "dusty", "flash", "flood", "foggy", "fresh", "front", "frost", "glaze", "gusty", "heavy", "humid", "jetty", "light", "lunar", "marsh", "moist", "muggy", "nippy", "point", "radar", "rainy", "shiny", "shore", "sleet", "smoky", "snowy", "soggy", "solar", "spark", "storm", "sunny", "swamp", "thick", "tidal", "vapor", "whirl", "windy", "albedo", "arctic", "aurora", "biting", "bitter", "boreal", "breeze", "breezy", "bright", "chilly", "cirrus", "cloudy", "corona", "drafty", "drippy", "drying", "fogbow", "freeze", "freezy", "frigid", "frosty", "frozen", "funnel", "haboob", "harbor", "icecap", "icicle", "kelvin", "lagoon", "levant", "meteor", "mirage", "nimbus", "pillar", "powder", "puddle", "rumble", "runoff", "season", "shiver", "shovel", "shower", "slushy", "sparse", "splash", "squall", "static", "steamy", "stormy", "stream", "strike", "sultry", "sundog", "system", "tundra", "vortex", "warmth", "wintry", "zephyr", "beaming", "boiling", "cascade", "celsius", "chinook", "climate", "cumulus", "current", "cyclone", "degrees", "drizzle", "drought", "eclipse", "estuary", "flooded", "frosted", "glacial", "glacier", "graupel", "iceberg", "icycold", "melting", "mistral", "monsoon", "oldsnow", "pouring", "radiant", "rainbow", "riptide", "rivulet", "searing", "showery", "snowcap", "snowman", "soaking", "spatter", "splashy", "stratus", "sunbeam", "thermal", "thunder", "tornado", "trickle", "typhoon", "updraft", "weather", "wetness", "whisper", "anabatic", "ashcloud", "blackice", "blizzard", "borealis", "chubasco", "clearsky", "climatic", "downpour", "easterly", "electric", "eruption", "forecast", "freezing", "frosting", "heatwave", "humidity", "icebound", "icelayer", "icesheet", "icesleet", "meteoric", "moisture", "moonbeam", "mudslide", "overcast", "pressure", "raincoat", "raindrop", "raindrum", "rainfall", "raingear", "rainwear", "seashore", "snowball", "snowfall", "snowline", "snowmelt", "snowpack", "snowplow", "snowshoe", "spectrum", "sunlight", "sunshine", "thundery", "tropical", "undertow", "volcanic", "westerly", "whiteout", "windfarm", "australis", "barometer", "blizzardy", "celestial", "clearance", "cloudbank", "coldfront", "condition", "downdraft", "driftwood", "dustcloud", "duststorm", "evalanche", "freshsnow", "frostbite", "frostline", "frostwork", "galeforce", "hailstone", "hailstorm", "hurricane", "iceflower", "jetstream", "katabatic", "lightning", "meltwater", "meteorite", "northerly", "radiating", "raincheck", "raincloud", "rainmeter", "rainproof", "rainslick", "rainstorm", "reservoir", "sandstorm", "satellite", "scorching", "snowdrift", "snowfield", "snowflake", "southerly", "sparkling", "starlight", "supercell", "temperate", "warmfront", "whirlwind", "windblown", "windspeed", "windstorm", "anemometer", "atmosphere", "breakwater", "cloudburst", "cloudscape", "fahrenheit", "floodplain", "glaciation", "glistening", "hygrometer", "icecrystal", "overheated", "prediction", "refraction", "sweltering", "waterspout", "altocumulus", "anticyclone", "frostflower", "meteorology", "seismometer", "temperature", "thermometer", "weathercock", "cirrocumulus", "cirrostratus", "cyclogenesis", "nimbostratus", "thunderstorm", "precipitation"}

func init() {
	if err := validDict(defaultDict); err != nil {
		panic(err)
	}
}
