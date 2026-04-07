package security
import "testing"
func TestHelperHashSmoke(t *testing.T) { if len(Digest([]byte("seed"))) == 0 { t.Fatal("expected bytes") } }
