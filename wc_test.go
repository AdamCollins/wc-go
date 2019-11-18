package main

import "testing"

func TestWC1Word( t *testing.T) {
	efn := "testfiles/1word.txt"
	elc := 0
	ewc := 1
	ebc := 5

	lc,wc,bc,fn := wc(efn)

	if lc!=elc || wc!=ewc || bc!=ebc || fn!=efn {
		t.Errorf("wc(\"%s\") failed, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}else{
		t.Logf("wc(\"%s\") success, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}
}

func TestWC2Words( t *testing.T) {
	efn := "testfiles/2words.txt"
	elc := 0
	ewc := 2
	ebc := 11

	lc,wc,bc,fn := wc(efn)

	if lc!=elc || wc!=ewc || bc!=ebc || fn!=efn {
		t.Errorf("wc(\"%s\") failed, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}else{
		t.Logf("wc(\"%s\") success, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}
}

func TestWCEmpty( t *testing.T) {
	efn := "testfiles/empty.txt"
	elc := 0
	ewc := 0
	ebc := 0

	lc,wc,bc,fn := wc(efn)

	if lc!=elc || wc!=ewc || bc!=ebc || fn!=efn {
		t.Errorf("wc(\"%s\") failed, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}else{
		t.Logf("wc(\"%s\") success, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}
}

func TestWCLoremIpsum( t *testing.T) {
	efn := "testfiles/loremipsum.txt"
	elc := 0
	ewc := 74
	ebc := 507

	lc,wc,bc,fn := wc(efn)

	if lc!=elc || wc!=ewc || bc!=ebc || fn!=efn {
		t.Errorf("wc(\"%s\") failed, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}else{
		t.Logf("wc(\"%s\") success, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}
}

func TestWCLoremIpsumLineBreaks( t *testing.T) {
	efn := "testfiles/loremipsum-linebreaks.txt"
	elc := 4
	ewc := 74
	ebc := 511

	lc,wc,bc,fn := wc(efn)

	if lc!=elc || wc!=ewc || bc!=ebc || fn!=efn {
		t.Errorf("wc(\"%s\") failed, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}else{
		t.Logf("wc(\"%s\") success, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}
}

func TestWCBeyondGoodAndEvil( t *testing.T) {
	efn := "testfiles/bge.txt"
	elc := 6506
	ewc := 65692
	ebc := 408778

	lc,wc,bc,fn := wc(efn)

	if lc!=elc || wc!=ewc || bc!=ebc || fn!=efn {
		t.Errorf("wc(\"%s\") failed, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}else{
		t.Logf("wc(\"%s\") success, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}
}

func TestWCAsciiArt( t *testing.T) {
	efn := "testfiles/asciiart.txt"
	elc := 15
	ewc := 35
	ebc := 350

	lc,wc,bc,fn := wc(efn)

	if lc!=elc || wc!=ewc || bc!=ebc || fn!=efn {
		t.Errorf("wc(\"%s\") failed, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}else{
		t.Logf("wc(\"%s\") success, expected %d %d %d %s got %d %d %d %s",fn, elc,ewc,ebc,efn, lc,wc,bc,fn)
	}
}


func TestWCUtf8( t *testing.T) {
	efn := "testfiles/UTF8.txt"
	elc := 0
	ewc := 2
	ebc := 16

	lc,wc,bc,fn := wc(efn)

	if lc!=elc || wc!=ewc || bc!=ebc || fn!=efn {
		t.Errorf("wc(\"%s\") failed, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}else{
		t.Logf("wc(\"%s\") success, expected %d %d %d %s got %d %d %d %s",fn,elc,ewc,ebc,efn, lc,wc,bc,fn)
	}
}
