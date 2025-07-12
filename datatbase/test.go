package main

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/jl-sky/grom/golangNotes/datatbase/models"
	log "github.com/sirupsen/logrus"
)

func changelogAdminTest() {
	req := mockReq()
	ChangeLogByAdmin(context.Background(), req)
}
func mockReq() *models.ChangelogAdminReq {
	rowData := "{\"c_modifier\":\"cheersjiang\",\"c_strategy_id\":\"9d115c1bad99493db72245b27df89885\",\"c_family_id\":\"3019164517a844e89e7f23699744f3ee\",\"c_parent_id\":\"3019164517a844e89e7f23699744f3ee\",\"\":\"platform\",\"c_data_type_value\":\"harmony_phone\",\"c_status\":1,\"c_mtime\":\"2025-07-10 20:27:58\",\"c_page_uuid\":\"cc54892bd8964078900eeff6b1bdc969\",\"c_creator\":\"sennazang\",\"c_version\":\"81\",\"c_version_type\":\"release\",\"c_page_class\":\"\",\"c_relation_id\":\"\",\"c_is_patch\":false,\"c_type\":\"default\",\"c_config\":{\"ad_params\":{\"param_pkg_uuid\":\"\",\"params\":{}},\"pagination_config\":[1],\"folding_pagination_config\":[1],\"pageprocess_rpc\":[{\"data_src_uuid\":\"f645769540a38dcc\",\"data_src_conf\":{\"technology_conf\":{\"dynamic_timeout_type\":1,\"origin_timeout\":\"163\",\"max_timeout\":\"266\",\"avg_cost_cfg\":{\"hourly_cost_config\":[{\"hour\":0,\"avg_cost\":300},{\"hour\":1,\"avg_cost\":400},{\"hour\":15,\"avg_cost\":600}],\"default_avg_cost\":1000}},\"operation_conf\":{}}},{\"data_src_uuid\":\"6deb231e45888339\",\"data_src_conf\":{\"technology_conf\":{},\"operation_conf\":{}}}],\"layout_conf\":{\"param_pkg_uuid\":\"\",\"params\":{}},\"page_properties\":{\"param_pkg_uuid\":\"\",\"params\":{}},\"extend_params\":[{\"param_pkg_uuid\":\"6460f5dc4a9ba754\",\"params\":{\"param_usage\":\"ad_dynamic_timeout\",\"enable_dynamic_timeout\":1,\"origin_timeout\":\"150\",\"max_timeout\":\"250\"}}],\"modules\":[{\"diff_id\":\"db819_131e9\",\"mod_index\":1,\"min_version\":\"\",\"max_version\":\"\",\"is_lock\":false,\"repeated_num\":10000,\"ad_params\":{\"param_pkg_uuid\":\"2095337e453ab6bb\",\"params\":{\"ad_type\":\"\"}},\"card_conf\":[{\"card_id\":\"f849b1484cdb862e\",\"card_name\":\"un_proto_hold\",\"type\":1,\"adaptor_uuid\":\"\",\"general_conf\":{\"param_pkg_uuid\":\"\",\"params\":{}},\"custom_conf\":{\"param_pkg_uuid\":\"\",\"params\":{}},\"layout_conf\":{\"param_pkg_uuid\":\"\",\"params\":{}},\"card_list\":[]}],\"data_src_uuid\":\"d09c76ed4a62abf2\",\"data_src_conf\":{\"technology_conf\":{\"accessID\":\"discover_120121\",\"scene\":\"rec\",\"wayToReadVuid\":\"ExtractFromMap\",\"control_has_next_page\":true,\"repeated_num\":\"0\",\"ext_data\":\"enable_immersive_star_float=true&enable_paid_album_tag=true&paid_album_tag_text=付费专辑&enable_immersive_resume_tag=true&enable_invite_tag=true&enable_second_creation_card=true\",\"is_fliter_empty_data\":\"1\",\"host_access_id\":\"\",\"feeds_title\":\"\",\"extend_config_params\":\"\",\"set_section_id\":\"\",\"is_pack_single_module\":\"\",\"title_extra_keys\":\"\",\"req_para_set_to_module\":\"\",\"feeds_platform_set\":\"default.gz.2\",\"is_skip_cid_check\":\"0\",\"item_to_module_params\":\"\",\"subscription_insert_title\":\"\",\"keys_from_page_params\":\"\",\"persist_keys_from_page_params\":\"\",\"no_setting_paging\":\"\",\"build_multi_recallitem_with_entity_sublist_key\":\"\",\"replace_title_key\":\"\",\"exp_toast_bubble\":\"\",\"block_style_type\":\"5\",\"is_float_card\":\"0\",\"is_hot_spot_card\":\"0\",\"feed_data_key\":\"is_auto_play=1&is_auto_play_next=0&is_loop_play=1&play_end_type=0&is_auto_mute=0&play_delay=0&is_show_mute_btn=0\",\"user_nick_key\":\"user_nick\",\"user_head_key\":\"user_head\",\"user_outline_logo_key\":\"user_outline_logo\",\"title_key\":\"title_new\",\"poster_subtitle_key\":\"vid_episodes\",\"image_key\":\"pic_new\",\"hori_image_key\":\"hori_image\",\"aspect_key\":\"aspect\",\"type_key\":\"type\",\"upload_time_key\":\"upload_time\",\"cid_key\":\"cid\",\"cid_title_key\":\"cid_title\",\"cid_sub_title_key\":\"cid_second_title\",\"praise_count_key\":\"video_like_num\",\"praise_status_key\":\"video_like_status\",\"duration_key\":\"duration\",\"duration_str_key\":\"duration_str\",\"feed_back_key\":\"item_datakey_info\",\"desc_key\":\"desc\",\"vcuid_key\":\"vcuid\",\"comment_count_key\":\"video_comment_num\",\"pid_key\":\"live_last_pid\",\"stream_url_key\":\"stream_url\",\"stream_ratio_key\":\"stream_ratio\",\"live_room_status_key\":\"playing_status\",\"tag_list_key\":\"tag\",\"longrec_tag_lottie_progress\":\"30\",\"is_traffic_spread\":\"1\",\"vuid_user_nick_key\":\"vuid_user_nick\",\"vuid_account_id_key\":\"vuid_account_id\",\"vuid_key\":\"vuid\",\"poster_subtitle_chain_key\":\"\",\"cid_subtitle\":\"\",\"float_icon\":\"http://puui.qpic.cn/media_img/lena/PICz6xdpe_72_72/0\",\"is_show_hot_event_rank_info\":\"1\",\"is_topic_float_card\":\"\",\"is_tag_show_next_video\":\"\",\"is_no_render_mzc_feedback\":\"0\",\"material_starting_time_key\":\"material_starting_time\",\"unable_share\":\"0\",\"unable_comment\":\"0\",\"unable_danmu\":\"0\",\"unable_show_avatar\":\"0\",\"unable_subscribe\":\"0\",\"is_ranklist_tag\":\"0\",\"pub_time_switch\":false,\"tag_config_switch\":\"0\",\"pad_new_style\":false,\"is_ip_topic_float_card\":\"\",\"live_immersive_board_unify_style\":false,\"is_show_resume_toast_tip\":false,\"enable_middle_album_card\":false,\"live_chid\":false,\"allow_fake_live\":false,\"live_immersive_board_optimize\":false,\"enable_middle_album_card_skip_start\":false,\"enable_middle_album_action_selections\":false,\"enable_middle_album_card_extend\":false,\"enable_use_rec_cacheable\":false,\"cacheable_cate_list\":[],\"enable_open_cate_resume\":false,\"resume_cate_list\":[],\"enable_open_episode_list\":false,\"is_immersive_album_second_page\":false,\"immersive_album_top_word\":\"0\",\"more_short_videos_guide_entry\":false,\"close_resume_action\":false,\"recr_mid_resume\":false,\"recr_mid_small_tag\":false,\"short_drama_new_ui\":false,\"module_item_num\":\"10\"},\"operation_conf\":{\"module_item_num\":\"10\"}},\"selected_sub_module_id\":\"\",\"hide_in_data_mode\":[],\"is_pagination\":true,\"ad_type\":\"no_ad\",\"id\":\"db819_131e9\",\"mod_uuid\":\"db819afe4307a056\",\"mod_type\":\"feeds_page_service_direct_render\",\"mod_name\":\"feeds流中台适配模块(UN无脚本层)\",\"description\":\"feeds流中台适配模块，走UN无脚本层协议的模块\",\"is_container\":false,\"wrapper_type\":\"normal\",\"control_paging\":true,\"sub_modules\":[],\"is_from_cms\":false,\"tab_module\":\"\"}],\"strategy_id\":\"9d115c1bad99493db72245b27df89885\",\"strategy_name\":\"发现120121-默认配置-默认策略\"}}"
	// rowData := "{\"c_modifier\":\"cheersjiang\",\"c_strategy_id\":\"9d115c1bad99493db72245b27df89885\",\"c_family_id\":\"3019164517a844e89e7f23699744f3ee\",\"c_parent_id\":\"3019164517a844e89e7f23699744f3ee\",\"c_data_type\":\"platform\",\"c_data_type_value\":\"harmony_phone\",\"c_status\":1,\"c_mtime\":\"2025-07-10 20:27:58\",\"c_page_uuid\":\"cc54892bd8964078900eeff6b1bdc969\",\"c_creator\":\"sennazang\",\"c_version\":\"81\",\"c_version_type\":\"release\",\"c_page_class\":\"\",\"c_relation_id\":\"\",\"c_is_patch\":false,\"c_type\":\"default\",\"c_config\":\"config\"}"
	changeLogAdmin := &models.ChangelogAdminReq{
		TableName:  "t_output",
		ChangeUser: "mvl_debug",
		RowData:    rowData,
	}
	return changeLogAdmin
}

