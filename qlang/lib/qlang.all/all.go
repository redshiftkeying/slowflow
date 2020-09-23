package qall

import (
	"github.com/redshiftkeying/slowflow-server/qlang/lib/bufio"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/bytes"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/crypto/md5"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/encoding/hex"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/encoding/json"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/eqlang"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/errors"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/io"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/io/ioutil"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/math"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/meta"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/net/http"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/os"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/path"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/reflect"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/runtime"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/strconv"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/strings"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/sync"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/terminal"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/tpl/extractor"
	"github.com/redshiftkeying/slowflow-server/qlang/lib/version"
	qlang "github.com/redshiftkeying/slowflow-server/qlang/spec"

	// qlang builtin modules
	_ "github.com/redshiftkeying/slowflow-server/qlang/lib/builtin"
	_ "github.com/redshiftkeying/slowflow-server/qlang/lib/chan"
)

// -----------------------------------------------------------------------------

// Copyright prints qlang copyright information.
//
func Copyright() {
	version.Copyright()
}

// InitSafe inits qlang and imports modules.
//
func InitSafe(safeMode bool) {

	qlang.SafeMode = safeMode

	qlang.Import("", math.Exports) // import math as builtin package
	qlang.Import("", meta.Exports) // import meta package
	qlang.Import("bufio", bufio.Exports)
	qlang.Import("bytes", bytes.Exports)
	qlang.Import("md5", md5.Exports)
	qlang.Import("io", io.Exports)
	qlang.Import("ioutil", ioutil.Exports)
	qlang.Import("hex", hex.Exports)
	qlang.Import("json", json.Exports)
	qlang.Import("errors", errors.Exports)
	qlang.Import("eqlang", eqlang.Exports)
	qlang.Import("math", math.Exports)
	qlang.Import("os", os.Exports)
	qlang.Import("", os.InlineExports)
	qlang.Import("path", path.Exports)
	qlang.Import("http", http.Exports)
	qlang.Import("reflect", reflect.Exports)
	qlang.Import("runtime", runtime.Exports)
	qlang.Import("strconv", strconv.Exports)
	qlang.Import("strings", strings.Exports)
	qlang.Import("sync", sync.Exports)
	qlang.Import("terminal", terminal.Exports)
	qlang.Import("extractor", extractor.Exports)
}

// -----------------------------------------------------------------------------
