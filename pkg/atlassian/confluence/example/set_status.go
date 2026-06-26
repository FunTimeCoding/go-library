package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page/page_put"
)

func SetStatus() {
	c := confluence.NewEnvironment()

	// 1. Fetch the draft page we created earlier (160497673)
	fmt.Println("=== Created draft page ===")
	draftPage, e := c.Basic().GetV2(
		c.Basic().Base().Copy().Path(
			"%s/%s",
			constant.Page,
			"160497673",
		).Set(constant.BodyFormat, constant.StorageFormat).
			Set(constant.Status, constant.DraftStatus).String(),
	)

	if e != nil {
		fmt.Printf("draft page fetch error: %v\n\n", e)
	} else {
		fmt.Printf("draft page:\n%s\n\n", draftPage)
	}

	// 2. Publish the draft page (draft → current)
	fmt.Println("=== Publish draft page ===")
	published, f := c.Basic().PutV2Path(
		fmt.Sprintf("%s/%s", constant.Page, "160497673"),
		page_put.NewWithStatus(
			"160497673",
			"Draft Test",
			"<p>This page was created as a draft.</p>",
			1,
			"",
			constant.CurrentStatus,
		).Encode(),
	)

	if f != nil {
		fmt.Printf("publish error: %v\n\n", f)
	} else {
		fmt.Printf("published:\n%s\n\n", published)
	}

	// 3. Fetch it now as current
	fmt.Println("=== Fetch as current after publish ===")
	afterPublish, g := c.Basic().GetV2(
		c.Basic().Base().Copy().Path(
			"%s/%s",
			constant.Page,
			"160497673",
		).Set(constant.BodyFormat, constant.StorageFormat).String(),
	)

	if g != nil {
		fmt.Printf("fetch after publish error: %v\n\n", g)
	} else {
		fmt.Printf("after publish:\n%s\n\n", afterPublish)
	}

	// 4. Create a draft overlay with different content
	fmt.Println("=== Create draft overlay with changed content ===")
	overlay, h := c.Basic().PutV2Path(
		fmt.Sprintf("%s/%s", constant.Page, "160497673"),
		page_put.NewWithStatus(
			"160497673",
			"Draft Test",
			"<p>This is the draft overlay with different content.</p>",
			1,
			"",
			constant.DraftStatus,
		).Encode(),
	)

	if h != nil {
		fmt.Printf("overlay error: %v\n\n", h)
	} else {
		fmt.Printf("overlay:\n%s\n\n", overlay)
	}

	// 5. Try to fetch the draft overlay
	fmt.Println("=== Fetch draft overlay ===")
	draftOverlay, i := c.Basic().GetV2(
		c.Basic().Base().Copy().Path(
			"%s/%s",
			constant.Page,
			"160497673",
		).Set(constant.BodyFormat, constant.StorageFormat).
			Set(constant.Status, constant.DraftStatus).String(),
	)

	if i != nil {
		fmt.Printf("draft overlay fetch error: %v\n\n", i)
	} else {
		fmt.Printf("draft overlay:\n%s\n\n", draftOverlay)
	}

	// 6. Fetch the published version (should still have old content)
	fmt.Println("=== Published version while overlay exists ===")
	currentWhileOverlay, j := c.Basic().GetV2(
		c.Basic().Base().Copy().Path(
			"%s/%s",
			constant.Page,
			"160497673",
		).Set(constant.BodyFormat, constant.StorageFormat).String(),
	)

	if j != nil {
		fmt.Printf("current fetch error: %v\n\n", j)
	} else {
		fmt.Printf("current:\n%s\n\n", currentWhileOverlay)
	}

	// 7. Publish the overlay (should apply the changed content)
	fmt.Println("=== Publish the overlay ===")
	publishOverlay, k := c.Basic().PutV2Path(
		fmt.Sprintf("%s/%s", constant.Page, "160497673"),
		page_put.NewWithStatus(
			"160497673",
			"Draft Test",
			"<p>This is the draft overlay with different content.</p>",
			2,
			"",
			constant.CurrentStatus,
		).Encode(),
	)

	if k != nil {
		fmt.Printf("publish overlay error: %v\n\n", k)
	} else {
		fmt.Printf("published overlay:\n%s\n\n", publishOverlay)
	}

	// 8. Final state
	fmt.Println("=== Final state ===")
	final, l := c.Basic().GetV2(
		c.Basic().Base().Copy().Path(
			"%s/%s",
			constant.Page,
			"160497673",
		).Set(constant.BodyFormat, constant.StorageFormat).String(),
	)

	if l != nil {
		fmt.Printf("final fetch error: %v\n\n", l)
	} else {
		fmt.Printf("final:\n%s\n\n", final)
	}
}
