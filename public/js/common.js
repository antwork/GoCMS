/**
 * tips提示消息
 * 
 * @param content
 */
function notice_tips(content) {
	art.dialog.tips(content, 1.5);
}

/**
 * 页面跳转
 * 
 * @param url
 */
function redirect(url) {
	location.href = url;
}

/**
 * 正确提示对话框
 * 
 * @param content
 */
function notice(content) {
	art.dialog({
		lock : true, // 锁屏
		icon : 'succeed',
		title : '提示',
		content : content
	}).time(3);
}

/**
 * 错误提示对话框
 * 
 * @param content
 */
function notice_error(content) {
	art.dialog({
		lock : true, // 锁屏
		icon : 'error',
		title : '提示',
		content : content
	}).time(3);
}

/**
 * 后台地图
 */
function OpenMap() {
	art.dialog.open('/Map/', {
		id : 'map',
		title : '后台地图',
		width : 700,
		height : 500,
		lock : true
	});
}

/**
 * 退出
 */
function Logout() {
	art.dialog.confirm('你确认退出吗?', function() {
		redirect("/Logout/")
	}, function() {
		
	});
}

/**
 * 右侧内容页刷新
 */
function right_refresh() {
	art.dialog.top.$('#rightMain').attr('src',art.dialog.top.$('#rightMain').attr("src"));
}

/**
 * 设置iframe
 * @param id
 * @param src
 */
function set_iframe(id, src) {
	$("#" + id).attr("src", src);
}

// 滚动条
$(function() {
	$(":text").addClass('input-text');
});

/**
 * 全选checkbox,注意：标识checkbox id固定为为check_box
 * 
 * @param string
 *            name 列表check名称,如 uid[]
 * 
 */
function selectall(name) {
	if ($("#check_box").attr("checked") == false
			|| $("#check_box").attr("checked") == undefined) {
		$("input[name='" + name + "']").each(function() {
			this.checked = false;
		});
	} else {
		$("input[name='" + name + "']").each(function() {
			this.checked = true;
		});
	}
}

/**
 * 上传
 * 
 * @param type
 * @param textareaid
 */
function upload(type, textareaid) {
	art.dialog.open('/Attachment/index/' + type + '/' + textareaid + '/', {
		id : "thumb_images",
		title : "附件上传",
		width : 500,
		height : 420,
		lock : true,
		ok : function() {
			var iframe = this.iframe.contentWindow;
			
			var fsUploadProgress = iframe.$('#fsUploadProgress');

			if (fsUploadProgress == '') {
				art.dialog.alert('参数错误!');
				return false;
			}
			
			var img_src = fsUploadProgress.find(".icon").prev("img").attr("src");
			window.top.$("#thumb_preview").attr("src", img_src);
			
			var small_img_src = fsUploadProgress.find(".icon").prev("img").attr("small_src");
			window.top.$("#" + textareaid).val(small_img_src);
			
			return true;
		},
		cancel : true
	});
}

function openwinx(url, name, w, h) {
	if (!w)
		w = screen.width - 4;
	if (!h)
		h = screen.height - 95;

	window.open(url,name,"top=100,left=400,width="
		+ w
		+ ",height="
		+ h
		+ ",toolbar=no,menubar=no,scrollbars=yes,resizable=yes,location=no,status=no");
}

// 弹出对话框
function omnipotent(id, linkurl, title, close_type, w, h) {
	if (!w) {
		w = 700;
	}

	if (!h) {
		h = 500;
	}

	art.dialog.open(linkurl, {
		id : id,
		title : title,
		width : w,
		height : h,
		lock : true
	});
}