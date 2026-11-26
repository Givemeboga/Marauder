package banner

import (
	"fmt"
	"os"
)

// Banner is the marauder startup banner. Printed to stderr by default,
// suppressed when --json or --quiet is set. Pure ASCII only — must
// render on reverse shells and narrow terminals.
const Banner = `
         .                                                       .
      .*.    __  __    _    ____      _   _   _  ____  _____ ____     .*.
       .    |  \/  |  / \  |  _ \    / \ | | | ||  _ \| ____|  _ \     .
            | |\/| | / _ \ | |_) |  / _ \| | | || | | |  _| | |_) |
            | |  | |/ ___ \|  _ <  / ___ \ |_| || |_| | |___|  _ <
            |_|  |_/_/   \_\_| \_\/_/   \_\___/ |____/|_____|_| \_\

                Aperit Semitas - it opens the paths | v0.1
`

// Print writes the banner to stderr. Call from main() unless the user
// asked for --json or --quiet.
func Print() {
	fmt.Fprint(os.Stderr, Banner)
	fmt.Fprintln(os.Stderr)
}
