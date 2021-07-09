package goutils

import "testing"

func TestPaginateByPageNumSize(t *testing.T) {
	p := PaginateByPageNumSize(100, 1, 10)
	if p.PagesCount != 10 || p.HasNext != true || p.HasPrev != false {
		t.Fatalf("paginate error %+v", p)
	}
	t.Logf("%+v", p)
	p = PaginateByPageNumSize(100, 3, 13)
	if p.PagesCount != 8 || p.HasNext != true || p.HasPrev != true {
		t.Fatalf("paginate error %+v", p)
	}
	p = PaginateByPageNumSize(100, 10, 10)
	if p.PagesCount != 10 || p.HasNext != false || p.HasPrev != true {
		t.Fatalf("paginate error %+v", p)
	}
}

func TestPaginateByOffsetLimit(t *testing.T) {
	p := PaginateByOffsetLimit(100, 0, 10)
	if p.HasNext != true || p.HasPrev != false || p.PageNum != 1 {
		t.Fatalf("paginate error %+v", p)
	}
	p = PaginateByOffsetLimit(100, 10, 10)
	if p.HasNext != true || p.HasPrev != true || p.PageNum != 2 {
		t.Fatalf("paginate error %+v", p)
	}
	p = PaginateByOffsetLimit(100, 99, 10)
	if p.HasNext != false || p.HasPrev != true || p.PageNum != 10 {
		t.Fatalf("paginate error %+v", p)
	}
}
