package api

import (
	"log"
	"strconv"
	"testing"
)

func TestVK_WallPost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	post, vkErr := vkUser.WallPost(map[string]string{
		"message": "Test post",
	})
	if vkErr.Code != 0 {
		log.Fatal(vkErr)
	}

	_, gotVkErr := vkUser.WallPin(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallPin() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.WallUnpin(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallUnpin() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.WallCloseComments(map[string]string{
		"post_id":  strconv.Itoa(post.PostID),
		"owner_id": strconv.Itoa(vkUserID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallCloseComments() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.WallOpenComments(map[string]string{
		"post_id":  strconv.Itoa(post.PostID),
		"owner_id": strconv.Itoa(vkUserID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallOpenComments() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.WallEdit(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
		"message": "Test post edited",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallEdit() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.WallDelete(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallDelete() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.WallRestore(map[string]string{
		"post_id": strconv.Itoa(post.PostID),
		"message": "Test post edited",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallRestore() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_WallGet(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	params := map[string]string{
		"owner_id": "-86529522",
		"count":    "100",
	}

	_, gotVkErr := vkService.WallGet(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallGet() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkService.WallGetExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallGetExtended() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_WallGetByID(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	params := map[string]string{
		"posts": "-1_340393",
	}

	_, gotVkErr := vkService.WallGetByID(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallGetByID() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkService.WallGetByIDExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallGetByIDExtended() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_WallGetComment(t *testing.T) {
	// BUG(VK): https://github.com/SevereCloud/vksdk/issues/55
	t.Skip("BUG")
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	params := map[string]string{
		"owner_id":   "66559",
		"comment_id": "73674",
	}

	_, gotVkErr := vkUser.WallGetComment(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallGetComment() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.WallGetCommentExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallGetCommentExtended() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_WallGetComments(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	params := map[string]string{
		"owner_id":           "85635407",
		"post_id":            "3199",
		"need_likes":         "1",
		"count":              "100",
		"thread_items_count": "10",
	}

	_, gotVkErr := vkService.WallGetComments(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallGetComments() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkService.WallGetCommentsExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallGetCommentsExtended() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_WallGetReposts(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	_, gotVkErr := vkService.WallGetReposts(map[string]string{
		"owner_id": "1",
		"post_id":  "2442097",
		"count":    "1000",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallGetReposts() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_WallCreateComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	comment, vkErr := vkUser.WallCreateComment(map[string]string{
		"owner_id": "117253521",
		"post_id":  "2342",
		"message":  "Test comment",
	})
	if vkErr.Code != 0 {
		log.Fatal(vkErr)
	}

	_, gotVkErr := vkUser.WallEditComment(map[string]string{
		"owner_id":   "117253521",
		"comment_id": strconv.Itoa(comment.CommentID),
		"message":    "Test comment edited",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallEditComment() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.WallDeleteComment(map[string]string{
		"owner_id":   "117253521",
		"comment_id": strconv.Itoa(comment.CommentID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallDeleteComment() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.WallRestoreComment(map[string]string{
		"owner_id":   "117253521",
		"comment_id": strconv.Itoa(comment.CommentID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallRestoreComment() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_WallPostAdsStealth(t *testing.T) {
	if vkUser.AccessToken == "" || vkGroupID == 0 {
		t.Skip("USER_TOKEN empty or GroupID not found")
	}

	post, gotVkErr := vkUser.WallPostAdsStealth(map[string]string{
		"owner_id": strconv.Itoa(-vkGroupID),
		"message":  "Test AdsStealth",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallPostAdsStealth() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkUser.WallEditAdsStealth(map[string]string{
		"owner_id": strconv.Itoa(-vkGroupID),
		"post_id":  strconv.Itoa(post.PostID),
		"message":  "Test AdsStealth edited",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallEditAdsStealth() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_WallReportComment(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.WallReportComment(map[string]string{
		"owner_id":   "66748",
		"comment_id": "4136",
		// "reason":   "3",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallReportComment() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_WallReportPost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.WallReportPost(map[string]string{
		"owner_id": "66748",
		"post_id":  "3821",
		"reason":   "3",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallReportPost() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_WallRepost(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.WallRepost(map[string]string{
		"object": "wall85635407_3133",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallRepost() gotVkErr = %v", gotVkErr)
	}
}

func TestVK_WallSearch(t *testing.T) {
	if vkService.AccessToken == "" {
		t.Skip("SERVICE_TOKEN empty")
	}

	params := map[string]string{
		"owner_id": "6",
		"query":    "vk",
		"count":    "100",
	}

	_, gotVkErr := vkService.WallSearch(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallSearch() gotVkErr = %v", gotVkErr)
	}

	_, gotVkErr = vkService.WallSearchExtended(params)
	if gotVkErr.Code != 0 {
		t.Errorf("VK.WallSearchExtended() gotVkErr = %v", gotVkErr)
	}
}