// TestConcurrentUpdates 测试不同平台的并发变更
func TestConcurrentUpdates() {
	// 准备基础测试数据
	baseData := map[string]interface{}{
		"c_family_id":    "family_123", // 固定familyID
		"c_data_type":    "platform",
		"c_status":       1,
		"c_mtime":        time.Now().Format("2006-01-02 15:04:05"),
		"c_page_uuid":    "page_123",
		"c_creator":      "test_user",
		"c_version":      "1.0",
		"c_version_type": "release",
		"c_config":       map[string]interface{}{"key": "value"},
	}

	// 定义要测试的不同平台
	platforms := []string{"mac", "pc", "harmony", "android"}

	// 准备并发测试的请求
	var requests []*models.ChangelogAdminReq
	for _, platform := range platforms {
		// 复制基础数据
		data := make(map[string]interface{})
		for k, v := range baseData {
			data[k] = v
		}
		// 设置不同的dataTypeValue
		data["c_data_type_value"] = platform

		// 转换为JSON字符串
		rowData, err := json.Marshal(data)
		if err != nil {
			log.Errorf("Failed to marshal data for platform %s: %v", platform, err)
			continue
		}

		requests = append(requests, &models.ChangelogAdminReq{
			TableName:  "t_output",
			ChangeUser: "tester",
			RowData:    string(rowData),
		})
	}

	// 模拟并发请求
	var wg sync.WaitGroup
	for i, req := range requests {
		wg.Add(1)
		go func(id int, request *models.ChangelogAdminReq) {
			defer wg.Done()

			ctx := context.Background()
			err := ChangeLogByAdmin(ctx, request)
			if err != nil {
				log.Errorf("Goroutine %d (%s) failed: %v",
					id, request.RowData, err)
			} else {
				log.Infof("Goroutine %d (%s) succeeded",
					id, request.RowData)
			}
		}(i, req)
	}

	wg.Wait()
	log.Info("All concurrent updates completed")
}
