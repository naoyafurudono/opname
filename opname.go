package opname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Generator interface {
	// 最大maxSize文字の文字列を生成する。実行時刻によって異なる値を返す.
	// 先頭にはPrefix()を含む.
	Gen() string
	// このジェネレータが返す文字列のprefix.
	Prefix() string
}

type generator struct {
	prefix string
}

var _ Generator = &generator{}

func New(prefix string) (Generator, error) {
	if !validPrefix(prefix) {
		return nil, errors.New("prefix length must satisfy 1 <= length <= 4")
	}
	return &generator{prefix}, nil
}

func (g *generator) Gen() string {
	t := time.Now()
	source := rand.NewSource(t.UnixNano())
	r := rand.New(source)
	pretty := dict[r.Intn(len(dict))]

	return format(g.prefix, t, pretty)
}

func format(prefix string, t time.Time, pretty string) string {
	return fmt.Sprintf("%s%s%s%s",
		prefix,
		t.Format("20060102"),
		t.Format("150405"),
		pretty,
	)
}

func (g *generator) Prefix() string {
	return g.prefix
}

const maxSize = 28
const datetimeSize = 4 + 4
const maxPrefixSize = 4
const safePrettySup = maxSize - datetimeSize - maxPrefixSize

func validPrefix(s string) bool {
	return (1 <= len(s) && len(s) <= maxPrefixSize)
}

func staticallySafePretty(s string) bool {
	return len(s) <= safePrettySup
}

var dict = []string{"dry", "dew", "bay", "hot", "icy", "fen", "wet", "dew", "icy", "dry", "fog", "bog", "wet", "wind", "calm", "pier", "dock", "port", "peal", "dune", "boom", "snow", "mild", "cool", "hail", "warm", "cold", "lake", "halo", "gust", "gale", "pond", "hazy", "mire", "rain", "heat", "flow", "airy", "mist", "haze", "smog", "tide", "thaw", "rime", "warm", "dewy", "arid", "flow", "cool", "roar", "boom", "bolt", "drip", "wave", "soak", "surf", "heat", "flash", "frost", "moist", "point", "tidal", "front", "jetty", "flash", "drift", "frost", "rainy", "whirl", "cloud", "clear", "clear", "chill", "flood", "humid", "clime", "flood", "radar", "solar", "draft", "glaze", "lunar", "swamp", "gusty", "spark", "vapor", "cloud", "snowy", "brisk", "balmy", "beach", "nippy", "shore", "muggy", "fresh", "marsh", "crisp", "blowy", "dusty", "dense", "soggy", "heavy", "foggy", "light", "smoky", "thick", "sunny", "crack", "shiny", "windy", "rainy", "storm", "chill", "drift", "coast", "sleet", "squall", "mirage", "freezy", "powder", "shovel", "arctic", "stormy", "static", "bright", "rumble", "breeze", "lagoon", "freeze", "wintry", "sparse", "breezy", "sultry", "frigid", "chilly", "frozen", "bitter", "steamy", "harbor", "drippy", "biting", "stormy", "drying", "cloudy", "frosty", "fogbow", "sundog", "shower", "pillar", "corona", "season", "icicle", "icecap", "aurora", "nimbus", "shiver", "puddle", "system", "funnel", "haboob", "levant", "cirrus", "nimbus", "meteor", "runoff", "mirage", "aurora", "kelvin", "bright", "zephyr", "puddle", "squall", "boreal", "tundra", "vortex", "icecap", "albedo", "stream", "icicle", "drafty", "static", "strike", "warmth", "splash", "slushy", "eclipse", "updraft", "current", "icycold", "frosted", "riptide", "climate", "drizzle", "monsoon", "estuary", "cyclone", "beaming", "tornado", "pouring", "flooded", "soaking", "splashy", "typhoon", "stratus", "cumulus", "rivulet", "cascade", "snowman", "sunbeam", "radiant", "trickle", "spatter", "celsius", "iceberg", "thunder", "rainbow", "celsius", "degrees", "eclipse", "drought", "glacier", "cyclone", "showery", "oldsnow", "chinook", "drizzle", "graupel", "melting", "glacial", "boiling", "thunder", "searing", "wetness", "snowcap", "whisper", "rainbow", "weather", "sunbeam", "tornado", "typhoon", "cyclone", "monsoon", "mistral", "cumulus", "thermal", "stratus", "windfarm", "sunlight", "undertow", "easterly", "forecast", "pressure", "climatic", "westerly", "downpour", "seashore", "blizzard", "clearsky", "snowfall", "overcast", "mudslide", "anabatic", "snowshoe", "chubasco", "tropical", "snowplow", "snowball", "sunshine", "overcast", "snowmelt", "meteoric", "snowfall", "heatwave", "icelayer", "icesheet", "humidity", "humidity", "frosting", "freezing", "icebound", "forecast", "blackice", "icesleet", "freezing", "whiteout", "volcanic", "eruption", "ashcloud", "blizzard", "heatwave", "raincoat", "moonbeam", "tropical", "sunlight", "electric", "snowline", "raindrum", "thundery", "raindrop", "borealis", "raingear", "raindrop", "spectrum", "moisture", "snowpack", "rainwear", "rainfall", "frostbite", "northerly", "rainslick", "driftwood", "iceflower", "cloudbank", "scorching", "lightning", "rainproof", "snowdrift", "australis", "celestial", "blizzardy", "temperate", "windblown", "starlight", "supercell", "meteorite", "dustcloud", "condition", "downdraft", "raincheck", "hailstorm", "hailstone", "satellite", "barometer", "meltwater", "freshsnow", "windspeed", "rainstorm", "jetstream", "windstorm", "rainmeter", "snowflake", "warmfront", "reservoir", "evalanche", "frostwork", "sparkling", "frostbite", "radiating", "lightning", "hurricane", "raincloud", "whirlwind", "frostline", "katabatic", "clearance", "coldfront", "sandstorm", "duststorm", "snowfield", "galeforce", "snowflake", "hurricane", "southerly", "icecrystal", "cloudburst", "sweltering", "fahrenheit", "refraction", "floodplain", "cloudscape", "anemometer", "fahrenheit", "hygrometer", "atmosphere", "glaciation", "waterspout", "breakwater", "overheated", "prediction", "glistening", "frostflower", "temperature", "thermometer", "weathercock", "altocumulus", "meteorology", "temperature", "seismometer", "anticyclone", "nimbostratus", "cyclogenesis", "thunderstorm", "cirrostratus", "cirrocumulus", "precipitation"}

func init() {
	for i, p := range dict {
		if len(p) > safePrettySup {
			panic(fmt.Sprintf("too long pretty in the dict %d / %d: %s\n", i, len(dict), p))
		}
	}
}
