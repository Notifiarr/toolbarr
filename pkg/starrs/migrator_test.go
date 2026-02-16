package starrs

import "testing"

func TestRewritePathForTablePathColumn(t *testing.T) {
	table := TableColumn{Table: "Series", Column: "Path"}

	updated, ok := rewritePathForTable(
		"/data/media/TV/Show",
		table,
		trailingSlash("/data/media"),
		trailingSlash("/mnt/storage"),
	)
	if !ok {
		t.Fatal("expected path rewrite to apply for Path column")
	}

	if updated != "/mnt/storage/TV/Show" {
		t.Fatalf("unexpected rewritten path: %s", updated)
	}
}

func TestRewritePathForTableRelativePathWinToLinux(t *testing.T) {
	table := TableColumn{Table: "EpisodeFiles", Column: "RelativePath"}

	updated, ok := rewritePathForTable(
		`Season 01\Show - S01E01.mkv`,
		table,
		`D:\\Media\\TV\\Show`,
		`/mnt/media/TV/Show`,
	)
	if !ok {
		t.Fatal("expected path rewrite to apply for RelativePath column")
	}

	if updated != "Season 01/Show - S01E01.mkv" {
		t.Fatalf("unexpected rewritten relative path: %s", updated)
	}
}

func TestRewritePathForTableRelativePathLinuxToWin(t *testing.T) {
	table := TableColumn{Table: "SubtitleFiles", Column: "RelativePath"}

	updated, ok := rewritePathForTable(
		"Season 01/Show - S01E01.en.srt",
		table,
		"/mnt/media/TV/Show",
		`D:\Media\TV\Show`,
	)
	if !ok {
		t.Fatal("expected path rewrite to apply for RelativePath column")
	}

	if updated != `Season 01\Show - S01E01.en.srt` {
		t.Fatalf("unexpected rewritten relative path: %s", updated)
	}
}

func TestAppTablesIncludeSonarrRelativePathTables(t *testing.T) {
	sonarr := AppTables("Sonarr")
	if _, ok := sonarr["Series"]; !ok {
		t.Fatal("Sonarr table map missing Series")
	}
	if _, ok := sonarr["ImportLists"]; !ok {
		t.Fatal("Sonarr table map missing ImportLists")
	}

	for _, name := range []string{"EpisodeFiles", "MetadataFiles", "SubtitleFiles"} {
		column, ok := sonarr[name]
		if !ok {
			t.Fatalf("Sonarr table map missing %s", name)
		}
		if column.Column != "RelativePath" {
			t.Fatalf("expected %s column to be RelativePath, got %s", name, column.Column)
		}
	}
}

func TestAppTablesIncludeWhisparrRelativePathTables(t *testing.T) {
	whisparr := AppTables("Whisparr")
	for _, name := range []string{"EpisodeFiles", "MetadataFiles", "SubtitleFiles"} {
		column, ok := whisparr[name]
		if !ok {
			t.Fatalf("Whisparr table map missing %s", name)
		}
		if column.Column != "RelativePath" {
			t.Fatalf("expected %s column to be RelativePath, got %s", name, column.Column)
		}
	}
}
