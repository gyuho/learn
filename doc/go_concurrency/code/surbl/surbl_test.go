package surbl

import "testing"

var goodSlice = []string{
	"google.com", "amazon.com", "yahoo.com", "gmail.com", "walmart.com",
	"stanford.edu", "intel.com", "github.com", "surbl.org", "wikipedia.org",
}

// Example spam websites
// (http://freakipedia.net/index.php5?title=Spam_Site_List)
var spamSlice = []string{
	"safepillstock.com", "pharm4all.com",
	"proofofpayday.com", "addacheckingaccount.com",
	"addacheckingaccount.com", "creditcardofcourse.com",
}

func TestSurblGood(t *testing.T) {
	rmap1 := checkWithLockBatch(goodSlice...)
	for k, v := range rmap1 {
		if v.IsSpam {
			t.Logf("checkWithLockBatch | Unexpected %+v %+v but it's ok", k, v)
		}
	}
	rmap2 := Check(goodSlice...)
	for k, v := range rmap2 {
		if v.IsSpam {
			t.Logf("Check | Unexpected %+v %+v but it's ok", k, v)
		}
	}
}

func TestSurblSpam(t *testing.T) {
	rmap1 := checkWithLockBatch(spamSlice...)
	for k, v := range rmap1 {
		if !v.IsSpam {
			t.Logf("checkWithLockBatch | Unexpected %+v %+v but it's ok", k, v)
		}
	}
	rmap2 := Check(spamSlice...)
	for k, v := range rmap2 {
		if !v.IsSpam {
			t.Logf("Check | Unexpected %+v %+v but it's ok", k, v)
		}
	}
}
