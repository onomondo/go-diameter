// Copyright 2013 Alexandre Fiori
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package dict

import "testing"

func TestFindAVP(t *testing.T) {
	if _, err := testParser.FindAVP(0, 263); err != nil {
		t.Fatal(err)
	}
}

func TestScanAVP(t *testing.T) {
	if avp, err := testParser.ScanAVP("Session-Id"); err != nil {
		t.Error(err)
	} else if avp.Code != 263 {
		t.Fatalf("Unexpected code %d for Session-Id AVP", avp.Code)
	}
}

func TestFindCmd(t *testing.T) {
	if cmd, err := testParser.FindCmd(0, 257); err != nil {
		t.Error(err)
	} else if cmd.Short != "CE" {
		t.Fatalf("Unexpected command: %#v", cmd)
	}
}

func TestEnum(t *testing.T) {
	if item, err := testParser.Enum(0, 274, 1); err != nil {
		t.Fatal(err)
	} else if item.Name != "AUTHENTICATE_ONLY" {
		t.Errorf(
			"Unexpected value %s, expected AUTHENTICATE_ONLY",
			item.Name,
		)
	}
}

func TestRule(t *testing.T) {
	if rule, err := testParser.Rule(0, 284, "Proxy-Host"); err != nil {
		t.Fatal(err)
	} else if !rule.Required {
		t.Errorf("Unexpected rule %#v", rule)
	}
}